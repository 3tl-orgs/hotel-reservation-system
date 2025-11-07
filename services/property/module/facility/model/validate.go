package model


func checkFacilityName(name string) error {
	if name == "" {
		return ErrFacilityNameEmpty
	}
	return nil
}
