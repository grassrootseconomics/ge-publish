package contract

type Contract[T any] interface {
	Version() string
	Bytecode(T) ([]byte, error)
}
