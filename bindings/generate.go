//go:build generate
// +build generate

//go:generate go run -tags generate generate.go
package main

import (
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/unpackdev/solgo/standards"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Lets have nice logging of what is going on here instead of just print lines...
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		panic(err)
	}

	// First we need to load standards from the standards package as there we have all of the necessary information.
	if err := standards.LoadStandards(); err != nil {
		logger.Fatal("failed to load standards", zap.Error(err))
	}

	// Second step is to get all of the supported and registered standards.
	regStandards := standards.GetRegisteredStandards()

	// It's just too much of the time to implement the exact mechanism for binding generation utilising bind package from go-ethereum.
	// As this is not something affecting the production and performance we can just call command line abigen tool to deal with
	// this problematic fast. Wanna do it better? Feel free to contribute.
	for name, contract := range regStandards {
		logger.Info("Generating golang bindings", zap.String("name", name.String()), zap.Any("standard", contract))

		tmpFile, err := ioutil.TempFile("", "abi_*.json")
		if err != nil {
			logger.Error("Failed to create temporary file for ABI", zap.Error(err))
			continue
		}
		defer tmpFile.Close()

		if _, err = tmpFile.WriteString(contract.GetABI()); err != nil {
			logger.Error("Failed to write ABI to temporary file", zap.Error(err))
			continue
		}

		pkg := contract.GetPackageName()

		if err := os.MkdirAll(pkg, 0700); err != nil {
			logger.Error("Failed to create package directory", zap.String("directory", pkg), zap.Error(err))
			continue
		}

		// Construct and execute abigen command
		cmd := exec.Command(
			"abigen",
			"--abi", tmpFile.Name(),
			"--type", name.String(),
			"--pkg", pkg,
			"--out", contract.GetPackageOutputPath(),
		)

		if err := cmd.Run(); err != nil {
			logger.Error("abigen command failed", zap.String("command", cmd.String()), zap.Error(err))
			continue
		}

		logger.Info("Generated golang bindings", zap.String("output", contract.GetPackageOutputPath()))
	}

	logger.Info("Finished generating golang bindings")
}
