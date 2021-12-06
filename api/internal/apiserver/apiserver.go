package apiserver

import (
	"os"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	httpPort string
	logLevel string
	logger *logrus.Logger
}

func New() *APIserver {
	httpPort := os.Getenv("PORT")
	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
	    lvl = "debug"
	}
	return &APIserver{
		httpPort: httpPort,
		logLevel: lvl,
		logger: logrus.New(),
	}
}

func (s *APIserver) Start() error {
    if err := s.configureLogLevel(); err != nil {
        return err
    }
    s.logger.Info("Starting api server")
	return nil
}

func (s *APIserver) configureLogLevel () error {
    level, err := logrus.ParseLevel(s.logLevel)
    if err != nil {
        return err
    }
    s.logger.SetLevel(level)
    return nil
}
