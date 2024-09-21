package web

import (
	"github.com/RagOfJoes/findthesniper.io/internal/config"
	"github.com/RagOfJoes/findthesniper.io/services"
	"github.com/sirupsen/logrus"
)

type WebServices struct {
	oauth   services.OAuth2Config
	session services.Session
	user    services.User
}

func NewWebServices(cfg config.Configuration, repositories WebRepositories) (WebServices, error) {
	logrus.Infoln("")
	logrus.Info("[Web] Setting up services...")

	return WebServices{
		oauth: services.NewOAuth2Config(services.OAuth2ConfigDependencies{
			Config: cfg,
		}),
		session: services.NewSession(services.SessionDependencies{
			Repository: repositories.Session(),
		}),
		user: services.NewUser(services.UserDependencies{
			Repository: repositories.User(),
		}),
	}, nil
}

func (w WebServices) OAuth2Config() services.OAuth2Config {
	return w.oauth
}

func (w WebServices) Session() services.Session {
	return w.session
}

func (w WebServices) User() services.User {
	return w.user
}
