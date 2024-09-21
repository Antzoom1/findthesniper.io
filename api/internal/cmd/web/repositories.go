package web

import (
	"github.com/RagOfJoes/findthesniper.io/internal/config"
	"github.com/RagOfJoes/findthesniper.io/mysql"
	"github.com/RagOfJoes/findthesniper.io/repositories"
	"github.com/sirupsen/logrus"
)

type WebRepositories struct {
	session repositories.Session
	user    repositories.User
}

func NewWebRepositories(cfg config.Configuration) (WebRepositories, error) {
	logrus.Info("[Web] Setting up repositories...")

	var repositories WebRepositories

	db, err := mysql.Connect(cfg)
	if err != nil {
		return repositories, err
	}

	repositories = WebRepositories{
		session: mysql.NewSession(db),
		user:    mysql.NewUser(db),
	}

	return repositories, nil
}

func (w *WebRepositories) Session() repositories.Session {
	return w.session
}

func (w *WebRepositories) User() repositories.User {
	return w.user
}
