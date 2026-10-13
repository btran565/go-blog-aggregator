package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/btran565/go-blog-aggregator/internal/config"
	"github.com/btran565/go-blog-aggregator/internal/database" //all database directory files are in 'package database'
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error //key: name of command, value: function mapped to name
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
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("user has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register requires a name argument")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("user was created: %v\n", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Reset(context.Background())
	if err != nil {
		fmt.Println("database reset failed")
		os.Exit(1)
	}
	fmt.Println("database reset successfully")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("failed to return users")
		os.Exit(1)
	}
	for _, val := range users {
		if val.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", val.Name)
		} else {
			fmt.Printf("* %v\n", val.Name)
		}
	}
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

}

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	dbQueries := database.New(db) //New() function from database/db.go
	s := &state{
		cfg: &cfg, //assign s.cfg to the address of cfg assigned above
		db:  dbQueries,
	}
	cmds := &commands{
		cmdMap: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 { //os.Args is a []string
		fmt.Printf("Error: less than 2 arguments")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(s, cmd) //run command the user called
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
