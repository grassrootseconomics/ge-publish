package contract

type Contract[T any] interface {
	Version() string
	GasLimit() uint64
	Bytecode(T) ([]byte, error)
}
