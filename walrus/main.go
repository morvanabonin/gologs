package main

import (
	"github.com/morvanabonin/gologs/walrus/logger"
)

func main() {
	// logs
	logger.Debug("Eu logo esse exemplo de debug")
	logger.Trace("Eu logo esse exemplo de trace")
	logger.Info( "Eu logo esse exemplo de info")
	logger.Warn("Eu logo esse exemplo de warn")
	logger.Error("Eu logo esse exemplo de error")
}
