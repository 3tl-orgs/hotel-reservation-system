package model

func checkCountryName(countryName string) error {
	if countryName == "" {
		return ErrCountryNameIsEmpty
	}

	return nil
}

func checkCountryCode(countryCode string) error {
	if countryCode == "" {
		return ErrCountryCodeIsEmpty
	}

	return nil
}
