package main

import "github.com/ghanithan/goBilling/instrumentation"

func main() {

	logger := instrumentation.InitInstruments()

	logger.Info("Started goBilling")

}
