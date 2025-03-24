package verify

import (
	"os"
)

func (json *ToJson) Write(content []byte) {
	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		panic(err)
	}
}
