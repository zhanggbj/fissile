// Code generated by go-bindata.
// sources:
// scripts/dockerfiles/Dockerfile-base
// scripts/dockerfiles/Dockerfile-role
// DO NOT EDIT!

package dockerfiles

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _scriptsDockerfilesDockerfileBase = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x8c\x3d\x0a\x83\x40\x10\x85\xfb\x39\xc5\x40\xc0\x4e\xbd\x42\x0c\x26\x60\xa1\x81\x25\x39\xc0\xb2\x19\xd7\x01\xf7\x27\xee\x58\x04\xf1\xee\x49\x65\xb6\x78\xc5\xf7\xde\xe3\xbb\xa9\x7b\x8f\xdb\x86\xd5\x45\x27\xea\x9c\xb6\x84\xfb\x0e\xd0\x37\xdd\xf0\xf8\xe5\xaa\x70\x32\xe3\x79\x8a\x54\x99\xe0\x00\x4e\xd8\xf9\x24\x7a\x9e\x31\x2e\xb4\xd0\x7b\xe5\xc4\x42\x29\x1f\x5c\xf0\x2c\x00\xea\x39\xa0\x8e\x52\x5a\x12\x5c\xe3\x4b\x0b\x61\x51\x1c\x0d\xe7\x67\x2c\x3f\xb9\xc0\x04\x3f\xb2\xb5\xec\x01\x9a\xb6\xfd\x23\xd6\x21\x4a\x7d\x60\x0d\xdf\x00\x00\x00\xff\xff\x9e\x78\xad\x05\xbc\x00\x00\x00")

func scriptsDockerfilesDockerfileBaseBytes() ([]byte, error) {
	return bindataRead(
		_scriptsDockerfilesDockerfileBase,
		"scripts/dockerfiles/Dockerfile-base",
	)
}

func scriptsDockerfilesDockerfileBase() (*asset, error) {
	bytes, err := scriptsDockerfilesDockerfileBaseBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/dockerfiles/Dockerfile-base", size: 188, mode: os.FileMode(420), modTime: time.Unix(1444296679, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsDockerfilesDockerfileRole = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x72\x0b\xf2\xf7\x55\xa8\xae\x56\xd0\x0b\xca\xcf\x49\x75\x4a\x2c\x4e\xf5\xcc\x4d\x4c\x4f\x55\xa8\xad\xe5\xe2\xf2\x75\xf4\xf4\x0b\x01\x62\xd7\x20\x85\x8c\xe4\x34\x87\x8c\x82\x54\xbd\xe4\xfc\x5c\x2e\x2e\x1f\x47\x27\x57\x1f\x05\xa5\x22\xa0\x06\x25\x5b\xa5\xea\x6a\xb0\xd6\xda\x5a\x25\xa0\x50\x6a\x4e\x2a\xd0\x08\xdd\xb2\xd4\xa2\xe2\xcc\xfc\x3c\xa8\x2c\x44\x50\x2f\x0c\x22\x08\x56\x88\xa2\x00\x21\xc1\xc5\xc5\x05\x08\x00\x00\xff\xff\x64\x0f\x2d\xe9\x90\x00\x00\x00")

func scriptsDockerfilesDockerfileRoleBytes() ([]byte, error) {
	return bindataRead(
		_scriptsDockerfilesDockerfileRole,
		"scripts/dockerfiles/Dockerfile-role",
	)
}

func scriptsDockerfilesDockerfileRole() (*asset, error) {
	bytes, err := scriptsDockerfilesDockerfileRoleBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/dockerfiles/Dockerfile-role", size: 144, mode: os.FileMode(420), modTime: time.Unix(1444291273, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"scripts/dockerfiles/Dockerfile-base": scriptsDockerfilesDockerfileBase,
	"scripts/dockerfiles/Dockerfile-role": scriptsDockerfilesDockerfileRole,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"scripts": &bintree{nil, map[string]*bintree{
		"dockerfiles": &bintree{nil, map[string]*bintree{
			"Dockerfile-base": &bintree{scriptsDockerfilesDockerfileBase, map[string]*bintree{
			}},
			"Dockerfile-role": &bintree{scriptsDockerfilesDockerfileRole, map[string]*bintree{
			}},
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

