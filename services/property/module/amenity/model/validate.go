package model

func checkAmenityName(name string) error {
	if name == "" {
		return ErrAmenityNameIsEmpty
	}
	return nil
}
