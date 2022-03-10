package main

import (
	"context"
	"flag"
	"fmt"

	"demo-go-kit/account/controller"
	db_config "demo-go-kit/account/database"
	"demo-go-kit/account/repository"
	"demo-go-kit/account/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"
)

const dbsource = "postgresql://postgres:postgres@localhost:5432/gokitexample?sslmode=disable"

func main() {

	var httpAddr = flag.String("http", ":8080", "http listen address")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	ctx := context.Background()
	db := db_config.GetDatabase(ctx, "gokitexample")

	flag.Parse()

	var srv controller.Service
	{
		repository := repository.NewRepo(db, logger)

		srv = controller.NewService(repository, logger)
	}

	errs := make(chan error)

	// gracefull shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := server.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := server.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
