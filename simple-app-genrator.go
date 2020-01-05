package main

import (
	"fmt"
	"text/template"
	"os"
//	"io/ioutil"
	"path/filepath"
	"strings"
)
// FilePathWalkDir - FilePathWalkDir
func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
       if path != root {
		   files = append(files, path)			
        }
        return nil
    })
    return files, err
}

// IsDir - IsDir
func IsDir(file string) bool {
	
	fi, err := os.Stat(file)
    if err != nil {
        fmt.Println(err)
        return false
    }
    switch mode := fi.Mode(); {
    case mode.IsDir():
        // do directory stuff
        return true
    case mode.IsRegular():
        // do file stuff
        return false
	}
	return false
}

// CreateFiles - CreateFiles
func CreateFiles(appname string, apptype string) {
	path, err := os.Getwd()
	templateFolder := "template\\"+apptype
	outputFolder := path + "\\" + "output\\"+apptype+"\\"
	files, err := FilePathWalkDir(templateFolder)
    if err != nil {
        fmt.Println(err)
    }
	
	for _, file := range files {
		
		if IsDir(file) {
			if strings.Contains(file,apptype){
				outputFolder := strings.Replace(file,"template","output",-1)
				fmt.Println(strings.Replace(outputFolder,"appname",appname,-1))
				err := os.MkdirAll(path+"\\"+strings.Replace(outputFolder,"appname",appname,-1),0777)
				if err != nil {
					fmt.Println("create file: ", err)
					return
				}
			}
		} else {

			t, err := template.ParseFiles(file)
			if err != nil {
				fmt.Println(err)
				return
			}
			outputFile1 := strings.Replace(file,"template","output",-1)
			outputFile := strings.Replace(outputFile1,"appname",appname,-1)
			f, err := os.Create(path+"\\"+ outputFile)
			if err != nil {
				fmt.Println("create file: ", err)
				return
			}

			config := map[string]string{
				"appname":      appname,				
			}

			err = t.Execute(f, config)
			if err != nil {
				fmt.Println("execute: ", err)
				return
			}
			f.Close()
		}		
	}
	ZipWriter(appname,outputFolder,apptype)
}
