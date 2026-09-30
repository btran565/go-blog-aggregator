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

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	value, ok := c.cmdMap[cmd.name]
	if ok == false {
		return fmt.Errorf("command doesn't exist")
	}
	return value(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login requires a username argument")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("user has been set")
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	s := &state{
		cfg: &cfg, //assign s.cfg to the address of cfg assigned above
	}
	cmds := &commands{
		cmdMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
}
