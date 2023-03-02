package internal

const DefaultHeaderSize = 4

const (
	StartJob = iota
	StopJob
	ContinueJob
)

type Package struct {
	Type int
	Body []byte
}
