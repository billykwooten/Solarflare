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
    firstRun = app.Flag("first", "Enables Ecobee Pin functions to authenticate your application on the first run.").Envar("ECOBEE_FIRST").Default("false").Bool()
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

	/// Only used when setting up the program for the first time in https://www.ecobee.com/consumerportal
	/// to get the bearer token
	if *firstRun {
		Code, err := ecobee.Get_ecobee_pin(*ecobeeApiKey)
		if err != nil {
			panic(err)
		}

		err = ecobee.GetToken(Code, *ecobeeApiKey)
		if err != nil {
			panic(err)
		}
	}


}