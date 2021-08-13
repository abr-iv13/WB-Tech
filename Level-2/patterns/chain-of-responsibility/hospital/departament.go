package hospital

type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
