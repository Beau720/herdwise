package main

import (
	"herdwise/http"
	"herdwise/service/database"
	"herdwise/service/device"
	"herdwise/service/farmer"
	"log"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	//Connect to the DB
	db_config := &database.Config{
		Username:       "root",
		Password:       "FireStore@321$",
		ConnectionType: "tcp",
		Host:           "127.0.0.1",
		Port:           3306,
		Name:           "herdwise",
	}

	http_config := &http.Config{
		Host:          "",
		Port:          "8085",
		FarmerService: farmer.NewFarmer(db_config),
		DeviceService: device.NewDevice(db_config),
	}

	http.Start(http_config)
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}

func main() {

	svcConfig := &service.Config{
		Name:        "Herd_Wise",
		DisplayName: "Herd Wise",
		Description: "Tracking lifestock project that listens to port 8082.",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)

	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}

}
