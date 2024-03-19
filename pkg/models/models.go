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
	PasswordNumber     int64      `json:"password_number" gorm:"password_number"`
	DateOfIssue        *time.Time `json:"date_of_issue" gorm:"date_of_issue"`
	PlaceOfIssue       string     `json:"place_of_issue" gorm:"place_of_issue"`
	ResidentialAddress string     `json:"residential_address" gorm:"residential_address"`
	PriorityId         int64      `json:"priority_id" gorm:"priority_id"`
}

func (Client) TableName() string {
	return "clients"
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

type Services struct {
	Id          int64      `json:"id" gorm:"id"`
	Name        string     `json:"name" gorm:"name"`
	Description string     `json:"description" gorm:"description"`
	Active      bool       `json:"active" gorm:"active"`
	TypeId      int64      `json:"type_id" gorm:"type_id"`
	RuleId      int64      `json:"rule_id" gorm:"rule_id"`
	CreatedAt   *time.Time `json:"created_at" gorm:"created_at"`
	CreatedBy   int64      `json:"created_by" gorm:"created_by"`
	UpdatedAt   *time.Time `json:"updated_at" gorm:"updated_at"`
	UpdatedBy   int64      `json:"updated_by" gorm:"updated_by"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"deleted_at"`
	DeletedBy   int64      `json:"deleted_by" gorm:"deleted_by"`
}
