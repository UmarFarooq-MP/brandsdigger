package client

type MessagesGenerator interface {
	GenerateNames(message string) ([]string, error)
}
