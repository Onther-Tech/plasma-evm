// Code generated by go-bindata.
// sources:
// faucet.html
// DO NOT EDIT!

package main

import (
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x6d\x93\xdb\x36\x92\xfe\x3c\xfe\x15\x1d\x9e\xbd\x92\xce\x43\x52\x33\x63\x7b\x7d\x12\xa9\x94\xd7\x9b\xdd\xf3\xd5\x5d\x92\x4a\x9c\xba\xdb\xca\xa6\xae\x40\xb2\x25\xc2\x03\x02\x0c\x00\x4a\xa3\x4c\xe9\xbf\x5f\x35\x40\x52\xd4\xcb\x4c\xec\xb5\xaf\x6a\xfd\x61\x4c\x02\x8d\x46\xa3\xfb\x69\xf4\x0b\x95\x7c\xf5\xe7\xef\xde\xbe\xff\xdb\xf7\xdf\x40\x69\x2b\xb1\x78\x92\xd0\x7f\x20\x98\x5c\xa5\x01\xca\x60\xf1\xe4\x22\x29\x91\x15\x8b\x27\x17\x17\x49\x85\x96\x41\x5e\x32\x6d\xd0\xa6\x41\x63\x97\xe1\xeb\x60\x3f\x51\x5a\x5b\x87\xf8\x6b\xc3\xd7\x69\xf0\x3f\xe1\x4f\x6f\xc2\xb7\xaa\xaa\x99\xe5\x99\xc0\x00\x72\x25\x2d\x4a\x9b\x06\xef\xbe\x49\xb1\x58\xe1\x60\x9d\x64\x15\xa6\xc1\x9a\xe3\xa6\x56\xda\x0e\x48\x37\xbc\xb0\x65\x5a\xe0\x9a\xe7\x18\xba\x97\x4b\xe0\x92\x5b\xce\x44\x68\x72\x26\x30\xbd\x0a\x16\x4f\x88\x8f\xe5\x56\xe0\xe2\xfe\x3e\xfa\x16\xed\x46\xe9\xdb\xdd\x6e\x06\x6f\x1a\x5b\xa2\xb4\x3c\x67\x16\x0b\xf8\x0b\x6b\x72\xb4\x49\xec\x29\xdd\x22\xc1\xe5\x2d\x94\x1a\x97\x69\x40\xa2\x9b\x59\x1c\xe7\x85\xfc\x60\xa2\x5c\xa8\xa6\x58\x0a\xa6\x31\xca\x55\x15\xb3\x0f\xec\x2e\x16\x3c\x33\xb1\xdd\x70\x6b\x51\x87\x99\x52\xd6\x58\xcd\xea\xf8\x26\xba\x89\xfe\x18\xe7\xc6\xc4\xfd\x58\x54\x71\x19\xe5\xc6\x04\xa0\x51\xa4\x81\xb1\x5b\x81\xa6\x44\xb4\x01\xc4\x8b\x7f\x6c\xdf\xa5\x92\x36\x64\x1b\x34\xaa\xc2\xf8\x45\xf4\xc7\x68\xea\xb6\x1c\x0e\x3f\xbe\x2b\x6d\x6b\x72\xcd\x6b\x0b\x46\xe7\x1f\xbd\xef\x87\x5f\x1b\xd4\xdb\xf8\x26\xba\x8a\xae\xda\x17\xb7\xcf\x07\x13\x2c\x92\xd8\x33\x5c\x7c\x16\xef\x50\x2a\xbb\x8d\xaf\xa3\x17\xd1\x55\x5c\xb3\xfc\x96\xad\xb0\xe8\x76\xa2\xa9\xa8\x1b\xfc\x62\xfb\x3e\x64\xc3\x0f\xc7\x26\xfc\x12\x9b\x55\xaa\x42\x69\xa3\x0f\x26\xbe\x8e\xae\x5e\x47\xd3\x6e\xe0\x94\xbf\xdb\x80\x8c\x46\x5b\x5d\x44\x6b\xd4\x84\x5c\x11\xe6\x28\x2d\x6a\xb8\xa7\xd1\x8b\x8a\xcb\xb0\x44\xbe\x2a\xed\x0c\xae\xa6\xd3\x67\xf3\x73\xa3\xeb\xd2\x0f\x17\xdc\xd4\x82\x6d\x67\xb0\x14\x78\xe7\x87\x98\xe0\x2b\x19\x72\x8b\x95\x99\x81\xe7\xec\x26\x76\x6e\xcf\x5a\xab\x95\x46\x63\xda\xcd\x6a\x65\xb8\xe5\x4a\xce\x08\x51\xcc\xf2\x35\x9e\xa3\x35\x35\x93\x27\x0b\x58\x66\x94\x68\x2c\x1e\x09\x92\x09\x95\xdf\xfa\x31\xe7\xcd\xc3\x43\xe4\x4a\x28\x3d\x83\x4d\xc9\xdb\x65\xe0\x36\x82\x5a\x63\xcb\x1e\x6a\x56\x14\x5c\xae\x66\xf0\xaa\x6e\xcf\x03\x15\xd3\x2b\x2e\x67\x30\xdd\x2f\x49\xe2\x4e\x8d\x49\xec\x2f\xae\x27\x17\x49\xa6\x8a\xad\xb3\x61\xc1\xd7\x90\x0b\x66\x4c\x1a\x1c\xa9\xd8\x5d\x48\x07\x04\x74\x0f\x31\x2e\xbb\xa9\x83\x39\xad\x36\x01\xb8\x8d\xd2\xc0\x0b\x11\x66\xca\x5a\x55\xcd\xe0\x8a\xc4\x6b\x97\x1c\xf1\x13\xa1\x58\x85\x57\xd7\xdd\xe4\x45\x52\x5e\x75\x4c\x2c\xde\xd9\xd0\xd9\xa7\xb7\x4c\xb0\x48\x78\xb7\x76\xc9\x60\xc9\xc2\x8c\xd9\x32\x00\xa6\x39\x0b\x4b\x5e\x14\x28\xd3\xc0\xea\x06\x09\x47\x7c\x01\xc3\xeb\xef\x81\xdb\xaf\xbc\xea\xe4\x8a\x0b\xbe\x6e\x8f\x35\x78\x3c\x3a\xe1\xc3\x87\x78\x0d\xed\x83\x5a\x2e\x0d\xda\x70\x70\xa6\x01\x31\x97\x75\x63\xc3\x95\x56\x4d\xdd\xcf\x5f\x24\x6e\x14\x78\x91\x06\x8d\x16\x41\x7b\xfd\xbb\x47\xbb\xad\x5b\x55\x04\xfd\xc1\x95\xae\x42\xb2\x84\x56\x22\x80\x5a\xb0\x1c\x4b\x25\x0a\xd4\x69\xf0\xa3\xca\x39\x13\x20\xfd\x99\xe1\xa7\x1f\xfe\x13\x5a\x93\x71\xb9\x82\xad\x6a\x34\x7c\x63\x4b\xd4\xd8\x54\xc0\x8a\x82\xe0\x1a\x45\xd1\x40\x10\x87\xdd\x53\x51\xc3\xcc\xca\x3d\xd5\x45\x92\x35\xd6\xaa\x9e\x30\xb3\x12\x32\x2b\xc3\x02\x97\xac\x11\x16\x0a\xad\xea\x42\x6d\x64\x68\xd5\x6a\x45\x91\xce\x1f\xc2\x2f\x0a\xa0\x60\x96\xb5\x53\x69\xd0\xd1\x76\x36\x64\xa6\x56\x75\x53\xb7\x56\xf4\x83\x78\x57\x33\x59\x60\x41\x36\x17\x06\x83\xc5\x5f\xf9\x1a\xa1\x42\x7f\x96\x8b\x63\x48\xe4\x4c\xa3\x0d\x87\x4c\x4f\x80\x91\xc4\x5e\x18\x7f\x24\x68\xff\x25\x8d\xe8\x38\xf5\x47\xa8\x50\x36\x70\xf0\x16\x6a\xba\x57\x82\xc5\xfd\xbd\x66\x72\x85\xf0\x94\x17\x77\x97\xf0\x94\x55\xaa\x91\x16\x66\x29\x44\x6f\xdc\xa3\xd9\xed\x0e\xb8\x03\x24\x82\x2f\x12\xf6\x18\xbc\x41\xc9\x5c\xf0\xfc\x36\x0d\x2c\x47\x9d\xde\xdf\x13\xf3\xdd\x6e\x0e\xf7\xf7\x7c\x09\x4f\xa3\x1f\x30\x67\xb5\xcd\x4b\xb6\xdb\xad\x74\xf7\x1c\xe1\x1d\xe6\x8d\xc5\xf1\xe4\xfe\x1e\x85\xc1\xdd\xce\x34\x59\xc5\xed\xb8\x5b\x4e\xe3\xb2\xd8\xed\x48\xe6\x56\xce\xdd\x0e\x62\x62\x2a\x0b\xbc\x83\xa7\xd1\xf7\xa8\xb9\x2a\x0c\x78\xfa\x24\x66\x8b\x24\x16\x7c\xd1\xae\x3b\x54\x52\xdc\x88\x3d\x5e\x62\x02\x4c\x8f\x73\xe7\x36\x4e\xd4\xa1\xa4\x67\xbc\x60\x15\xf6\xd2\xb7\x78\x30\xdc\xe2\x2d\x6e\xd3\xe0\xfe\x7e\xb8\xb6\x9d\xcd\x99\x10\x19\x23\xbd\xf8\xa3\xf5\x8b\x7e\x43\xc2\xe9\x9a\x1b\x97\x52\x2d\x3a\x09\xf6\x62\x7f\xa4\x5b\x1f\x5d\x5c\x56\xd5\x33\xb8\xb9\x1e\xdc\x5a\xe7\x3c\xfe\xd5\x91\xc7\xdf\x9c\x25\xae\x99\x44\x01\xee\x6f\x68\x2a\x26\xba\xe7\xd6\x5b\x06\xce\x77\xbc\x28\xa4\x3b\xba\x17\xad\xbf\xeb\xa7\x73\x50\x6b\xd4\x4b\xa1\x36\x33\x60\x8d\x55\x73\xa8\xd8\x5d\x1f\xef\x6e\xa6\xd3\xa1\xdc\x94\x0a\xb2\x4c\xa0\xbb\x5d\x34\xfe\xda\xa0\xb1\xa6\xbf\x4b\xfc\x94\xfb\x4b\x57\x4a\x81\xd2\x60\x71\xa4\x0d\xda\x91\x54\xeb\xa8\x06\xa6\xef\x95\x79\x56\xf6\xa5\x52\x7d\x08\x19\x8a\xd1\xb2\x1e\x44\xbb\x60\x91\x58\xbd\xa7\xbb\x48\x6c\xf1\x49\x21\x40\x53\x8a\xf7\x50\x04\xf0\x37\x1a\x9d\xbd\x46\xd4\x3e\xbf\x20\xc8\x82\x7b\x4d\x62\x5b\x7c\xc6\xce\x04\xc2\x8c\x19\xfc\x98\xed\x5d\xa4\xdf\x6f\xef\x5e\x3f\x77\xff\x12\x99\xb6\x19\x32\xfb\x31\x02\x2c\x1b\x59\x0c\xce\xef\xee\xce\xcf\x15\xa0\x91\x7c\x8d\xda\x70\xbb\xfd\x58\x09\xb0\xd8\x8b\xe0\xdf\x0f\x45\x48\x62\xab\x1f\xc7\xda\xf0\xe5\x0b\x39\xf7\xef\xa5\x24\x37\x8b\x7f\x57\x1b\x28\x14\x1a\xb0\x25\x37\x40\xc1\xf5\xeb\x24\x2e\x6f\x7a\x92\x7a\xf1\x9e\x26\xde\xab\x5b\x56\xb1\x5b\x68\x93\x8e\x91\xf1\x6a\x86\xa5\x4b\x36\x80\x1b\xd0\x8d\x74\xb1\x58\x49\xb0\x25\x1e\x26\x28\x6d\xd8\x8e\xe0\xbd\xa2\x24\x6f\x8d\xd2\x42\xc5\x04\xcf\xb9\x6a\x0c\xb0\xdc\x2a\x6d\x60\xa9\x55\x05\x78\x57\xb2\xc6\x58\x62\x44\x17\x0a\x5b\x33\x2e\x9c\x77\x39\x23\x83\xd2\xc0\xf2\xbc\xa9\x1a\x4a\x52\xe5\x0a\x50\xaa\x66\x55\xb6\xb2\x58\x05\x3e\x54\x09\x25\x57\xbd\x3c\xa6\x66\x15\x30\x6b\x59\x7e\x6b\x2e\xa1\xbb\x27\x80\x69\x04\xcb\xb1\xa0\x55\xb9\xaa\x2a\x25\xe1\x46\x17\x50\x33\x6d\xb7\x60\x0e\xb3\x0d\x96\xe7\x2e\xee\x45\xf0\x46\x6e\x95\x44\x28\xd9\xda\x49\x08\xef\x7d\x81\x41\x72\xfd\x85\xe5\x98\x29\xd5\x53\x43\xc5\xb6\xdd\x76\xad\xf4\x1b\x6e\x4b\xee\xd5\x53\xa3\xae\x68\x69\x01\x82\x57\xdc\x9a\x28\x89\xeb\xfd\x1d\xbb\x8f\xd6\x22\x2c\x95\xe6\xbf\x51\xaa\x23\x86\x17\xaa\x3d\xba\x6e\xba\xdb\xd2\xe1\x40\xe0\xd2\xce\xe0\x85\xbf\x2d\x8f\x91\xdd\xd6\x44\xe7\x60\xdd\xf1\x74\xb5\x26\x85\xa0\x19\xdc\xf8\x04\xd7\xa7\x16\x85\x1d\x48\x50\x1c\x81\xcf\x6f\xfa\xfa\x75\x7d\xd7\xcb\xd1\x67\xc9\xd3\x9e\x09\x21\xe0\x50\x29\x6b\xde\xab\xf1\x12\x2a\x76\x8b\xc0\x20\x61\x47\x35\x73\x2b\xb4\xab\xb8\xb8\xeb\x18\xc4\x76\x83\x68\xbf\x26\x67\x4e\x7f\xf0\x0c\xb9\x5c\x3d\xbb\x9e\x7a\x44\xd2\x03\xb1\x7f\x76\x3d\xe5\xd2\xaa\x67\xd7\xd3\xe9\xdd\xf4\x23\xff\x3d\xbb\x9e\x2a\xf9\xec\x7a\x6a\x4b\x7c\x76\x3d\x7d\x76\x7d\x33\xc4\xb2\x1f\x69\xdd\x81\x88\xd0\xd0\x66\x1d\xc2\x03\xb0\x4c\xaf\xd0\xa6\xc1\xff\xb2\x4c\x35\x76\x96\x09\x26\x6f\x83\x85\x93\x96\xd2\x0f\x07\x82\xf3\x09\x2b\xd4\xcc\x10\x22\x48\x60\x07\x92\xb6\x39\x62\x60\x6c\x1a\xad\x55\x23\x29\x4c\x02\x1d\xd9\xb9\xac\x1c\x11\xc8\x48\x2f\x93\x28\xc9\x74\xbc\x78\xab\xea\x6d\xe8\x98\xb8\xe5\x27\x5a\x34\x4d\x5d\x2b\x6d\xa3\xa1\x36\x19\x15\x46\x02\x4d\xfc\x7a\xfa\xf2\xf5\xab\x47\xc5\x37\x94\x76\xbb\x33\xf4\x12\xb2\x4c\xad\x11\x7c\x92\x9f\xa9\x3b\x60\xb2\x80\x25\xd7\x08\x6c\xc3\xb6\x5f\x25\x71\xe1\x4a\xb2\xcf\x07\xed\xb2\x75\xae\x7f\x2a\xd4\x76\x1e\x7f\x09\x75\x93\x09\x6e\x4a\x60\x20\x71\x03\x89\xb1\x5a\xc9\xd5\xc2\x8d\xe6\x54\xa3\xba\x57\xa8\x95\xb1\x8f\x99\x1f\xab\x0c\x8b\xe2\x0c\x00\xbe\x94\xfd\x37\x9b\x4d\xd4\x69\xd2\x19\xbf\x44\x51\xc7\x74\xfb\x35\x92\xdb\x6d\xec\xbd\x48\xc9\xf8\x6b\x5e\xa4\xd7\xaf\xaf\x5f\xbd\xba\x7e\xf1\x6f\xaf\x5f\xbe\xbc\x7e\xfd\xe2\xe5\x43\xc8\xa0\x43\x7d\x26\x30\x7c\x5e\xfd\xad\xa2\x32\xb6\x4f\xaa\x3d\x5e\xba\x64\x8e\x42\x76\x41\x45\x89\x0e\xfe\x61\x0c\x35\x92\x32\x93\x90\x89\xb3\x49\xc5\x27\xa0\xc8\xc1\xe8\x11\xc9\x3e\x13\x5a\x1d\x7c\x08\x29\xaa\xb1\x74\xc2\xae\xba\xe7\x4a\xf6\x70\xba\x04\xc3\xab\x5a\x6c\x21\xdf\x5b\xfd\x3c\xae\x1e\x34\xca\xef\xc2\xea\xd0\x6c\x1e\x64\x2e\x1d\xa8\x54\x81\x14\xf4\x4d\x63\x72\xac\x5d\xdb\x97\x02\xe9\x9f\xb6\xbf\x31\x69\xb9\xc4\x2e\xe0\x46\xf0\x9d\x14\x5b\x68\x0c\xc2\x52\x69\x28\x30\x6b\x56\x2b\x97\x25\x68\xa8\x35\x5f\x33\x8b\x5d\x94\x35\x2d\x2a\x7a\x50\x0c\x4a\x1d\xca\x81\xc4\x20\x25\xf9\x9b\x6a\x20\x67\x12\xac\x66\xf9\xad\xf7\x94\x46\x6b\xf2\x94\x1a\xfd\x69\xfa\x38\x9f\xa1\x50\x1b\x47\xe2\xcf\xbd\xe4\x28\x5c\xd0\x37\x88\x50\xaa\x0d\x54\x4d\xee\x1c\x92\x82\xba\x3b\xc4\x86\x71\x0b\x8d\xb4\x5c\x78\x7d\xda\x46\x4b\x4a\x11\xf0\x20\x48\x9f\x14\x83\x09\x56\x8b\xf7\x25\x9e\xc9\x88\xfa\x32\x0e\x34\xbe\xf5\xe4\x50\x6b\x65\x31\x27\x83\x02\x5b\x31\x2e\x0d\x59\xc4\xa5\x01\x58\x7d\x44\x99\xd7\x3f\xb5\x0f\xfb\x96\xa5\x9b\x8e\x63\xf8\xab\x50\x19\x13\xb0\x26\xa4\x67\x82\xf2\x3b\x05\xa5\xa2\xa3\x0f\xb4\x65\x2c\xb3\x8d\x01\xb5\x74\xa3\x5e\x72\x5a\xbf\x66\x9a\x2c\x88\x55\x6d\x21\x6d\x1b\x6e\x34\x66\x50\xaf\xdb\x36\x22\xbd\x52\x29\x7f\x30\xdf\x6b\x3d\x85\x9f\x7f\x99\x3f\x69\x45\xf9\x33\x2e\x1d\x24\x08\xdf\xfe\xc8\xb6\x64\x16\x72\x8d\xcc\xa2\x81\x5c\x28\xd3\x68\x2f\x61\xa1\x55\x0d\x24\x65\xc7\xa9\xe3\x4c\x13\xb5\xdb\xad\x63\x32\x2e\x99\x29\x27\x6d\xbf\x50\xa3\xb3\x52\x3f\xd7\x8d\x5f\x10\xea\xc6\xc4\x80\xa7\xd3\x39\xf0\xa4\xe3\x1b\x09\x94\x2b\x5b\xce\x81\x3f\x7f\xde\x13\x5f\xf0\x25\x8c\x3b\x8a\x9f\xf9\x2f\x91\xbd\x8b\x68\x17\x48\x53\x18\xee\xe6\x36\x6c\xf9\x98\x5a\xf0\x1c\xc7\xfc\x12\xae\x26\xf3\x6e\x36\xd3\xc8\x6e\xbb\xb7\xd6\x8e\xfe\x3f\xf7\x77\x37\x3f\xd4\x8c\x53\xfe\x81\x6e\x7c\x33\xc0\x00\x83\x15\x37\x16\x1a\x2d\xa0\xf5\x61\x6f\x82\xde\x20\x8e\x6e\xa8\x95\x13\x5c\xb6\x0f\x2d\xa6\xba\x23\x78\x36\x91\x41\x59\x8c\xff\xe3\xc7\xef\xbe\x8d\x8c\xd5\x5c\xae\xf8\x72\x3b\xbe\x6f\xb4\x98\xc1\xd3\x71\xf0\x2f\x8d\x16\xc1\xe4\xe7\xe9\x2f\xd1\x9a\x89\x06\x2f\x9d\xbd\x67\xee\xef\xc9\x2e\x97\xd0\x3e\xce\xe0\x70\xc3\xdd\x64\x32\x3f\xdf\x38\x19\xf4\x79\x34\x1a\xb4\x63\x22\xec\x81\x7f\xac\x23\x06\x15\xda\x52\x39\xd7\xd5\x98\x2b\x29\x31\xb7\xd0\xd4\x4a\xb6\x2a\x01\xa1\x8c\xd9\x03\xb1\xa3\x48\x4f\x41\xd1\xd2\xa7\x2e\x58\xff\x37\x66\x3f\xaa\xfc\x16\xed\x78\x3c\xde\x70\x59\xa8\x4d\x24\x94\xbf\x6a\x23\x72\x52\x95\x2b\x01\x69\x9a\x42\x1b\x45\x83\x09\x7c\x0d\xc1\xc6\x50\x3c\x0d\x60\x46\x8f\xf4\x34\x81\xe7\x70\xbc\xbc\xa4\x78\xff\x1c\x82\x98\xd5\x3c\x98\x78\x77\xe8\x14\xaf\x64\x85\xc6\xb0\x15\x0e\x05\x74\x85\x51\x0f\x32\x3a\x47\x65\x56\x90\x82\x33\x50\xcd\xb4\x41\x4f\x12\x51\x79\xde\xa1\x8d\x30\xeb\xc8\xd2\x14\x64\x23\xc4\x1e\xa4\xde\x29\xe6\x1d\xfc\x0e\xc8\x23\x1f\x6b\xbe\x4a\x53\xa0\x5a\x95\x54\x5c\xec\x57\x92\xf1\x7d\x55\x3d\x89\x28\x2e\xec\x57\x4c\xe6\x43\x34\x1f\x70\xc3\xe2\xf7\xd8\x61\x71\xcc\x0f\x8b\x07\x18\xba\x26\xc6\x63\xfc\x7c\xd3\x63\xc0\xce\x0d\x3c\xc0\x4d\x36\x55\x86\xfa\x31\x76\xbe\x89\xd1\xb2\x73\xaa\x7e\x27\xed\x60\xed\x25\x5c\xbd\x9a\x3c\xc0\x1d\xb5\x56\x0f\x32\x97\xca\x6e\xc7\xf7\x82\x6d\x29\x67\x82\x91\x55\xf5\x5b\xd7\x73\x18\x5d\xba\x88\x3b\x83\x9e\xc3\xa5\xeb\x26\xcf\x60\xe4\xde\x68\x9e\x57\xe8\x56\xbd\x9c\x4e\xa7\x97\xd0\x7d\x86\xf9\x13\x23\x27\xd4\x0d\xee\x1e\x90\xc7\x34\x79\x4e\x71\xff\x73\x24\x6a\x79\xf4\x32\xb5\xef\x9f\x21\x55\x1f\x1b\x0e\xc4\x82\x3f\xfc\x01\x4e\x66\x0f\x61\x1c\xc7\xf0\x5f\x8c\xaa\x70\x21\x5c\xf3\xc0\xf5\x0c\x7a\xfa\x8a\x1b\xe3\x6a\x71\x03\x85\x92\xd8\xae\xf9\xb4\x6b\xff\x44\xc6\x96\x0c\x16\x30\x3d\x16\x90\xae\xc3\x41\x58\x38\x13\x2d\x06\x7c\x0f\x03\xc1\xc5\x6e\xb8\xdf\xc1\x4a\x5e\x21\x7c\x95\x42\x10\x0c\x17\x9f\x50\x10\x41\xcf\xec\xc2\xa0\x7d\xef\x6d\x31\x6e\xa3\xe3\xb9\xd8\x35\xb9\x84\x9b\xe9\x74\x3a\x39\x11\x62\xb7\x57\xef\x9b\x9a\xd2\x26\x60\x72\xeb\xae\xc4\x5e\xb7\x2e\x71\xa4\x14\x88\xae\x34\x01\xb9\x12\xc2\xe7\x2c\xed\x52\x52\x70\xdb\x3b\x49\x21\xbc\x9a\x9f\x89\xa2\x03\x4d\x0e\x8e\x76\x6c\x9e\x33\xba\x3f\x36\xd1\xa1\xce\x8e\x88\xc3\xab\x03\xa3\x1c\xd8\xeb\xbc\x61\x2e\x7a\xb9\xf9\x5e\xa3\x47\xe6\xda\xdb\xeb\x58\x67\x03\xf9\x3d\x9f\xe7\x57\x1f\x79\x8c\x7e\xba\x6e\x4c\x39\x3e\x12\x74\x32\x3f\xb5\xcd\x3b\x8b\x9a\xb2\x64\x45\x21\x8b\x6c\x41\xa5\x80\xc6\x13\x93\xb8\x54\x5d\x63\xa8\x51\x16\xa8\xbb\x94\xc2\x67\xf6\x94\x00\x1e\x98\xcc\x57\x95\x43\x38\x7d\xa2\xc3\xb8\x94\x4c\x49\x04\x00\x38\x72\x02\x07\xd4\x03\xa4\x12\x31\x0a\x56\x1b\x2c\x20\x05\xff\x55\x7c\x3c\x89\x1a\xc9\xef\xc6\x93\xb0\x7d\x3f\xe6\xd1\xcd\xcf\xfb\x32\xb1\x13\xfb\x79\x0a\x41\x62\x35\xf0\x22\x1d\x05\xf0\xfc\x9c\x0b\x52\xd4\x1d\x2d\xf6\x12\x0c\x97\x02\x24\xb6\x58\xb8\xc6\xa8\xaf\xd7\xfe\x1e\x64\x2c\xbf\x5d\xb9\x42\x68\x46\xa9\xd6\xf8\x84\x2d\x5b\x33\xcb\xb4\xe3\x3a\x99\xc3\x9e\xbc\x2d\x14\x73\x32\xce\x1c\x7c\x45\xea\xfa\xaf\xd0\x7f\xb3\x70\x6f\x99\xd2\x05\xea\x50\xb3\x82\x37\x66\x06\x2f\xea\xbb\xf9\xdf\xbb\x6f\x3a\xae\x4b\xfc\xa8\xa8\xb5\xc6\xc5\x89\x44\x6d\x93\xf1\x39\x04\x49\x4c\x04\xbf\xc7\xa6\x3f\xec\xf0\x6b\x3c\x9c\xe9\x85\x43\xff\xad\xbc\x1d\xaf\x78\x51\x08\x24\x81\xf7\xec\xc9\x19\xc9\xfe\x43\x97\x3a\xdc\x12\xda\x26\xf8\x7e\xcd\x0e\x50\x18\x7c\x64\x41\xdf\x4f\x1f\x11\x00\x42\x3a\x32\x77\x3a\x6f\x8b\x6d\x37\xac\x47\x4e\x17\xed\x6f\x2b\x8a\x46\xbb\x5c\x6b\x1c\xb6\x00\xbb\x84\x91\xa1\xdc\xaf\x30\xa3\x49\x54\x36\x15\x93\xfc\x37\x1c\x53\x5c\x9a\x78\x5d\xb9\x06\x7d\x70\x7a\x25\x9f\x08\xb3\xef\x9c\x8f\xba\x18\x37\x6a\x95\x38\xea\xac\xfb\x62\x5f\xdb\xcf\x60\x3a\x1f\x7d\xa2\x86\xce\xef\x12\x66\x4c\xc3\xf0\x25\xec\x82\x2f\x68\x45\xbb\x77\x73\x19\xd3\x23\xdf\xc9\x70\xf9\xb9\x54\x9b\x74\x74\x33\xed\x85\xf4\x86\x76\x76\x1e\xb5\x58\x3b\x31\x06\x49\xd9\xb9\xe6\x02\x6e\xa6\x5f\x42\x5a\xdf\x0d\x39\x3a\x81\xd5\xbc\xc6\x02\x58\x6e\xf9\x1a\xff\x1f\x0e\xf2\x05\x94\xfc\xc9\x22\x12\x0e\x3b\xe5\x39\x98\x1e\xc8\x4b\xb3\xbd\x6e\xff\x95\xfc\x0d\x62\xa7\xe1\xe7\x10\x9c\x3d\xc8\x83\x48\x3c\x22\x3c\x72\xed\x87\xfd\xde\x7d\x71\x0a\x8e\x63\x0a\x65\xbb\xfd\xd7\xd2\x49\x54\xda\x4a\x8c\x83\xc4\xba\x5f\xcd\x90\xcc\x3d\x07\xc7\xc0\x0f\x1f\xa6\x74\xbb\xc3\x42\x86\xea\x77\x3c\xaa\xb3\x60\x90\x9c\xf4\xb5\x58\x97\x89\xc0\x6e\xff\xe3\xa2\x38\x86\x1f\x2d\xd3\x16\x18\xfc\xf4\x0e\x9a\xba\x60\xd6\x7f\xc9\xa1\xf8\xe8\xbf\x94\x74\xbf\x3e\xca\x98\x36\xb0\x54\x7a\xc3\x74\xd1\xf6\x67\x6c\x89\x5b\xf7\x25\xa7\x4b\xfd\x0c\xda\x77\x74\x8b\xad\x99\x18\x9f\xd4\x7d\x4f\xc7\xa3\x68\x68\xf2\xd1\x24\x42\x96\x97\xa7\x84\x2e\x62\xf5\xfb\xa6\xf0\xad\x2b\x01\xc6\x4f\xc7\xb6\xe4\x66\x12\x31\x6b\xf5\x78\x74\x00\x86\xd1\x84\xec\x7a\x35\x28\xc9\xfa\xe5\xc9\x81\x5b\x3d\xc6\x63\x9f\x4c\xf7\x89\x40\x47\x9e\x1b\x33\xf6\xb8\x1a\x5d\x0e\x78\x1f\xc2\x6a\xf4\x6c\xd4\x1b\x6a\xef\xde\xfb\x73\xa4\x67\x25\x39\x60\x3d\x22\x2f\x1b\x9d\x6c\xcf\x8a\xe2\x2d\xf9\xcf\x38\x38\xe3\xe9\xc7\xe8\x98\xf4\xca\xf6\xf7\xf5\xa3\x5a\xf6\xbf\xd3\x78\x40\xc5\xbc\x18\x4d\x22\xd3\x64\xbe\x37\x31\x7e\xd9\x17\x60\x1d\x99\x03\xef\x71\x28\x38\x49\x28\x68\x8b\xc3\xa4\x22\x3c\x4a\x42\x1e\x89\x1a\xed\x96\xfe\x54\xbb\x4b\x52\xf8\x74\xd2\xb7\xb6\xbe\x31\x94\x5c\xf9\xd6\xff\x06\x33\xe3\x3a\x09\xd0\xe2\xdd\x75\x73\x7c\xd7\xe6\xcd\xf7\xef\x06\x9d\x9b\xde\x23\xc6\x8e\x7b\xff\xc3\xc0\x73\x7d\x92\xb3\xbf\x44\xdc\x6c\x36\xd1\x4a\xa9\x95\xf0\xbf\x41\xec\x1b\x29\x31\xab\x79\xf4\xc1\x04\xc0\xcc\x56\xe6\x50\xe0\x12\xf5\x62\xc0\xbe\xed\xae\x24\xb1\xff\x8d\x5c\x12\xfb\x9f\x01\xff\x5f\x00\x00\x00\xff\xff\xd5\x72\xf5\x28\x17\x2c\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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
	"faucet.html": faucetHtml,
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
	"faucet.html": {faucetHtml, map[string]*bintree{}},
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
