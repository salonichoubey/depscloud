// Package swagger Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// v1alpha/tracker/tracker.swagger.json
// v1alpha/extractor/extractor.swagger.json
// v1alpha/schema/schema.swagger.json
// v1alpha/deps/deps.swagger.json
// v1alpha/store/store.swagger.json
package swagger

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}


type assetFile struct {
	*bytes.Reader
	name            string
	childInfos      []os.FileInfo
	childInfoOffset int
}

type assetOperator struct{}

// Open implement http.FileSystem interface
func (f *assetOperator) Open(name string) (http.File, error) {
	var err error
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}
	content, err := Asset(name)
	if err == nil {
		return &assetFile{name: name, Reader: bytes.NewReader(content)}, nil
	}
	children, err := AssetDir(name)
	if err == nil {
		childInfos := make([]os.FileInfo, 0, len(children))
		for _, child := range children {
			childPath := filepath.Join(name, child)
			info, errInfo := AssetInfo(filepath.Join(name, child))
			if errInfo == nil {
				childInfos = append(childInfos, info)
			} else {
				childInfos = append(childInfos, newDirFileInfo(childPath))
			}
		}
		return &assetFile{name: name, childInfos: childInfos}, nil
	} else {
		// If the error is not found, return an error that will
		// result in a 404 error. Otherwise the server returns
		// a 500 error for files not found.
		if strings.Contains(err.Error(), "not found") {
			return nil, os.ErrNotExist
		}
		return nil, err
	}
}

// Close no need do anything
func (f *assetFile) Close() error {
	return nil
}

// Readdir read dir's children file info
func (f *assetFile) Readdir(count int) ([]os.FileInfo, error) {
	if len(f.childInfos) == 0 {
		return nil, os.ErrNotExist
	}
	if count <= 0 {
		return f.childInfos, nil
	}
	if f.childInfoOffset+count > len(f.childInfos) {
		count = len(f.childInfos) - f.childInfoOffset
	}
	offset := f.childInfoOffset
	f.childInfoOffset += count
	return f.childInfos[offset : offset+count], nil
}

// Stat read file info from asset item
func (f *assetFile) Stat() (os.FileInfo, error) {
	if len(f.childInfos) != 0 {
		return newDirFileInfo(f.name), nil
	}
	return AssetInfo(f.name)
}

// newDirFileInfo return default dir file info
func newDirFileInfo(name string) os.FileInfo {
	return &bindataFileInfo{
		name:    name,
		size:    0,
		mode:    os.FileMode(2147484068), // equal os.FileMode(0644)|os.ModeDir
		modTime: time.Time{}}
}

// AssetFile return a http.FileSystem instance that data backend by asset
func AssetFile() http.FileSystem {
	return &assetOperator{}
}

