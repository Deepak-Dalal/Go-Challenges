package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Config struct {
	AppName   string `json:"app_name"`
	Version   string `json:"version"`
	DebugMode bool   `json:"debug_mode"`
}

func main() {
	var lastModified time.Time
	var config Config
	fileFound := true
	for {
		fileInfo,err:=os.Stat("./config.json");
		if err!=nil {
			if fileFound == true {

				fmt.Println("Error in accessing the file at the given path.",err)
				fileFound=false
			}
			time.Sleep(2*time.Second)
			continue;
		}
		if !fileFound {
			fileFound=true;
			fmt.Println("Found file at the path")
		}
		modificationTime := fileInfo.ModTime()

		if modificationTime.After(lastModified) {
			fileContent, readfileError:= os.ReadFile("./config.json")
			if readfileError != nil {
				fmt.Println("Error in reading file at the path")
				time.Sleep(2*time.Second) 
        		continue
			}
			unmarshalError:=json.Unmarshal(fileContent, &config)
			if unmarshalError != nil {
				fmt.Println("Error in parsing config json file")
			} else {
				fmt.Printf("Config updated: [appName:%s, version:%s, debugMode:%t]", config.AppName, config.Version, config.DebugMode)
				fmt.Println("")
			}
			lastModified = modificationTime
		}

		time.Sleep(2*time.Second)
	}
}