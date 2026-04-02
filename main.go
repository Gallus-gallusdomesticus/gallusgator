package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Gallus-gallusdomesticus/gallusgator/internal/config"
	"github.com/Gallus-gallusdomesticus/gallusgator/internal/database" //importing for side effect not used directly
	_ "github.com/lib/pq"
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

	progCmds.register("login", handlerLogin)       //register login handler function
	progCmds.register("register", handlerRegister) //register register handler function
	progCmds.register("users", handlerUsers)       //register users handler function
	progCmds.register("reset", handlerReset)       //register reset handler function

	if len(os.Args) < 2 { //check lengths of the command
		log.Fatal("Not enough arguments provided.")
	}

	cmdName := os.Args[1] //the command name
	cmdArg := os.Args[2:] //the command argument

	cmd := command{ //make new instance of command from the args
		name: cmdName,
		args: cmdArg,
	}

	db, err := sql.Open("postgres", cfg.DbURL) //load in database URL to config struct

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db) //make a new database queries using generated database package

	progState.db = dbQueries

	if err := progCmds.run(progState, cmd); err != nil {
		log.Fatal(err)
	} //run the program

}

type state struct {
	db  *database.Queries //struct that hold the database queries
	cfg *config.Config    //struct that hold pointer to config (to give our handler the application state)
}
