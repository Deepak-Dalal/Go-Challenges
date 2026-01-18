/*
The "JSON Config Watcher"
Now that you've mastered basic concurrency, let's look at File I/O, JSON handling, and Long-running loops. This is a common pattern for background workers or "sidecars" in cloud-native apps. ☁️

📋 The Goal:
Build a program that reads a configuration file and automatically reloads it whenever you save changes to the file, without stopping the program.

📝 Step 1: Create a config.json
Create a simple file in your project folder:

# JSON

	{
	    "app_name": "GoWatcher",
	    "version": "1.0.0",
	    "debug_mode": true
	}

🛠️ Step 2: The Coding Requirements:
Define the Struct: Create a Config struct. Use Struct Tags (e.g., `json:"app_name"`) to map the JSON keys to your struct fields. 🏷️

State Management: Store the "Last Modified" time of the file in a variable. 🕒

The Watcher Loop:

Use an infinite loop for { ... }.

Use os.Stat() to check the file's metadata.

Compare the ModTime() with your stored version.

The Reload Logic:

If the file is newer, use os.ReadFile() and json.Unmarshal() to update your struct. 🔄

Print: "Config Updated: [New Values]"

Efficiency: Use time.Sleep(2 * time.Second) so you don't burn 100% of your CPU checking the file millions of times per second. 😴
*/
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