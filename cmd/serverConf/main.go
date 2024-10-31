package main

import (
	"flag"
	"log"
	"os"

	"github.com/farinap5/yalbaf/pkg/server"
	"gopkg.in/yaml.v2"
)


var DefaultConf string = `
server:
  http: 8080
  https: 4433
  host: 0.0.0.0
  key: server.key
  crt: server.crt

upstream:
  root: https://farinap5.com
  timeout: 5
  path: /
`


type Server struct {
	Server struct {
		HTTP  string `yaml:"http"`
		HTTPS string `yaml:"https"`
		Host  string `yaml:"host"`
		Key   string `yaml:"key"`
		Crt   string `yaml:"crt"`
	} `yaml:"server"`
	Upstream struct {
		Root    string `yaml:"root"`
		Timeout int    `yaml:"timeout"`
		Path    string `yaml:"path"`
	} `yaml:"upstream"`
}

func main() {
	var conf = flag.String("c", "none", "Set configuration file: `-c server.yml` ")
	flag.Parse()

	var yamlData []byte
	var err error
	var serverConf Server

	if *conf != "none" {
		yamlData, err = os.ReadFile(*conf)
		if err != nil {
			log.Println(err.Error())
			return
		}
	} else {
		yamlData = []byte(DefaultConf)
		log.Println("Using default configuration. Playing for tests only!")
	}

	err = yaml.Unmarshal(yamlData, &serverConf)
	if err != nil {
		log.Println(err.Error())
		return
	}


	s := server.New(serverConf.Upstream.Root)
	s.SetCertificate(serverConf.Server.Key, serverConf.Server.Crt)
	log.Printf("Using Key from %s and certificate from %s\n", serverConf.Server.Key, serverConf.Server.Crt)

	s.SetHTTPHost(serverConf.Server.Host+":"+serverConf.Server.HTTP)
	log.Printf("HTTP host on %s\n", serverConf.Server.Host+":"+serverConf.Server.HTTP)
	s.SetHTTPSHost(serverConf.Server.Host+":"+serverConf.Server.HTTPS)
	log.Printf("HTTPS host on %s\n", serverConf.Server.Host+":"+serverConf.Server.HTTPS)

	s.SetPath(serverConf.Upstream.Path)
	log.Printf("Path to %s\n", serverConf.Upstream.Path)

	s.StartServer()
}