package main

import (
	"io"
	"log"
	"os"

	"github.com/pengyeng/batch103"
)

type MyReader struct {
	batch103.FileReader
}

func (r *MyReader) Read() ([]batch103.BatchData, error) {

	var result []batch103.BatchData
	var myAWSUtil = &batch103.AWSUtils{}
	r.FileReader.SetFileName(os.Getenv("FILE_NAME"))
	myAWSUtil.SetRegion(os.Getenv("REGION"))
	err := myAWSUtil.DownloadFileFromS3Bucket(os.Getenv("BUCKET_NAME"), r.FileReader.GetFileName())
	if err != nil {
		return result, err
	}
	csvFileReader, err := r.OpenCSVFile()
	if err != nil {
		return result, err
	}

	for {

		record, err := csvFileReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		//Inserting File Record into Result set
		input := []string{record[0], record[1], record[2], record[3]}
		var batchData = &batch103.BatchData{}
		batchData = batchData.Create(input)
		result = append(result, *batchData)
	}
	return result, nil
}
