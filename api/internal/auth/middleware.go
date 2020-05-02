package auth

import (
	"api/internal/barber"
	"api/jwt"
	"context"
	"net/http"
	"strconv"
)

// How and why?
var barberCtxKey = &contextKey{"barber"}

type contextKey struct {
	name string
}

// Middleware returns a function of type http.Handler with a return of type
// http.Handler. A callback function?
func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Is this 'finding' a cookie with key words token?
			c, err := r.Cookie("token")

			// Allow unauthenticated users some initial access to our api.
			// Has no authentication context
			if err != nil || c == nil {
				// No modification to context. Un authenticated?
				next.ServeHTTP(w, r)
				return
			}

			//validate jwt token. What IS tokenStr?
			tokenStr := c.Value
			username, err := jwt.ParseToken(tokenStr)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			// GetBarberIDByUsername is a method of the package barber. It's diff.
			// from SaveOne() and Get() because those are directly connected
			// to the structure being passed as parameters.
			dbBarber := barber.Barber{UserName: username}
			id, err := barber.GetBarberIDByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			// Why barberID and not username?
			dbBarber.BarberID = strconv.Itoa(id)

			// Put it in context with barber information?
			ctx := context.WithValue(r.Context(), barberCtxKey, dbBarber)

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
