# go-hasura

A simple Go clien for Hasura JSON APIs

### Installation

```bash
go get github.com/shahidhk/go-hasura
```

### Usage

```go
// import the packages
import (
    "fmt"
    "log"

    "github.com/shahidhk/go-hasura" // imported as 'hasura'
)

// response type
type author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
    // create a client
    c := hasura.NewClient("http://localhost:8080/v1/query", nil)
    // if you need to add access key or other headers,
    // c := hasura.NewClient("http://localhost:8080/v1/query", map[string]string{
    //     "X-Hausra-Access-Key": "xyz",
    // })

    // build a query object
	q := hasura.Query{
		Type: "select",
		Args: hasura.Args{
			Table: "author",
			Columns: []string{
				"id",
				"name",
			},
		},
    }
    
    // create response object
	var authors []author
    
    // execute the query
    err := c.Execute(q, &response)
    // error is thrown for non-200 responses also
    if err != nil {
        log.Fatal(err)
    }

    // print the response
    fmt.Println(authors)
}
```