package handlers

import (
	"net/http"
	"resources-ms/repositories"

	"github.com/rodriez/restface"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	authRepo := &repositories.AuthRepository{}

	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			authenticator := &restface.BasicAuthenticator{
				Request:   req,
				ValidUser: authRepo.ValidateUser,
			}

			if err := authenticator.Authenticate(); err != nil {
				presenter := restface.Presenter{Writer: res}
				presenter.PresentError(restface.Forbidden())

				return
			}

			next.ServeHTTP(res, req)
		})
}
