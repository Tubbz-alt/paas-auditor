// Code generated by go-bindata.
// sources:
// sql/001_create_to_seconds_function.sql
// sql/002_create_iso8601_function.sql
// sql/010_create_app_usage_events.sql
// sql/020_create_service_usage_events.sql
// sql/025_create_pricing_plans.sql
// sql/026_create_pricing_functions.sql
// sql/027_create_validate_formula_trigger.sql
// sql/030_create_billable.sql
// sql/040_create_billable_range_function.sql
// DO NOT EDIT!

package db

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

var _sql001_create_to_seconds_functionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xcd\x3d\x4b\x43\x31\x14\xc6\xf1\xd9\x7c\x8a\x67\xc8\x90\x0c\x77\x70\x6d\xa7\x18\x62\xb9\x90\xa6\x9a\x97\x59\xae\xf1\xd4\x0a\xed\xcd\x25\x49\x51\xfc\xf4\x22\xb5\xd2\xed\xc0\x79\xf8\xfd\xd9\x30\xa0\x97\x97\x46\xb9\xcc\x6f\x0d\xf4\xd5\xeb\x94\x7b\x43\x3f\x10\xe6\xf3\xe9\x95\x2a\xca\x1e\xd7\xf7\xbe\x96\x13\x26\xf4\xd6\xbf\xeb\x34\xbf\x13\xd3\xde\xa8\x68\xb0\xf3\xf0\xe6\xc9\x2a\x6d\xf0\x98\x9c\x8e\xe3\xce\xdd\xa8\xe2\x7f\x2f\xe1\x4d\x4c\xde\x85\x5f\x9b\xea\x47\x86\x0a\xe0\x9c\xdd\x35\x3a\x52\xee\xd7\xbc\xa0\xa5\xe4\xc3\xa5\x26\xce\xcb\x42\x55\xf0\x7b\x89\x01\xc7\xf2\x79\xb9\xa5\x5c\xad\xfe\x88\x35\xe3\x1c\x56\xb9\x4d\x52\x1b\x83\xf0\x6c\x31\x6e\xb7\x29\xaa\x07\x6b\x10\xa2\x1f\x75\x5c\xb3\x9f\x00\x00\x00\xff\xff\xec\xd8\x3f\xd6\xe6\x00\x00\x00")

func sql001_create_to_seconds_functionSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql001_create_to_seconds_functionSql,
		"sql/001_create_to_seconds_function.sql",
	)
}

func sql001_create_to_seconds_functionSql() (*asset, error) {
	bytes, err := sql001_create_to_seconds_functionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/001_create_to_seconds_function.sql", size: 230, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql002_create_iso8601_functionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\xcc\x41\x4b\xc3\x30\x18\x87\xf1\xb3\xfd\x14\x7f\xca\xa0\x1b\x58\x50\x11\x91\xf5\x14\xd3\xd8\x15\xd6\x54\xd2\xe4\x50\x2f\x12\x66\x36\x0b\x5b\x52\x93\x77\x20\xfb\xf4\x32\x76\x7f\x9e\x5f\x59\x62\x7f\xf6\x3b\x9a\x82\xc7\x3e\x44\x44\x37\x47\x97\x9c\xa7\xc9\x1f\xf0\x6d\xc9\x25\xd8\x84\x76\xe8\xf1\xfa\xf2\xf0\x08\xa3\x79\xc6\x95\x60\x5a\xa0\x57\x50\xe2\x63\xcb\xb8\xc0\xbb\x91\x5c\xb7\xbd\xc4\x94\xc2\x35\x5b\x12\x68\x3a\xb9\x44\xf6\x34\xd3\x65\x85\xe8\xe8\x1c\x7d\x02\xb9\x3f\x02\x1b\xb0\x58\xe0\x4d\x34\xad\xcc\xee\x94\xd0\x46\x49\x50\xf8\xda\xfd\xd8\xb8\x24\xd8\xdb\x8a\x4b\xf0\x0e\x85\xd1\xbc\xb8\x47\x31\x8e\xe3\x58\x76\x5d\x59\xd7\xb9\xce\x37\x9b\xa7\xe7\x75\xd7\xae\x87\x21\xff\xcc\x8b\x55\x95\x09\x59\x57\x57\x72\xcb\x64\x63\x58\x23\x30\x1f\xe7\x43\xfa\x3d\x56\xd9\x7f\x00\x00\x00\xff\xff\xf6\x3f\xf6\x71\xdd\x00\x00\x00")

func sql002_create_iso8601_functionSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql002_create_iso8601_functionSql,
		"sql/002_create_iso8601_function.sql",
	)
}

