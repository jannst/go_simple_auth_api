// This file is safe to edit. Once it exists it will not be overwritten

package api

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"github.com/go-openapi/swag"
	"github.com/jannst/go_start/auth_service/src/apimodel"
	"github.com/jannst/go_start/auth_service/src/middlewares"
	"github.com/jannst/go_start/auth_service/src/store"
	"github.com/jannst/go_start/auth_service/src/util"
	"github.com/rs/zerolog/log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/jannst/go_start/auth_service/src"
	"github.com/jannst/go_start/auth_service/src/api/operations"
	"github.com/jannst/go_start/auth_service/src/api/operations/user"
)

//g o:ge nerate swagger generate server --target ../../src --name Auth --spec ../../spec/auth_service.yml --model-package apimodel --server-package api --principal github.com/jannst/go_start/auth_service/src.Session

func configureFlags(api *operations.AuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AuthAPI) http.Handler {
	api.ServeError = errors.ServeError

	cfg := LoadConfig()

	redisConnPool := store.NewPool(cfg.RedisHost, cfg.RedisPort, cfg.RedisUser, cfg.RedisPassword)
	sessionService := store.NewSessionService(redisConnPool)

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-TOKEN" header is set
	api.APIKeyAuthAuth = func(token string) (*src.Session, error) {
		fmt.Printf("token: %s\n", token)
		session, err := sessionService.FetchSession(token)
		if err != nil {
			log.Error().Err(err).Msg("could not fetch token")
			return nil, errors.New(http.StatusUnauthorized, "invalid token")
		} else {
			return session, nil
		}
	}

	api.UserInfoHandler = user.InfoHandlerFunc(func(params user.InfoParams, principal *src.Session) middleware.Responder {
		userModel := apimodel.User{
			BaseUser:   apimodel.BaseUser{
				Email:    util.CreateMail(principal.UserEmail),
				Username: swag.String(principal.UserName),
			},
			ID:         swag.Uint32(principal.UserId),
			UserRole: swag.String(principal.UserRole),
		}
		return user.NewInfoOK().WithPayload(&userModel)
	})

	api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
		token, _ := createNewSessionToken()
		if *params.Body.Email == "admin@lol.de" && *params.Body.Password == "12345678" {
			result := &apimodel.AccessTokenResponse{
				Token: swag.String(token),
				User:  &apimodel.User{
					BaseUser: apimodel.BaseUser{
						Email:    util.CreateMail("test@lol.de"),
						Username: swag.String("JanST"),
					},
					ID:       swag.Uint32(100),
					UserRole: swag.String("ADMIN"),
				},
			}
			session := src.Session{
				SessionToken: *result.Token,
				UserId:       *result.User.ID,
				UserRole:     *result.User.UserRole,
				UserName:     *result.User.Username,
				UserEmail:    string(*result.User.Email),
			}
			sessionService.PersistSession(session)
			return user.NewLoginOK().WithPayload(result)
		}
		if *params.Body.Email == "user@lol.de" && *params.Body.Password == "qwertzui" {
			result := &apimodel.AccessTokenResponse{
				Token: swag.String(token),
				User:  &apimodel.User{
					BaseUser: apimodel.BaseUser{
						Email:    util.CreateMail("user@lol.de"),
						Username: swag.String("SomeUser"),
					},
					ID:       swag.Uint32(200),
					UserRole: swag.String("USER"),
				},
			}
			session := src.Session{
				SessionToken: *result.Token,
				UserId:       *result.User.ID,
				UserRole:     *result.User.UserRole,
				UserName:     *result.User.Username,
				UserEmail:    string(*result.User.Email),
			}
			sessionService.PersistSession(session)
			return user.NewLoginOK().WithPayload(result)
		}

		return nil
		//return middleware.NotImplemented("operation user.Login has not yet been implemented")
	})

	api.UserLogoutHandler = user.LogoutHandlerFunc(func(params user.LogoutParams, principal *src.Session) middleware.Responder {
		err := sessionService.DeleteSession(principal.SessionToken)
		if err != nil {
			return nil
		}
		return nil
	})

	if api.UserRegisterHandler == nil {
		api.UserRegisterHandler = user.RegisterHandlerFunc(func(params user.RegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation user.Register has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return middlewares.ZerologMiddleware(middlewares.SimpleCors(handler))
}

func createNewSessionToken() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(buf), nil
}
