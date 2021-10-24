package memory

import (
	"reflect"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	testCache := New()
	if reflect.TypeOf(testCache) != reflect.TypeOf(Cache{}) {
		t.Errorf("New method does not return Cache instance true")
	}
}

func TestCache_Set(t *testing.T){
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	if testCache.Get(testKey) != testValue{
		t.Errorf("Set initialy method does not work properly")
	}
	testCache.Set(testKey,testValue+testValue)
	if testCache.Get(testKey) != testValue+testValue{
		t.Errorf("Set old value method does not work properly")
	}
}

func TestCache_Get(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	if testCache.Get(testKey) != testValue{
		t.Errorf("Get value method does not work properly")
	}
}

func TestCache_Delete(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	testCache.Delete(testKey)
	if testCache.Get(testKey) == testValue{
		t.Errorf("Delete method does not work right")
	}
}

func TestCache_DeleteAll(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	testCache.Set(testKey+testKey,testValue)
	testCache.DeleteAll()
	if testCache.Get(testKey) == testValue || testCache.Get(testKey+testKey) == testValue{
		t.Errorf("Delete all method does not work properly")
	}
}

func TestCache_GetFrequency(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	if testCache.GetFrequency() != 60 {
		t.Errorf("Get Frequency method does not work properly")
	}
}

func TestCache_SetFrequency(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	newFreq := 28
	testCache.SetFrequency(newFreq)
	if testCache.GetFrequency() != newFreq {
		t.Errorf("Get Frequency method does not work properly")
	}
}

func TestCache_GetPath(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	if testCache.GetPath() != "test.json" {
		t.Errorf("Set Path method does not work properly")
	}
}

func TestCache_SetPath(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	newPath := "harun.json"
	testCache.SetPath(newPath)
	if testCache.GetPath() != newPath {
		t.Errorf("Get Path method does not work properly")
	}
}

func TestCache_GetValue(t *testing.T) {
	testKey := "test"
	testValue := "test value"
	testCache :=New()
	testCache.Set(testKey,testValue)
	getValue := testCache.GetValue()
	if !strings.Contains(getValue, testKey)  || !strings.Contains(getValue, testValue){
		t.Errorf("Get Path method does not work properly")
	}
}