package i2pd

import (
    "log"
    "os"
)

type fsi interface{
    IsDir() bool
    ReadDir(int) ([]os.FileInfo, error)
}

func FindAllDirectories(filesystem fsi) ([]string, error) {
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

func FindAllFiles(filesystem *fsi) ([]string, error) {
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