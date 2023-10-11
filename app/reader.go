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
	log.Println("Running Reader")
	var result []batch103.BatchData
	r.SetFileName("test1.csv")
	err := r.DownloadFileFromS3Bucket(os.Getenv("REGION"), os.Getenv("BUCKET_NAME"))
	//err := r.DownloadFileFromS3Bucket("ap-southeast-1", "aws-batch-file-repo1")
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
	log.Println("No Of Record Retrieved ", len(result))
	return result, nil
}
