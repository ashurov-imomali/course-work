package service

import (
	"back-end/pkg/models"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"log"
	"math/big"
	"net/http"
	"net/smtp"
	"regexp"
	"strconv"
	"time"
)

const SaltLength = 7

type Service struct {
	rClient *redis.Client
	Repo    ReposMeths
	Auth    smtp.Auth
}

func GetService(repo ReposMeths, r *redis.Client) SrvMeths {
	return &Service{Repo: repo, rClient: r}
}

func (s *Service) GetServices(typeId, id int64, name string) ([]models.Service, error) {
	return s.Repo.SelectServices(typeId, id, name)
}

func (s *Service) GetService(id int64) (*models.ServiceResponse, error) {
	return s.Repo.ServiceById(id)
}

func (s *Service) Login(login *models.Login) (string, *Error) {
	client, err := s.Repo.GetClientByLogin(login.Login)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", s.internalError(err, "can't get client bu login")
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", s.badRequest(errors.New("has not client with this login"), "wrong login")
	}
	ok, err := s.checkPasswordAndHash(login.Password, client.Password, client.Salt)
	if err != nil {
		return "", s.internalError(err, "can't check password")
	}
	if !ok {
		return "", s.badRequest(errors.New("wrong password"), "wrong password")
	}
	jwtClaims := make(map[string]interface{})
	jwtClaims["id"] = client.Id
	jwtClaims["priority_id"] = client.PriorityId
	jwtClaims["expiration_time"] = time.Now().Add(time.Hour).Unix()
	_, err = s.generateJWT(jwtClaims, client.Salt)
	if err != nil {
		return "", s.internalError(err, "can't generate jwt token")
	}
	return strconv.FormatInt(client.Id, 10), s.Ok()

}

func (s *Service) Registration(newClient *models.Client) *Error {
	_, err := s.Repo.GetClientByLogin(newClient.Login)
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return s.internalError(err, "can't get client by login")
	}
	if !errors.Is(gorm.ErrRecordNotFound, err) {
		return s.badRequest(errors.New("current user"), "this login already uses")
	}

	if ok, strErr := s.validatePassword(newClient.Password); !ok {
		return s.badRequest(errors.New(strErr), strErr)
	}
	salt := make([]byte, SaltLength)
	if _, err := rand.Read(salt); err != nil {
		return s.internalError(err, "can't generate rand salt")
	}
	newClient.Salt = hex.EncodeToString(salt)
	hash, err := s.getPasswordHash(newClient.Password, hex.EncodeToString(salt))
	if err != nil {
		return s.internalError(err, "can't get hash")
	}
	newClient.Password = hash
	if err := s.Repo.CreateClient(newClient); err != nil {
		return s.internalError(err, "can't create new client")
	}
	if err := s.sendOtp(newClient.Email); err != nil {
		return s.internalError(err, err.Error())
	}
	return s.Ok()
}

func (s *Service) sendOtp(email string) error {
	smtpServer := "smtp.gmail.com"
	port := "587"
	senderEmail := "lashurov.l000@gmail.com"
	password := "uwzr hwhf rhdc xdky"

	otp, err := generateNumericOTP(6)
	if err != nil {
		return err
	}
	if err := s.rClient.Set(email, otp, 7*time.Minute).Err(); err != nil {
		return err
	}
	subject := "Subject: Test Email\n"
	body := "Некому не сообщайте данный код \n" + otp
	message := []byte(subject + "\n" + body)
	auth := smtp.PlainAuth("", senderEmail, password, smtpServer)
	if err := smtp.SendMail(smtpServer+":"+port, auth, senderEmail, []string{email}, message); err != nil {
		return err
	}
	return nil
}

func generateNumericOTP(length int) (string, error) {
	const digits = "0123456789"
	otp := make([]byte, length)

	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[index.Int64()]
	}

	return string(otp), nil
}

func (s *Service) validatePassword(password string) (bool, string) {
	hasNum := regexp.MustCompile(`[0-9]`).MatchString
	hasSpecialSymbol := regexp.MustCompile(`[!@#~$%^&*()+|_{}<>?:;,.]`).MatchString
	hasLowerCase := regexp.MustCompile(`[a-z]`).MatchString
	hasUpperCase := regexp.MustCompile(`[A-Z]`).MatchString

	switch {
	case len(password) < 8:
		return false, "Too Shor Password :("
	case !hasNum(password):
		return false, "Password must contain one number"
	case !hasUpperCase(password):
		return false, "Password must contain one upper latin letter"
	case !hasLowerCase(password):
		return false, "Password must contain one lower latin letter"
	case !hasSpecialSymbol(password):
		return false, "Password must contain one special symbol"
	}
	return true, "Password is valid"

}

func (s *Service) internalError(err error, message string) *Error {
	return &Error{
		Err:     err,
		Message: message,
		Status:  http.StatusInternalServerError,
	}
}
func (s *Service) badRequest(err error, message string) *Error {
	return &Error{
		Err:     err,
		Message: message,
		Status:  http.StatusBadRequest,
	}
}
func (s *Service) otherError(err error, message string, status int) *Error {
	return &Error{
		Err:     err,
		Message: message,
		Status:  status,
	}
}
func (s *Service) Ok() *Error {
	return &Error{
		Err:     nil,
		Message: "ok",
		Status:  http.StatusOK,
	}
}

func (s *Service) getPasswordHash(password string, salt string) (string, error) {
	var passHash []byte
	hash := md5.New()
	passHash = append(passHash, []byte(password)...)
	passHash = append(passHash, salt...)
	if _, err := hash.Write(passHash); err != nil {
		return "", err
	}
	log.Println(hex.EncodeToString(hash.Sum(nil)))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func (s *Service) checkPasswordAndHash(password string, hash string, salt string) (bool, error) {
	passwordHash, err := s.getPasswordHash(password, salt)
	if err != nil {
		return false, err
	}
	log.Println(passwordHash, hash)
	return passwordHash == hash, nil
}

func (s *Service) generateJWT(jwtClaims map[string]interface{}, secretKey string) (string, error) {
	mp := jwt.MapClaims{}
	mp = jwtClaims
	return jwt.NewWithClaims(jwt.SigningMethodHS256, mp).SignedString([]byte(secretKey))
}

func (s *Service) VerifyLogin(otp *models.VerifyLogin) *Error {
	user, err := s.Repo.GetClientByLogin(otp.Login)
	if err != nil {
		return s.badRequest(err, "invalid login")
	}
	cmd := s.rClient.Get(user.Email)
	if cmd.Err() != nil {
		return s.internalError(err, "[REDIS] GET")
	}
	if cmd.Val() == otp.Otp {
		user.Verify = true
		return s.otherError(s.Repo.CreateClient(user), "", 500)
	}
	return s.badRequest(errors.New("invalid otp"), "invalid otp")
}

func (s *Service) CreateSrvForClient(srv *models.SrvReq) *Error {
	service, err := s.Repo.GetServiceById(srv.ServiceId)
	if err != nil {
		return s.internalError(err, "Database ERROR")
	}
	otp, err := generateNumericOTP(12)
	if err != nil {
		return s.internalError(err, "couldn't generate otp")
	}
	return s.otherError(s.Repo.CreateCreditSrv(srv, service.TypeId, otp), "database error", 500)
}
