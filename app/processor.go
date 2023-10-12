package main

import (
	"log"

	"github.com/pengyeng/batch103"
)

type MyProcessor struct {
	batch103.BaseProcessor
}

type ProcessError struct{}

func (m *ProcessError) Error() string {
	return "Empty Record to Process"
}

func (r *MyProcessor) Process(data []batch103.BatchData) ([]batch103.BatchData, error) {
	var processRecords []batch103.BatchData
	if len(data) == 0 {
		return processRecords, &ProcessError{}
	}
	for i := 0; i < len(data); i++ {
		if data[i].IsActive() {
			var record = data[i].GenericData
			if record[1] == "TREASURE" {
				log.Println("TREASURE FOUND : ", record[0])
			} else {
				data[i].Reject(batch103.StgProcess)
			}
			processRecords = append(processRecords, data[i])
		}

	}
	return processRecords, nil
}
