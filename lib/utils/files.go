// utils pertaining to creating/reading files

package utils

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// write an obj to a gob file
func WriteGob[DataT any](filename string,data *DataT) error {
    var wfile *os.File
    var e error
    wfile,e=os.Create(filename)

    if e!=nil {
        return e
    }
    defer wfile.Close()

    var encoder *gob.Encoder=gob.NewEncoder(wfile)
    e=encoder.Encode(data)

    if e!=nil {
        return e
    }

    return nil
}

// read gob file and return the data
func ReadGob[DataT any](filename string) (DataT,error) {
    var rfile *os.File
    var e error
    rfile,e=os.Open(filename)

    var data DataT
    if e!=nil {
        return data,e
    }
    defer rfile.Close()

    var decoder *gob.Decoder=gob.NewDecoder(rfile)
    e=decoder.Decode(&data)

    if e!=nil {
        return data,e
    }

    return data,nil
}

// read a yaml file and return result
func ReadYaml[DataT any](filename string) (DataT,error) {
	var data []byte
	var e error
	data,e=os.ReadFile(filename)

	if errors.Is(e,fs.ErrNotExist) {
		log.Info().Msgf("file not found: %s",filename)
		var def DataT
		return def,e
	}

	if e!=nil {
		var def DataT
		return def,e
	}

	var parsedData DataT
	yaml.Unmarshal(data,&parsedData)

	return parsedData,nil
}

// overwrite target yml file with a new file
func WriteYaml(filename string,data any) error {
	var wfile *os.File
	var e error
	wfile,e=os.Create(filename)

	if e!=nil {
		panic(e)
	}

	defer wfile.Close()

	var ymldata []byte
	ymldata,e=yaml.Marshal(data)

	if e!=nil {
		panic(e)
	}

	wfile.Write(ymldata)
	return nil
}