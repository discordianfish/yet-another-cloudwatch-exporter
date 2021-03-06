package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type conf struct {
	Jobs []job `yaml:"jobs"`
}

type job struct {
	Discovery discovery `yaml:"discovery"`
	Metrics   []metric  `yaml:"metrics"`
}

type discovery struct {
	Region       string   `yaml:"region"`
	Type         string   `yaml:"type"`
	SearchTags   []tag    `yaml:"searchTags"`
	ExportedTags []string `yaml:"exportedTags"`
}

type metric struct {
	Name       string `yaml:"name"`
	Statistics string `yaml:"statistics"`
	Period     int    `yaml:"period"`
	Length     int    `yaml:"length"`
	NilToZero  bool   `yaml:"nilToZero"`
}

type tag struct {
	Key   string `yaml:"Key"`
	Value string `yaml:"Value"`
}

func (c *conf) getConf(file *string) *conf {
	yamlFile, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	for _, job := range c.Jobs {
		if !stringInSlice(job.Discovery.Type, supportedServices) {
			log.Fatalf("Service is not in known list!: %v", job.Discovery.Type)
		}
	}

	return c
}