func sql002_create_iso8601_functionSql() (*asset, error) {
	bytes, err := sql002_create_iso8601_functionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/002_create_iso8601_function.sql", size: 221, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql010_create_app_usage_eventsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\x41\x6b\x83\x30\x18\xc6\xf1\x73\xf2\x29\xde\x9b\x06\xdc\x69\xb0\x8b\x20\x44\xf7\x8e\x65\x68\xdc\x4c\x04\x6f\x21\x2c\x41\x72\xd8\x10\x13\xb7\x7e\xfc\xd2\x52\x68\x29\x5e\x7a\x7f\xf8\x3f\xbf\x66\x40\xae\x11\x34\xaf\x5b\x04\xf1\x06\xb2\xd7\x80\x93\x50\x5a\x81\x5d\x16\xb3\x45\x3b\x7b\xe3\xff\xfc\x6f\x8a\x90\x53\x12\x1c\x28\x1c\x04\x6f\x0b\x4a\xe6\x2d\x38\x68\xde\xf9\x90\x3f\xbf\x30\x18\xa5\xf8\x1a\xb1\xa0\xe4\x7b\xf5\x36\x79\x67\x6c\x02\x2d\x3a\x54\x9a\x77\x9f\x05\x25\xab\xfd\x37\x3f\x3e\x9e\x7a\xf0\xa1\x7a\x59\x53\x56\xd2\xcb\xbb\x90\xaf\x38\xdd\xbd\x07\x77\x30\x57\x41\x70\xd0\xcb\x1d\x51\x70\x8f\x54\x62\xb2\xc9\xef\x87\x20\xbf\x01\x3e\x55\x55\x76\xde\x66\x0c\x58\x49\x8f\x01\x00\x00\xff\xff\xaa\xf2\x69\x45\x24\x01\x00\x00")

func sql010_create_app_usage_eventsSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql010_create_app_usage_eventsSql,
		"sql/010_create_app_usage_events.sql",
	)
}

func sql010_create_app_usage_eventsSql() (*asset, error) {
	bytes, err := sql010_create_app_usage_eventsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/010_create_app_usage_events.sql", size: 292, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql020_create_service_usage_eventsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x8e\xc1\x4a\xc4\x30\x14\x45\xd7\xc9\x57\xbc\xdd\xb4\x50\x57\x82\x1b\x61\x20\x33\x3e\x31\x32\x4d\xb5\x49\xa1\xbb\x10\x9a\x47\xc9\xc2\x22\x4d\x5a\xf5\xef\x45\xdb\x85\x0e\x74\xd1\xfd\xbd\xe7\x9c\x73\x8d\xc2\x20\x18\x71\xba\x20\xc8\x47\x50\x95\x01\x6c\xa5\x36\x1a\x22\x8d\x73\xe8\xc8\x4e\xd1\xf5\x64\x69\xa6\x21\x45\xc8\x38\x0b\x1e\x34\xd6\x52\x5c\x0a\xce\xfa\x29\x78\x38\x3f\x89\x3a\xbb\xbd\xcb\xa1\x51\xf2\xb5\xc1\x82\xb3\x6e\x24\x97\xc8\x5b\x97\xc0\xc8\x12\xb5\x11\xe5\x4b\xc1\xd9\xe8\x3e\xec\x1b\xc5\x1f\x1e\x3c\xeb\x4a\x9d\x78\x7e\xcf\xd7\x02\xa9\x1e\xb0\xbd\x2a\x08\xfe\xd3\xfe\xaf\x08\x1e\x2a\xb5\x51\x16\xfc\x5e\x5a\x4c\x2e\xd1\x36\x10\xb2\x3f\xc1\x37\xc7\xe3\xe1\x77\x7f\xc8\x61\xaf\x27\x7d\xbd\xef\xd2\xac\xb3\x30\xc4\xe4\x86\x6e\xf9\x2f\xda\xef\x00\x00\x00\xff\xff\xc6\x29\xc3\x37\xaf\x01\x00\x00")

func sql020_create_service_usage_eventsSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql020_create_service_usage_eventsSql,
		"sql/020_create_service_usage_events.sql",
	)
}

