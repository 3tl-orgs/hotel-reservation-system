package propertymodel

func checkAddress(address string) error {
	if address == "" {
		return ErrAddressIsEmpty
	}
	return nil
}

func checkName(name string) error {
	if name == "" {
		return ErrNameIsEmpty
	}
	return nil
}

func checkWardId(wardId string) error {
	if wardId == "" {
		return ErrWarIdIsEmpty
	}
	return nil
}

func checkProvinceId(provinceId string) error {
	if provinceId == "" {
		return ErrProvinceIdIsEmpty
	}
	return nil
}

func checkCountryId(countryId string) error {
	if countryId == "" {
		return ErrCountryIdIsEmpty
	}
	return nil
}

func checkPropertyTypeId(propertyTypeId string) error {
	if propertyTypeId == "" {
		return ErrPropertyTypeIdIsEmpty
	}
	return nil
}
