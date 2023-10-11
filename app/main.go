package main

import (
	"github.com/pengyeng/batch103"
)

func main() {

	var myJobLauncher batch103.JobLauncher
	var myReader batch103.ReaderType
	myReader = &MyReader{}
	myProcessorList := []batch103.ProcessorType{}

	//Prepare Writer List
	myWriterList := []batch103.WriterType{}
	myJobLauncher.Run(myReader, myProcessorList, myWriterList)

}
