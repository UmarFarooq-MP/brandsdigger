package dns

type DomainValidator interface {
	ValidateDomain(domains []string) (map[string]bool, error)
}
