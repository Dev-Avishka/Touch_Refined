package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var count_file string = "count.txt"

func incrementCount() {
	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}
	exeDir := filepath.Dir(exePath)
	countFilePath := filepath.Join(exeDir, count_file)

	file, err := os.OpenFile(countFilePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error opening count file:", err)
		return
	}
	defer file.Close()

	var count int
	_, err = fmt.Fscanf(file, "%d", &count)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error reading count file:", err)
		return
	}

	count++
	if count == 5 {
		fmt.Println("Remember to donate to me so I can buy treats for my dog")
		fmt.Println("And a sandwich for me to keep coding...")
		fmt.Println("But ultimately, your donation will be fueling my coding session at the nearby cafe!")
		fmt.Println("https://buymeacoffee.com/devavishka")
		count = 0
	}
	file.Seek(0, 0)
	_, err = fmt.Fprintf(file, "%d", count)
	if err != nil {
		fmt.Println("Error writing to count file:", err)
		return
	}
}

func main() {
	incrementCount()
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file name")
		return
	}
	file_name := os.Args[1]
	if strings.ToLower(file_name) == "minors" {
		fmt.Print("Are you sure you want to continue? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			fmt.Println("Creating file")
			fmt.Println("You sure you're not Diddy?")
			fmt.Println("Adding you to a watchlist for safety")
			file_name += ".txt"

		} else {
			fmt.Println("Operation cancelled")
			return
		}
	} else if strings.ToLower(file_name) == "god" {
		fmt.Println("Thou need heavenly approval to create such a file")
		fmt.Println("Thou is not worthy")
		fmt.Println("But thy wallet is worthy to donate to me")
		fmt.Println("https://buymeacoffee.com/devavishka")
		return
	} else if strings.ToLower(file_name) == "hawkings" {
		fmt.Print("Are you sure you want to continue? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			fmt.Println("Creating file")
			fmt.Println("Be warned, this file has nothing to do with hawkings")
			file_name += ".txt"

		} else {
			fmt.Println("Operation cancelled")
			return
		}
	} else if strings.ToLower(file_name) == "epstein" {
		fmt.Print("Are you over 18? (y/n): ")
		var response string
		fmt.Scanln(&response)
		if response == "y" || response == "Y" {
			fmt.Println("Creating file")
			fmt.Println("Nothing to see here")
			file_name += ".txt"

		} else {
			fmt.Println("You need to be over 18 to create this file")
			return
		}

	} else if strings.ToLower(file_name) == "devil" {
		fmt.Print("Thou hast embraced darkness")
		fmt.Println("proceed with caution")
		file_name += ".txt"
	} else if strings.ToLower(file_name) == "satan" {
		fmt.Print("Thou hast embraced darkness")
		fmt.Println("proceed with caution")
		file_name += ".txt"

	} else if strings.ToLower(file_name) == "grass" {
		fmt.Println("Impossible ")
		fmt.Println("You must go outside")
		fmt.Println("But you can touch your wallet and donate to me")
		fmt.Println("https://buymeacoffee.com/devavishka")
		return
	} else if strings.ToLower(file_name) == "donate" {
		fmt.Println("Donate to me so I can buy a sandwich and continue coding")
		fmt.Println("https://buymeacoffee.com/devavishka")
		return
	}

	file_extension := ""
	if i := strings.LastIndex(file_name, "."); i != -1 && i < len(file_name)-1 {
		file_extension = file_name[i:]
	}

	if file_extension == ".txt" {
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		fmt.Println("Created empty .txt file:", file_name)
	} else if file_extension == ".py" {
		//add python code to file
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		_, err = file.WriteString("print('Hello World!')")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			file.Close()
			return
		}

		err = file.Sync()
		if err != nil {
			fmt.Println("Error syncing file:", err)
		}
		file.Close()
		fmt.Println("Created file ", file_name)
	} else if file_extension == ".c" {
		//add c code to file
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		_, err = file.WriteString("#include <stdio.h>\n\nint main() {\n    printf(\"Hello, World!\\n\");\n    return 0;\n}")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			file.Close()
			return
		}

		err = file.Sync()
		if err != nil {
			fmt.Println("Error syncing file:", err)
		}
		file.Close()
		fmt.Println("Created file ", file_name)
	} else if file_extension == ".js" {
		//add javascript code to file
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		_, err = file.WriteString("console.log('Hello World!');")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			file.Close()
			return
		}

		err = file.Sync()
		if err != nil {
			fmt.Println("Error syncing file:", err)
		}
		file.Close()
		fmt.Println("Created file ", file_name)
	} else if file_extension == ".html" {
		//add html code to file
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		_, err = file.WriteString("<!DOCTYPE html>\n<html>\n<head>\n    <title>Hello World</title>\n</head>\n<body>\n    <h1>Hello World!</h1>\n</body>\n</html>")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			file.Close()
			return
		}

		err = file.Sync()
		if err != nil {
			fmt.Println("Error syncing file:", err)
		}
		file.Close()
		fmt.Println("Created file ", file_name)

	} else if file_extension == ".java" {
		//add java code to file
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		className := strings.TrimSuffix(file_name, ".java")
		_, err = file.WriteString("public class " + className + " {\n    public static void main(String[] args) {\n        System.out.println(\"Hello, World!\");\n    }\n}")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			file.Close()
			return
		}

		err = file.Sync()
		if err != nil {
			fmt.Println("Error syncing file:", err)
		}
		file.Close()
		fmt.Println("Created file ", file_name)
	} else {
		file, err := os.Create(file_name)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		fmt.Println("Created file with no special handling:", file_name)
	}
}
