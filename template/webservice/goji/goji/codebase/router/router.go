package router

import (
//	"codebase/controllers"

	"goji.io"
	"net/http"
	
)

// RegisterRoutes creates router and routes requests
func RegisterRoutes(mux *goji.Mux) {
	
	//mux.HandleFunc(pat.Get("/books"), allBooks)
	//mux.HandleFunc(pat.Get("/books/:isbn"), bookByISBN)
	//mux.Use(logging)

/*	m.Group("/v1", func(v1 martini.Router) {
		user := new(controllers.UserController)
		v1.Post("/user", user.Create)
		v1.Get("/user/:id", user.Get)
		v1.Get("/user", user.Find)
		v1.Put("/user/:id", user.Update)
		v1.Delete("/user/:id", user.Delete)
	})*/
	http.ListenAndServe("localhost:8081", mux)

}
