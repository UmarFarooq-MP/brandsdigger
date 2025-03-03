package client

type MessagesGenerator interface {
	GenerateNames(message string) ([]string, error)
}

type DomainValidator interface {
	ValidateDomain(domain []string) (map[string]bool, error)
}
