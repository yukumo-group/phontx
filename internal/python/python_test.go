package python

import (
	"encoding/hex"
	"testing"
)

// TestExecPython tests the execution of python script
func TestExecPython(t *testing.T) {
	t.Parallel()
	const scriptPath string = "./testdata/testdata.py"
	output, err := ExecutePython(scriptPath)
	if err != nil {
		t.Fatal(err)
	}
	expectedData := "Hello World!"
	strOutput := string(output)
	if strOutput != expectedData {
		t.Errorf(
			"e: %v, g: %v",
			[]byte(expectedData),
			output,
		)
		t.Fatalf(
			"expected %v, got %v",
			expectedData,
			strOutput,
		)
	}
}

// TestGetByte tests getting byte data through hex output of python script
func TestGetByte(t *testing.T) {
	t.Parallel()
	const scriptPath string = "./testdata/testdata2.py"
	result, err := ExecutePython(
		scriptPath,
	)
	if err != nil {
		t.Fatal(err)
	}
	decodedData, err := hex.DecodeString(
		string(result),
	)
	if err != nil {
		t.Fatal(err)
	}
	expectedData := "114514"
	if string(decodedData) != expectedData {
		t.Fatalf(
			"expected %s, got %s",
			string(decodedData),
			expectedData,
		)
	}
}
