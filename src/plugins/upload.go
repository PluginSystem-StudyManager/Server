package plugins

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"server/db"
)

const (
	pluginsPath    = "../plugins/all"
	pluginsTmpPath = "../plugins/tmp"
)

// upload handles the upload request.
// upload parses all data from the request.
// The final logic is implemented in ``upload_impl()``
func upload(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	// Parse arguments
	token, errToken := request.Cookie("token")
	pluginId, errId := request.Cookie("pluginName")
	if errToken != nil || errId != nil {
		respond(writer, http.StatusBadRequest, "No cookie with authentication data specified")
		return
	}

	// TODO: handle wrong formatted upload: file name, too big, ...
	// max 1000 MB
	_ = request.ParseMultipartForm(1000 << 20)
	form := request.MultipartForm
	if form == nil {
		respond(writer, http.StatusBadRequest, "No Multipart form uploaded")
		return
	}
	if _, ok := form.File["file"]; !ok {
		respond(writer, http.StatusBadRequest, "Multipart form without field `file` uploaded")
		return
	}
	if len(form.File["file"]) != 1 {
		respond(writer, http.StatusBadRequest, "Multipart form with not exactly one file")
		return
	}

	file := request.MultipartForm.File["file"][0]
	fileHandle, err := file.Open()
	if err != nil {
		respond(writer, http.StatusBadRequest, fmt.Sprintf("Error opening zip file: %s", err))
		return
	}
	content := make([]byte, file.Size)
	_, _ = fileHandle.Read(content)
	_ = fileHandle.Close()
	err = uploadImpl(token.Value, pluginId.Value, content)
	if err != nil {
		respond(writer, http.StatusBadRequest, err.Error())
		return
	}
	respond(writer, http.StatusOK, "Successfully uploaded plugin")
}

func uploadImpl(token string, pluginId string, fileContent []byte) error {
	log.Printf("token: %s, pluginId: %s", token, pluginId)
	// Authenticate
	userId, err := db.UserIdByToken(token)
	if err != nil {
		log.Printf("Not authenticated: %v\n", err)
		return err
	}
	// save and update
	zipPath := filepath.Join(pluginsTmpPath, fmt.Sprintf("%s.zip", pluginId))
	_ = ioutil.WriteFile(zipPath, fileContent, os.ModePerm)
	pluginPath := filepath.Join(pluginsPath, pluginId)
	err = unzip(zipPath, pluginPath)
	_ = os.Remove(zipPath)
	if err != nil {
		log.Printf("Error unzipping path (%v): %v -->%v\n", err, zipPath, pluginPath)
		return errors.New(fmt.Sprintf("zip: error unzipping: %s", err))
	}
	// TODO: validate correct uploaded files
	// read plugin_info.yaml and add entry in db
	infoFile := filepath.Join(pluginPath, "plugin_info.yaml")
	data, err := ioutil.ReadFile(infoFile)
	if err != nil {
		log.Printf("Error reading file: \n" + infoFile)
		return err
	}
	pluginInfo := PluginInfo{}
	err = yaml.Unmarshal(data, &pluginInfo)
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = db.PluginIdByName(pluginInfo.Name)
	if err != nil {
		// Add
		err = db.AddPlugin(db.PluginData{
			Name:             pluginId, // TODO: use Name from File or better ensure both are the sane
			ShortDescription: pluginInfo.ShortDescription,
			Tags:             []string{},
			UserIds:          []int{userId},
		})
		if err != nil {
			// TODO: Delete upload?
			log.Printf("Error adding plugin to db: %v\n", err)
			return err
		}
	} else {
		// Update
		// TODO: update shortDescription, tags, ...
	}

	log.Printf("Successfully uploaded plugin: %s", pluginId)
	return nil
}

func respond(writer http.ResponseWriter, statusCode int, body string) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(statusCode)
	_, _ = writer.Write([]byte(body))
}

func unzip(src string, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()
	_ = os.MkdirAll(dest, os.ModeDir)

	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(path, f.Mode())
		} else {
			_ = os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}
	return nil
}
