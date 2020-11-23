package settings

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var data = `
url: etcd.cwxstat.io:2379
testurl: localhost:2379
certs:
  directory: /certs
  ca: ca.pem
  client: client.pem
  clientKey: client-key.pem
  usernamePassword: [root, password]
tls: true
`

type T struct {
	URL   string
	TestURL string `yaml:"testurl"`
	Certs struct {
		Directory string `yaml:"directory"`
		Ca        string `yaml:"ca"`
		Client    string `yaml:"client"`
		ClientKey string `yaml:"clientKey"`
	}
	UsernamePassword []string `yaml:",flow"`
	TLS              bool
}

func CreateDefault() {

	home, err := os.UserHomeDir()
	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	d, err := yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(home+"/"+".go.etcd.yaml", d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadConfig() (T, error) {
	t := T{}
	home, err := os.UserHomeDir()

	b, err := ioutil.ReadFile(home + "/" + ".go.etcd.yaml")
	if err != nil {
		return t, err
	}

	err = yaml.Unmarshal(b, &t)
	if err != nil {
		return t, err
	}
	return t, err

}

func TestRead() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
