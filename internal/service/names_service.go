package service

import "brandsdigger/internal/factory"

type NamesService struct {
}

func (ns *NamesService) GetNames(message string) ([]string, error) {

	suggestedName, err := factory.Generate.GenerateNames(message)
	if err != nil {
		return nil, err
	}

	m, err := factory.DomainValidator.ValidateDomain(suggestedName)
	if err != nil {
		return nil, err
	}

	var availableDomain []string
	for key, value := range m {
		if value == true {
			availableDomain = append(availableDomain, key)
		}
	}
	return availableDomain, nil
}
