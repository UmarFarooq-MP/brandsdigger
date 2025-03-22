package names

type Generator interface {
	GenerateNames(message string) ([]string, error)
}
