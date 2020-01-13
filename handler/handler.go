package handler

import (
	"archive/zip"

	"fmt"
	"go-initializer/consts"

	"go-initializer/utils"
	"io"
	
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

//GenerateTemplateRequest request payload for generate template
type GenerateTemplateRequest struct {
	AppName             string `form:"appname" json:"appname" xml:"appname"  binding:"required"`
	AppType             string `form:"apptype" json:"apptype" xml:"apptype"  binding:"required"`
	Library             string `form:"library" json:"library" xml:"library"  binding:"required"`
	DependencyManagment string `form:"dependencies" json:"dependencies" xml:"dependencies" `
	requestTime         string
	outputFolder        string
	sourceFolder	string
	outputZip	string
}

//GenerateTemplateResponse for future use
type GenerateTemplateResponse struct {
	path    string
	message string
}


//Cleanup perfoming cleanup activities
func (request *GenerateTemplateRequest)Cleanup() error{

	//cleaning output folder

	err := os.RemoveAll(request.outputFolder)
		if err != nil{
			return fmt.Errorf("Error cleaning up output folder for %s",request.outputFolder)
		}
	// removing zip file

	err = os.RemoveAll(request.outputZip)
		if err != nil{
			return fmt.Errorf("Error cleaning up Zip file for %s",request.outputZip)
		}
		return nil

}



//Validate request payload TODO must improve this
func (request *GenerateTemplateRequest) Validate() error {
	if request.AppName == "" {
		return fmt.Errorf("appname cannot be empty")
	}
	if !utils.HasElem(consts.SupportedAppType, request.AppType) {
		return fmt.Errorf("apptype %s is not supported", request.AppType)
	}
	switch request.AppType {
	case "cli":
		if !utils.HasElem(consts.SupportedCliLib, request.Library) {
			return fmt.Errorf("library %s is not supported ", request.Library)
		}
	}
	return nil
}

//GenerateTemplate Create a zip file of a template code
func GenerateTemplate(ctx *gin.Context) {

	var request GenerateTemplateRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//TODO: later include this 
	// if err := request.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	request.requestTime = fmt.Sprintf("%d", time.Now().Unix())

	_, err := generateOutput(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		
		//TODO: .zip should be made optional
		ctx.Header("Content-Disposition", "attachment; filename="+request.AppName+".zip")
		ctx.Header("Content-Type", "application/zip")
		ctx.File(request.outputZip)
		
		err  = request.Cleanup()
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println("cleanup finished  ")
	}

}

func generateOutput(request *GenerateTemplateRequest) (*GenerateTemplateResponse, error) {
	sourcePath, _ := utils.GetWorkingDir()
	request.outputFolder = sourcePath + "/" + consts.OUTPUT_FOLDER + "/" + request.AppName + "-" + request.requestTime
	request.outputZip = sourcePath +"/"+consts.OUTPUT_ZIP+"/"+request.AppName + ".zip"
	
	if !utils.AppTypeExists(request.AppType){
		return nil,fmt.Errorf("requested apptype does not exists")
	}

	if !utils.LibExists(request.AppType+"/"+request.Library){
		return nil,fmt.Errorf("request library does not exists")
	}

	request.sourceFolder = sourcePath + "/template/"+request.AppType +"/" + request.Library +"/codebase"

	err := createOuputFolder(request)
	if err != nil {
		return nil, err
	}

	err = createZip(request)

	if err != nil {
		return nil, err
	}
	response := &GenerateTemplateResponse{
		path:    request.AppName,
		message: "Thanks for downloading",
	}
	return response, nil
}

func createOuputFolder(request *GenerateTemplateRequest) error {

	err := os.Mkdir(request.outputFolder, 0777)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = filepath.Walk(request.sourceFolder, func(filePath string, info os.FileInfo, err error) error {

		outputFileName := strings.TrimPrefix(filePath, request.sourceFolder)
		
		if outputFileName == ""{
			return nil
		}
		if info.IsDir() {
			err := os.Mkdir(request.outputFolder+outputFileName, 0777)
			if err != nil {
				return err
			}
		} else {
			t, err := template.ParseFiles(filePath)
			if err != nil {
				fmt.Println(err)
				return err
			}

			f, err := os.Create(request.outputFolder + outputFileName)
			if err != nil {
				fmt.Println("create file: ", err)
				return err
			}
			config := map[string]string{
				"appname": request.AppName,
			}
			err = t.Execute(f, config)
			if err != nil {
				fmt.Println("execute: ", err)
				return err
			}
			f.Close()

		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil

}

func createZip(request *GenerateTemplateRequest) error {

	zipfile, err := os.Create(request.outputZip)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	err = filepath.Walk(request.outputFolder, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}
		if request.outputFolder == path {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, request.outputFolder)[1:]

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)

		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)

		file.Close()
		return nil

	})
	if err != nil {
		return err
	}
	return nil
}
