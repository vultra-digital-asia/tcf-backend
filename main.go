package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"tcfback/internal/db"
	"tcfback/internal/handlers"
	"tcfback/internal/repositories"
)

type Server struct {
	Queries *db.Queries
	Conn    *pgx.Conn
}

func ConnectDB() (*Server, error) {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=aws-0-ap-southeast-1.pooler.supabase.com user=postgres.fpeuaykmjlszvofokbcn dbname=postgres password=n4kb03ank sslmode=require")
	if err != nil {
		return nil, err
	}

	queries := db.New(conn)

	server := &Server{
		Queries: queries,
		Conn:    conn,
	}

	return server, nil

	//defer conn.Close(ctx)
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
