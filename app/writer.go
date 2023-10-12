package main

import (
	"os"

	"github.com/pengyeng/batch103"
)

type MyWriter struct {
	batch103.BaseWriter
}

func (w *MyWriter) Write(data []batch103.BatchData) error {

	var concatenatedContent string
	for i := 0; i < len(data); i++ {
		if data[i].IsActive() {
			var record = data[i].GenericData
			concatenatedContent = concatenatedContent + record[0] + "," + record[1] + "," + record[2] + "\n"
		}
	}
	fileContent := []byte(concatenatedContent)
	err := os.WriteFile("output.txt", fileContent, 0644)
	if err != nil {
		return err
	}
	var myAWSUtil = &batch103.AWSUtils{}
	myAWSUtil.SetRegion(os.Getenv("REGION"))
	uploadErr := myAWSUtil.UploadFileToS3Bucket(os.Getenv("BUCKET_NAME"), os.Getenv("OUTPUT_FILE"))
	if uploadErr != nil {
		return uploadErr
	}
	return nil
}
