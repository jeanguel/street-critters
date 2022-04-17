package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"

	"github.com/jeanguel/street-critters/api/config"
)

func main() {
	args := struct {
		PerformReset bool
		AddData      bool
		ShowHelp     bool
	}{}

	flag.BoolVar(&args.PerformReset, "perform-reset", false, "Resets all the data, tables, types, and functions in the database")
	flag.BoolVar(&args.AddData, "add-data", false, "Adds the core data to the database")
	flag.BoolVar(&args.ShowHelp, "h", false, "Shows the script help and usage")
	flag.Parse()

	flag.Usage = scriptHelp
	if args.ShowHelp {
		flag.Usage()
		flag.PrintDefaults()
		return
	}

	if !args.PerformReset && !args.AddData && !args.ShowHelp {
		panic("There is nothing to do!")
	}

	config.InitializeApplication()
	logger := config.CreateWorkerLogger("database-script")
	logger.Info.Println("database script initialized", args)

	defer func() {
		logger.Info.Printf("completed execution\n\n")
		config.CloseApplication()
	}()

	dbConfig := struct {
		CoreScripts []string `toml:"core"`
	}{}

	c, err := ioutil.ReadFile("database/db.config.toml")
	if err != nil {
		logger.Error.Panicln("db config read error:", err.Error())
	}
	if _, err := toml.Decode(string(c), &dbConfig); err != nil {
		logger.Error.Panicln("db config decode error:", err.Error())
	}

	if args.PerformReset {
		for _, script := range dbConfig.CoreScripts {
			scriptBytes, err := ioutil.ReadFile(filepath.Join("database", script))
			if err != nil {
				logger.Error.Panicf("[%s] db script read error: %s\n", script, err.Error())
			}

			err = config.ExecRawPSQLQuery(string(scriptBytes))
			if err != nil {
				logger.Error.Panicf("[%s] db script execution error: %s\n", script, err.Error())
			}

			logger.Info.Printf("[%s] executed successfully\n", script)
		}
	}
}

func scriptHelp() {
	fmt.Println("Usage:")
	fmt.Println("\tgo run database/main.go [-perform-reset] [-add-data] [-h]")

	fmt.Println("\nStreet Critters database utility script")
	fmt.Println("\nParameters:")
}
