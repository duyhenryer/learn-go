package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Service struct {
	Name           string   `yaml:"name"`
	LongName       string   `yaml:"longName"`
	Description    string   `yaml:"description"`
	Kind           string   `yaml:"kind"`
	Maintainer     string   `yaml:"maintainer"`
	Team           string   `yaml:"team"`
	ExecSponsor    string   `yaml:"exec_sponsor"`
	ProductManager string   `yaml:"product_manager"`
	Repo           string   `yaml:"repo"`
	TeamSlack      string   `yaml:"team_slack"`
	Sev2           Sev      `yaml:"sev2"`
	Sev3           Sev      `yaml:"sev3"`
	Qos            string   `yaml:"qos"`
	Tier           int      `yaml:"tier"`
	Dependencies   []string `yaml:"dependencies"`
}

type Sev struct {
	Issue string `yaml:"issue"`
	TAA   string `yaml:"taa"`
	Slack string `yaml:"slack"`
}

type File struct {
	Service        string   `yaml:"service"`
	CodeOwnwes     []string `yaml:"codeOwnwes"`
	MaintainerTeam string   `yaml:"maintainer_team"`
	ReviewTeams    []string `yaml:"review_teams"`
}

type Config struct {
	SHA            string          `yaml:"sha"`
	ServiceCatalog ServiceCatalog  `yaml:"service_catalog"`
	Files          map[string]File `yaml:"files"`
}

type ServiceCatalog struct {
	Version   int       `yaml:"version"`
	Ownership []Service `yaml:"ownership"`
}

func main() {
	data, err := os.ReadFile("data.yaml")
	if err != nil {
		log.Fatal("Error reading YAML file: %v", err)
	}
	// Unmarshal YAML into a Config struct
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("Error unmarshalling YAML: %v", err)
	}
	fmt.Printf("%+v\n", config)
}
