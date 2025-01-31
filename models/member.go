package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	SlackID          string
	Username         string
	Fullname         string
	Enrollment       string
	DOB              string
	Contact          string
	Email            string
	Branch           string
	YearOfGraduation string
	GSuiteID         string
	Address          string
	IsAlum           bool
	Image            string
}
