package auth

import (
	"api/internal/barber"
	"api/jwt"
	"context"
	"log"
	"net/http"
)

// How and why?
var barberCtxKey = &contextKey{"barber"}

type contextKey struct {
	name string
}

// Middleware closure.
func Middleware() func(http.Handler) http.Handler {
	// Inner anonymous function. next is our i:=0?
	return func(next http.Handler) http.Handler {
		// Inner return. w and r are to be used by next it seems.
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := ""
			if r != nil {
				tokenStr = r.Header.Get("auth-token")
			}

			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			// Authenticate user via username.
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				log.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			// It's not necessary to save a whole struct in the context...
			dbBarber := barber.Barber{UserName: username}

			// Place barber object in context.
			// ForContext returns a barber pointer, so need to pass addr of barber.
			ctx := context.WithValue(r.Context(), barberCtxKey, &dbBarber)

			// and call the next with our new context. Context has to do with
			// requesting something from api with the proper auth verified.
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the barber from the context. REQUIRES Middleware to have run.
// Returns a barber struct.
func ForContext(ctx context.Context) *barber.Barber {
	// Casting result as barber struct?
	raw, _ := ctx.Value(barberCtxKey).(*barber.Barber)
	return raw
}
