package hospital

type Department interface {
	execute(*Patient)
	setNext(Department)
}
