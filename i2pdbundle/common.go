package i2pd

import (
    "log"
    "net/http"
)

func FindAllDirectories(filesystem http.FileSystem) ([]string, error) {
    if filesystem.IsDir() {
        filelist, err := filesystem.Readdir(0)
        if err != nil {
            return nil, err
        }
        var rlist []string
        for index, file := range filelist {
            if file.IsDir(){
                rlist = append(rlist, file.Name())
                log.Println(index, file.Name())
            }
        }
    }
    return nil, nil
}

func FindAllFiles(filesystem http.FileSystem) ([]string, error) {
    if filesystem.IsDir() {
        filelist, err := filesystem.Readdir(0)
        if err != nil {
            return nil, err
        }
        var rlist []string
        for index, file := range filelist {
            if !file.IsDir(){
                rlist = append(rlist, file.Name())
                log.Println(index, file.Name())
            }
        }
    }
    return nil, nil
}