package chainresponsibility

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
