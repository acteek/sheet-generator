package generate

type FileGenerator interface {
	Name() string
	Init()
	InsertRecord(record []string)
	Complete()
}
