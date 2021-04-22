package store

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jannst/go_start/auth_service/src"
	"regexp"
)

const sessionStoreFormat = "sessions:%s"

type SessionService interface {
	FetchSession(token string) (*src.Session, error)
	PersistSession(session src.Session) error
	DeleteSession(token string) error
}

type sessionService struct {
	Pool *redis.Pool
}

func NewSessionService(pool *redis.Pool) SessionService {
	return &sessionService{Pool: pool}
}

func (s sessionService) FetchSession(token string) (*src.Session, error) {
	err := ValidateTokenFormat(token)
	if err == nil {
		conn := s.Pool.Get()
		defer conn.Close()
		reply, err := redis.Bytes(conn.Do("GET", fmt.Sprintf(sessionStoreFormat, token)))
		if err != nil {
			return nil, err
		} else {
			var session src.Session
			err := json.NewDecoder(bytes.NewReader(reply)).Decode(&session)
			if err != nil {
				return nil, err
			} else {
				return &session, nil
			}
		}
	} else {
		return nil, err
	}
}

func (s sessionService) PersistSession(session src.Session) error {
	token := session.SessionToken
	err := ValidateTokenFormat(token)
	if err == nil {
		buf := &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(session)
		if err != nil {
			return err
		} else {
			conn := s.Pool.Get()
			defer conn.Close()
			_, err := conn.Do("SET", fmt.Sprintf(sessionStoreFormat, token), buf.Bytes())
			return err
		}
	} else {
		return err
	}
}

func (s sessionService) DeleteSession(token string) error {
	conn := s.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", fmt.Sprintf(sessionStoreFormat, token))
	return err
}

func ValidateTokenFormat(token string) error {
	match, err := regexp.MatchString("^[a-f0-9]{32}$", token)
	if err != nil {
		return err
	} else if !match {
		return errors.New("tokes does not match required format")
	} else {
		return nil
	}
}

