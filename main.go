package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"os"
	"tcfback/internal/db"
	"tcfback/internal/handlers"
	"tcfback/internal/repositories"
)

type Server struct {
	Queries *db.Queries
	Conn    *pgx.Conn
}

func ConnectDB() (*Server, error) {
	// Load environment variables from .env file if needed
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using system environment variables.")
	}

	ctx := context.Background()

	dbURL := os.Getenv("GOOSE_DBSTRING")
	if dbURL == "" {
		return nil, fmt.Errorf("database connection string is not set in environment variables")
	}

	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse database URL: %w", err)
	}

	conn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	queries := db.New(conn)

	server := &Server{
		Queries: queries,
		Conn:    conn,
	}

	return server, nil
}

func main() {
	ctx := context.Background()

	//ConnectDB
	dbServer, err := ConnectDB()

	if err != nil {
		log.Fatalf("error connected to db %v", err)
	}

	defer func(Conn *pgx.Conn, ctx context.Context) {
		err := Conn.Close(ctx)
		if err != nil {
			fmt.Println("can't close db")
		}
	}(dbServer.Conn, ctx)

	e := echo.New()

	api := e.Group("/api")

	userRepo := repositories.NewUserRepository(dbServer.Queries)
	userHandler := handlers.NewUserHandler(&userRepo)
	userHandler.Router(api)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3001"))
}
