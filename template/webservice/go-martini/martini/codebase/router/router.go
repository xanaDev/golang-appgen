package router

import (
	"codebase/controllers"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	//"net/http"
	
)

// RegisterRoutes creates router and routes requests
func RegisterRoutes(m *martini.ClassicMartini) {
	var data forms.CreateUserCommand

	m.Group("/v1", func(v1 martini.Router) {
		user := new(controllers.UserController)
		v1.Post("/user", binding.Json( data ), user.Create)
		v1.Get("/user/:id", user.Get)
		v1.Get("/user", user.Find)
		v1.Put("/user/:id", user.Update)
		v1.Delete("/user/:id", user.Delete)
	})


	m.Run()
}
