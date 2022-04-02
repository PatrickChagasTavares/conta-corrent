package main

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	emiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/patrickchagastavares/StoneTest/api"
	"github.com/patrickchagastavares/StoneTest/api/middleware"
	"github.com/patrickchagastavares/StoneTest/app"
	"github.com/patrickchagastavares/StoneTest/model"
	"github.com/patrickchagastavares/StoneTest/store"
	"github.com/patrickchagastavares/StoneTest/utils/logger"
	"github.com/patrickchagastavares/StoneTest/utils/session"
	"golang.org/x/sync/errgroup"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	godotenv.Load(".env")
	var (
		e         *echo.Echo
		errs      errgroup.Group
		sqlWriter *sqlx.DB
		sqlReader *sqlx.DB
	)

	e = echo.New()
	e.Debug = os.Getenv("ENV") != "prod"

	loggerConf := emiddleware.LoggerConfig{
		Format:           "${id} ${time_custom} ${remote_ip} ${method} ${status} ${uri} ${latency_human}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}
	e.Use(emiddleware.CORS())
	e.Use(emiddleware.RequestID())
	e.Use(emiddleware.LoggerWithConfig(loggerConf))
	e.Use(emiddleware.BodyLimit("2M"))
	e.Use(emiddleware.Recover())

	errs.Go(func() (err error) {
		sqlWriter, err = startSql(os.Getenv("DATABASE_WRITE_URL"))
		return
	})
	errs.Go(func() (err error) {
		sqlReader, err = startSql(os.Getenv("DATABASE_READ_URL"))
		return
	})

	defer func() {
		sqlWriter.Close()
		sqlReader.Close()
	}()

	if err := errs.Wait(); err != nil {
		logger.Fatal(err)
	}

	//
	runMigrations(os.Getenv("DATABASE_WRITE_URL"))

	// instanciando camada de repository/store
	stores := store.New(store.Options{
		Writer: sqlWriter,
		Reader: sqlReader,
	})

	// instanciando camada sessão
	session := session.NewSession(os.Getenv("SESSION_SECRET"))

	// instanciando camada de aplicação
	apps := app.New(app.Options{
		Stores:    stores,
		StartedAt: time.Now(),
		Session:   session,
	})

	api.Register(api.Options{
		Group: e.Group(""),
		Apps:  apps,
		Middleware: middleware.Register(middleware.Options{
			SessionJwt: session,
		}),
	})

	// funcão padrão pra tratamento de erros da camada http
	e.HTTPErrorHandler = createHTTPErrorHandler()

	if err := e.Start(":" + os.Getenv("PORT")); err != nil {
		logger.Error(err)
	}

	logger.Info("api: started")
}

func startSql(url string) (*sqlx.DB, error) {
	logger.Info("database: connecting")
	sql, err := sqlx.Connect("postgres", url)
	if err != nil {
		logger.Error(err)
		return sql, err
	}
	logger.Info("database: connected")

	return sql, nil
}

func createHTTPErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		if err := c.JSON(model.GetHTTPCode(err), model.Response{Err: err}); err != nil {
			logger.ErrorContext(c.Request().Context(), err)
		}
	}
}

func runMigrations(url string) {
	logger.Info("migration: running")
	if err := getMigration(url).Up(); err != nil && err != migrate.ErrNoChange {
		burnError(err)
	}
	logger.Info("migration: ended")
}

func getMigration(url string) *migrate.Migrate {
	dir, _ := os.Getwd()
	m, err := migrate.New(
		fmt.Sprintf("file://%s/db/migrations", dir),
		url,
	)
	if err != nil {
		burnError(err)
	}
	return m
}

func burnError(err error) {
	logger.Error(err)
	os.Exit(1)
}