var _v1alphaTrackerTrackerSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x5b\x4b\x6f\xe3\x36\x10\xbe\xfb\x57\x10\x6a\x0f\x2d\xb0\x1b\xa7\xe9\x2d\xb7\x45\x5f\x28\xd0\xbd\x34\xb9\x15\x41\x40\x4b\x63\x99\x5b\x89\x64\xc8\x51\x5a\x77\xe1\xff\x5e\x88\x94\xac\xb7\xac\xa7\x9d\x87\x7c\x59\xad\x45\x8e\xe6\xf1\x7d\x33\xc3\xb1\xf2\x75\x45\x88\xa3\xff\xa1\xbe\x0f\xca\xb9\x25\xce\xcd\xd5\xb5\xf3\x21\xfe\x8e\xf1\xad\x70\x6e\x49\x7c\x9f\x10\x07\x19\x06\x10\xdf\x7f\xfe\x81\x06\x72\x47\xd7\xa8\xa8\xfb\x37\xa8\xf4\xdf\x2b\xa9\x04\x0a\xb3\x93\x10\xe7\x19\x94\x66\x82\x9b\xf5\xf6\x92\x70\x81\x44\x03\x3a\x2b\x42\x0e\x46\xbe\x76\x77\x10\x82\x76\x6e\xc9\x5f\x76\xd3\x0e\x51\xa6\x02\xe2\x6b\x1d\xaf\x7d\x30\x6b\x5d\xc1\x75\x54\x58\x4c\xa5\x0c\x98\x4b\x91\x09\xbe\xfe\xa2\x05\xcf\xd6\x4a\x25\xbc\xc8\xed\xb8\x96\xe2\x4e\x67\x46\xae\x53\xe3\x7c\x45\xe5\x6e\xfd\x35\xa0\xdc\x8f\xa8\x0f\x87\xb5\x07\x12\xb8\x07\xdc\x65\x90\xad\x27\xc4\xf1\x01\x73\xff\x25\xc4\x11\x12\x94\x79\xd2\xef\x5e\x6c\xfd\x6f\x80\x3f\xe7\xb7\x7e\xc8\x56\x2a\xd0\x52\x70\x5d\x90\x67\x6e\xdc\x5c\x5f\x97\xbe\x22\xc4\xf1\x40\xbb\x8a\x49\x4c\xbc\xfa\x89\xe8\xc8\x75\x41\xeb\x6d\x14\x90\x54\xd2\xd5\x77\x1a\x15\xd0\x90\x71\xff\xf8\x9d\xfe\x3e\xf7\x4c\x23\xc9\xf8\x9d\x56\x9e\x40\x88\xf3\xad\x82\x6d\x2c\xfc\x9b\xf5\xbf\x1f\xad\xa4\x8f\x1e\x6c\x19\x67\xf1\x53\x75\xea\x9c\x24\xe0\x47\xb3\xf6\x4e\x41\xce\x61\x55\x77\x7d\xc8\x19\x2e\xa9\xa2\x21\x20\xa8\x2c\x44\xf6\x53\x32\x99\xd3\xd0\x20\x2e\x8d\x42\xd9\x10\x66\x3c\x11\x87\xb0\x7c\x47\xc1\x53\xc4\x14\xc4\x11\x40\x15\x41\xe9\x2e\xee\xa5\x91\xab\x51\x31\xee\xe7\xb5\x3f\x7c\x38\xad\x8d\x50\x3e\xe5\xec\x3f\x13\xe3\x7a\x8d\x9e\x22\x50\xfb\x16\x95\xb6\x34\xd0\xd3\xea\x14\x0a\x2f\x0a\x1a\xfc\x33\xb1\x36\xc7\xeb\x87\x5c\x44\x91\xfa\xe5\x58\x3a\x19\x3e\xee\x40\x3d\x33\x17\x32\x31\x0f\xab\xbc\xb0\xc4\xc0\x8e\xe4\x5b\xa3\x90\x22\x10\xfe\x7e\x04\x0b\xef\x53\x11\x0b\x1b\x17\x36\xbe\x0f\x36\xa6\x90\x9f\x87\x8b\x6b\x64\x60\xf5\x1f\x4b\xc9\x7b\x2b\xe8\x35\x13\x33\x61\x64\xde\xa2\x85\x93\x45\x8d\x16\x4e\xce\xc2\x49\x1c\xda\x9a\xe2\xd2\x98\x2e\xb4\x7b\x2f\xb4\x9b\xa1\x31\xc5\xd1\x6d\x29\x2e\x4d\xe9\xc2\xc4\xf7\xc6\xc4\x19\x0a\xe0\x44\x2d\x29\x2e\x0d\xe9\xc2\xc7\x85\x8f\xc7\x8d\xed\x7c\xb4\xb6\xf4\x6a\x3e\xff\x60\x1a\x67\x26\xd5\x10\xfe\xd4\xd0\x26\xd6\xf4\xb3\x31\xf0\xcf\x44\xf2\xdc\xe4\x91\x8d\xc4\x19\x03\x0c\xc6\x11\x7c\x50\xe5\xdd\x5b\xa1\x42\x8a\xc9\x82\x1f\x6f\xfa\xc2\xd8\x15\x11\xc7\x17\xa1\x6c\x67\x94\xdb\x60\x0e\xc4\xf8\x3a\xa4\x9c\xfa\xbd\xeb\xca\xe7\x64\xd7\x1b\x28\x23\x89\x29\xd6\x8b\x73\x53\x21\x52\xc1\xcb\x4a\x91\xe3\xc0\xa3\x45\xa4\x5c\xe8\x89\x9d\x3b\xbb\xe9\xed\x40\x27\x31\xe8\xa2\x2d\xc8\xf9\x0b\xfe\xd2\x84\xd8\xcf\x3c\x0c\xb3\xcc\x7a\xc3\x2d\x88\x25\xcd\xd2\x82\x54\x75\x7d\x85\x2d\x88\x0d\xe6\x40\x8c\x5b\x54\xe4\x91\x2e\x85\x6e\x87\xfa\xbd\xd9\xf1\x2a\xb0\x6e\x54\x3d\x17\xcc\x37\xc2\xab\xc0\xc3\x22\xa7\xee\x4e\xfb\x11\x75\xa0\xbd\x29\xaf\x9f\x22\xd0\xd8\xc5\xde\x69\xb0\x75\x7c\x0b\x27\xa7\x52\xf6\x1e\x8c\x07\x52\x67\xc3\x3b\x5b\xb6\x43\xe0\xf8\x2b\x0b\x0a\xdd\x4b\xca\x15\xb1\xf9\x02\x6e\xc6\x41\x47\xaa\x18\x80\xc8\x4a\xf8\xca\x6a\x72\x09\x75\x4d\x75\x23\x1f\x5b\xbd\xd7\x08\xe1\x90\x9d\x85\xca\x3b\x60\x7f\x52\x25\x07\xec\xcc\x5e\x83\xea\xbd\xb5\xe1\x85\xa3\xc2\x7e\xaa\x14\x2d\xa2\xd4\x61\x08\x61\x79\x7d\x23\x04\x93\xc4\x52\x8c\x76\x7d\x46\x3b\xd4\xa6\x26\xf3\xce\xd7\x26\xda\x7e\xe2\xfb\x31\xb0\x88\xd7\x3f\xc6\xad\xfe\x10\x17\xd3\x20\x3a\x15\x9b\x82\x87\xb2\x04\xbe\xd9\x63\x8e\x20\xf5\x16\xaa\x88\x23\x0b\xe1\xce\xf4\xb4\xbf\x28\x25\xd4\x18\x43\x7d\x25\xdd\x47\x57\x78\x8d\x0a\xd7\x15\x9d\xe6\x92\x93\xf7\xc3\x0e\x51\xce\x24\x3a\x04\xad\x07\x92\xd6\x68\xa5\x91\x62\xd4\x08\xe2\x76\x12\x20\x65\xc1\x6c\xf8\xcf\xc3\xb7\x07\xea\x6d\xa6\xb7\x8c\xd1\x97\x4a\x87\x49\x6a\x79\x74\x05\xd7\xa8\x28\xe3\x38\x28\xa9\xba\x42\x4e\x91\x5f\xba\x34\xfe\x6d\xce\xb4\x25\xe6\x62\xce\x1c\x5e\x5b\xfa\x66\xf8\x36\x1f\x94\xcb\xcc\x59\x5d\x70\xe6\x22\xd9\xe2\x87\xbb\xca\x90\xa4\xaf\x1f\xfa\x94\x92\x5a\x45\x9a\x47\xbf\x23\xb4\x92\x2d\x91\x19\x9e\x9c\xed\xe9\x67\xfa\x9c\x5f\x19\xe9\x93\x29\x33\x6f\x01\xf1\x3d\xb2\x45\xf3\x81\xf8\xbd\x04\xa6\x3a\xe8\x98\x21\x30\xd5\x09\x5d\xc7\xc0\x14\x87\xc3\x23\x62\x12\x56\x0a\xc2\x49\xa5\xd3\x1a\xd2\x23\x3d\xf5\x01\x67\x17\xbb\xc7\xa7\xae\xca\x84\xb8\x77\xac\x0a\x56\x4f\xe1\xc5\x36\xbb\x8b\xa7\xd7\x17\x66\x77\x7a\x62\x9d\x8d\x2c\xad\xc7\xe4\xfe\xe4\x29\xfc\x40\x3f\xe6\x38\x55\xdc\x3f\xa9\xc9\x5d\x5e\xf0\xe9\x6a\x6e\x61\xd2\x33\xc6\xde\x58\x50\x5c\xd5\x1b\x6c\xde\x08\x11\x00\xe5\x8d\xc7\xc0\xe4\xf6\x09\xad\xeb\x0f\xcb\x23\xd4\xbe\xf8\x58\xe2\xa7\x57\x7e\x74\x68\xc4\xe2\x88\xa0\x78\x95\x53\x1d\x39\x95\x88\xd2\x83\xe0\x79\xea\xce\x71\x70\x57\xf7\x43\x5b\x36\xc1\x9b\xba\x1e\x2b\xd0\x51\x50\xc1\x49\xeb\x64\xb3\xe1\x87\xe2\xbc\x77\xa0\x34\x50\x69\x11\x5a\x33\x88\xa9\xba\xe8\x68\xc6\xf1\xaf\x15\xed\x7a\x62\xd5\x27\x62\x4b\x9a\x95\x9b\xab\xa4\x0f\x77\x5d\x5b\x89\xbb\xb8\xeb\xf2\xca\xcd\x54\xc8\x06\x78\xae\xfe\x25\xb7\x97\xe1\xb8\xaa\x6e\x73\x66\xb2\x3e\xce\x3b\x5d\xd5\x2f\xec\xc1\x76\x05\x4d\x52\x5c\x1d\x56\xff\x07\x00\x00\xff\xff\xb0\xc6\x9f\xe0\xcb\x3c\x00\x00")

