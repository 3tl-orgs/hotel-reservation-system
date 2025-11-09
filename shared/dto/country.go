package dto

type CountryResponse struct {
	Id   int
	Code string
	Name string
}

type ProvinceResponse struct {
	Id        int
	Code      string
	Name      string
	CountryId int
}

type WardResponse struct {
	Id         int
	Code       string
	Name       string
	ProvinceId int
}
