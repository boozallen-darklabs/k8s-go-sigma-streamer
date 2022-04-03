package main

import (
	"context"

	"github.com/Jeffail/benthos/v3/public/service"

	// Import all Benthos components
	_ "github.com/Jeffail/benthos/v3/public/components/all"
	_ "github.com/boozallen-darklabs/k8s-go-sigma-streamer/build/engine/processor"
)

func main() {
	service.RunCLI(context.Background())
}