func v1alphaTrackerTrackerSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaTrackerTrackerSwaggerJson,
		"v1alpha/tracker/tracker.swagger.json",
	)
}

func v1alphaTrackerTrackerSwaggerJson() (*asset, error) {
	bytes, err := v1alphaTrackerTrackerSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/tracker/tracker.swagger.json", size: 15563, mode: os.FileMode(420), modTime: time.Unix(1571198791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaExtractorExtractorSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xe4\x56\x4d\x6f\xdb\x30\x0c\xbd\xfb\x57\x08\xda\x8e\x45\xd3\xf5\x98\xdb\xd0\x75\xc0\x0e\x05\x8a\x5d\x87\x1c\x58\x89\xb1\x55\xd8\x92\x2a\xd2\xdd\xb2\x21\xff\x7d\x90\x62\xd7\x1f\x8d\x8b\xc4\x19\x16\x0c\xcb\x49\xb1\xf8\x28\xbe\xa7\x67\xd2\xbf\x32\x21\x24\x7d\x87\x3c\xc7\x20\x97\x42\x5e\x5f\x5e\xc9\x8b\xf8\xcc\xd8\xb5\x93\x4b\x11\xf7\x85\x90\x6c\xb8\xc4\xb8\xff\xfc\x01\x4a\x5f\xc0\x02\x7f\x70\x00\xc5\x2e\x74\xab\x4b\x1f\x1c\xbb\x84\x16\x42\x3e\x63\x20\xe3\x6c\xc2\xec\x96\xc2\x3a\x16\x84\x2c\x33\x21\xb6\xe9\x0c\x52\x05\x56\x48\x72\x29\xbe\xed\x40\x05\xb3\x6f\x13\xc4\x35\xc5\xd8\x55\x8a\x55\xce\x52\x3d\x08\x06\xef\x4b\xa3\x80\x8d\xb3\x8b\x47\x72\xb6\x8b\xf5\xc1\xe9\x5a\x1d\x18\x0b\x5c\x50\x47\x74\xd1\x12\xd4\xe8\xd1\x6a\xb4\xca\x20\xb5\x1c\x5f\xc2\x22\xce\x51\xff\xbf\x10\xd2\x79\x0c\xe9\x84\x2f\x3a\xb2\xbe\x6d\x30\x17\x5d\x44\x40\xf2\xce\x12\xd2\x00\x28\x84\xbc\xbe\xba\x1a\x3d\x12\x42\x6a\x24\x15\x8c\xe7\x46\xc5\x8f\x82\x6a\xa5\x90\x68\x5d\x97\xa2\xcd\x74\xd9\x4b\x9f\x40\x49\x52\x78\x95\x4c\x08\xf9\x3e\xe0\x3a\xe6\x79\xb7\xd0\xb8\x36\xd6\xc4\xbc\xd4\xdd\x5e\x53\xee\xd7\x26\xb1\x1c\xc0\xb7\xd9\xbe\xf5\xb6\x47\xcd\x43\x80\x0a\x19\x43\x27\xfa\xee\x37\x22\x65\xa1\x4a\x3e\x7a\x70\x7a\x33\xae\xdd\xd8\xa9\x9d\x80\x4f\xb5\x09\x18\x75\xe5\x50\xe3\x1f\xe7\xfc\x54\x23\xf1\x21\x94\x57\x3d\xca\x0c\xf9\x98\xac\xfc\xd4\xba\x66\x73\xdb\x1e\xd2\xa5\x5d\x65\xfd\x74\x8d\x7a\x13\x8e\xab\x80\x55\x71\x94\xdf\xee\x12\xe2\x1f\x71\x5b\x2a\xf6\xbf\xf2\x5a\xc3\xf8\x2c\x4e\x7b\x69\xb8\xbd\xca\xba\x96\xa7\xd1\x53\x97\xad\xef\x39\xde\xf8\x24\xa0\x7b\x78\xc4\x5e\x27\x8b\xfd\xd5\x63\x60\x33\xb2\x96\x74\x21\x07\x6b\x7e\x42\x63\xa2\x81\xe9\xda\x5c\xc4\xc1\xd8\x5c\xee\xbd\xd8\xca\xe9\x3a\x8d\x99\xa3\x91\xcd\x88\xb9\x71\x96\x38\x80\xb1\x3c\x27\x09\x29\xe7\x5f\xbf\x2d\x2d\x12\x42\x80\xa1\x59\xa4\x61\xac\xc6\xf1\x6f\x9c\x35\xf0\xf3\xde\x5e\x30\xbc\x8a\x3b\xb0\x90\x63\x85\x96\x3f\x9b\x81\x2c\x47\x5f\x4c\x09\x36\xaf\x21\x9f\x25\x2d\x6d\x88\xb1\x9a\x83\x3c\xbb\x1d\xe6\x40\xfb\x5d\xf8\x74\x2b\xec\x6f\x0a\xa3\x37\xee\x08\x83\x4c\x8d\xae\x13\xbc\x41\x18\xfb\x69\x6c\x1f\x33\xd4\x5a\x9b\x12\x6f\x9c\x65\xb4\x3c\xa9\xd6\xa8\x90\xb4\x07\x5a\x27\x2d\xa0\xbc\xdf\x5f\xd6\x9b\x15\x4c\x0d\x86\x3a\x94\x07\x93\x38\x50\xde\x66\x42\x9d\xa0\x6f\x35\x78\x8b\xff\x92\xa7\x46\xad\x63\x8e\xc3\x06\x03\xeb\x5c\xfe\x1a\x7e\x9d\xcf\x96\x6c\x7e\x47\x9e\xf8\x62\x39\xc9\x0f\xac\x0a\xd4\xf7\x67\x63\x16\xbf\x06\xb2\x6d\xf6\x3b\x00\x00\xff\xff\xbe\x6f\x55\xf4\xfa\x0d\x00\x00")

