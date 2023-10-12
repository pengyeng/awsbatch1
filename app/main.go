package main

import (
	"github.com/pengyeng/batch103"
)

func main() {

	var myJobLauncher batch103.JobLauncher
	var myReader batch103.ReaderType
	var myProcessor batch103.ProcessorType
	var myWriter batch103.WriterType

	myReader = &MyReader{}
	myProcessor = &MyProcessor{}
	myProcessorList := []batch103.ProcessorType{}
	myProcessorList = append(myProcessorList, myProcessor)

	//Prepare Writer List
	myWriter = &MyWriter{}
	myWriterList := []batch103.WriterType{}
	myWriterList = append(myWriterList, myWriter)

	myJobLauncher.Run(myReader, myProcessorList, myWriterList)

}
