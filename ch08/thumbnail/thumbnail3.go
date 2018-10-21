package main
import (
    "fmt"
    "thumbnail"
)

func makeThumbnails3(filenames []string) {
    ch := make(chan struct{})
    for _, f := range filenames {
        go func(f string) {
            thumbnail.ImageFile(f)
            ch <- struct(){}{}
        }(f)
    }
    for range filenames {
        <- ch
    }
}


