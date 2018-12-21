package hasura

import (
	"testing"
)

type author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestSampleRequest(t *testing.T) {
	c := NewClient("http://localhost:8080/v1/query", nil)

	q := Query{
		Type: "select",
		Args: Args{
			Table: "author",
			Columns: []string{
				"id",
				"name",
			},
		},
	}

	var response []author
	err := c.Execute(q, &response)
	if err != nil {
		t.Fatalf("expected no error, encountered %v, forgot to setup db?", err)
	}

	if len(response) < 1 {
		t.Fatalf("expected atleast one elemet, got %v, forgot to setup db?", response)
	}
}
