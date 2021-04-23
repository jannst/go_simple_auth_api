// This file is safe to edit. Once it exists it will not be overwritten

package api

import (
	"crypto/tls"
	"fmt"
	"github.com/rs/zerolog/log"
	"haw-hamburg.de/cloudWP/src/middlewares"
	"haw-hamburg.de/cloudWP/src/store"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"haw-hamburg.de/cloudWP/src"
	"haw-hamburg.de/cloudWP/src/api/operations"
	"haw-hamburg.de/cloudWP/src/api/operations/user"
)

//go:generate swagger generate server --target ../../src --name Auth --spec ../../spec/auth_service.yml --model-package apimodel --server-package api --principal github.com/jannst/go_start/auth_service/src.Session

func configureFlags(api *operations.AuthAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AuthAPI) http.Handler {
	api.ServeError = errors.ServeError

	//create our session service
	sessionService := store.NewSessionService()

	api.UseSwaggerUI()
	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "X-API-TOKEN" header is set
	api.APIKeyAuthAuth = func(token string) (*src.Session, error) {
		session, err := sessionService.FetchSession(token)
		if err != nil {
			log.Error().Err(err).Msg("could not fetch token")
			return nil, errors.New(http.StatusUnauthorized, "invalid token")
		} else {
			return session, nil
		}
	}

	api.UserInfoHandler = user.InfoHandlerFunc(func(params user.InfoParams, principal *src.Session) middleware.Responder {
		fmt.Printf("userId: %d", principal.UserId)
		userById, err := sessionService.GetUserById(principal.UserId)
		if err != nil {
			return user.NewInfoOK().WithPayload(userById)
		} else {
			fmt.Println(err)
			return user.NewInfoInternalServerError()
		}
	})

	api.UserLoginHandler = user.LoginHandlerFunc(func(params user.LoginParams) middleware.Responder {
		session, err := sessionService.Login(params.Body)
		if err != nil {
			return user.NewLoginBadRequest()
		} else {
			fmt.Println(err)
			return user.NewLoginOK().WithPayload(session)
		}
	})

	api.UserLogoutHandler = user.LogoutHandlerFunc(func(params user.LogoutParams, principal *src.Session) middleware.Responder {
		err := sessionService.Logout(principal)
		if err != nil {
			return user.NewLogoutUserOK()
		} else {
			fmt.Println(err)
			return user.NewLogoutInternalServerError()
		}
	})

	api.UserRegisterHandler = user.RegisterHandlerFunc(func(params user.RegisterParams) middleware.Responder {
		err := sessionService.CreateUser(*params.Body)
		if err != nil {
			return user.NewRegisterNoContent()
		} else {
			fmt.Println(err)
			return user.NewRegisterInternalServerError()
		}
	})

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