func sql020_create_service_usage_eventsSql() (*asset, error) {
	bytes, err := sql020_create_service_usage_eventsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/020_create_service_usage_events.sql", size: 431, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql025_create_pricing_plansSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x51\x6b\xb3\x30\x18\x85\xaf\xcd\xaf\x38\x97\x0a\x0a\xdf\x7d\xe1\x83\xae\xcb\x98\xd4\xd9\xcd\x2a\xb4\x57\x21\x68\xea\x02\x49\x94\x18\x47\xb7\x5f\x3f\x74\x76\xae\xdd\xc6\x9a\xab\x10\xde\xf7\x21\xe7\x3c\x51\x84\xd6\xca\x52\x9a\x1a\xad\xe2\xa6\x43\xd9\x18\xc7\xa5\xc1\xa1\xb1\xba\x57\xbc\x1b\x2e\x28\xb9\x2a\x7b\xc5\xdd\x38\x66\x65\x29\x3a\xb2\xca\xe8\x32\xa7\xc8\x97\x37\x09\x45\x7c\x87\x74\x93\x83\xee\xe2\x6d\xbe\x3d\x01\xd9\x08\xf4\x89\x27\x2b\x74\xc2\x4a\xae\xf0\x98\xc5\x0f\xcb\x6c\x8f\x35\xdd\x87\xc4\x33\x5c\x0b\x38\x71\x74\xe3\x72\x5a\x24\x49\x48\xbc\x17\xae\x64\xc5\x0e\xb6\xd1\x70\x52\x8b\xce\x71\xdd\xba\xb7\xaf\x13\x03\x96\xd5\xbd\xac\xbe\xed\x4e\x7f\xbe\x7c\x27\xde\xea\x9e\xae\xd6\xf0\x95\x30\xb5\x7b\xf6\x9d\x95\xda\xff\xa4\x04\x01\xfe\xe3\x5f\x10\xe2\xb7\x13\x45\x30\x8d\x83\xd0\xad\x7b\xfd\x11\x35\xe4\xf8\x93\x72\x1d\x6a\x4a\x30\xd1\xae\x46\x05\x0b\x72\x12\x52\xa4\xf1\x53\x41\x11\xa7\xb7\x74\x77\xe1\x45\x56\x47\x76\xe6\x86\x7d\x94\xcd\x4d\xc5\xe6\x52\x37\xe9\xb9\x40\xf8\xb3\x92\x10\x73\x6d\x0b\xf2\x1e\x00\x00\xff\xff\xdd\x23\x2b\xd5\x3c\x02\x00\x00")

func sql025_create_pricing_plansSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql025_create_pricing_plansSql,
		"sql/025_create_pricing_plans.sql",
	)
}

func sql025_create_pricing_plansSql() (*asset, error) {
	bytes, err := sql025_create_pricing_plansSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/025_create_pricing_plans.sql", size: 572, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql026_create_pricing_functionsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x4f\x6b\xe3\x3c\x10\xc6\xcf\x9a\x4f\x31\x07\x83\x6c\x70\xa0\xef\x7b\x8c\xd9\x83\x9b\x6a\x43\x20\x75\x17\x37\xb9\x2d\x04\x45\x99\xa6\x02\xc9\x72\xf5\xa7\xcd\x2e\xfd\xf0\x8b\x9b\xd8\x84\xee\xb2\xac\x2f\xd6\xcc\xe8\x99\x79\xf4\x9b\xd9\x0c\xe9\x44\x0a\x25\x3e\x39\x6f\x93\x91\xf8\xa6\xe3\x33\xbe\x4a\xaf\xe5\xde\x50\xc0\x90\xf6\x21\xea\x98\x22\x1d\x60\xd1\x8a\x7a\x23\xf0\xa1\xc5\x56\x7c\x5b\xd7\x0b\x81\x5f\xb7\xcd\x62\xb3\x7a\x68\x90\x5e\xa5\xd9\x5d\x5a\xe4\xc0\xec\x1e\xbb\x64\xc9\x6b\x55\x02\x3b\x24\x2f\xa3\x76\x1d\xc6\x10\x7f\x7a\xd9\x1d\xa9\x04\x36\x8e\x8b\x74\x8a\x50\xa0\xa7\x98\x7c\x17\x46\x15\xd6\x8f\x98\x65\x70\x27\x16\xeb\xba\x15\xc0\x5c\x8a\x63\xa9\x82\x5b\xb1\x5c\x35\xc0\x06\xdf\x29\x12\x2a\x67\x7b\x6d\x68\x9a\x7e\xf9\x17\xa8\xbb\xe8\x70\x50\xa6\xa0\xbb\x23\x30\xa6\x9c\x34\x14\x14\xe5\x76\x5f\xe2\x4d\x51\x02\x63\x74\x8a\x5e\xaa\x98\x53\xef\xd4\x33\x3e\x79\x67\x31\x4f\x7d\x4f\x3e\x1f\x5d\x17\x38\x43\xe3\xde\xae\x33\x45\x51\x01\x3b\x3b\x1e\xfa\x57\x20\x9a\xbb\x0a\xb3\x0c\xd7\x75\xb3\xdc\xd6\x4b\x81\xbd\xe9\x8f\xe1\xc5\xe0\xea\xfe\x7e\xbb\xa9\x6f\xd7\xa2\x02\x00\x98\xcd\x46\xb3\x13\xed\x0f\x8f\xe1\xc5\xfc\x8d\xed\xe7\x07\xe2\x35\x3b\x2c\xb0\x15\x9b\x6d\xdb\x3c\x9e\xc3\x3f\x80\x1b\xf2\x13\xb5\x21\x31\xff\x82\x79\x20\x43\x2a\x02\x63\xfc\x7c\xc2\x9c\x03\x63\xec\xfd\x1d\x3d\x1d\xe9\xd4\xef\x3c\xf5\x46\x2a\xca\x3f\x85\x13\xc3\x33\x93\x11\x76\x89\xfc\x86\x7f\x10\x65\x8c\x7f\xcf\x2c\x59\xe7\x7f\xec\x74\xb7\xb3\x7b\x5e\xe2\xf0\xf1\x3c\xfb\x6f\x3e\xbf\x2c\xb1\xe0\x25\xf2\xe3\x95\x20\x6a\x4b\xc3\xf5\x40\xca\x75\x87\x30\x68\x78\x9e\xfd\xff\x9b\xe0\xe2\x91\x17\x53\xa5\xe2\xc0\xfe\x75\x1f\x15\xfc\x0a\x00\x00\xff\xff\x3f\x83\x71\xf1\xf0\x02\x00\x00")

