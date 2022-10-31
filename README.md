# zipfs

Open and browsing zip file as golang io/fs 

## Quick start


```go

f, err := os.OpenFile("path/to/file.zip")

if err != nil {
    // ....
}

zipReader, err := zip.NewReader(f, sizeOfZipFile)

if err != nil {
    // ....
}

fsys := zipfs.NewFS(zipReader)

fs.WalkDir(fsys, "", func(path string, d fs.DirEntry, err error) error {
    // do something with file
    return nil
})

```