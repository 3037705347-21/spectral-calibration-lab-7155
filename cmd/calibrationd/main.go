package main

import (
	"log"
	"net/http"
	"os"
	"spectralcalibrationlab/internal/httpapi"
	"spectralcalibrationlab/internal/lab"
)

func main() {
	address := os.Getenv("CALIBRATION_ADDR")
	if address == "" {
		address = "127.0.0.1:18084"
	}
	server := httpapi.New(lab.NewEngine())
	log.Printf("spectral calibration service listening on %s", address)
	if err := http.ListenAndServe(address, server.Handler()); err != nil {
		log.Fatal(err)
	}
}
