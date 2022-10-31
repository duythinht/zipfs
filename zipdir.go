package zipfs

import (
	"io/fs"
	"time"
)

// dirEntry implement fs.DirEntry, which present of zip file members list
type dirEntry struct {
	fs.FileInfo
	path string
}

// IsDir check a entry is directory or not
func (d dirEntry) IsDir() bool {
	return d.FileInfo.IsDir()
}

// Type return type of FileMode
func (d dirEntry) Type() fs.FileMode {
	return d.FileInfo.Mode().Type()
}

// Info return entry FileInfo
func (d dirEntry) Info() (fs.FileInfo, error) {
	return d.FileInfo, nil
}

// Name return entry path
func (d dirEntry) Name() string {
	return d.path
}

// zipInfo presentation of virtual zip root dir directory, which implement fs.FileInfo
type zipInfo struct {
	name string
}

// Name return zip name
func (d *zipInfo) Name() string {
	return d.name
}

// IsDir alway return true for zip file
func (d *zipInfo) IsDir() bool {
	return true
}

// Type return FileMode type
func (d *zipInfo) Type() fs.FileMode {
	return fs.ModeDir.Perm().Type()
}

// ModTime sumulation modify time as now
func (d *zipInfo) ModTime() time.Time {
	return time.Now()
}

// Mode return fs mod
func (d *zipInfo) Mode() fs.FileMode {
	return 0755
}

// Sys return nil
func (d *zipInfo) Sys() interface{} {
	return nil
}

// Size of zip file dir
func (d *zipInfo) Size() int64 {
	return 0
}
