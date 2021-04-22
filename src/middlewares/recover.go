package middlewares

import (
	"fmt"
	"github.com/go-chi/render"
	"net/http"
	"runtime/debug"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				stack := debug.Stack()
				f := "PANIC: %s\n%s"
				fmt.Printf(f, err, stack)

				render.JSON(w,r, struct {
					message string
					code int
				}{
					message: "internal server error",
					code:    http.StatusInternalServerError,
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}