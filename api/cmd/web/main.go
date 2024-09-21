package main

import (
	"github.com/RagOfJoes/findthesniper.io/internal/cmd/web"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := web.Run(); err != nil {
		logrus.Fatal(err)
	}
}
