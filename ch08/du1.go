package main

import(
    "flag"
    "fmt"
    "io/ioutil"
    "path/filepath"
    "os"
)

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes
func walkDir(dir string, fileSizes chan<- int64) {
    for _, entry := range dirents(dir) {
        // For each subdirectory, walkDir recursively calls itself,
        // and for each file, walkDir sends a message on the fileSizes channel.
        if entry.IsDir() {
            subdir := filepath.Join(dir, entry.Name())
            walkDir(subdir, fileSizes)
        } else {
            fileSizes <- entry.Size()
        }
    }
}

// Dirents return the entries of directory dir
func dirents(dir string)[]os.FileInfo {
    // ioutil.ReadDir function returns a slice of os.FileInfo
    entries, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Fprintf(os.Stderr, "du1:%v\n", err)
        return nil
    }
    return entries
}

func printDiskUsage(nfiles, nbytes int64) {
    fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func main() {
    // Determine the initial directories
    flag.Parse()
    roots := flag.Args()
    if len(roots) == 0 {
        roots = []string{"."}
    }

    fileSizes := make(chan int64)
    // Traverse the file tree
    go func() {
        for _, root := range roots {
            walkDir(root, fileSizes)
        }
        close(fileSizes)
    }()

    // Print the results
    var nfiles, nbytes int64
    for size := range fileSizes {
        nfiles++
        nbytes += size
    }
    printDiskUsage(nfiles, nbytes)
}

