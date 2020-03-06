package main

import (
	"github.com/gappy023/inventory-srv/handler"
	"github.com/gappy023/inventory-srv/subscriber"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	inventory "github.com/gappy023/inventory-srv/proto/inventory"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("mu.micro.book.srv.inventory"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	inventory.RegisterInventoryHandler(service.Server(), new(handler.Inventory))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("mu.micro.book.srv.inventory", service.Server(), new(subscriber.Inventory))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
