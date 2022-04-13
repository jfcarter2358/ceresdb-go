package connection

import (
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	Initialize("ceresdb", "ceresdb", "localhost", 7437)

	expectedData := []map[string]interface{}{{"name": "_auth"}}

	data, err := Query("GET DATABASE")

	if !reflect.DeepEqual(data, expectedData) {
		t.Errorf("Data was incorrect, got: %v, want: %v", data, expectedData)
	}
	if err != nil {
		t.Errorf("Error was incorrect, got: %v, want: %v", err, "<nil>")
	}
}