func sql026_create_pricing_functionsSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql026_create_pricing_functionsSql,
		"sql/026_create_pricing_functions.sql",
	)
}

func sql026_create_pricing_functionsSql() (*asset, error) {
	bytes, err := sql026_create_pricing_functionsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/026_create_pricing_functions.sql", size: 752, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql027_create_validate_formula_triggerSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x95\xd1\x4b\xe3\x4a\x18\xc5\x9f\x3b\x7f\xc5\xb9\xd8\x4b\x12\x6d\xbc\xfa\x78\x1b\xe4\x52\xd3\xa9\x06\x24\x29\xd3\x14\x85\xeb\x6e\x88\xed\x67\x1c\x9c\x4c\xb2\xc9\xc4\xea\xe2\x1f\xbf\xb4\xb6\x52\xab\xc2\xb2\x9b\x7d\x99\x96\x61\xce\xef\x7c\x5f\x1e\xce\x71\x5d\x2c\xee\xa4\x21\x25\x6b\x83\x42\xab\x27\xe6\x0b\x3e\x88\x39\x22\x01\xc1\xc7\x17\x03\x9f\x63\x34\x0d\xfd\x38\x88\x42\x3c\xa4\x4a\xce\x53\x43\xc9\x6d\x51\xe5\x8d\x4a\x6d\x07\x82\xc7\x53\x11\x4e\x60\x2a\x99\x65\x54\x61\x30\x41\xb7\xcb\x86\xdc\xbf\x18\x08\xce\x3a\x52\xaf\x34\x1b\x01\x0c\x3d\x1a\x8f\x75\xa4\x52\x94\xa5\x2a\x31\xc5\x3d\xe9\xcd\xe5\xbc\xc9\xf3\xa7\xa4\xac\xe4\x8c\xa0\x9b\x9c\x2a\x39\xf3\xd8\x29\x3f\x0b\xc2\xf7\x9c\xfe\x09\x54\xb1\xa0\xca\x0e\xf9\xe5\xe1\xfa\xd2\xf1\x3e\x7c\x67\xd7\xa4\x68\x66\x50\x51\x46\x8f\x65\x52\x51\xa9\xd2\x19\xd9\x3b\x2f\x7b\xb0\xfa\x7d\x5b\x6a\x43\x19\x55\xcf\x37\x32\x93\xda\x3c\xaf\xa7\x70\xac\x1e\xac\xbd\xe5\x91\x59\xce\xef\xba\xd8\xff\x1f\xb9\xff\x7e\x39\x70\xfe\xbb\x3e\xdc\xfc\xfd\x13\xf8\x36\x99\xd7\xdd\x9c\xf2\xa2\x7a\x4a\xa4\x4e\xf2\x9b\x76\xc9\x46\xe6\xb4\xe4\xd6\x34\x2b\xf4\xbc\x6e\x15\x6e\x3f\x5f\xb7\xfb\x1d\xf6\x5b\xa5\xb9\xad\xd2\x0e\x5a\xa5\xfd\xd3\x2a\xed\x6b\xab\xb4\xba\xd5\x55\xf7\x56\xb4\x2d\x58\x30\xc2\xee\x33\xfc\x75\x02\xcb\x72\x10\x9f\xf3\x90\x75\x76\xd2\x6b\xcb\x6c\x1f\xb7\x55\x91\x6f\x3c\xeb\x52\x49\x93\x98\x22\x31\xe9\x8d\xfa\x6c\x13\x07\x4a\xe6\xd2\xe0\x78\x69\xdd\x11\x83\x60\xc2\xc1\xaf\x7c\x3e\x5e\x05\xae\xb5\xb6\xc2\x8b\x95\xd4\x58\xab\xfb\xf8\xdb\xea\xe1\xcd\x20\x1e\xeb\xf0\x70\x88\x60\xe4\xb1\x8e\xeb\x22\x35\x86\xf2\xd2\xc0\x14\x68\x6a\x82\xb9\x23\xbc\x86\x70\x01\xd2\x75\x53\x11\xa4\xc1\xa2\xa8\xee\x6b\x2c\xa4\xb9\xc3\xac\xc8\xf3\x42\x83\xe6\x19\x61\x96\xd6\x04\xa9\xcb\xc6\xd4\x6f\x93\x79\x6b\x5d\x7a\x48\xd5\x6b\x15\x1c\xf5\x4c\x6d\xbe\x57\xa9\xce\xc8\xd6\xc5\xc2\x76\x7a\x58\xfd\x38\x3d\x6c\x67\xb4\xe3\xfd\x1c\xee\xf8\x13\x1c\x0e\x60\x1d\xe3\x25\x2d\xac\x5f\x64\xeb\x46\xa9\xde\xea\x78\x27\x7f\xa9\xb3\xe5\xb5\xc7\x78\x38\xf4\x58\xb7\x0b\x95\xea\xac\x49\x33\x42\xa9\xca\xac\xfe\xa6\x3c\xc6\x5c\x17\x35\x99\xa6\xdc\xd4\x1e\x1b\x8a\x68\x8c\x58\x04\x67\x67\x5c\x20\x18\x81\x5f\x05\x93\x78\x02\x93\x55\xc9\x6e\x69\x22\x0a\xb1\x9c\x4f\xea\x2c\x29\x55\xaa\x6b\x6f\x53\xb9\x1b\xfd\x87\xaa\x53\x3e\x8a\x04\x47\x10\x4e\xb8\x88\x97\xed\x3c\x1d\x0f\x57\x3d\xbd\x43\x63\xa3\x48\x80\x0f\xfc\x73\x88\xe8\x12\xfc\x8a\xfb\xd3\x98\x63\x2c\x22\x9f\x0f\xa7\x82\x7f\xd0\xe1\x1e\x63\x3f\x02\x00\x00\xff\xff\x00\xb3\x17\x19\x06\x08\x00\x00")

