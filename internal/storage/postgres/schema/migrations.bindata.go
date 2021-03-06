// Code generated by go-bindata.
// sources:
// migrations/1551106041_init.down.sql
// migrations/1551106041_init.up.sql
// DO NOT EDIT!

package schema

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

var __1551106041_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xf0\x74\x53\x70\x8d\xf0\x0c\x0e\x09\x56\x48\x2c\x2a\xc9\x4c\xce\x49\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xc8\xa2\xce\x28\x1f\x00\x00\x00")

func _1551106041_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1551106041_initDownSql,
		"1551106041_init.down.sql",
	)
}

func _1551106041_initDownSql() (*asset, error) {
	bytes, err := _1551106041_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1551106041_init.down.sql", size: 31, mode: os.FileMode(420), modTime: time.Unix(1551381926, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1551106041_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcc\xb1\x4e\xc3\x30\x10\xc6\xf1\xf9\xee\x29\x6e\x6c\xab\x48\x65\x61\x62\x3a\xc2\x55\x58\x38\xa1\x3a\x5f\x50\x3b\x55\xa6\xf6\x60\xa9\x11\x51\x62\x06\xde\x1e\x89\x21\x62\xee\xfa\xe9\xfb\xfd\x5b\x15\x36\x21\xe3\x67\x2f\xe4\x0e\xd4\xbf\x1b\xc9\xc9\x05\x0b\x14\xe7\x5a\xae\xb7\xbc\xd0\x06\xa1\x24\x02\x08\xa2\x8e\x3d\x1d\xd5\x75\xac\x67\x7a\x93\x73\x83\x50\x4b\xbd\x65\x80\x0f\xd6\xf6\x95\x95\x36\x8f\x0f\xdb\xbf\x48\x3f\x78\xdf\x20\x7c\x7e\xa5\x1f\x00\x93\x93\xfd\x5b\x11\xf6\x3b\xaa\x65\xcc\x4b\x8d\xe3\xb4\xd0\x6e\x8f\x70\x9d\x73\xac\x39\x5d\x62\x05\x73\x9d\x04\xe3\xee\xb8\x12\x7a\x91\x03\x0f\xde\xa8\x1d\x54\xa5\xb7\xcb\x7a\x69\x10\xbe\xa7\x74\x8f\xc4\xed\x13\x22\xfe\x06\x00\x00\xff\xff\x84\x79\xff\x55\x00\x01\x00\x00")

func _1551106041_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1551106041_initUpSql,
		"1551106041_init.up.sql",
	)
}

func _1551106041_initUpSql() (*asset, error) {
	bytes, err := _1551106041_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1551106041_init.up.sql", size: 256, mode: os.FileMode(420), modTime: time.Unix(1551200047, 0)}
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
	"1551106041_init.down.sql": _1551106041_initDownSql,
	"1551106041_init.up.sql": _1551106041_initUpSql,
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
	"1551106041_init.down.sql": &bintree{_1551106041_initDownSql, map[string]*bintree{}},
	"1551106041_init.up.sql": &bintree{_1551106041_initUpSql, map[string]*bintree{}},
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

