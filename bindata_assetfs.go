// Code generated by go-bindata.
// sources:
// html/index.css
// html/index.html
// html/index.js
// DO NOT EDIT!

package main

import (
	"github.com/elazarl/go-bindata-assetfs"
	"bytes"
	"compress/gzip"
	"fmt"
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
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _htmlIndexCss = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x52\xcd\x6e\xdb\x3c\x10\xbc\xeb\x29\xf6\xa6\x38\xf8\x28\xc5\x46\xf0\xa1\x90\x2e\x05\x8a\x04\x30\xda\x02\x45\x73\xe8\x99\x92\x56\xd2\x22\x14\x29\x2c\xd7\x90\xe5\x22\xef\x5e\x88\xf2\x4f\xec\xb8\x87\x1e\x7c\x13\x87\xc3\x99\xd9\xd5\x7c\xae\x9d\x15\x55\xeb\x12\xe1\x77\x04\x00\xb0\x3f\x77\x64\xc6\x0c\xe2\xef\x5a\x90\x49\x1b\x58\x97\xce\xfa\x38\x3f\x51\xbc\x8c\x06\x33\xb0\x8e\x3b\x6d\xde\xe1\x03\x52\xd3\x4a\x06\x8f\x0f\x0f\x33\xea\xb9\xcc\x60\xc3\xe6\xae\x15\xe9\xb3\x34\x2d\x2b\x9b\x14\xce\x49\xe9\x7d\x52\xba\x2e\xed\xf6\x16\xaa\x42\x4f\x8d\x55\x34\x39\xa5\xab\x64\x99\x2c\xd3\xe9\x7b\x52\x4d\x0f\x39\x42\x0c\xf5\x13\x9b\x8d\xd1\x9c\xa0\x93\x45\x0e\xe9\x3d\x3c\x3b\x86\xf5\xd3\xff\xea\x13\xdc\xa7\x27\x53\xe3\x4a\x6d\xee\x2e\x67\x58\xfc\x17\x18\x17\x97\x67\xc2\x07\xce\x6d\x52\x0f\xae\xae\x57\x0b\xa8\xa7\xcd\xc9\x5d\x1c\x8e\xb7\x77\x3c\x37\xbc\xad\x9f\xc8\x3b\x3b\xe1\x0d\xca\xd8\x63\xbc\xc8\xa3\xb7\x28\x4a\x8e\xba\x41\xf0\x1f\x5b\x77\x68\xd7\x87\xda\xfd\xad\x8e\x9e\x76\x98\xc1\xea\xb1\xdf\xe6\x30\x35\xe5\x07\x63\x8d\xcc\x58\xc1\x64\x0f\xd3\xf5\xa1\x33\x15\xf9\xde\xe8\x31\x03\xb2\x86\x2c\xaa\xc2\xb8\xf2\x75\x96\x1a\xa8\x92\x36\x83\x25\x76\xf3\xb9\xdd\xc7\x38\x02\xe1\xc1\x11\x9d\x31\xc1\xad\x28\x61\x6d\xfd\xb4\x8a\x29\x9b\xc5\x3d\x1b\x45\x90\x95\xef\x75\x49\xb6\x39\x4f\x3d\x38\xae\xd4\xc0\xba\xbf\x80\x5b\x12\x0c\x2f\xc2\x94\x13\x21\xdf\xa7\x66\x2c\x85\x9c\xcd\xc0\x08\xe7\x51\x00\xd3\x7b\x78\xd9\xf4\xbd\x63\x99\x7e\x03\x68\x63\xe0\x17\x16\x5f\x49\xa0\x60\x37\x78\x64\x9f\x1c\xa6\x56\x03\x16\xaf\x24\x6a\xde\x56\xe7\x9c\xb4\x21\x93\xb6\x42\xda\x90\xf6\x58\xe5\xd7\x34\x5f\x74\xad\x99\x40\xdb\x0a\xbe\xb4\xec\x3a\x3c\x2a\x86\xb9\x19\x6d\x85\x1c\x94\x5c\x2f\xd4\xd1\x0e\xbf\x61\x43\x05\x19\x92\xf1\x7a\xca\x67\x62\xac\xdd\xf6\x94\xac\x73\x3b\xe5\xfc\xf6\x43\xb4\x86\xf5\xe8\x4b\x6d\xf0\xba\xce\xfa\xe9\x28\x31\xd7\x0a\xb5\x6c\x18\x95\x47\x11\xb2\x8d\xcf\x20\x36\xd4\xe8\x38\x8f\xde\xfe\x04\x00\x00\xff\xff\xde\x37\xd6\x77\xfe\x04\x00\x00")

func htmlIndexCssBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexCss,
		"html/index.css",
	)
}

