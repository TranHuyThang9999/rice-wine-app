package main

import (
	"context"
	"errors"
	"flag"
	"net"
	"rice-wine-shop/api/routers"
	log "rice-wine-shop/common/log"
	"rice-wine-shop/core/adapters/interfaces"
	"rice-wine-shop/core/configs"
	"rice-wine-shop/core/generator"
	"rice-wine-shop/fxloader"

	"net/http"
	"os"
	"os/signal"

	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func init() {
	log.LoadLogger()
	var pathConfig string
	flag.StringVar(&pathConfig, "configs", "common/configs/configs.json", "path config")
	flag.Parse()
	configs.LoadConfig(pathConfig)

}

func main() {
	app := fx.New(
		fx.Provide(configs.Get),
		fx.Options(fxloader.Load()...),
		fx.Invoke(serverLifecycle),
		fx.Options(),
		fx.Invoke(startGRPCServer),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatal(err, "Error starting application")
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	if err := app.Stop(context.Background()); err != nil {
		log.Fatal(err, "Error stopping application")
	}
}

func serverLifecycle(lc fx.Lifecycle, apiRouter *routers.ApiRouter, cf *configs.Configs) {
	server := &http.Server{
		Addr:    ":" + cf.Port,
		Handler: apiRouter.Engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Fatal(err, "Cannot start server,address")
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Infof("Stopping backend server.", cf.Port)
			return server.Shutdown(ctx)
		},
	})
}
func startGRPCServer(lc fx.Lifecycle) {
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Error(err, "error")
	}

	grpcServer := grpc.NewServer()
	generator.RegisterOrderServiceServer(grpcServer, interfaces.NewOrderServerService())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := grpcServer.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
					log.Error(err, "error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Stopping gRPC server.")
			grpcServer.GracefulStop()
			return nil
		},
	})
}
