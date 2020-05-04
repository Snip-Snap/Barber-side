package directive

import (
	"api/generated"
	"api/internal/auth"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
)

// VerifyAuth implements checkAuth directive.
func VerifyAuth(c *generated.Config) {
	// c needs to be * because you're changing(assigning) to it!
	c.Directives.CheckAuth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		barber := auth.ForContext(ctx)

		if barber != nil {
			if barber.UserName != "" {
				// Let barber proceed with api calls.
				return next(ctx)
			}
		}
		return nil, errors.New("Unauthorised")
	}
}