func sql027_create_validate_formula_triggerSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql027_create_validate_formula_triggerSql,
		"sql/027_create_validate_formula_trigger.sql",
	)
}

func sql027_create_validate_formula_triggerSql() (*asset, error) {
	bytes, err := sql027_create_validate_formula_triggerSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/027_create_validate_formula_trigger.sql", size: 2054, mode: os.FileMode(420), modTime: time.Unix(1517918552, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql030_create_billableSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x96\x51\x6f\xdb\x36\x10\xc7\x9f\xa5\x4f\x71\xf0\x8b\xac\xc1\xf2\x9a\xd5\xde\x5a\x07\x2b\x10\xcc\xde\x60\xa0\x4b\x82\xd8\xdd\x8a\xbd\x08\x34\x79\xb2\xb9\x49\xa4\x40\x52\x56\x92\x4f\x3f\x90\x32\x2d\x4b\x75\xe2\x3c\x0c\xcb\x4b\x64\xe9\xfe\xc7\xe3\xdd\xef\x78\x9c\x3f\xdc\xdd\xc3\xef\x37\xeb\xc5\xc3\xf2\xe6\xf3\xf2\xaf\xc5\x1c\xfe\x58\x2e\xfe\x84\xe5\xaf\xb0\xf8\xba\x5c\xad\x57\xb0\xe1\x79\x4e\x36\x39\x5e\x87\xbf\x3c\x2c\x6e\xd6\x8b\xf3\xc6\xb7\x77\xeb\xbe\x00\x6e\x56\x50\x73\xb3\x0b\xc3\x24\x01\x7c\x34\x8a\x50\x03\x95\xc6\xac\xca\x41\x9b\x2a\xcb\x20\x53\xb2\x80\x4a\x93\x2d\x02\xee\x51\x18\x6d\x2d\x6b\x04\xa3\x90\x18\xd8\x48\xb3\x03\x52\x96\x1a\x88\x60\xa0\x51\xed\x39\x45\x0d\x44\xb7\x2b\x0c\x14\x6a\x59\x29\x8a\x7a\x00\x5a\x82\x90\xaa\x20\x39\x7f\x46\x30\x3b\x84\x8c\x63\xce\xbc\xcb\xf6\x93\x36\xc4\xa0\x06\x23\xe1\xef\x4a\x1b\x58\xad\x6f\x1e\xd6\x8b\xf9\xf7\xab\xf5\xdd\xfd\xfd\x62\x0e\x1b\xa4\xa4\xd2\xd8\x86\x41\xa5\xd0\x48\x2b\xc3\xf7\xe8\x8d\xad\xb8\x40\x22\x60\x50\x95\x8c\x18\x1c\x84\x4d\xf8\x36\xb6\x61\x18\x0c\xc3\x20\xd0\x98\x23\x35\x61\x10\x04\x9c\x8d\xec\x3f\x6a\x9d\x21\x4b\x89\x99\xcd\x0c\x2f\x50\x1b\x52\x94\xe6\xd9\x4a\xda\x4f\xce\x72\xa8\x48\x9d\x16\xa8\x6d\x5e\x92\x4f\x9f\x22\x52\x96\xe9\xb6\xe2\x2c\x8a\xad\xb1\x7d\x7a\xd1\x4c\x90\x02\x1b\x33\xfb\x74\xd6\x4c\xaa\xed\x89\x37\xff\xeb\xac\xa9\x2e\x09\xc5\x13\xe3\xf6\xb7\x33\x8f\xb2\x09\x9b\x6c\x3e\x4e\x49\x92\x4d\xa7\x98\x4c\xa6\x1f\xdf\x27\x1f\xd8\x74\x92\xbc\x7f\xff\xe3\x84\xfe\x30\xfd\xe9\xe3\x07\x3a\x89\x66\x33\x83\x8f\xc6\xca\xcb\x9c\x88\x46\x0d\x49\x02\x19\xf9\x07\xdd\x2b\xe0\x0c\x32\xa9\x80\xca\xa2\xac\x4c\xf3\x4e\x3b\xff\x0c\x33\x52\xe5\x26\x39\x7c\xe9\xbb\x72\x5b\x84\xb3\x7f\x1d\xff\xd6\xee\xfc\x0a\x54\x92\x1c\x35\xc5\xfe\xc6\xb9\xd0\x86\x08\x8a\x29\x95\x95\x30\xd1\x08\xa2\x77\x51\x3c\x9b\x89\xaa\x40\xc5\xa9\x0d\xc0\x5a\x34\x5f\x47\xaf\x39\x2a\xb0\x90\xea\x29\xe5\x22\x2d\x36\x69\x89\x2a\xf5\x9e\xcf\xf9\x3c\x35\x76\x5e\xfb\xe5\xb0\xe0\x46\xae\x10\xf6\x29\x0c\x02\xdb\x3f\xd6\xd0\x96\xde\xb5\x51\x7a\x68\xa3\x20\xa8\x77\xa8\xf0\x65\x27\x3f\x43\x74\x60\x39\xb2\x46\x52\xc1\x2b\x76\xae\x31\xa2\x30\x88\xa1\x12\x5c\x0a\x20\x79\x0e\xff\x31\xe3\x87\xee\x3e\xe6\xe7\x2d\xc0\x7f\xa3\xf9\x5f\xe9\x7f\x29\x9a\x23\xe5\x8d\xaa\x85\xfe\xa2\xa8\x0d\xbf\xc5\xdb\xf5\xc1\x55\xf4\x1a\x7b\xd1\xbb\xe8\x75\x8c\x28\xd1\x8e\x04\xcb\x84\x80\x6f\x42\x70\x85\x8e\x6d\xa5\x9b\xf3\x7d\x1e\xd9\xd3\x53\x74\x01\xb9\xa8\x9d\x2f\x3e\x2f\x3a\x5a\x0f\xcd\x1b\xb4\x5f\xee\xe7\x2f\xad\x8b\x82\x9d\x03\xde\xe7\xed\xad\xd0\xf7\x49\x31\x4f\x65\x03\x77\x41\x04\xd9\x22\x4b\xfb\x16\x96\xf6\x30\x1e\xb9\xd1\x45\x65\xb1\xe1\x02\xdd\x14\xa2\x52\xec\x51\x19\x3b\x00\x94\xac\xb5\x3b\x56\x90\xd0\x9d\x9f\x0c\x89\x9f\x10\x52\x9d\xbc\x6a\x26\x4b\x49\xb8\x72\x03\x11\x58\xa5\x88\xb1\x9d\xa4\x88\xd8\xa2\x5d\x83\x67\x40\xe0\x38\xcc\x80\x49\xd4\x20\xa4\x81\x1d\xd9\x23\x10\xf0\x2e\xdc\x4e\xe1\x09\xcd\xa8\xc9\x95\x1d\x73\x5e\x15\xe9\x9e\x5f\xa8\x79\x9e\x83\xcb\xa0\x71\x96\xb6\x25\xdd\xc3\x9e\x63\x0d\x5c\x83\xc2\x4c\xa1\xde\x21\x6b\xe6\x57\xea\x64\x7e\x8a\x1d\xbb\xbb\x61\xd7\x33\x7c\xda\x33\xdd\x96\xf0\xc0\x76\x90\xef\xa0\xdc\x25\xb7\x0f\xaa\xd1\xe6\xd9\x45\x30\x3c\x39\x2c\x20\x47\xc2\x3a\x2f\xae\x46\x20\x64\x3d\x8c\x63\x90\x7b\x54\xc7\xed\xa7\xcd\x5c\x77\xfd\xe3\xf3\xe0\x62\x6c\xd0\x39\x90\xe3\x51\xa9\xb9\x60\xb2\x0e\x83\xa0\x27\x77\x7b\x2f\x89\x32\xdc\xe5\x71\xf3\xe4\x8e\x1f\x90\x8a\xa1\xb2\xbf\x38\x6b\xea\xbe\x41\x53\x23\x0a\xa0\x95\x52\xb6\x24\x4a\xd6\x8e\x8f\x2b\xc8\x64\x9e\xcb\x9a\x8b\x6d\x1c\x06\x5e\xe6\xd3\x07\x9c\x79\xa8\xb6\x28\x50\x11\x83\x27\x18\xed\x51\x3d\xc1\xc0\x13\x38\x00\x99\x01\x11\xf6\x02\x64\xed\xcd\xce\xd5\x4b\x57\xb9\xb1\x27\x80\xa5\x45\xd6\x50\xa2\x4a\xbc\x20\x6c\xf1\xe9\x15\xd0\x8c\xbf\x73\x15\x3c\x2c\x69\x61\xe7\xa8\x87\x57\x23\x30\xe3\xb6\x24\x71\x37\x49\x1e\x06\x13\xc6\x61\xe8\x5d\xa9\xb1\xab\xaa\x1a\x1f\xca\xab\xc6\x87\xd2\xaa\xf1\x09\x17\x6a\xdc\x01\x43\x8d\x4f\x89\x50\xe3\x5e\xdd\xd5\xd8\x97\x2b\x6c\xd6\x6f\xf7\xa1\xc2\x43\x47\xab\xb1\x2b\x4f\x67\x6c\xb5\xc9\xb5\x51\x85\xd7\xa1\xbf\x9f\x2e\x6f\xe7\x8b\xaf\xbd\x3b\x29\x67\x8f\x29\x67\x70\x77\xdb\xde\x1d\x87\x9c\xc5\xd7\x97\x34\x52\x6d\xbb\x22\xbf\xcd\xcb\x52\x97\x83\xae\xb8\x4d\xcb\x65\xf9\xb1\x97\x4f\x3d\x7c\x59\x2d\x6f\x7f\x83\x2d\xd7\x06\x86\xde\xe0\xb2\x2b\x77\x0b\xea\x04\x72\x2c\x48\x7c\x1d\xfe\x1b\x00\x00\xff\xff\xb0\xa6\xaa\x51\x01\x0c\x00\x00")

