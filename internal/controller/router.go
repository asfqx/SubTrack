package controller

import "github.com/go-chi/chi/v5"

type SubscriptionsHandler interface{}

func RegisterRoutes(mux *chi.Mux, subsHandler SubscriptionsHandler) {
	mux.HandleFunc()
}
