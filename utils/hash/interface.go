package hash

type HashInterface interface {
	GenerateHash(password string) (string, error)
	ComparePassword(hash, password string) (bool, error)
}
