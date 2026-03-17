package main

import (
	"log"
	"os"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/config"
)

func main() {
	cfg, err := config.Read() //read config file

	if err != nil {
		log.Fatal(err)
	}

	progState := &state{ //store the config in new state struct
		cfg: &cfg,
	}

	progCmds := commands{ //make a commands struct with initialized map of handler function
		handlers: make(map[string]func(*state, command) error),
	}

	progCmds.register("login", handlerLogin) //register login handler function

	if len(os.Args) < 2 { //check lengths of the command
		log.Fatal("not enough arguments provided")
	}

	cmdName := os.Args[1] //the command name
	cmdArg := os.Args[2:] //the command argument

	cmd := command{ //make new instance of command from the args
		name: cmdName,
		args: cmdArg,
	}

	if err := progCmds.run(progState, cmd); err != nil {
		log.Fatal(err)
	} //run the program

}

type state struct {
	cfg *config.Config //struct that hold pointer to config (to give our handler the application state)
}
