// TEST FILE pro načtení configu a otestovaní knihovny
package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Cars struct {
		Count          int64
		ArrivalTimeMin string `yaml:"arrival_time_min"`
		ArrivalTimeMax string `yaml:"arrival_time_max"`
	}

	Stations struct {
		Name       string
		Parameters struct {
		}
	}
}

func testLoad(filename string) (*Configuration, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Configuration{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}

func main() {
	fmt.Println("Píčo")
	c, err := testLoad("./configs/prod.yaml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", c)
}
