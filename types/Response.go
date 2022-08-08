package types

type FormattedResponse struct {
	StatusCode int
	Message    string
	Data       any
	Errors     any
}

type ListProvince struct {
	Provinces any
}
type ShowProvince struct {
	Province any
}

type ListCity struct {
	Cities any
}
type ShowCity struct {
	City any
}

type ListSubdistrict struct {
	Subdistricts any
}
type ShowSubdistrict struct {
	Subdistrict any
}
