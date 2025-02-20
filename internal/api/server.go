package api

import (
	"financial/internal/api/router"
	"financial/internal/db"
	"net/http"

	"github.com/rs/cors"
	"gorm.io/gorm"
)

type Config struct {
	ServerPort string
	DBConfig   *map[string]string
}

type Server struct {
	config *Config
	router *routes.Router
	dbCon  *gorm.DB
}

func NewServer(config *Config) *Server {
	dbCon := db.InitDatabase(db.SqliteDriver, *config.DBConfig)

	return &Server{
		config: config,
		router: routes.NewRouter(dbCon, "tmp-jwt-key"),
		dbCon:  dbCon,
	}
}

func (s *Server) Run() error {
	// Create a CORS middleware instance
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposedHeaders:   []string{"X-Total-Count", "Location"},
		AllowCredentials: true,
		MaxAge:           300,
		Debug:            true,
	})

	srv := &http.Server{
		Addr:    ":" + s.config.ServerPort,
		Handler: corsHandler.Handler(s.router),
	}

	println("Server running on port: ", s.config.ServerPort)
	return srv.ListenAndServe()
}
