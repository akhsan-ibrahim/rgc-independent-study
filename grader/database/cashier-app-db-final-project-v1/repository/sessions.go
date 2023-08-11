package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	err := u.db.Model(model.Session{}).Where("token = ?", tokenTarget).Delete(&model.Session{}).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	err := u.db.Table("sessions").Where("username = ?", session.Username).Updates(session).Error
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}
	return session, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	result := model.Session{}
	err := u.db.Table("sessions").Select("*").Where("username = ?", name).Find(&result).Error
	if err != nil {
		return model.Session{}, err
	}
	if result.ID == 0 {
		return model.Session{}, fmt.Errorf("not found")
	}
	return result, nil // TODO: replace this
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	result := model.Session{}
	err := u.db.Table("sessions").Select("*").Where("token = ?", token).Find(&result).Error
	if err != nil {
		return model.Session{}, err
	}
	if result.ID == 0 {
		return model.Session{}, fmt.Errorf("not found")
	}
	return result, nil // TODO: replace this
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
