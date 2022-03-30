package sport

import (
	"reflect"
	"testing"
)

func TestWrite_file(t *testing.T) {

	arr := []byte("Here is a string....")
	// t.Error("here", arr)
	_, err := Write_file("uj.txt", arr)
	if err != nil {
		t.Errorf("Test write error occured")
	}

}

func TestProcess_file(t *testing.T) {
	str := "Rohan,12,M,[cricket,football],5.9,50"
	data, _ := Process_file(str)

	k := []string{"Rohan", "12", "M", "", "5.9", "50"}
	// t.Error(data, k)
	// t.Error(reflect.DeepEqual(data, k))

	// if data[0] != k[0] {
	if !(reflect.DeepEqual(data, k)) {
		t.Errorf("process file error occured")
	}

}
