package store

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/go-openapi/strfmt"
	"haw-hamburg.de/cloudWP/src"
	"haw-hamburg.de/cloudWP/src/api/operations/user"
	"haw-hamburg.de/cloudWP/src/apimodel"
	"sync"
)

type AuthService interface {
	FetchSession(token string) (*src.Session, error)
	Login(loginParams user.LoginBody) (*apimodel.AccessTokenResponse, error)
	Logout(session *src.Session) error
	CreateUser(user apimodel.NewUser) error
	GetUserById(userId uint32) (*apimodel.User, error)
}

type internalUser struct {
	passwordHash string
	username     string
	id           uint32
	email        strfmt.Email
}

type authServiceImpl struct {
	users      map[uint32]internalUser
	sessions   map[string]internalUser
	nextUserId uint32
	mu         sync.Mutex
}

func NewSessionService() AuthService {
	return &authServiceImpl{
		users: map[uint32]internalUser{},
		sessions: map[string]internalUser{},
	}
}

func (s *authServiceImpl) GetUserById(userId uint32) (*apimodel.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if userById, ok := s.users[userId]; ok {
		return &apimodel.User{
			BaseUser: apimodel.BaseUser{
				Email:    &userById.email,
				Username: &userById.username,
			},
			ID: &userById.id,
		}, nil
	} else {
		return nil, errors.New("could not find userById for userId")
	}
}

func (s *authServiceImpl) CreateUser(newUser apimodel.NewUser) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	//check for duplicate email
	for _, user := range s.users {
		if user.email == *newUser.Email {
			return errors.New("email already in use")
		}
	}
	//hash passwordHash
	password, err := HashPassword(*newUser.Password)
	if err != nil {
		return err
	}

	//add user to "database"
	s.nextUserId++
	user := internalUser{
		passwordHash: password,
		username:     *newUser.Username,
		id:           s.nextUserId,
		email:        *newUser.Email,
	}
	s.users[s.nextUserId] = user
	return nil
}

func (s *authServiceImpl) FetchSession(token string) (*src.Session, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if user, ok := s.sessions[token]; ok {
		return &src.Session{
			SessionToken: token,
			UserId:       user.id,
		}, nil
	} else {
		return nil, errors.New("session not found")
	}
}

func (s *authServiceImpl) Login(loginParams user.LoginBody) (*apimodel.AccessTokenResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.users {
		if user.email == *loginParams.Email {
			if CheckPasswordHash(*loginParams.Password, user.passwordHash) {
				sessionToken, err := createNewSessionToken()
				if err != nil {
					return nil, err
				}
				s.sessions[sessionToken] = user
				return &apimodel.AccessTokenResponse{
					Token: &sessionToken,
					User: &apimodel.User{
						BaseUser: apimodel.BaseUser{
							Email:    &user.email,
							Username: &user.username,
						},
						ID: &user.id,
					},
				}, nil
			}
		}
	}
	return nil, errors.New("email not found or invalid password")
}

func (s *authServiceImpl) Logout(session *src.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.sessions[session.SessionToken]; ok {
		delete(s.sessions, session.SessionToken)
		return nil
	} else {
		return errors.New("session not found")
	}
}

func createNewSessionToken() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
