package main

import (
	"fmt"
	"net/http"

	"github.com/manoj-H-C/rssagg/internal/auth"
	"github.com/manoj-H-C/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %s", err))
			return
		}
		user, err := apicfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("couldn't get user: %s", err))
			return
		}
		handler(w, r, user)
	}
}
