package gadget

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type Binding struct {
	name string
	path string

	config   *Config
	function Function
}

func CreateBinding(c *Config, f Function, name string) *Binding {
	functionPath := filepath.Join(f.Path(), f.Name())
	configPath := filepath.Join(c.path, c.name)
	linkPath := filepath.Join(configPath, name)

	fmt.Println(functionPath, linkPath)

	binding := &Binding{
		name:     name,
		path:     configPath,
		config:   c,
		function: f,
	}

	err := os.Symlink(functionPath, linkPath)
	if err != nil {
		log.Fatal(err)
	}

	return binding
}
