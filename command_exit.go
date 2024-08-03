package main

import "os"

func commandExit(cfg *config, areaName *string) error{
	os.Exit(0)
	return nil
}