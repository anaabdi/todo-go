package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anaabdi/todo-go/middleware"

	"github.com/anaabdi/todo-go/config"
	"github.com/anaabdi/todo-go/handler"
	"github.com/anaabdi/todo-go/repository"
	"github.com/go-chi/chi"
)

func main() {
	repository.InitDB()

	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1Route chi.Router) {
			v1Route.Route("/users", func(userRoute chi.Router) {
				userRoute.Post("/", handler.Register)
				userRoute.With(middleware.AuthMiddleware).Get("/me", handler.Me)
			})
			v1Route.Route("/auth", func(authRoute chi.Router) {
				authRoute.Post("/", handler.Login)
			})
		})
	})

	port := config.GetInt("PORT", 1000)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))

}
