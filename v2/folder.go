package jotform

import (
	"encoding/json"
	"fmt"
	"os"
)

// GetFolder
// folderID (int64): You can get a list of folders from /user/folders.
// Returns a list of forms in a folder, and other details about the form such as folder color.
func (client jotformAPIClient) GetFolder(folderID string) ([]byte, error) {
	return client.executeHttpRequest("folder/"+folderID, "", "GET")
}

// AddFormsToFolder
// folderID (string): You can get the list of folders from /user/folders.
// formIDs ([]string): You can get the list of forms from /user/forms.
// Returns status of the request.
func (client jotformAPIClient) AddFormsToFolder(folderID string, formIDs []string) ([]byte, error) {
	formattedFormIDs, err := json.Marshal(map[string][]string{
		"forms": formIDs,
	})

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return client.executeHttpRequest("folder/"+folderID, formattedFormIDs, "PUT")
}

// CreateFolder
// folderProperties (map[string]string): Properties of new folder.
// Returns folder details.
func (client jotformAPIClient) CreateFolder(folderProperties map[string]string) ([]byte, error) {
	return client.executeHttpRequest("folder", folderProperties, "POST")
}

// DeleteFolder
// folderID (string): You can get the list of folders from /user/folders.
// Returns status of the request.
func (client jotformAPIClient) DeleteFolder(folderID string) ([]byte, error) {
	return client.executeHttpRequest("folder/"+folderID, nil, "DELETE")
}

// UpdateFolder
// folderID (string): You can get the list of folders from /user/folders.
// folderProperties ([]byte): Properties of folder.
// Returns status of the request.
func (client jotformAPIClient) UpdateFolder(folderID string, folderProperties []byte) ([]byte, error) {
	return client.executeHttpRequest("folder/"+folderID, folderProperties, "PUT")
}

// AddFormToFolder
// folderID (string): You can get a list of folders from /user/folders.
// formID (string): You can get the list of forms from /user/forms.
// Returns status of the request.
func (client jotformAPIClient) AddFormToFolder(folderID string, formID string) ([]byte, error) {
	formattedFormID, err := json.Marshal(map[string][]string{
		"forms": {formID},
	})

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	return client.executeHttpRequest("folder/"+folderID, formattedFormID, "PUT")
}
