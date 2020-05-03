package auth

import (
	"api/internal/barber"
	"api/jwt"
	"context"
	"log"
	"net/http"
	"strconv"
)

// How and why?
var barberCtxKey = &contextKey{"barber"}

type contextKey struct {
	name string
}

// Middleware returns a function of type http.Handler with a return of type
// http.Handler. A callback function? Is the WHOLE thing run whenever there
// is any request made?
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := ""
			if r != nil {
				tokenStr = r.Header.Get("auth")
			}

			if tokenStr == "" {
				next.ServeHTTP(w, r)
				return
			}

			//Validate jwt token.
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				log.Println(err)
				next.ServeHTTP(w, r)
				return
			}

			// It's not necessary to save a whole struct in the context...
			dbBarber := barber.Barber{UserName: username}
			id, err := barber.GetBarberIDByUsername(username)
			if err != nil {
				log.Println(err)
				next.ServeHTTP(w, r)
				return
			}
			dbBarber.BarberID = strconv.Itoa(id)

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
