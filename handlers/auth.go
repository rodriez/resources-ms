package handlers

import (
	"log"
	"net/http"

	"github.com/rodriez/restface"
)

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(res http.ResponseWriter, req *http.Request) {
			log.Println("BasicAuthMiddleware")

			authenticator := restface.BasicAuthenticator{
				Request: req,
				ValidUser: func(name, pass string) bool {
					return name == "john" && pass == "johnS.67"
				},
			}

			if err := authenticator.Authenticate(); err != nil {
				presenter := restface.Presenter{Writer: res}
				presenter.PresentError(restface.Forbidden())

				return
			}

			next.ServeHTTP(res, req)
		})
}
