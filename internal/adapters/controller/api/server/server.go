package server

import (
	"AvitoTechTask/internal/adapters/app"
	"AvitoTechTask/internal/adapters/controller/api/v1/pr"
	"AvitoTechTask/internal/adapters/controller/api/v1/team"
	"AvitoTechTask/internal/adapters/controller/api/v1/user"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Setup(app *app.App) {
	app.Server.Logger.SetOutput(io.Discard)
	app.Server.HideBanner = true
	app.Server.Debug = false

	//app.Server.Use(middleware.Recover())
	app.Server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://localhost:8080"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"},
		// Важно для WebSocket!
		AllowCredentials: true,
	}))

	app.Server.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:      true,
		LogStatus:   true,
		LogMethod:   true,
		HandleError: true,
		LogError:    true,
		LogRemoteIP: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				app.ServiceProvider.Logger().Infow("request completed",
					"ip", v.RemoteIP,
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
				)
			} else {
				app.ServiceProvider.Logger().Errorw("request failed",
					"ip", v.RemoteIP,
					"method", v.Method,
					"uri", v.URI,
					"status", v.Status,
					"error", v.Error.Error(),
				)
			}
			return nil
		},
	}))

	route(app)
}

func route(app *app.App) {
	server := app.Server
	serviceProvider := app.ServiceProvider

	apiV1 := server.Group("/api/v1")

	userHandler := user.NewHandler(serviceProvider.UserService(), serviceProvider.Validator(), serviceProvider.Decoder())
	userHandler.Setup(apiV1)

	teamHandler := team.NewHandler(serviceProvider.TeamService(), serviceProvider.Validator(), serviceProvider.Decoder())
	teamHandler.Setup(apiV1)

	prHandler := pr.NewHandler(serviceProvider.PRService(), serviceProvider.Validator())
	prHandler.Setup(apiV1)

}
