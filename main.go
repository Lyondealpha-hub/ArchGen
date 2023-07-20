package main

import (
	"fmt"
	"os"
)

func main() {

	var userinput string

	// Taking input from user to name the componenet
	fmt.Println("What name will you like to use : ")
	fmt.Scanln(&userinput)

	// Project structure layouts

	projectStruct := map[string]interface{}{
		userinput: map[string]interface{}{
			"assests": map[string]interface{}{
				"images": "images.tsx",
				"brands": "logo.jsx",
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
		if err != nil {
			return err
		}

		// Checking if there is a subdirectory ie. another map directory
		if subDirectory, ok := value.(map[string]interface{}); !ok {
			// creating an instruction file
			// errFiles := os.Mkdir(filespath, 0755)

			// if errFiles != nil {
			// 	return errFiles
			// }

		} else {

			// creates subdirectories
			err = createDirectories(subDirectory, dirPath)

			if err != nil {
				return err
			}
		}

	}
	return nil
}
