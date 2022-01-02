package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is johnny"}
	diff2 := &String{Value: "My name is johnny"}
	if hello1.DictKey() != hello2.DictKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if diff1.DictKey() != diff2.DictKey() {
		t.Errorf("strings with same content have different hash keys")
	}
	if hello1.DictKey() == diff1.DictKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
