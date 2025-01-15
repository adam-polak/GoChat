package main

import (
	controller "GoChat/controllers"
	"GoChat/lib"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}

			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   "./public",
		Browse: true,
	}))

	// Create db connection
	connStr, err := lib.GetSecret("ConnectionString")
	if err != nil {
		e.Logger.Fatal(err)
	}

	dbPool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer dbPool.Close()

	c := &controller.Controller{DB: dbPool}

	// Routes
	e.GET("/ws", hello)
	e.POST("/login", c.Login)
	e.POST("/signup", c.SignUp)

	// Restart
	e.Logger.Fatal(e.Start(":1323"))
}
