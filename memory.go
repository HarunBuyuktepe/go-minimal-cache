package memory

import (
	"encoding/json"
	"sync"
)

var (
	once sync.Once
	Instance Cache
)

type Cache struct {
	*cache
}


type cache struct {
	items map[string]interface{}
	//to avoid race-condition
	mu    sync.RWMutex
}

//create once an instance and return
func New() Cache {
	once.Do(func() {
		c := &cache{
			items: make(map[string]interface{}),
		}
		Instance = Cache{c}
		go ActivateWriter()
	})
	ReadFileAndWriteToMemory()
	return Instance
}

//set key: create if not exists ; edit if exists
func (c *cache) Set(k string, x interface{}) {
	c.mu.Lock()
	c.items[k] = x
	c.mu.Unlock()
}

//get requested key
func (c *cache) Get(k string) interface{} {
	var result interface{}
	result = "No value available"
	c.mu.Lock()
	if val, ok := c.items[k]; ok {
		result = val
	}
	c.mu.Unlock()

	return result
}

//delete selected key
func (c *cache) Delete(k string) bool {
	result := false
	c.mu.Lock()
	_, ok := c.items [k]
	if ok {
		delete(c.items, k)
		result = true
	}
	c.mu.Unlock()
	return result
}

//flush all data
func (c *cache) DeleteAll() bool{
	result := false
	c.mu.Lock()
	for k := range c.items {
		delete(c.items, k)
	}
	c.mu.Unlock()
	return result
}

//set frequency(second) to write file
func (c *cache) SetFrequency(NewFrequency int) {
	Frequency = NewFrequency
}
//get frequency(second) to write file
func (c *cache) GetFrequency() int {
	return Frequency
}

//set temp store file path
func (c *cache) SetPath(NewPath string) {
	TempFilePath = NewPath
}
//get temp store file path
func (c *cache) GetPath() string {
	return TempFilePath
}

//return all value in memory in json format
func (c *cache) GetValue() string {
	jsonString, _ := json.MarshalIndent(Instance.items, "", " ")
	return string(jsonString)
}
