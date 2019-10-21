package models

import (
	"fmt"

	u "github.com/go-pg-mux-apis/medium/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserId uint   `json:"user-id"`
}

func (c *Contact) Validate() (map[string]interface{}, bool) {
	if c.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if c.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if c.UserId <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	return u.Message(true, "Success, Valid contact"), true
}

func (c *Contact) Create() map[string]interface{} {
	if resp, ok := c.Validate(); !ok {
		return resp
	}

	GetDB().Create(c)

	resp := u.Message(true, "Contact created")
	resp["contact"] = contact
	return resp
}

func GetContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).Find(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