func sql030_create_billableSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql030_create_billableSql,
		"sql/030_create_billable.sql",
	)
}

func sql030_create_billableSql() (*asset, error) {
	bytes, err := sql030_create_billableSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/030_create_billable.sql", size: 3073, mode: os.FileMode(420), modTime: time.Unix(1518608567, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sql040_create_billable_range_functionSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x54\xcd\x6e\xdb\x3c\x10\x3c\x93\x4f\xb1\x87\x20\xb6\x02\xc9\x88\x3f\xe0\x43\x81\x1a\x3d\x38\x89\x92\x1a\x08\x9c\xd4\x3f\x68\x7b\x12\x28\x6b\xad\xb0\xa0\x48\x86\xa2\xad\xa4\x4f\x5f\x90\xb2\x24\x5b\x49\x7b\x8a\x2f\xe6\xce\xae\x66\xb9\xb3\x23\x45\x11\x18\x2c\xd1\xd2\x9b\xc5\xc3\x23\xdc\xae\xe7\xd7\xab\xd9\xc3\x1c\x66\xb7\x10\xff\x98\x2d\x57\x4b\x48\xb9\x10\x2c\x15\x98\x18\x26\x73\x1c\xda\xd2\xfe\xf6\xa7\x60\x42\x69\x14\xf5\xd2\xf0\xbc\x43\xc3\xb1\x04\xfb\x84\x6d\x0a\xf6\x1c\x2b\xd8\x2a\x03\x0c\x72\xbe\x47\x09\xbe\x96\x46\x91\x23\xc0\x17\x56\x68\x81\x9f\x0f\x21\x00\x94\x28\x70\x63\xe1\x02\xb6\x46\x15\x7f\xed\x3f\x1c\xfc\x77\x39\xfe\x14\x5d\x8e\xa3\xcb\xf1\x20\x84\x2e\xfa\x7f\x10\x04\x13\xc7\x46\xaf\x17\xf1\x74\x15\xc3\xc3\x02\x16\xf1\xe3\xfd\xf4\x3a\xee\xe6\xeb\xb1\xd6\x2d\x31\x3b\x4c\xd1\x0d\x09\x8b\x78\xb5\x5e\xcc\x97\xb0\x9a\x5e\xdd\xc7\x43\x4a\x78\x06\x5c\xda\x90\x92\x7c\xc7\x33\xb0\xf8\xe2\xce\x92\x15\xd8\x9c\x95\xc9\x93\xe3\x5c\xa9\xd9\x06\x4f\x90\x02\x0b\x65\x5e\x13\x2e\x93\x22\x05\xb9\x2b\xd0\xf0\x4d\x48\x49\xb6\x33\xcc\x72\x25\xbb\xee\x21\x25\xda\xf0\x0d\x97\x79\xa2\x05\x93\x49\xdb\xfb\x04\x3d\x6e\xbe\x55\xa6\xd8\x09\xd6\x84\xae\x0e\x9b\x16\x34\x80\xe9\x12\xce\xce\xe0\x2a\xbe\x9b\xcd\x29\xa9\x27\x83\x6f\xeb\x78\xf1\x13\xbe\xcf\x56\x5f\x29\xd9\x33\xc1\xb3\xe4\x98\xbd\x04\x56\xc2\x90\x12\x52\x2b\x44\x09\x21\x5a\x8f\x2e\x42\x77\xe8\x56\x51\x3f\xe7\xd6\x15\x82\x40\x96\x9d\x00\xe3\x10\x06\x5c\x6e\xb9\xe4\xf6\x75\x10\x80\xda\xa3\x01\x4f\x1d\x38\xee\x43\xa5\x32\x94\x10\x57\xef\x3b\x9c\x5c\x40\x6b\x4a\x48\xc5\x65\xa6\x2a\x9f\x6c\x6f\xa5\x99\xb1\xdc\x2b\x96\xbe\x7a\xc6\x5a\x65\x65\x32\x34\x0e\xea\x2e\x01\x46\x55\x25\xa4\x68\x2b\x44\x09\x9b\x9d\x31\x28\xad\x03\x81\xc9\x0c\xc6\xb0\x55\x42\xa8\x8a\xcb\x3c\xa0\x24\xa0\xdd\xac\xe9\x88\x67\xa1\xff\x77\xcc\xf5\xc9\xc9\x5d\x9f\x9a\x4d\xd7\x51\xb7\xe7\x3a\x3e\xde\xb2\x43\x3a\xb5\xdc\x14\xb9\x41\x66\xb1\xb4\x43\xa1\x2a\x34\x3d\xff\x05\x21\xd4\xf0\x5e\xeb\x51\x2b\x50\x8b\xa6\xa3\xc6\x2a\x41\xe0\x17\x21\x90\x95\x76\xb8\xd3\xfa\x3d\xa6\x1a\xee\x33\xd5\xe8\x31\x13\x25\xc4\x2f\xa4\x41\x1c\xb3\x7b\x8a\x67\xce\x36\x3d\x1f\x36\x49\xef\xbd\x7e\xba\x51\xc8\x15\x1c\x0c\xe9\x42\xdc\x33\x91\x1c\x62\xaf\xc1\x5b\x8d\x7a\x22\x7d\xac\x4a\x1f\x26\x13\xa9\xf9\x8e\xc6\x6b\xb4\xf3\xef\x1b\x6d\x7c\xdc\x7e\xff\x52\x4a\x08\x97\x12\x0d\xfc\x52\x5c\x3a\x65\xde\x79\xcf\xf6\x5a\x03\x05\xff\x73\x8e\x1e\x75\x86\xfe\xe2\x72\x5d\x7c\x28\x72\xce\x3d\xb9\x2e\x9c\x9f\x43\x77\x53\xf8\x57\xd9\xe9\xf4\x94\x54\x4f\x68\xd0\xbb\xb6\x7d\xfc\x4d\xd5\x84\xc6\xf3\x9b\x89\xfb\x7e\xdc\x4f\xe7\x77\xeb\xe9\x5d\x0c\x5a\xe8\xbc\x7c\x16\x13\xfa\x27\x00\x00\xff\xff\xff\x94\x54\xe6\x45\x06\x00\x00")

