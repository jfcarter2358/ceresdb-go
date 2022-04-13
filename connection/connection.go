package connection

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"strings"
)

var Username string
var Password string
var Host string
var Port int

const DATA_TERMINATOR = "EOD"

func Initialize(username string, password string, host string, port int) {
	Username = username
	Password = password
	Host = host
	Port = port
}

func Query(queryString string) ([]map[string]interface{}, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", Host, Port))
	if err != nil {
		return nil, err
	}
	payload := map[string]interface{}{
		"_auth": fmt.Sprintf("%v:%v", Username, Password),
		"query": queryString,
	}
	data, _ := json.Marshal(payload)
	fmt.Fprint(conn, string(data)+"\n")

	outputString := ""

	reader := bufio.NewReader(conn)

	byteData := make([]byte, 65536)
	n, err := reader.Read(byteData)
	if err != nil {
		return nil, err
	}
	outputString += string(byteData[:n])
	for !strings.HasSuffix(outputString, DATA_TERMINATOR) {
		byteData := make([]byte, 65536)
		n, err = reader.Read(byteData[:n])
		if err != nil {
			return nil, err
		}
		outputString += string(byteData)
	}
	conn.Close()

	outputString = outputString[:len(outputString)-3]

	if outputString == "null" {
		return nil, nil
	}

	// We got back a dictionary not a list, so that means an error was thrown
	if strings.HasPrefix(outputString, "{") {
		var outputData map[string]interface{}
		err = json.Unmarshal([]byte(outputString), &outputData)
		if err != nil {
			return nil, err
		}
		err = errors.New(outputData["error"].(string))
		return nil, err
	}

	var outputData []map[string]interface{}
	err = json.Unmarshal([]byte(outputString), &outputData)
	if err != nil {
		return nil, err
	}

	return outputData, nil
}