func v1alphaExtractorExtractorSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaExtractorExtractorSwaggerJson,
		"v1alpha/extractor/extractor.swagger.json",
	)
}

func v1alphaExtractorExtractorSwaggerJson() (*asset, error) {
	bytes, err := v1alphaExtractorExtractorSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/extractor/extractor.swagger.json", size: 3578, mode: os.FileMode(420), modTime: time.Unix(1571198791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaSchemaSchemaSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\x31\xae\xc3\x20\x0c\x06\xe0\x9d\x53\x58\x9e\x9f\x92\xd7\x8e\xb9\x4a\xd5\x01\x11\x12\xa8\x52\x8c\xb0\x93\x0e\x51\xee\x5e\x19\x14\x55\xdd\x3a\x61\xd9\xdf\xff\x8b\xdd\x00\x20\xbf\xec\x3c\xfb\x82\x03\xe0\xb5\xfb\xc7\x3f\xdd\xc5\x34\x11\x0e\xa0\x77\x00\x94\x28\x8b\xd7\xfb\x76\xb1\x4b\x0e\xb6\x67\x17\xfc\xf3\x7c\xba\x5c\x48\xa8\xe6\x00\x70\xf3\x85\x23\xa5\xaa\xdb\x08\x89\x04\xd8\x0b\x1a\x80\xa3\xb6\xd7\x9c\x67\x1c\xe0\xd6\x42\x41\x24\x9f\x05\x3a\xb3\xda\x7b\xb5\x8e\x12\xaf\x5f\xd8\xe6\xbc\x44\x67\x25\x52\xea\x1f\x4c\xe9\x63\x73\xa1\x71\x75\x3f\x5a\x2b\x41\xe1\xde\xbe\x34\xfa\x29\xa6\xa8\xae\x2d\xcd\x61\xde\x01\x00\x00\xff\xff\x1c\xe4\x75\xba\x1d\x01\x00\x00")

