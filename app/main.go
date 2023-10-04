package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadFile(downloader *s3manager.Downloader, bucketName string, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = downloader.Download(
		file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(fileName),
		},
	)

	return err
}

type Processor struct{}

func (p Processor) Process() error {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String(os.Getenv("REGION")),
		},
	})

	if err != nil {
		return err
	}
	fileName := os.Getenv("FILE_NAME")
	bucketName := os.Getenv("BUCKET_NAME")
	downloader := s3manager.NewDownloader(sess)
	err = DownloadFile(downloader, bucketName, fileName)

	if err != nil {
		return err
	}
	file, _ := os.Open(fileName)
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()
	fmt.Println("==== AWS Batch Start ====")
	fmt.Println("Process File : ", fileName, " in ", bucketName)
	for _, eachrecord := range records {
		if eachrecord[1] == "TREASURE" {
			fmt.Println(fileName, " : ", eachrecord[0])
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("==== AWS Batch End ====")
	return nil
}

func main() {

	var myJobLauncher JobLauncher
	myProcessor := Processor{}
	myJobLauncher.Run(myProcessor)
}
