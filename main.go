package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/thegrumpylion/uop/op"
	"github.com/thegrumpylion/uop/storage"
	"gopkg.in/yaml.v3"
)

func main() {
	//we will run on :9998
	port := "9998"
	//which gives us the issuer: http://localhost:9998/
	issuer := fmt.Sprintf("http://localhost:%s/", port)

	// register the clients
	if err := registerClients("clients.yaml"); err != nil {
		log.Fatal(err)
	}

	// the OpenIDProvider interface needs a Storage interface handling various checks and state manipulations
	// this might be the layer for accessing your database
	// in this example it will be handled in-memory
	us, err := storage.NewUserStore("users.yaml")
	if err != nil {
		log.Fatal(err)
	}
	storage := storage.NewStorage(us)

	router := op.SetupServer(issuer, storage)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Printf("server listening on http://localhost:%s/", port)
	log.Println("press ctrl+c to stop")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func registerClients(s string) error {
	data, err := os.ReadFile(s)
	if err != nil {
		return err
	}

	clients := []*storage.ClientConfig{}
	if err := yaml.Unmarshal(data, &clients); err != nil {
		return err
	}

	storage.RegisterClients(clients...)

	return nil
}
