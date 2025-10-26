package model

func checkName(name string) error {
	if name == "" {
		return ErrAmenityNameIsEmpty
	}
	return nil
}
