package model

import (
	"github.com/TobyIcetea/fastgo/internal/pkg/rid"
	"github.com/onexstack/onexstack/pkg/authn"
	"gorm.io/gorm"
)

// AfterCreate 在创建数据库记录之后生成 postID
func (m *Post) AfterCreate(tx *gorm.DB) error {
	m.PostID = rid.PostID.New(uint64(m.ID))

	return tx.Save(m).Error
}

// BeforeCreate 在创建数据库记录之前加密明文密码
func (m *User) BeforeCreate(tx *gorm.DB) error {
	// Encrypt the user password
	var err error
	m.Password, err = authn.Encrypt(m.Password)
	if err != nil {
		return err
	}

	return nil
}

// AfterCreate 在创建数据库记录之后生成 userID
func (m *User) AfterCreate(tx *gorm.DB) error {
	m.UserID = rid.UserID.New(uint64(m.ID))

	return tx.Save(m).Error
}