func v1alphaSchemaSchemaSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaSchemaSchemaSwaggerJson,
		"v1alpha/schema/schema.swagger.json",
	)
}

func v1alphaSchemaSchemaSwaggerJson() (*asset, error) {
	bytes, err := v1alphaSchemaSchemaSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/schema/schema.swagger.json", size: 285, mode: os.FileMode(420), modTime: time.Unix(1571198791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaDepsDepsSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xc1\x0a\xc2\x30\x0c\x06\xe0\x7b\x9f\x22\xe4\x2c\x9b\x7a\xdc\xab\x88\x87\xd2\x65\x6b\x65\x36\xa5\xc9\xe6\x61\xec\xdd\xa5\x2d\x22\xde\xbc\x94\x90\x7c\xff\x4f\x77\x03\x80\xf2\xb2\xf3\x4c\x19\x07\xc0\x6b\x77\xc6\x53\xd9\x85\x38\x31\x0e\x50\xee\x00\xa8\x41\x17\x2a\xf7\xed\x62\x97\xe4\x6d\x3f\x52\x92\xfa\x74\x29\xb3\x72\xcd\x00\xe0\x46\x59\x02\xc7\x2a\xdb\x08\x91\x15\x84\x14\x0d\xc0\x51\x9b\xc5\x79\x7a\x92\xe0\x00\xb7\x16\xf2\xaa\xe9\x53\x50\x66\x29\xf6\x5e\xad\xe3\x28\xeb\x0f\xb6\x29\x2d\xc1\x59\x0d\x1c\xfb\x87\x70\xfc\xda\x94\x79\x5c\xdd\x9f\xd6\xaa\x2f\x70\x6f\x5f\x1a\x69\x0a\x31\x14\xd7\x96\xe6\x30\xef\x00\x00\x00\xff\xff\x08\xa8\x20\x09\x19\x01\x00\x00")

