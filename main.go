package main

import (
	"fmt"

	"github.com/btran565/go-blog-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		return
	}
	cfg.SetUser("brandon")
	cfg, err = config.Read()
	fmt.Print(cfg)
}
