package main

import (
	"github.com/d3v-friends/go-pure/fnPanic"
	"github.com/spf13/cobra"
)

func main() {
	var cmd = &cobra.Command{
		Use: "go-grpc",
	}

	fnPanic.On(cmd.Execute())
}

type Config struct {
	Proto   string   `yaml:"proto"`
	Out     string   `yaml:"out"`
	Package string   `yaml:"package"`
	Include []string `yaml:"include"`
	Import  []string `yaml:"import"`
}
