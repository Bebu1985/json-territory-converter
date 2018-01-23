package convert

import (
	"encoding/json"
	"io/ioutil"
)

//FileToObjects reads a json file and gives us back those objects
func FileToObjects(Path string, objects interface{}) error {
	JSON, readErr := ioutil.ReadFile(Path)
	if readErr != nil {
		return readErr
	}

	marsErr := json.Unmarshal(JSON, &objects)
	if marsErr != nil {
		return marsErr
	}
	return nil
}

//ObjectsToFile takes a slice of objects and writes them to the specified file
func ObjectsToFile(objects interface{}, Path string) error {
	JSON, marsErr := json.Marshal(objects)
	if marsErr != nil {
		return marsErr
	}

	writeErr := ioutil.WriteFile(Path, JSON, 0664)
	if writeErr != nil {
		return writeErr
	}

	return nil
}
