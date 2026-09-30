package main

import (
	"fmt"

	"github.com/btran565/go-blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login requires a username argument")
	}
	fmt.Println("user has been set")
	return s.cfg.SetUser(cmd.args[0]) //already returns error, so i can return it directly w/out check

}

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	cfg.SetUser("brandon")
	cfg, err = config.Read()
	fmt.Print(cfg)
}
