package main

import (
	"github.com/kubecost/cost-model/pkg/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	// runs the appropriate application mode using the default cost-model command
	// see: github.com/kubecost/cost-model/pkg/cmd package for details
	if err := cmd.Execute(nil); err != nil {
		log.Fatal().Err(err)
	}
}
