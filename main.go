package main

import (
	"fmt"
	"os"
)

func main() {
var userinfo string

// Taking input from user to name the componenet
	fmt.Print("What name will you like to use : ")
	fmt.Scanln(&userinfo)

	defer createStructure(userinfo);
	
	
}

func createStructure (structureName string, ) {

	// Project structure layouts
	projectStruct := map[string]interface{}{
		structureName: map[string]interface{}{
			"assests": map[string]interface{}{
				"images": "file",
				"brands": "file",
			},
			"screen":    nil,
			"templates": nil,
			"hooks":     nil,
			"setups": map[string]interface{}{
				"router-manager": nil,
				"auth-manager":   nil,
			},
		},
	}

	currentDir := "."
	err := createDirectories(projectStruct, currentDir)

	if err != nil {
		fmt.Println("Error creating ArchGen framework...", err)
	}

	fmt.Println("ArchGen run successfully completed")
}

func createDirectories(directoriesMap map[string]interface{}, pathname string) error {
	for key, value := range directoriesMap {
		dirPath := pathname + "/" + key

		fmt.Println("this is ", value)

		// Create directories
		err := os.Mkdir(dirPath, 0755)
		// if error found
		checkNilErr(err)

		// Checking if there is a subdirectory ie. another map directory
		if subDirectory, ok := value.(map[string]interface{}); !ok {
			// do nothing
		} else {
			// creates subdirectories
			err := createDirectories(subDirectory, dirPath)
			// creating files in subdirectories
			for i, v := range subDirectory {
				sub_dirPath := dirPath + "/" + i
				fmt.Println("------------- PATH --------------- \n " + sub_dirPath + "\n")
				// subs = os.Create([v])
				if(v != nil && v == "file") {
					fmt.Println(v)
					sub_file, err := os.Create(sub_dirPath + "./images.tsx")

					defer sub_file.Close();

					checkNilErr(err)
				}
				
			}
			checkNilErr(err)
		}

	}
	return nil
}


func checkNilErr (err error) error {
	if(err != nil) {
		panic(err)
	}
	return nil
}
