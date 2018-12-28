package discovery

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

func Exec() error {
	config := consul.DefaultConfig()
	config.Address = "localhost:8500"
	name := "discovery"
	cli, err := NewClient(config, "localhost", name, 8080)
	if err != nil {
		return err
	}
	if err := cli.Register([]string{"Go", "Awesome"}); err != nil {
		return err
	}
	entries, _, err := cli.Service(name, "Go")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Printf("%#v\n", entry.Service)
	}
	return nil
}