func v1alphaDepsDepsSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaDepsDepsSwaggerJson,
		"v1alpha/deps/deps.swagger.json",
	)
}

func v1alphaDepsDepsSwaggerJson() (*asset, error) {
	bytes, err := v1alphaDepsDepsSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/deps/deps.swagger.json", size: 281, mode: os.FileMode(420), modTime: time.Unix(1571198791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v1alphaStoreStoreSwaggerJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x54\x4d\x8f\xda\x30\x10\xbd\xe7\x57\x58\x6e\x8f\x08\x0a\x47\x6e\x95\xa0\x55\xab\xaa\x45\xb4\x52\x0f\x2b\x0e\x26\x99\x24\x66\x13\x8f\x65\x4f\x58\x21\x94\xff\xbe\xb2\xb3\xf9\x00\xc2\x6e\x56\xb0\x97\xc8\x9a\x79\xef\x79\xde\x8c\x33\xc7\x80\x31\x6e\x9f\x44\x92\x80\xe1\x73\xc6\x67\xe3\x2f\x7c\xe4\x62\x52\xc5\xc8\xe7\xcc\xe5\x19\xe3\x24\x29\x03\x97\xdf\x4f\x45\xa6\x53\x31\xb1\x84\x06\xaa\xef\x58\x1b\x24\xf4\x2c\xc6\xf8\x1e\x8c\x95\xa8\x3c\xb6\x3a\x32\x85\xc4\x2c\x10\x0f\x18\x2b\xbd\xb6\x0d\x53\xc8\xc1\xf2\x39\x7b\xa8\x48\x29\x91\xae\x05\xdc\xd9\x3a\xec\xc6\x63\x43\x54\xb6\x38\x01\x0b\xad\x33\x19\x0a\x92\xa8\x26\x3b\x8b\xaa\xc5\x6a\x83\x51\x11\x0e\xc4\x0a\x4a\x1d\xf0\x58\x95\x14\x41\x2c\x95\x74\x38\xdb\xba\xf6\xf6\x16\x90\x01\xc1\x1a\xac\x46\x65\xa1\x49\xba\xa6\x1c\xb4\xef\x09\x6e\x77\x10\x7a\x7b\x2f\x06\x6b\xea\x37\xa9\xa2\x01\xc4\x51\x1d\xd7\x06\x35\x18\x92\x60\x3b\x68\x5f\xab\x34\xa7\xa1\x8e\x88\x30\x46\x1c\x1a\x0d\x9f\x92\x04\xf9\x39\x9e\x31\xfe\xd9\x40\xec\x18\x9f\x26\x1d\xb7\xd5\x0c\xbf\x1b\xa1\xd3\x1f\x04\xf9\x4a\x48\xc3\x3b\xb4\x32\x38\x3f\x95\x97\x3e\x1b\xf6\x2d\x26\x93\x5a\xe4\x5f\x45\xec\x35\x6b\xc9\x48\x95\xb4\x05\x96\xad\x6f\xfe\x38\x7d\x83\x75\xd2\xa3\x18\x4d\x2e\xc8\x65\xb7\x07\x82\x2b\x8a\xb3\x7b\x2b\x82\x0a\x31\x72\xcc\x33\xdd\x21\xa3\x59\xd6\xdc\x5e\xe5\xa6\x7d\x0b\x41\xe2\x2e\x65\x0f\x98\xf7\xf2\xd2\xce\xb5\xbb\x38\xa8\x22\x6f\x7e\x4c\x1f\x59\x7f\xfd\xdf\x29\x85\xff\xfc\xfb\xe7\x77\x7d\xfd\xa6\xa1\x45\x10\x8b\x22\xf3\xf5\x39\xfc\x2b\xc5\xf8\xa7\x7b\xc3\x03\x84\x28\xb9\x78\x77\x43\xe6\xd2\x3f\x0f\x85\xd1\xad\x6a\xd7\x07\xf0\x4b\x5a\xba\xc7\x62\xe9\x5b\x14\x1f\xb4\x58\xde\xbd\x54\x56\xc5\x10\x8b\x41\xcd\x2e\x83\x32\x78\x0e\x00\x00\xff\xff\x9b\x51\xf8\x1c\xcd\x06\x00\x00")

