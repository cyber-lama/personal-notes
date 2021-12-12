package apiserver

import (
	"github.com/cyber-lama/personal-notes/api/internal/store"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type APIserver struct {
	confHttpPort string
	confLogLevel string
	confStore    *store.DBConfig
	logger       *logrus.Logger
	router       *mux.Router
	store        *store.Store
}

func New() *APIserver {
	//the port we will listen to
	confHttpPort := os.Getenv("API_PORT")
	// debug level
	lvl := os.Getenv("LOG_LEVEL")
	if lvl == "" {
		lvl = "debug"
	}
	// db conf obj
	cnfStr := &store.DBConfig{}
	cnfStr.Host = os.Getenv("DB_HOST")
	cnfStr.Port = os.Getenv("DB_PORT")
	cnfStr.User = os.Getenv("DB_USER")
	cnfStr.Password = os.Getenv("DB_PASSWORD")
	cnfStr.DBname = os.Getenv("DB_NAME")

	return &APIserver{
		confHttpPort: ":" + confHttpPort,
		confLogLevel: lvl,
		confStore:    cnfStr,
		logger:       logrus.New(),
		router:       mux.NewRouter(),
	}
}

func (s *APIserver) Start() error {
	if err := s.configureLogLevel(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Server up and listen port" + s.confHttpPort)
	return http.ListenAndServe(s.confHttpPort, s.router)
}

func (s *APIserver) configureLogLevel() error {
	level, err := logrus.ParseLevel(s.confLogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *APIserver) configureRouter() {
	s.router.HandleFunc("/test", s.handleHello())
}

func (s *APIserver) configureStore() error {

	st := store.New(s.confStore)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st

	s.logger.Info("DB connected and ready!")

	return nil
}

func (s *APIserver) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
