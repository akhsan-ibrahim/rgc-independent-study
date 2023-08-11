package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	sesion, err := u.TokenExist(tokenTarget)
	if err != nil {
		return err
	}

	var ind int
	for i, ses := range listSessions {
		if ses == sesion {
			ind = i
		}
	}

	lsSessions := append(listSessions[:ind], listSessions[ind+1:]...)
	// Select target token and delete from listSessions
	// TODO: answer here

	jsonData, err := json.Marshal(lsSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", jsonData)
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	lsSesion, err := u.ReadSessions()
	if err != nil {
		return err
	}

	lsSesion = append(lsSesion, session)
	jsonContent, err := json.Marshal(lsSesion)
	if err != nil {
		return err
	}

	u.db.Save("sessions", jsonContent)

	return nil // TODO: replace this
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	sesion, err := u.TokenExist(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(sesion) {
		u.DeleteSessions(token)
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return sesion, nil // TODO: replace this
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
