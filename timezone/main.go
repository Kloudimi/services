package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/services/timezone/handler"
	pb "github.com/micro/services/timezone/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("timezone"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterTimezoneHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}