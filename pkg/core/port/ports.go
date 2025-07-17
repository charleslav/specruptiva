package ports

type Validator interface {
	Validate() error
	SetSchema(schema string) error
	SetData(data string) error
}
