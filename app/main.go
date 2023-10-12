package main

import (
	"github.com/pengyeng/batch103"
)

func main() {

	var myJobLauncher batch103.JobLauncher
	var myReader batch103.ReaderType
	var myProcessor batch103.ProcessorType

	myReader = &MyReader{}
	myProcessor = &MyProcessor{}
	myProcessorList := []batch103.ProcessorType{}
	myProcessorList = append(myProcessorList, myProcessor)

	//Prepare Writer List
	myWriterList := []batch103.WriterType{}
	myJobLauncher.Run(myReader, myProcessorList, myWriterList)

}
