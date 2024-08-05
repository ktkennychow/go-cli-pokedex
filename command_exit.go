package main

import "os"

func commandExit(_ *config, _ *string, _ *string) error{
	os.Exit(0)
	return nil
}