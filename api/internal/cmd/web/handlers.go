package web

import (
	"strings"

	"github.com/RagOfJoes/findthesniper.io/handlers"
	"github.com/RagOfJoes/findthesniper.io/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func SetupHandlers(config config.Configuration, services WebServices) *chi.Mux {
	logrus.Infoln("")
	logrus.Info("[Web] Setting up handlers...")

	// Setup Handlers
	router := handlers.New(config)

	session := handlers.Session(handlers.SessionDependencies{
		Config: config,

		Service: services.Session(),
	})

	handlers.Auth(handlers.AuthDependencies{
		Config: config,

		OAuth2: services.OAuth2Config(),
		User:   services.User(),

		Session: session,
	}, router)
	handlers.User(handlers.UserDependencies{
		Service: services.User(),

		Session: session,
	}, router)

	var routes []string
	for _, r := range router.Routes() {
		routes = append(routes, r.Pattern)
	}
	logrus.Infof("Attached %s to HTTP Server", strings.Join(routes, ", "))

	return router
}

func RunHandlers(cfg config.Configuration, router *chi.Mux) error {
	return handlers.Run(cfg, router)
}
