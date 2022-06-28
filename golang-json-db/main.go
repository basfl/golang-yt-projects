package main

import (
	"encoding/json"
	"fmt"
	"golang-json-db/mydb"
)

const VERSION string = "1.0.0"

/////////////////////////////////////////////
// type Logger interface {
// 	Fatal(string, ...interface{})
// 	Info(string, ...interface{})
// 	Debug(string, ...interface{})
// 	Error(string, ...interface{})
// }

// type Driver struct {
// 	mutex   sync.Mutex
// 	mutexes map[string]*sync.Mutex
// 	dir     string
// 	logger  *zap.Logger
// }

// type Options struct {
// 	Logger *zap.Logger
// }

// func New(dir string, options *Options) (*Driver, error) {
// 	dir = filepath.Clean(dir)
// 	opts := Options{}
// 	if options != nil {
// 		opts = *options
// 	}
// 	//init logger

// 	if opts.Logger == nil {
// 		opts.Logger = getZipLoggerConfig()
// 	}

// 	driver := Driver{
// 		mutexes: make(map[string]*sync.Mutex),
// 		dir:     dir,
// 		logger:  opts.Logger,
// 	}

// 	if _, err := os.Stat(dir); err == nil {
// 		opts.Logger.Debug(fmt.Sprintf("using %v database already exist\n", dir))
// 		return &driver, nil
// 	}
// 	opts.Logger.Debug(fmt.Sprintf("creating  database at %v ...\n", dir))

// 	return &driver, os.MkdirAll(dir, 0755)

// }

// func (d *Driver) Write(collection, resource string, v interface{}) error {
// 	d.logger.Info(fmt.Sprintf("start writing collection %v with resource %v", collection, resource))
// 	if collection == "" {
// 		return fmt.Errorf("Missing collection no place to save collection")
// 	}
// 	if resource == "" {
// 		return fmt.Errorf("Missing resource un-able to save resource")
// 	}
// 	mutex := d.getOrCreateMutex(collection)
// 	mutex.Lock()
// 	defer mutex.Unlock()
// 	dir := filepath.Join(d.dir, collection)
// 	fnPath := filepath.Join(dir, resource+".json")
// 	tmpPath := fnPath + ".tmp"

// 	if err := os.MkdirAll(dir, 0755); err != nil {
// 		return err
// 	}

// 	b, err := json.MarshalIndent(v, "", "\t")
// 	if err != nil {
// 		return err
// 	}

// 	b = append(b, byte('\n'))
// 	if err := ioutil.WriteFile(tmpPath, b, 0644); err != nil {
// 		return err
// 	}

// 	return os.Rename(tmpPath, fnPath)

// }

// func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {

// 	d.mutex.Lock()
// 	defer d.mutex.Unlock()
// 	m, ok := d.mutexes[collection]

// 	if !ok {
// 		m = &sync.Mutex{}
// 		d.mutexes[collection] = m
// 	}

// 	return m
// }

/////////////////////////////////////////////////
type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}
type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	fmt.Println("............")
	dir := "./"

	db, err := mydb.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{Name: "john", Age: "23", Contact: "1st ave ", Company: "company1", Address: Address{City: "nyc", State: "ny", Country: "US", Pincode: "123"}},
		{Name: "ben", Age: "28", Contact: "main street ", Company: "company2", Address: Address{City: "alb", State: "nm", Country: "US", Pincode: "123"}},
		{Name: "amir", Age: "32", Contact: "beesat ", Company: "company3", Address: Address{City: "syz", State: "fars", Country: "Iran", Pincode: "123"}},
		{Name: "jack", Age: "43", Contact: "west ave ", Company: "company4", Address: Address{City: "london", State: "london", Country: "Uk", Pincode: "1235"}},
		{Name: "marven", Age: "52", Contact: "5th ave ", Company: "company5", Address: Address{City: "bridgeport", State: "wv", Country: "US", Pincode: "123"}},
		{Name: "natasha", Age: "21", Contact: "east street ", Company: "company6", Address: Address{City: "sf", State: "ca", Country: "US", Pincode: "123"}},
	}

	fmt.Printf("employees are %v\n", employees)

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}
	res, err := db.Read("users", "natasha")
	if err != nil {
		//fmt.Printf("error " + err.Error())
		panic(err)
	}
	fmt.Printf("result is " + res)
}
