package propertytypemodel

func checkPropertyName(name string) error {
	if name == "" {
		return PropertyTypeNameNotEmptyErr
	}
	return nil
}
