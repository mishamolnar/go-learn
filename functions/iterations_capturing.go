package functions

import (
	"fmt"
	"os"
)

func ShowFunctionClousure() {
	strs := []string{"str1", "str2", "str3"}
	var prints []func() string
	for _, str := range strs {
		fmt.Println(&str)
		strCurr := str
		prints = append(prints, func() string {
			return strCurr
		})
	}
	fmt.Println(prints[0]()) //what does it print
}

func CorrectCreateDirs() {
	dirs := []string{"functions/temp1", "functions/temp2", "functions/temp3", "functions/temp4", "functions/temp5"}
	var rmdirs []func() error
	for _, dir := range dirs {
		d := dir
		os.Mkdir(d, 777)
		fmt.Println(&d)
		rmdirs = append(rmdirs, func() error {
			return os.Remove(d)
		})
	}
	fmt.Println("doing some work")
	for _, rmdir := range rmdirs {
		err := rmdir()
		if err != nil {
			fmt.Printf("Error: %v, \n", err)
		}
	}
}

func CreateDirsByIndex() {
	dirs := []string{"functions/temp1", "functions/temp2", "functions/temp3", "functions/temp4", "functions/temp5"}
	var rmdirs []func() error
	for i := 0; i < len(dirs); i++ {
		os.Mkdir(dirs[i], 777)
		fmt.Println(&dirs[i])
		rmdirs = append(rmdirs, func() error {
			//after the end of the loop - all functions will have a reference to dirs[5] - out of bound
			//because function refers to variable dirs[i] which is managed by the for loop
			return os.Remove(dirs[i])
		})
	}
	fmt.Println("doing some work")
	for _, rmdir := range rmdirs {
		err := rmdir()
		if err != nil {
			fmt.Printf("Error: %v, \n", err)
		}
	}
}

func CreateDirs() {
	dirs := []string{"functions/temp1", "functions/temp2", "functions/temp3", "functions/temp4", "functions/temp5"}
	var rmdirs []func() error
	for _, d := range dirs {
		os.Mkdir(d, 777)
		fmt.Println(&d)
		rmdirs = append(rmdirs, func() error {
			return os.Remove(d) //after the end of the loop - all functions will have a reference to last dir, not to 1, 2, 3, 4, 5
		})
	}
	fmt.Println("doing some work")
	for _, rmdir := range rmdirs {
		err := rmdir()
		if err != nil {
			fmt.Printf("Error: %v, \n", err)
		}
	}
}
