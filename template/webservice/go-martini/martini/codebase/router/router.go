package router

import (
	"{{ .appname }}/controllers"
	"github.com/go-martini/martini"
	
)

// RegisterRoutes creates router and routes requests
func RegisterRoutes(m *martini.ClassicMartini) {
	
	m.Group("/v1", func(r martini.Router) {
		user := new(controllers.UserController)
		r.Post("/user",  user.Create)
		r.Get("/user/:id", user.Get)
		r.Get("/user", user.Find)
		r.Put("/user/:id", user.Update)
		r.Delete("/user/:id", user.Delete)
	})

	m.Run()
}
