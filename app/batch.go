package main

import (
	"fmt"
	"os"
)

type JobLauncher struct{}

func (j *JobLauncher) Run(p ProcessorType) {
	var err error
	err = p.Process()
	if err != nil {
		fmt.Println("Print Error ", err)
		os.Exit(1)
	}

}