func htmlIndexCss() (*asset, error) {
	bytes, err := htmlIndexCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.css", size: 1278, mode: os.FileMode(420), modTime: time.Unix(1498736861, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x58\x5d\x6b\xe4\x36\x17\xbe\xcf\xaf\xd0\xea\x85\x77\x12\x58\xdb\xd9\x49\x0b\x9b\x60\x7b\x29\xbb\x2d\x14\x16\x5a\xf6\xa3\x50\x4a\x59\x64\xeb\xcc\x8c\x36\xb2\xe4\x4a\xb2\x9d\x69\x08\xec\x55\xdb\x50\xda\xa5\x14\x7a\xd5\xab\x42\xe9\x6d\xef\xda\x5e\x6c\x7f\x4d\x37\x6c\xfe\x45\x91\xbf\x62\xcf\xd8\x93\x0c\xdb\xdc\x8c\x2d\x1d\x9d\xf3\x9c\x47\x47\x8f\x8e\xe3\xdf\x7a\xf0\xd1\xfd\x27\x9f\x7e\xfc\x3e\x5a\x98\x84\x87\x3b\xbe\xfd\x41\x9c\x88\x79\x80\x41\x60\x74\x92\x70\xa1\x8f\x72\x27\x62\x82\x06\x78\x61\x4c\x7a\xe4\x79\x45\x51\xb8\xc5\x81\x2b\xd5\xdc\xbb\x73\x78\x78\xe8\x9d\xd8\x45\x57\xb6\x52\x5c\x63\xb9\x83\xca\xbf\xc6\x3e\x91\x14\xf8\x35\x4b\x2c\x32\x20\x34\x2c\x97\xfa\x09\x18\x82\xe2\x05\x51\x1a\x4c\x80\x9f\x3e\xf9\xc0\xb9\x8b\xeb\x29\xce\xc4\x31\x5a\x28\x98\x05\xd8\xf3\x62\x2a\xdc\x48\x4a\x13\x6b\xed\xc6\x32\xf1\x12\x62\x40\x31\xc2\xd9\x97\xe0\xed\xbb\x87\x77\xdd\xa9\x17\x6b\xdd\x1d\x76\x13\x26\xdc\x58\x6b\x8c\x94\xc5\xa4\xcd\x92\x83\x5e\x00\x98\xad\xfc\x3b\x14\x34\x9b\x0b\x87\xc5\x52\xb0\xd8\x99\x49\x61\xbc\xa9\x3b\x75\xf7\x7b\xe1\x06\xac\xb6\x09\xcf\x04\x85\x93\x8d\xc6\x3a\x56\x2c\x35\x48\xab\x78\x08\xec\xf3\x2f\x32\x50\x4b\xef\xc0\x9d\xba\x77\xea\x97\x32\xfc\x73\x8d\x43\xdf\xab\xd6\xde\xc8\xd1\x00\xab\xcf\xd7\x49\xdd\xda\x6d\x9e\x81\x37\x75\x0f\xdc\x77\xec\xd3\x80\x8b\x1d\xdf\xab\x4a\xc2\x8f\x24\x5d\x86\x3b\x3e\x65\x39\x62\x34\xc0\x09\x61\xa2\xa1\x40\x90\x1c\xc5\x9c\x68\x1d\x60\x41\x72\x07\x4e\x0c\x08\x0a\x14\x45\x3c\x83\xda\xa6\xb4\xb3\x6b\x3b\x76\x85\x22\x69\x0a\xaa\x63\x51\x5a\x91\x9a\xfa\xff\xe1\xc6\x38\x52\x44\x50\x87\xcb\xb9\xc4\xe1\x9b\xdf\xbe\xba\xfc\xe5\x87\x8b\xf3\xbf\x2e\xbf\x7e\xf9\xcf\x9f\xdf\xbe\x79\xf5\xca\xf7\xc8\x8a\x87\x8c\x97\x10\x6d\x8c\x44\x46\x8c\x43\xeb\x49\xb1\xf9\xc2\xa0\x05\xa3\xe0\x48\xe1\x24\x40\x1d\xeb\x99\xca\x42\xac\xa0\xa8\xeb\x00\xa1\xdc\x99\x49\x15\xe0\x5c\x23\x26\x50\xae\xf5\x80\x5d\x8d\xda\x9e\xc9\xa3\x98\xb3\xf8\x38\xc0\x3a\x8b\x76\x73\xed\x32\xba\x87\x51\x75\xae\x8f\x6a\x04\x9f\x69\xe0\x10\x9b\x4f\x74\x10\x94\x06\xf7\x26\x96\x25\x44\x89\x3a\x06\xe1\x4c\x27\x47\x93\xc9\xe7\x38\x3c\x3d\xcd\xb5\x6b\x98\xe1\x70\x76\xb6\x96\x5f\x19\xcf\xe3\x6c\x25\x6b\x2f\xe3\x1d\xae\x3d\xca\xf2\x71\xea\x63\x29\x0c\x08\xb3\x4a\x7d\xc7\x2a\x96\x1c\xe9\x3b\xd3\x21\x5a\xa2\xcc\x18\x29\x1a\x66\x74\x16\x59\x6a\x74\x16\x69\xdc\xe3\x80\x33\x6d\x76\x75\x16\x0d\xb2\x30\x89\x8c\x70\x66\x9c\x98\xc9\xed\x92\x80\xc9\xed\x49\xb1\x60\x06\x1c\x03\x27\x66\x72\xbb\x22\xe9\x71\x16\x05\x41\xe5\xe1\xde\xa4\x66\xe8\xa0\x65\xc8\x4e\xb4\x14\x55\x98\x56\x19\xe9\x73\x70\xf5\xea\x7b\x82\x34\x8f\xbd\x9c\x85\x21\x4c\x80\xb2\x70\x59\x79\xfa\x99\x01\x3a\x52\xc2\x4a\x16\xd6\x4e\x2f\x64\x11\xe0\x5b\xb9\x4d\x97\x4b\x42\x57\x39\x8d\xd4\x7a\x75\x5e\x71\x6c\xd3\x64\x72\xac\xfa\xd6\xec\x1c\x66\x20\xc1\x0d\xf5\xbc\xac\x49\x1b\x78\xac\x2a\x75\x4a\x44\x7b\x8a\x08\x9d\xc3\x88\xa1\xfd\x3b\x3d\xe5\x6e\xca\xc9\xf2\xbe\xcc\x84\x39\x3b\x1b\xb5\x2b\x4b\xbd\xdc\xcc\xea\x9c\x72\x57\xcb\x4c\xc5\xf0\xf4\xd1\x43\x8c\x0c\x51\x73\x7b\x5d\xe4\x0c\x8a\x67\x05\x13\x54\x16\x38\x44\x7e\x9b\x4a\x2b\xcb\x56\x8f\x35\x2a\xcf\x23\x0e\xad\xe0\xfa\x1e\x0b\xd1\x60\xad\x5f\x05\xbe\xd6\x53\xaf\x00\x6d\x32\xbb\xdc\xcd\x19\x05\xf9\xe1\x83\x3d\x1c\xda\x81\x67\x44\x29\x59\x6c\x0e\xe6\x7b\x96\xb7\x91\x39\x26\xd2\xcc\x20\xb3\x4c\x21\xc0\xf1\x02\xe2\xe3\x48\x9e\xd8\xb8\xf5\x05\x9b\x33\xaa\xdb\x6a\xb7\x2a\xd4\x02\x68\x47\x73\xc2\x33\xe8\x4d\x08\x92\x40\xbd\xd4\x1b\x09\xcb\x49\x04\xbc\xf1\x50\xed\x7e\xbb\x3e\xb4\x5b\xd7\x1e\x85\xd2\x72\x7b\xc5\x40\x83\x02\xb0\x3f\x50\x30\xff\x17\x91\x4e\xd7\xfd\x93\xb6\xd0\x8c\xa8\xa4\xbf\xb7\x1b\x33\x22\x1e\x97\xa7\x1a\x87\xaf\x5f\x7e\x77\xf9\xe2\x7c\x58\xd7\xae\xf1\x52\x8a\xca\x1e\x0e\x5f\x7f\xf3\xc7\xc5\x4f\xbf\xaf\x4b\x7f\x75\xc6\x37\x27\x35\xad\x8a\xc5\x21\x9c\xcd\x07\x4f\x5e\x2d\x70\x1d\x20\x0a\x68\x1f\x47\x79\x65\x34\xdb\x41\x99\x26\x11\x07\x1a\x4c\xec\x16\xba\x1c\xc4\xdc\x2c\x82\x60\x7f\x12\x36\x37\xd4\x76\xf2\xb4\x8a\xba\xa7\x34\x37\x17\x9a\xb5\x81\x8e\xcb\x54\xc9\xb9\x82\xc1\xeb\xac\x6b\x66\xdb\x1f\x03\x2a\x61\x82\x18\xb0\x7d\x41\x0f\xe3\xe6\x2c\xba\x7b\x31\x90\x0d\x70\x0d\x9d\x75\x5d\xb0\xfd\x97\x95\xed\x93\xb3\x99\x06\xe3\xe8\x77\x37\x5c\x5e\xa9\x02\xcb\x10\xa8\xa6\xc7\x40\x11\x9b\x23\x12\x1b\x96\x0f\x29\x60\x77\xa9\x4e\x99\x10\xa0\x1c\x4e\x96\xa0\x50\xf3\x66\x0b\xd1\x91\x82\x2f\xc7\x84\xb6\x0b\x92\xa9\x98\x83\x13\x73\x56\x46\xe6\x30\x1b\x93\xe7\xe1\x95\x83\x34\xf7\xf9\xed\x2e\x9a\x93\xd4\x49\x89\x89\x17\x37\x86\x36\x1a\x60\xdd\xf9\x4a\x2e\xb5\x5c\xff\x37\x71\x6e\x32\xbc\x45\x75\x59\xad\xb5\x02\xdf\xf6\x7a\x89\xa4\x84\x8f\xdc\xdd\xe5\xdc\x58\x17\x54\xca\x6a\xd5\x30\x5f\x35\x8e\xa0\x53\x29\x34\xcb\xc1\x29\xa7\x71\x19\xaf\x7e\xb4\x7e\x94\xe4\x3a\xf4\xbd\x72\xe4\xe6\x80\x2b\x19\x79\x5b\xc0\x2b\x87\x6b\x73\x81\x8f\x37\x76\x68\x58\xfb\xd6\x45\xd8\xa2\x7e\x8f\x73\x5c\xcb\xdb\xc5\xf9\x8b\x8b\x9f\xcf\x2f\x7f\xfc\xfb\xe2\xfb\x5f\x77\x4f\x4f\xed\xec\xd3\x47\x0f\x1b\x1d\x3c\x3b\xdb\x1b\xd6\xbf\x4d\x95\xb0\x2a\x5d\x5b\x0e\x0e\xa4\xdb\xf4\x4c\x35\x3a\xdb\x39\x35\x40\x47\xa8\x68\x33\x19\xe9\x85\x86\x43\xbf\x4d\x42\xd7\x56\x7b\xfb\x53\x7f\x8c\x75\xbf\xed\xaa\x2f\xd5\xfe\xd7\x9b\xef\x95\xff\x70\xf8\x37\x00\x00\xff\xff\xac\xd1\x37\xbe\x80\x10\x00\x00")

func htmlIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexHtml,
		"html/index.html",
	)
}

func htmlIndexHtml() (*asset, error) {
	bytes, err := htmlIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.html", size: 4224, mode: os.FileMode(420), modTime: time.Unix(1498742799, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlIndexJs = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x5f\x6f\xe2\x38\x10\x7f\x86\x4f\xe1\xa6\xd5\xc5\x01\x1a\xe0\x4e\xba\x93\x72\x8a\x4e\x77\xbd\x7d\xa8\xb4\xd2\x4a\xed\xb6\x2f\x88\x87\x24\x9e\x2c\x96\x1c\x3b\xc2\x76\x58\x84\xf2\xdd\x57\x36\x04\x08\x71\x28\x2b\xad\x9f\x92\x99\xdf\x8c\xe7\x37\xff\x92\xe9\x68\x34\x44\x23\xf4\xb4\x86\x44\x01\x41\xe9\x16\x15\xdb\x82\x21\xc1\xd1\xfc\xaf\xc7\x3f\x1f\x7f\xff\x23\x1c\xa2\xd1\x74\x98\x94\x65\xcc\x61\x83\xde\x35\xe0\xdd\x10\x21\x84\x80\x45\xc8\xbf\x2f\x12\xca\xfd\x89\x15\x90\x44\x25\x11\xda\x2b\xcd\xa9\xa4\x8c\x76\xf5\xe4\xf8\x2e\x81\x41\xa6\xde\x65\xe4\x79\x67\x42\x9d\xca\x68\xb1\xbc\x44\xbd\xea\xb4\x05\xab\x18\x95\xaa\xe5\xad\xa2\xa4\x6d\x48\xc4\x86\xbf\xbd\x7c\x6e\x0b\xad\x1d\x13\x09\x89\xf2\x84\x49\x38\x29\x28\xa7\x0a\x5a\xd2\x83\xef\x02\xd4\x4a\x10\x19\xed\x5a\xd0\x28\xd7\x3c\x53\x54\x70\x84\x83\x33\x8a\xe6\x3c\x84\xdf\x40\x61\x6f\x25\xa4\xf2\x26\x27\x98\xc9\xc6\x25\xd4\x9c\xa4\x2c\xc3\x4a\xca\xd8\xe8\x9d\x4a\xa9\x53\x6b\xbc\x98\x2d\x43\x4a\x82\x16\xa6\x3e\xbd\xd6\xad\x1c\x9e\xc5\x57\xc9\xcb\x6b\xed\x95\x4d\x22\x62\xb5\xd6\xd0\x51\x9b\x32\xc4\x8b\x65\x57\x7e\xa8\x59\x5c\xc9\x3e\xd2\x53\x6f\x5c\xc9\x1b\x89\xdb\x6b\x7a\x99\x9b\x10\x7b\xa9\x37\xa0\x7d\xe1\xba\x2c\xdc\xa9\xb1\x6d\xd3\xc4\x86\xa5\x4e\x83\x9f\x4a\x0d\xcd\x8d\xcd\x5d\xac\x39\x81\x9c\x72\xe8\x89\xe9\xd8\xb3\xb1\xd4\x69\x0b\x01\x4c\x42\xc7\x44\xea\x34\x6e\x99\x5d\x49\xed\x79\x11\xc6\xde\x54\xea\xb4\x25\x7c\xd5\xe9\xad\x3d\x67\x58\xc6\xbb\xda\xad\xa4\xa4\x53\x7e\x73\x72\xb1\xc6\x14\x51\x6e\x67\x3b\xe8\x3a\x6e\x39\x5f\xd8\xd2\xd1\xa5\xf1\x06\xe2\xf9\xff\x65\x7c\x10\xf4\x9b\x51\x22\xc3\x52\xcb\x15\xbe\x30\xed\xa6\x19\x75\x8a\x65\x47\xb7\x83\xab\x6f\x68\x8a\x3c\xe1\xaf\x36\x7b\x57\xa6\xfa\xc8\x5c\xad\xa8\xdc\xdf\xea\xe0\x5f\x51\x12\x9f\x00\x67\x04\x3a\x48\xca\x09\x7c\xff\x92\x1f\xd0\x86\xf7\x41\x82\x2b\x57\xa7\xd3\x1c\x1f\xf4\x77\xf1\xe3\xbc\x27\xf3\x27\x5f\xb2\x64\x34\x83\xc6\x64\x32\xef\x3a\xac\x4d\x27\x7e\xe4\xc6\x96\xc2\x19\xcf\x45\x5a\x5d\x59\x2d\x59\xb2\x3d\x5f\x43\x94\x5c\xd9\x94\xdd\xce\xb6\x89\x33\xbb\x84\x92\x1b\x3a\xda\xa2\x63\x22\x32\x5d\x00\x57\xc6\xef\x27\x06\xe6\xf1\xbf\xed\x33\xc1\xbe\x55\xfb\x5d\x1a\x56\x1e\xca\x75\x16\x7b\xf6\xf1\x1f\xfb\x38\x06\x9e\x09\x02\x6f\x2f\xcf\x4f\xa2\x28\x05\x07\xbe\x5f\x43\x06\xd8\xe7\xc3\xb0\xc5\x5d\xe5\x03\xf6\xef\x8d\xca\x0f\xc2\x42\x90\x84\x61\x5f\x94\xc0\xfd\x9b\xb6\xb8\x12\xe5\x95\x86\xfc\x88\x6b\x58\x26\x5a\x02\x76\xba\x36\x9f\xc6\x2b\xae\x4d\x25\x9a\xaf\xe7\xe5\x16\x30\x84\x8c\xee\x2a\xa1\xe3\xb0\x34\x43\xed\x1a\x95\x64\x6d\xc7\xa5\x81\xb8\xf6\x82\xc1\xf0\xa4\x80\xf8\xb4\x53\x2a\x4a\x96\xa1\xa2\x8a\xc1\xdf\x1d\x38\x6e\x75\xdb\xc4\x58\xba\x7a\x05\xfd\xda\xd6\x73\x65\x6d\x3f\x39\x4c\x64\x89\xb1\x0d\x57\x6b\xc8\xc7\x37\x77\xd8\xd8\xfb\xcd\x92\x76\xa2\x2c\x29\xf7\x3a\xac\x1d\x63\x1a\x9c\x32\xf1\xf1\xc4\x9a\xe8\xff\x65\xec\xe6\x2d\xd8\xb0\x0d\x76\xc3\xc1\x60\x30\x90\xa0\xbe\xd2\x02\x84\x56\xd8\xdb\x50\x4e\xc4\x26\x34\xbd\x81\x7d\x6f\xdc\x82\x2f\xe8\x72\xec\xf9\x81\x37\x99\xcf\x66\xb3\x11\xed\x0d\xac\xf9\x0b\xab\x83\xe1\x03\xf6\xf7\xed\x76\x6c\xbb\x7d\x5c\x99\x28\x4a\x06\x0a\x7a\x42\xb6\x75\x55\xa2\x3c\x4c\x41\x6d\x5c\x35\xff\x0c\x38\xf8\x11\x00\x00\xff\xff\xb4\xcb\x9a\x5f\xea\x0a\x00\x00")

func htmlIndexJsBytes() ([]byte, error) {
	return bindataRead(
		_htmlIndexJs,
		"html/index.js",
	)
}

func htmlIndexJs() (*asset, error) {
	bytes, err := htmlIndexJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/index.js", size: 2794, mode: os.FileMode(420), modTime: time.Unix(1499658232, 0)}
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
	"html/index.css": htmlIndexCss,
	"html/index.html": htmlIndexHtml,
	"html/index.js": htmlIndexJs,
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
	"html": &bintree{nil, map[string]*bintree{
		"index.css": &bintree{htmlIndexCss, map[string]*bintree{}},
		"index.html": &bintree{htmlIndexHtml, map[string]*bintree{}},
		"index.js": &bintree{htmlIndexJs, map[string]*bintree{}},
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


func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}
