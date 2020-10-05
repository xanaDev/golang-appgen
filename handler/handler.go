package handler

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"

	"fmt"
	"go-initializer/consts"

	"go-initializer/types"
	"go-initializer/utils"
	"io"

	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"
	"os/exec"
	"io/ioutil"
	"github.com/gin-gonic/gin"

)

//GenerateTemplateRequest request payload for generate template
type GenerateTemplateRequest struct {
	AppName             string `form:"appname" json:"appname" xml:"appname"  binding:"required"`
	AppType             string `form:"apptype" json:"apptype" xml:"apptype"  binding:"required"`
	Library             string `form:"library" json:"library" xml:"library"  binding:"required"`
	DependencyManagment string `form:"dependencies" json:"dependencies" xml:"dependencies" `
	LoggingFramework    string `form:"loggingframework" json:"loggingframework"`
	OutputFormat        string `form:"outputformat" json:"outputformat"`
	requestTime         string
	outputFolder        string
	sourceFolder        string
	outputArchive       string
}

//GenerateTemplateResponse for future use
type GenerateTemplateResponse struct {
	path    string
	message string
}

// GetSupportedLibrariesRequest request payload for the GET /libs API call
type GetSupportedLibrariesRequest struct {
	AppType string `form:"apptype" json:"apptype" xml:"apptype" binding:"required"`
}

// GetSupportedLibraries get supported libraries by AppType
func GetSupportedLibraries(ctx *gin.Context)  {
	var request GetSupportedLibrariesRequest

	if	err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	if !utils.AppTypeExists(request.AppType) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid AppType!"})
		ctx.Abort()
		return
	}

	// Holds the collection of supported libraries for an AppType
	var supportedLibs []string

	// Get the path to the AppType base directory, ex. template/cli, template/webservice
	basePath := filepath.Join(utils.GetTemplateDir(), request.AppType)

	// Walk through the appType base directory and list all subfolders up to codebase
	err := filepath.Walk(basePath , func (path string, info os.FileInfo, err error) error {

		// If last part of the path is 'codebase' then add the relative path to the collection
		if info.IsDir() && strings.HasSuffix(path, "codebase") {
			// filepath.Dir removes 'codebase' from the path before trying to get the relative path
			relativeLibPath, _ := filepath.Rel(basePath, filepath.Dir(path))
			supportedLibs = append(supportedLibs, relativeLibPath)
		}

		return nil
	})

	if	err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"libraries": supportedLibs})
}

// Test : test function ...Must be removed 
func Test(ctx *gin.Context) {
	var request GenerateTemplateRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(request)

}

