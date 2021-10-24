package memory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//write file every given second
var Frequency = 60
var TempFilePath = "test.json"

func ActivateWriter() {
	for {
		fmt.Println(time.Now(),"- I will wait",Frequency,"seconds to rerun operation")
		time.Sleep(time.Duration(Frequency)*time.Second)
		fmt.Println(time.Now(),"- Now, I am parallelly running to write memory to file")
		WriteToFile()
		fmt.Println(time.Now(),"- I finished my job at ")

	}
}
//provide to write to file
func WriteToFile() {
	file:= Instance.GetValue()
	_ = ioutil.WriteFile(TempFilePath, []byte(file), 0644)
}


func ReadFileAndWriteToMemory(){
	jsonFile, err := os.Open(TempFilePath)
	// if an error occurs
	if err != nil {
		return
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	jsonMap := make(map[string]interface{})
	_ = json.Unmarshal(byteValue, &jsonMap)
	Instance.items = jsonMap
}
