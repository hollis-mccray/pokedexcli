package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	if cfg.Next == "" {
		fmt.Println("Already at the end of the list.")
		return nil
	}
	locationResp, err := cfg.pokeapiClient.ListlocationList(cfg.Next)
	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.Previous == "" {
		fmt.Println("Already at the beginning of the list.")
		return nil
	}
	locationResp, err := cfg.pokeapiClient.ListlocationList(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = locationResp.Next
	cfg.Previous = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