//Cleanup perfoming cleanup activities
func (request *GenerateTemplateRequest) Cleanup() error {

	//cleaning output folder

	err := os.RemoveAll(request.outputFolder)
	if err != nil {
		return fmt.Errorf("Error cleaning up output folder for %s", request.outputFolder)
	}
	// removing zip file

	err = os.RemoveAll(request.outputArchive)
	if err != nil {
		return fmt.Errorf("Error cleaning up Zip file for %s", request.outputArchive)
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

//Liveness sds
func Liveness(ctx *gin.Context) {

	ctx.JSON(200, gin.H{"message": "liveness", "active": "true"})
	ctx.Abort()
	return
}

// GenerateGitHubRepo: creating github repo 
func GenerateGitHubRepo(ctx *gin.Context) {
	var request GenerateTemplateRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(request)
	request.requestTime = fmt.Sprintf("%d", time.Now().Unix())
	_, err := generateOutput(&request)

	createRepo(&request)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		err = request.Cleanup()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("cleanup finished  ")
	}

}
//GenerateTemplate Create a zip file of a template code
func GenerateTemplate(ctx *gin.Context) {

	var request GenerateTemplateRequest

	if err := ctx.ShouldBind(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(request)
	//TODO: later include this
	// if err := request.Validate(); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// }

	request.requestTime = fmt.Sprintf("%d", time.Now().Unix())

	if request.OutputFormat == "tar" {
		ctx.Header("Content-Type", "application/octet-stream")
	} else {
		request.OutputFormat = "zip"
		ctx.Header("Content-Type", "application/zip")
	}

	_, err := generateOutput(&request)

	if request.OutputFormat == "tar" {
		err = createTar(&request)
	} else {
		err = createZip(&request)
	}
	
	if err != nil {
		fmt.Println(err)
	}

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")

		// zip is default output format . This is to support cli feature
	
		ctx.Header("Content-Disposition", "attachment; filename="+request.AppName+"."+request.OutputFormat)
		
		ctx.Header("File-name", request.AppName+"."+request.OutputFormat)
		ctx.File(request.outputArchive)

		err = request.Cleanup()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("cleanup finished  ")
	}

}

func generateOutput(request *GenerateTemplateRequest) (*GenerateTemplateResponse, error) {
	sourcePath, _ := utils.GetWorkingDir()
	request.outputFolder = filepath.Join(sourcePath, consts.OUTPUT_FOLDER, request.AppName+request.requestTime)
	request.outputArchive = filepath.Join(sourcePath, consts.OUTPUT_ZIP, request.AppName+"."+request.OutputFormat)

	if !utils.AppTypeExists(request.AppType) {

		return nil, fmt.Errorf("requested apptype does not exists")
	}
	request.Library = strings.ReplaceAll(request.Library, "/", string(os.PathSeparator))

	if !utils.LibExists(filepath.Join(request.AppType, request.Library)) {
		return nil, fmt.Errorf("request library does not exists")
	}

	request.sourceFolder = filepath.Join(sourcePath, "template", request.AppType, request.Library, "codebase")

	err := createOuputFolder(request)
	if err != nil {
		return nil, err
	}

	/*	cmd := exec.Command("bash", "-c", "gofmt -w "+request.AppName+request.requestTime)
		cmd.Dir = consts.OUTPUT_FOLDER
		fmt.Println("Running gofmt command and waiting for it to finish...")
		err = cmd.Run()
		if err != nil {
			fmt.Println("Command finished with error:", err)
		}
	*/
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

	config := getConfiguration(request)

	if request.LoggingFramework != "" {
		request.appendLogGoFile()
	}

	err = filepath.Walk(request.sourceFolder, func(filePath string, info os.FileInfo, err error) error {

		outputFileName := strings.TrimPrefix(filePath, request.sourceFolder)

		if outputFileName == "" {
			return nil
		}

		outputFileName = outputFileName[1:]
		if info.IsDir() {
			err := os.Mkdir(filepath.Join(request.outputFolder, outputFileName), 0777)
			if err != nil {
				return err
			}
		} else {
			t, err := template.ParseFiles(filePath)
			if err != nil {
				fmt.Println(err)
				return err
			}

			f, err := os.Create(filepath.Join(request.outputFolder, outputFileName))
			if err != nil {
				fmt.Println("create file: ", err)
				return err
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

func createRepo(request *GenerateTemplateRequest){
	repo_name := request.AppName
	token := os.Getenv("GITHUB_AUTH_TOKEN")	
	org_name := os.Getenv("GITHUB_ORG_NAME")
	
	url := "https://api.github.com/orgs/" + org_name + "/repos"

	payload := strings.NewReader("{\"name\":\""+repo_name+"\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("authorization", "token "+token)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "e1c5b0b0-a573-98dd-008b-72250586c558")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	//fmt.Printf("Successfully created new repo: %v\n", repo.GetName())

	cmd := exec.Command("git", "init", request.outputFolder)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Git init err :", err)
	}

	cmd = exec.Command("git", "add", "*")
	cmd.Dir = request.outputFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Git add err :", err)
	}

	cmd = exec.Command("git", "commit", "-m", "Initail Commit")
	cmd.Dir = request.outputFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Git commit err :", err)
	}

	cmd = exec.Command("git", "remote", "add", "origin", "https://github.com/" + org_name + "/" + request.AppName)
	cmd.Dir = request.outputFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Git remote add err :", err)
	}

	cmd = exec.Command("git", "push", "origin", "master")
	cmd.Dir = request.outputFolder
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("Git push err :", err)
	}

}
func createZip(request *GenerateTemplateRequest) error {

	zipfile, err := os.Create(request.outputArchive)
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
			header.Name += string(os.PathSeparator)
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

// tarrer walks paths to create tar file tarName
func createTar(request *GenerateTemplateRequest) (err error) {
	tarFile, err := os.Create(request.outputArchive)
	if err != nil {
		return err
	}
	defer func() {
		err = tarFile.Close()
	}()

	absTar, err := filepath.Abs(request.outputArchive)
	if err != nil {
		return err
	}

	// enable compression if file ends in .gz
	tw := tar.NewWriter(tarFile)
	if strings.HasSuffix(request.outputArchive, ".gz") || strings.HasSuffix(request.outputArchive, ".gzip") {
		gz := gzip.NewWriter(tarFile)
		defer gz.Close()
		tw = tar.NewWriter(gz)
	}
	defer tw.Close()

	var paths []string
	
	paths = append(paths, request.outputFolder)

	for _, path := range paths {
		// validate path
		path = filepath.Clean(path)
		absPath, err := filepath.Abs(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if absPath == absTar {
			fmt.Printf("tar file %s cannot be the source\n", request.outputArchive)
			continue
		}
		if absPath == filepath.Dir(absTar) {
			fmt.Printf("tar file %s cannot be in source %s\n", request.outputArchive, absPath)
			continue
		}

		walker := func(file string, finfo os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// fill in header info using func FileInfoHeader
			hdr, err := tar.FileInfoHeader(finfo, finfo.Name())
			if err != nil {
				return err
			}

			relFilePath := file
			if filepath.IsAbs(path) {
				relFilePath, err = filepath.Rel(path, file)
				if err != nil {
					return err
				}
			}
			// ensure header has relative file path
			hdr.Name = relFilePath

			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}
			// if path is a dir, dont continue
			if finfo.Mode().IsDir() {
				return nil
			}

			// add file to tar
			srcFile, err := os.Open(file)
			
			if err != nil {
				return err
			}
			defer srcFile.Close()
			_, err = io.Copy(tw, srcFile)
			if err != nil {
				return err
			}
			return nil
		}

		// build tar
		if err := filepath.Walk(path, walker); err != nil {
			fmt.Printf("failed to add %s to tar: %s\n", path, err)
		}
	}
	return nil
}

func getConfiguration(req *GenerateTemplateRequest) types.Configuration {
	var res types.Configuration
	res.AppName = req.AppName

	if req.LoggingFramework != "" {
		loggingframework, err := readLogJson(req.LoggingFramework)
		if err != nil {
			fmt.Println(err)
			res.Logging = &types.LoggingFramework{}
		} else {
			res.Logging = loggingframework
		}
	} else {
		res.Logging = &types.LoggingFramework{}
	}

	return res

}
