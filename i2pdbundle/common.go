package i2pd

import (
    "log"
)

func FindAllSubdirectories(filesystem *fs) ([]string, error) {
    if filesystem.IsDir() {
        filelist, err := filesystem.Readdir(0)
        if err != nil {
            return nil, err
        }
        var rlist []string
        for index, file := range filelist {
            rlist = append(rlist, file.Name())
            log.Println(index, file.Name())
        }
    }
    return nil, nil
}