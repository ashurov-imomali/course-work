package models

import "time"

type Role struct {
	Id          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

func (Role) TableName() string {
	return "roles"
}

type Permission struct {
	Id          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

func (Permission) TableName() string {
	return "permissions"
}

type Priority struct {
	Id          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

func (Priority) TableName() string {
	return "priorities"
}

type Department struct {
	Id           int64  `json:"id" gorm:"column:id"`
	Name         string `json:"name" gorm:"column:name"`
	Location     string `json:"location" gorm:"column:location"`
	ContactPhone string `json:"contact_phone" gorm:"column:contact_phone"`
}

type Employer struct {
	Id                 int64      `json:"id" gorm:"id"`
	FirstName          string     `json:"first_name" gorm:"first_name"`
	Surname            string     `json:"surname" gorm:"surname"`
	FatherName         string     `json:"father_name" gorm:"father_name"`
	DateOfBirth        *time.Time `json:"date_of_birth" gorm:"date_of_birth"`
	Gender             bool       `json:"gender" gorm:"gender"`
	Citizenship        string     `json:"citizenship" gorm:"citizenship"`
	PassportSeries     string     `json:"passport_series" gorm:"passport_series"`
	PasswordNumber     int64      `json:"password_number" gorm:"password_number"`
	DateOfIssue        *time.Time `json:"date_of_issue" gorm:"date_of_issue"`
	PlaceOfIssue       string     `json:"place_of_issue" gorm:"place_of_issue"`
	ResidentialAddress string     `json:"residential_address" gorm:"residential_address"`
	RoleId             int64      `json:"role_id" gorm:"role_id"`
	Email              string     `json:"email" gorm:"column:email"`
	PhoneNumber        string     `json:"phone_number" gorm:"column:phone_number"`
	Salary             float64    `json:"salary" gorm:"column:salary"`
	DepartmentId       int64      `json:"department_id" gorm:"column:department_id"`
}

func (Employer) TableName() string {
	return "employees"
}

type Client struct {
	Id                 int64      `json:"id" gorm:"id"`
	Name               string     `json:"name" gorm:"first_name"`
	Surname            string     `json:"surname" gorm:"surname"`
	FatherName         string     `json:"father_name" gorm:"father_name"`
	DateOfBirth        *time.Time `json:"date_of_birth" gorm:"date_of_birth"`
	Gender             bool       `json:"gender" gorm:"gender"`
	Citizenship        string     `json:"citizenship" gorm:"citizenship"`
	PassportSeries     string     `json:"passport_series" gorm:"passport_series"`
	PassportNumber     int64      `json:"password_number" gorm:"password_number"`
	DateOfIssue        *time.Time `json:"date_of_issue" gorm:"date_of_issue"`
	PlaceOfIssue       string     `json:"place_of_issue" gorm:"place_of_issue"`
	ResidentialAddress string     `json:"residential_address" gorm:"residential_address"`
	PriorityId         int64      `json:"priority_id" gorm:"priority_id"`
	Phone              string     `json:"phone" gorm:"column:phone"`
	Email              string     `json:"email" gorm:"column:email"`
	LastPaymentId      int64      `json:"last_payment_id" gorm:"column:last_payment_id"`
	Login              string     `json:"login" gorm:"column:login"`
	Password           string     `json:"password" gorm:"column:password"`
	Salt               string     `json:"salt" gorm:"column:salt"`
	Verify             bool       `json:"verify" gorm:"column:verify"`
}

func (Client) TableName() string {
	return "clients"
}

type Rule struct {
	Id         int64   `json:"id" gorm:"column:id"`
	Commission float64 `json:"commission" gorm:"column:commission"`
	Term       int64   `json:"term" gorm:"column:term"`
	Active     bool    `json:"active" gorm:"column:active"`
	Min        float64 `json:"min" gorm:"column:min"`
	Max        float64 `json:"max" gorm:"column:max"`
}

func (Rule) TableName() string {
	return "rules"
}

type Type struct {
	Id          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
	Active      bool   `json:"active" gorm:"active"`
}

func (Type) TableName() string {
	return "types"
}

type Service struct {
	Id          int64      `json:"id" gorm:"id"`
	Name        string     `json:"name" gorm:"name"`
	Description string     `json:"description" gorm:"description"`
	Active      bool       `json:"active" gorm:"active"`
	TypeId      int64      `json:"type_id" gorm:"type_id"`
	RuleId      int64      `json:"rule_id" gorm:"rule_id"`
	ImageSource string     `json:"image_source" gorm:"column:image"`
	CreatedAt   *time.Time `json:"created_at" gorm:"created_at"`
	CreatedBy   int64      `json:"created_by" gorm:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"updated_at"`
	UpdatedBy   int64      `json:"updated_by" gorm:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"deleted_at"`
	DeletedBy   int64      `json:"deleted_by" gorm:"deleted_by"`
}

func (Service) TableName() string {
	return "services"
}

type AccountType struct {
	Id          int64  `json:"id" gorm:"id"`
	Name        string `json:"name" gorm:"name"`
	Description string `json:"description" gorm:"description"`
}

func (AccountType) TableName() string {
	return "account_types"
}

type Account struct {
	Id        int64      `json:"id" gorm:"column:id"`
	Number    int64      `json:"number" gorm:"column:number"`
	TypeId    int64      `json:"type_id" gorm:"column:type_id"`
	ClientId  int64      `json:"client_id" gorm:"column:client_id"`
	ServiceId int64      `json:"service_id" gorm:"column:service_id"`
	Amount    float64    `json:"amount" gorm:"column:amount"`
	Active    bool       `json:"active" gorm:"column:active"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"deleted_at"`
	ClosedAt  *time.Time `json:"closed_at" gorm:"column:closed_at"`
}

func (Account) TableName() string {
	return "accounts"
}

type Payment struct {
	Id          int64     `json:"id" gorm:"column:id"`
	ClientId    int64     `json:"client_id" gorm:"column:client_id"`
	Amount      float64   `json:"amount" gorm:"column:amount"`
	PaymentTime time.Time `json:"payment_time" gorm:"column:payment_time"`
	Status      string    `json:"status" gorm:"column:status"`
}

func (Payment) TableName() string {
	return "payments"
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type VerifyLogin struct {
	Login string `json:"email"`
	Otp   string `json:"otp"`
}