func v1alphaStoreStoreSwaggerJsonBytes() ([]byte, error) {
	return bindataRead(
		_v1alphaStoreStoreSwaggerJson,
		"v1alpha/store/store.swagger.json",
	)
}

func v1alphaStoreStoreSwaggerJson() (*asset, error) {
	bytes, err := v1alphaStoreStoreSwaggerJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v1alpha/store/store.swagger.json", size: 1741, mode: os.FileMode(420), modTime: time.Unix(1571198791, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v1alpha/tracker/tracker.swagger.json":     v1alphaTrackerTrackerSwaggerJson,
	"v1alpha/extractor/extractor.swagger.json": v1alphaExtractorExtractorSwaggerJson,
	"v1alpha/schema/schema.swagger.json":       v1alphaSchemaSchemaSwaggerJson,
	"v1alpha/deps/deps.swagger.json":           v1alphaDepsDepsSwaggerJson,
	"v1alpha/store/store.swagger.json":         v1alphaStoreStoreSwaggerJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v1alpha": &bintree{nil, map[string]*bintree{
		"deps": &bintree{nil, map[string]*bintree{
			"deps.swagger.json": &bintree{v1alphaDepsDepsSwaggerJson, map[string]*bintree{}},
		}},
		"extractor": &bintree{nil, map[string]*bintree{
			"extractor.swagger.json": &bintree{v1alphaExtractorExtractorSwaggerJson, map[string]*bintree{}},
		}},
		"schema": &bintree{nil, map[string]*bintree{
			"schema.swagger.json": &bintree{v1alphaSchemaSchemaSwaggerJson, map[string]*bintree{}},
		}},
		"store": &bintree{nil, map[string]*bintree{
			"store.swagger.json": &bintree{v1alphaStoreStoreSwaggerJson, map[string]*bintree{}},
		}},
		"tracker": &bintree{nil, map[string]*bintree{
			"tracker.swagger.json": &bintree{v1alphaTrackerTrackerSwaggerJson, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
