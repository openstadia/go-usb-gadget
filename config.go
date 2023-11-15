package gadget

import (
	"fmt"
	o "github.com/openstadia/go-usb-gadget/option"
	"log"
	"os"
	"path/filepath"
)

const DefaultConfigLabel = "config"
const ConfigsDir = "configs"

type Config struct {
	name  string
	path  string
	label string
	id    int

	gadget   *Gadget
	bindings []*Binding
}

type ConfigAttrs struct {
	BmAttributes o.Option[uint8]
	BMaxPower    o.Option[uint8]
}

type ConfigStrs struct {
	Configuration string
}

func CreateConfig(gadget *Gadget, label string, id int) *Config {
	path := filepath.Join(gadget.path, gadget.name, ConfigsDir)
	name := fmt.Sprintf("%s.%d", label, id)

	config := &Config{
		gadget: gadget,
		name:   name,
		path:   path,
		label:  label,
		id:     id,
	}

	err := os.MkdirAll(filepath.Join(path, name), os.ModePerm)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	gadget.configs = append(gadget.configs, config)

	return config
}

func (c *Config) SetAttrs(attrs *ConfigAttrs) {
	if attrs.BMaxPower.IsSome() {
		WriteDec(c.path, c.name, "MaxPower", int(attrs.BMaxPower.Value()))
	}

	if attrs.BmAttributes.IsSome() {
		WriteHex8(c.path, c.name, "bmAttributes", attrs.BmAttributes.Value())
	}
}

func (c *Config) SetStrs(strs *ConfigStrs, lang int) {
	langStr := fmt.Sprintf("0x%x", lang)
	path := filepath.Join(c.path, c.name, StringsDir, langStr)

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	WriteString(path, "", "configuration", strs.Configuration)
}
