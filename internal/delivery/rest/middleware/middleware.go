package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/soerjadi/monsterdex/internal/delivery/rest"
	"github.com/soerjadi/monsterdex/internal/model"
	"github.com/soerjadi/monsterdex/internal/model/constant"
	"github.com/soerjadi/monsterdex/internal/usecase/access_token"
	"github.com/soerjadi/monsterdex/internal/usecase/user"
)

func CheckLoggedUser(tokenManagement access_token.Usecase, userManagement user.Usecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			var token string

			// it mean's that user not authorized
			if r.Header.Get("Authorization") == "" {
				next.ServeHTTP(w, r)
			}

			token, err := extractToken(w, r)
			if err != nil {
				return
			}

			at, err := tokenManagement.GetUserIDByToken(r.Context(), token)
			if err != nil {
				unauthorizedResp(w, r)

				return
			}

			_, err = userManagement.GetUserByID(r.Context(), at.UserID)
			if err != nil {
				unauthorizedResp(w, r)

				return
			}

			reqMap := make(map[string]interface{})
			reqMap["userID"] = at.UserID
			reqMap["req"] = r

			r = appendToContext(r.Context(), reqMap)

			next.ServeHTTP(w, r)

		})
	}
}

func OnlyLoggedInUser(tokenManagement access_token.Usecase, userManagement user.Usecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := extractToken(w, r)
			if err != nil {
				return
			}

			at, err := tokenManagement.GetUserIDByToken(r.Context(), token)
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			_, err = userManagement.GetUserByID(r.Context(), at.UserID)
			if err != nil {
				unauthorizedResp(w, r)

				return
			}

			reqMap := make(map[string]interface{})
			reqMap["userID"] = at.UserID
			reqMap["req"] = r

			r = appendToContext(r.Context(), reqMap)

			next.ServeHTTP(w, r)
		})
	}
}

func OnlyAdminLoggedIn(tokenManagement access_token.Usecase, userManagement user.Usecase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			token, err := extractToken(w, r)
			if err != nil {
				return
			}

			at, err := tokenManagement.GetUserIDByToken(r.Context(), token)
			if err != nil {
				unauthorizedResp(w, r)
				return
			}

			user, err := userManagement.GetUserByID(r.Context(), at.UserID)
			if err != nil {
				unauthorizedResp(w, r)

				return
			}

			if user.Role != model.ADMIN_ROLE {
				unauthorizedResp(w, r)

				return
			}

			reqMap := make(map[string]interface{})
			reqMap["userID"] = at.UserID
			reqMap["req"] = r

			r = appendToContext(r.Context(), reqMap)

			next.ServeHTTP(w, r)
		})
	}
}

func appendToContext(ctx context.Context, reqMap map[string]interface{}) *http.Request {
	userID := reqMap["userID"]
	r := reqMap["req"].(*http.Request)
	ctx = context.WithValue(ctx, constant.USER_ID_KEY, userID)

	r = r.WithContext(ctx)
	return r
}

func extractToken(w http.ResponseWriter, r *http.Request) (string, error) {
	authToken := strings.Split(r.Header.Get("Authorization"), " ")
	if len(authToken) != 2 {
		unauthorizedResp(w, r)

		return "", errors.New("invalid authorization")
	}

	return authToken[1], nil
}

func unauthorizedResp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := rest.Response{
		Message: "unauthorized access",
		Status:  "error",
	}

	w.WriteHeader(http.StatusUnauthorized)
	x, _ := json.Marshal(resp)
	w.Write(x)
	// return
}
