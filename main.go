package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"

	"os"
	"Solarflare/ecobee"
)

var (
	//Set your kingpin help flags for app var
	app              = kingpin.New("solarflare", "Program for Ecobee Thermostats").Author("Kyle Wooten")
	ecobeeApiKey     = app.Flag("apikey", "Ecobee API Key (Example Format: 4ok2o45k4o25k4o25kok) ").Envar("ECOBEE_APIKEY").Required().String()
)

var (
	Version     = "0.0.2"
	GitCommit   = "HEAD"
	BuildStamp  = "UNKNOWN"
	FullVersion = Version + "+" + GitCommit + "-" + BuildStamp
)

func init() {
	app.Version(FullVersion)
}

func main() {
	/// Parse kingpin app flags for --help option
	kingpin.MustParse(app.Parse(os.Args[1:]))

	err := ecobee.Get_ecobee_pin(*ecobeeApiKey)
	if err != nil {
		panic(err)
	}
}