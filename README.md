# CeresDB Go

CeresDB Go is a Go SDK for the CeresDB database. To use it, follow the below 
example:

```go
package main

import (
	"github.com/jfcarter2358/ceresdb-go/connection"
)

func main() {
	ceresDBUsername := "ceresdb"
	ceresDBPassword := "ceresdb"
	ceresDBHost := "localhost"
	ceresDBPort := 7437

	connection.Initialize(ceresDBUsername, ceresDBPassword, ceresDBHost, ceresDBPort)
	data, err := connection.Query("GET DATABASE")
}
```
    

Data will always be returned as `[]map[string]interface{}` with any errors during the query 
being returned as the second parameter. For more information on AQL queries, head 
[Here](https://ceresdb.readthedocs.io/en/latest/querying.html>)