func sql040_create_billable_range_functionSqlBytes() ([]byte, error) {
	return bindataRead(
		_sql040_create_billable_range_functionSql,
		"sql/040_create_billable_range_function.sql",
	)
}

func sql040_create_billable_range_functionSql() (*asset, error) {
	bytes, err := sql040_create_billable_range_functionSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sql/040_create_billable_range_function.sql", size: 1605, mode: os.FileMode(420), modTime: time.Unix(1518609853, 0)}
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
	"sql/001_create_to_seconds_function.sql": sql001_create_to_seconds_functionSql,
	"sql/002_create_iso8601_function.sql": sql002_create_iso8601_functionSql,
	"sql/010_create_app_usage_events.sql": sql010_create_app_usage_eventsSql,
	"sql/020_create_service_usage_events.sql": sql020_create_service_usage_eventsSql,
	"sql/025_create_pricing_plans.sql": sql025_create_pricing_plansSql,
	"sql/026_create_pricing_functions.sql": sql026_create_pricing_functionsSql,
	"sql/027_create_validate_formula_trigger.sql": sql027_create_validate_formula_triggerSql,
	"sql/030_create_billable.sql": sql030_create_billableSql,
	"sql/040_create_billable_range_function.sql": sql040_create_billable_range_functionSql,
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
	"sql": &bintree{nil, map[string]*bintree{
		"001_create_to_seconds_function.sql": &bintree{sql001_create_to_seconds_functionSql, map[string]*bintree{}},
		"002_create_iso8601_function.sql": &bintree{sql002_create_iso8601_functionSql, map[string]*bintree{}},
		"010_create_app_usage_events.sql": &bintree{sql010_create_app_usage_eventsSql, map[string]*bintree{}},
		"020_create_service_usage_events.sql": &bintree{sql020_create_service_usage_eventsSql, map[string]*bintree{}},
		"025_create_pricing_plans.sql": &bintree{sql025_create_pricing_plansSql, map[string]*bintree{}},
		"026_create_pricing_functions.sql": &bintree{sql026_create_pricing_functionsSql, map[string]*bintree{}},
		"027_create_validate_formula_trigger.sql": &bintree{sql027_create_validate_formula_triggerSql, map[string]*bintree{}},
		"030_create_billable.sql": &bintree{sql030_create_billableSql, map[string]*bintree{}},
		"040_create_billable_range_function.sql": &bintree{sql040_create_billable_range_functionSql, map[string]*bintree{}},
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

