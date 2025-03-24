package main

import (
	"duplicate-finder/cmd"
	"log/slog"
)

func main() {
	dpf := cmd.NewCommand()
	err := dpf.Execute()
	if err != nil {
		slog.Error("Execution command error: ", err)
	}
}
