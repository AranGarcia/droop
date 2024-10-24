package main

import (
	"fmt"

	"github.com/AranGarcia/droop/dnd/internal/adapters/secondary/characters"
	"github.com/AranGarcia/droop/dnd/internal/core/services"
)

func initDNDAPI() (*services.DND, error) {
	charactersConfig := characters.Config{
		Addr: charactersServiceAddr,
	}
	charactersService, err := characters.NewClient(charactersConfig)
	if err != nil {
		return nil, fmt.Errorf("character service client failure; %v", err)
	}

	dndService := services.NewDNDService(charactersService)
	return &dndService, nil
}
