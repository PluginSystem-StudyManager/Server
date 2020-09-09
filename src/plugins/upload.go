package plugins

//go:generate schema-generate -o upload.schema.go -p plugins ../../schemas/plugins/upload.schema.json

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
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

	// TODO: handle wrong formatted upload: file name, too big, ...
	// max 1000 MB
	_ = request.ParseMultipartForm(1000 << 20)
	form := request.MultipartForm
	if form == nil {
		respond(writer, http.StatusBadRequest, "No Multipart form uploaded")
		return
	}

	// Check Form parameters
	// This can't be refactored into a function, because 'File' and 'Value' are different types

	// File: file, token
	neededParametersFile := []string{"file"}
	for _, param := range neededParametersFile {
		if _, ok := form.File[param]; !ok {
			respond(writer, http.StatusBadRequest, "Multipart form without field `"+param+"` uploaded!")
			return
		}
	}
	// Value: plugin_data
	neededParametersValue := []string{"token", "plugin_data"}
	for _, param := range neededParametersValue {
		if _, ok := form.Value[param]; !ok {
			respond(writer, http.StatusBadRequest, "Multipart form without field `"+param+"` uploaded!")
			return
		}
	}

	// Get zip file content
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

	// Get token
	if len(form.Value["token"]) != 1 {
		respond(writer, http.StatusBadRequest, "Multipart form with not exactly one token")
		return
	}
	token := request.MultipartForm.Value["token"][0]

	// Get plugin data
	if len(form.Value["plugin_data"]) != 1 {
		respond(writer, http.StatusBadRequest, "Multipart form with not exactly one plugin_data")
		return
	}
	pluginDataString := request.MultipartForm.Value["plugin_data"][0]
	log.Printf("PluginData: %s", pluginDataString)
	var pluginData PluginDataUpload
	err = json.Unmarshal([]byte(pluginDataString), &pluginData)
	if err != nil {
		respond(writer, http.StatusBadRequest, "plugin_data is not properly formatted so unmarshal fails")
		return
	}

	err = uploadImpl(token, pluginData, content)
	if err != nil {
		respond(writer, http.StatusBadRequest, err.Error())
		return
	}
	respond(writer, http.StatusOK, "Successfully uploaded plugin")
}

func uploadImpl(token string, pluginData PluginDataUpload, fileContent []byte) error {
	// Authenticate
	userId, err := db.UserIdByPermanentToken(token)
	if err != nil {
		log.Printf("Not authenticated: %v\n", err)
		return err
	}
	// save and update
	zipPath := filepath.Join(pluginsTmpPath, fmt.Sprintf("%s.zip", pluginData.Id))
	_ = ioutil.WriteFile(zipPath, fileContent, os.ModePerm)
	pluginPath := filepath.Join(pluginsPath, pluginData.Id)
	err = unzip(zipPath, pluginPath)
	_ = os.Remove(zipPath)
	if err != nil {
		log.Printf("Error unzipping path (%v): %v -->%v\n", err, zipPath, pluginPath)
		return errors.New(fmt.Sprintf("zip: error unzipping: %s", err))
	}
	failedUpload := func(err error, msg string) error {
		log.Printf("%s (%v)", msg, err)
		_ = os.RemoveAll(pluginPath)
		return err
	}

	// TODO: validate correct uploaded files

	_, err = db.PluginIdByName(pluginData.Id)
	if err != nil {
		// Add
		err = db.AddPlugin(db.PluginData{
			// TODO: Add Plugin ID or use PluginDataUpload + userIDs
			Name:             pluginData.Name,
			ShortDescription: pluginData.ShortDescription,
			Tags:             []string{},
			UserIds:          []int{userId},
		})
		if err != nil {
			return failedUpload(err, "Error adding plugin to db: ")
		}
	} else {
		// Update
		// TODO: update shortDescription, tags, ...
	}

	log.Printf("Successfully uploaded plugin: %s", pluginData.Id)
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
