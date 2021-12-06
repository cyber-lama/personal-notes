package apiserver

import (
    "net/http"
	"os"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

type APIserver struct {
	httpPort string
	logLevel string
	logger *logrus.Logger
	router *mux.Router
}

func New() *APIserver {
	httpPort := os.Getenv("PORT")
	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
	    lvl = "debug"
	}
	return &APIserver{
		httpPort: ":" + httpPort,
		logLevel: lvl,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
    if err := s.configureLogLevel(); err != nil {
        return err
    }

    s.configureRouter()

    s.logger.Info("Server up and listen port" + s.httpPort)

    return http.ListenAndServe(s.httpPort, s.router)
}

func (s *APIserver) configureLogLevel () error {
    level, err := logrus.ParseLevel(s.logLevel)
    if err != nil {
        return err
    }
    s.logger.SetLevel(level)
    return nil
}

func (s *APIserver) configureRouter() {
    s.router.HandleFunc("/test", s.handleHello())
}

func (s *APIserver) handleHello() http.HandlerFunc{
    return func(w http.ResponseWriter, r *http.Request){

    }
}