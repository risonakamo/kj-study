// utils pertaining to creating/reading files

package utils

import (
	"encoding/gob"
	"os"
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