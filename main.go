package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Azamjon99/admin-api-gateway/src/application/controllers"
	grpcclient "github.com/Azamjon99/admin-api-gateway/src/application/grpc"
	logisticsstafftpb "github.com/Azamjon99/admin-api-gateway/src/application/protos/logistics_staff"
	logisticsvehicletpb "github.com/Azamjon99/admin-api-gateway/src/application/protos/logistics_vehicle"
	"github.com/Azamjon99/admin-api-gateway/src/application/services"
	"github.com/Azamjon99/admin-api-gateway/src/infrastructure/config"
	"github.com/Azamjon99/admin-api-gateway/src/infrastructure/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

func main() {
	config, err := config.Load()

	if err != nil {
		panic(err)
	}

	logger, err := config.NewLogger()

	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	tokenSvc := jwt.NewService(config.JWTSecret, config.JWTExpiresInSec)



	logisticsStaffConn := grpcclient.NewClient(config.LogisticStaffAddress)
	defer logisticsStaffConn.Close()

	logisticsvehicleConn := grpcclient.NewClient(config.LogisticsVehicleService)
	defer logisticsvehicleConn.Close()

	logisticsStafClient := logisticsstafftpb.NewStaffServiceClient(logisticsStaffConn)
	logiticvehicleClient := logisticsvehicletpb.NewVehicleServiceClient(logisticsvehicleConn)
	//Application
	driverAppSvc := services.NewDriverApplicationService(logisticsStafClient)

	// Controllers

	root := gin.Default()
	root.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	controllers.Init(
		root,
		tokenSvc,
		driverAppSvc,
		logiticvehicleClient,
	)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(osSignals)

	g, ctx := errgroup.WithContext(ctx)
	var httpServer *http.Server

	g.Go(func() error {
		httpServer = &http.Server{
			Addr:    config.HttpPort,
			Handler: root,
		}
		logger.Debug("main: started http server", zap.String("port", config.HttpPort))
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	select {
	case <-osSignals:
		logger.Info("main: received os signal, shutting down")
		break
	case <-ctx.Done():
		break
	}

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if httpServer != nil {
		httpServer.Shutdown(shutdownCtx)
	}

	if err := g.Wait(); err != nil {
		logger.Error("main: server returned an error", zap.Error(err))
	}
}
