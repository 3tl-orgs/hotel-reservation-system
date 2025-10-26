package provincemodel

func checkCountryName(name string) error {
	if name == "" {
		return ErrProvinceNameIsEmpty
	}

	return nil
}

func checkCountryCode(code string) error {
	if code == "" {
		return ErrProvinceCodeIsEmpty
	}

	return nil
}
