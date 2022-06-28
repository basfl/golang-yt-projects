package mydb

import (
	"encoding/json"
	"fmt"
	"golang-json-db/mydb/util"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"go.uber.org/zap"
)

type Logger interface {
	Fatal(string, ...interface{})
	Info(string, ...interface{})
	Debug(string, ...interface{})
	Error(string, ...interface{})
}

type Driver struct {
	mutex   sync.Mutex
	mutexes map[string]*sync.Mutex
	dir     string
	logger  *zap.Logger
}

type Options struct {
	Logger *zap.Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)
	opts := Options{}
	if options != nil {
		opts = *options
	}
	//init logger

	if opts.Logger == nil {
		opts.Logger = util.GetLogger()
	}

	driver := Driver{
		mutexes: make(map[string]*sync.Mutex),
		dir:     dir,
		logger:  opts.Logger,
	}

	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug(fmt.Sprintf("using %v database already exist\n", dir))
		return &driver, nil
	}
	opts.Logger.Debug(fmt.Sprintf("creating  database at %v ...\n", dir))

	return &driver, os.MkdirAll(dir, 0755)

}

func (d *Driver) Write(collection, resource string, v interface{}) error {
	d.logger.Info(fmt.Sprintf("start writing collection %v with resource %v", collection, resource))
	if collection == "" {
		return fmt.Errorf("Missing collection no place to save collection")
	}
	if resource == "" {
		return fmt.Errorf("Missing resource un-able to save resource")
	}
	mutex := d.getOrCreateMutex(collection)
	mutex.Lock()
	defer mutex.Unlock()
	dir := filepath.Join(d.dir, collection)
	fnPath := filepath.Join(dir, resource+".json")
	tmpPath := fnPath + ".tmp"

	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}

	b = append(b, byte('\n'))
	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
		return err
	}

	return os.Rename(tmpPath, fnPath)

}

func (d *Driver) Read(collection, resource string) (string, error) {
	var result string
	if collection == "" {
		return result, fmt.Errorf("Missing collection - unable to read!")
	}

	if resource == "" {
		return result, fmt.Errorf("Missing resource - unable to read record (no name)!")
	}
	record := filepath.Join(d.dir, collection, resource)
	if _, err := util.Stat(record); err != nil {
		return result, err
	}

	b, err := ioutil.ReadFile(record + ".json")
	if err != nil {
		return result, err
	}
	result = string(b)
	return result, nil

}

func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {

	d.mutex.Lock()
	defer d.mutex.Unlock()
	m, ok := d.mutexes[collection]

	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}
