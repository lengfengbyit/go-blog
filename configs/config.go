// Code generated for package configs by go-bindata DO NOT EDIT. (@generated)
// sources:
// configs/config.yaml
package configs

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

// Mode return file modify time
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

var _configsConfigYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x52\x4b\x6e\xdb\x30\x10\xdd\xeb\x14\x03\x74\xdb\xa8\x94\xe5\x2f\x57\x4d\x63\x07\x4d\x11\xb7\x46\x25\x23\xcb\x62\x6c\x8d\x19\x05\x94\x48\x93\x94\x23\xf7\x0e\xd9\x14\xe8\x19\x7a\x8a\xde\xa6\x45\xaf\x51\x90\x72\x6c\x17\xcd\x6e\xf0\x66\xf8\xe6\xbd\x37\xcc\xc8\xec\xc8\xf0\x08\xe0\x73\x53\xcf\x55\x41\x1c\x0a\x5a\x35\x02\xe0\x55\x57\xbc\x06\x43\x92\xd0\x52\x04\xf0\xde\x39\xbd\x50\xc6\x71\x18\xb3\x31\xf3\x6f\x08\x8b\xbc\xac\x48\x35\x8e\xc3\xd0\x23\x77\xa6\x74\x74\x0e\x5d\x6a\xed\xd9\xa7\xb4\xc1\x46\xba\x05\x0a\xca\xca\xaf\xc4\x21\xf1\xd3\x73\x6c\xcf\x11\x0f\xdd\x2a\x91\xe1\x8e\x16\xe8\xee\x39\x58\xa7\x0c\x0a\x7a\x23\x95\xb0\x5d\xef\xba\x94\xf4\x11\x2b\xe2\x80\x5a\x9f\xa0\x59\xeb\x38\xc4\x52\x89\x0e\x9a\xa2\xa3\xac\xd9\x6c\xca\x96\x83\x33\x8d\xd7\x7e\xa5\x6a\x47\xad\x3b\x4a\x4b\xbc\xc3\xdf\x4f\xdf\x7f\xfd\x7c\x82\x3f\x3f\xbe\x45\x53\x74\xb8\x42\x4b\x41\xec\xbb\x7c\xaf\x89\x43\xb5\xb7\x5b\x19\x01\x2c\x2d\x99\x6e\xa9\x51\xca\x45\x00\x0b\xb4\xf6\x51\x99\x82\x43\xd2\x4b\xfb\x83\xa1\xcf\x46\x59\xcf\xda\x1b\xc5\x2c\x66\x71\xe2\x87\x42\x52\x69\xca\x06\x81\xb2\x23\x58\x49\x25\xbe\x58\x32\xbb\x72\xed\x55\xe5\xb8\x92\xb4\x30\x14\x94\x86\x9e\x97\x7a\x8f\xc6\x92\xe3\xd0\xb8\xcd\x38\x6c\x33\x36\x64\xca\x21\xef\xcc\xcc\xb1\xbd\x29\x24\x5d\xa9\xba\xb6\xa7\x28\x3f\x69\xaa\x0f\x50\xca\xa2\xa5\x96\x0a\x0b\x6f\xe7\xff\x3c\x9b\xd0\xf3\x91\x76\xf7\x5f\x1a\xe9\x9b\xe8\xca\x75\x04\x70\x53\xa1\xa0\x39\xb6\x67\x87\x0a\xd0\xa5\x94\xea\x71\xd6\x3a\xeb\x49\x01\x2e\x40\x94\x9b\x43\xa5\x6b\x71\xa8\x1e\x34\x9d\x4a\x11\x7d\xb8\xcb\x83\x04\x5a\x1b\xef\x88\x8a\x62\xbf\x7e\xd8\x7b\x46\x6b\x1b\x32\x9d\xe9\x8b\x53\x20\xb3\x56\x97\x86\x38\x8c\x7a\x8c\x45\xb3\x0a\x4b\xc9\x8f\xe1\xda\xca\xe9\x78\xbb\x8d\xd7\xaa\x3a\xc6\xdb\x1f\x0e\xfe\x39\x50\xc2\x92\xc9\x24\xe9\x0d\x26\x93\xb7\xa7\xc9\xe7\x6b\x85\xb5\x59\x76\x7b\xfc\x14\xd7\x46\x55\x2f\xbf\xc9\xd5\xb3\xc9\x24\xed\x8f\x7b\x6c\x94\xb2\xf1\xa1\xfb\x37\x00\x00\xff\xff\x4f\xd2\x33\x16\x33\x03\x00\x00")

func configsConfigYamlBytes() ([]byte, error) {
	return bindataRead(
		_configsConfigYaml,
		"configs/config.yaml",
	)
}

func configsConfigYaml() (*asset, error) {
	bytes, err := configsConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "configs/config.yaml", size: 819, mode: os.FileMode(420), modTime: time.Unix(1606280069, 0)}
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
	"configs/config.yaml": configsConfigYaml,
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
	"configs": &bintree{nil, map[string]*bintree{
		"config.yaml": &bintree{configsConfigYaml, map[string]*bintree{}},
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
