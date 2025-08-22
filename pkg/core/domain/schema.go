package domain

type Schema struct {
	Id         string `json:"id"`
	Schema     string `json:"schema"`
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type Schemas []Schema
