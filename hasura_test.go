package hasura

import (
	"testing"
)

type author struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type insertAuthorResponse struct {
	AffectedRows int      `json:"affected_rows"`
	Returning    []author `json:"returning"`
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

func TestSampleInsert(t *testing.T) {
	c := NewClient("http://localhost:8080/v1/query", nil)

	q := Query{
		Type: "insert",
		Args: Args{
			Table: "author",
			Objects: []interface{}{
				map[string]interface{}{
					"name": "skm",
				},
			},
			OnConflict: OnConflict{
				Action:     "update",
				Constraint: "author_pkey",
			},
			Returning: []string{"id", "name"},
		},
	}

	var response insertAuthorResponse
	err := c.Execute(q, &response)
	if err != nil {
		t.Fatalf("expected no error, encountered %v, forgot to setup db?", err)
	}

	if response.AffectedRows < 1 {
		t.Fatalf("expected atleast one change, got %v, forgot to setup db?", response)
	}
}
