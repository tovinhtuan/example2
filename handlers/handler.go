package handlers

import (
	"context"
	"encoding/json"
	"ex2/repositories"
	"ex2/storages"
	"log"
	"net/http"
	// "github.com/go-chi/chi/v5"
)

type UserHandler struct {
	UserRepository repositories.UserRepository
}

type FailedRequest struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewUser(db *storages.PSQLManager) (*UserHandler, error) {
	return &UserHandler{
		UserRepository: repositories.NewUserRepository(db),
	}, nil
}

func(h *UserHandler) ResInformation (w http.ResponseWriter, r *http.Request ){
	ctx := r.Context()
	id := ctx.Value("UserID")
	user, err := h.UserRepository.ReadUserByUserId(id.(int64))
	if err != nil {
		failReq := FailedRequest{false, "User not exist"}
		fail, _ := json.Marshal(failReq)
		w.WriteHeader(http.StatusForbidden)
		w.Write(fail)
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	userJson, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error happened in JSON marshal. Err: %s", err)
		return
	}
	w.Write(userJson)
}

func (h *UserHandler )AuthenHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		md := r.Header.Get("Authorization")
		if md == "" {
			failReq := FailedRequest{false, "field authorization false"}
			fail, _:= json.Marshal(failReq)
			w.WriteHeader(http.StatusForbidden)
			w.Write(fail)
			return
		}
		token := md[7:]
		checkToken, _ := h.UserRepository.ReadTokenByToken(token)
		if checkToken == nil {
			failReq := FailedRequest{false, "authentication error"}
			fail, err := json.Marshal(failReq)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("Error happened in JSON marshal. Err: %s", err)
				return
			}
			// http.Error(w, "authentication error", http.StatusForbidden)
			w.WriteHeader(http.StatusForbidden)
			w.Write(fail)
			return
		}
		ctx :=context.WithValue(r.Context(),"UserID",checkToken.UserId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}