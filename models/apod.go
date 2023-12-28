package models

type ApodData struct {
	Copyright      string     `json:"copyright"`
	Date           CustomTime `json:"date"`
	Explanation    string     `json:"explanation"`
	MediaType      string     `json:"media_type"`
	ServiceVersion string     `json:"service_version"`
	Title          string     `json:"title"`
	URL            string     `json:"url"`
	HDURL          string     `json:"hdurl"`
}

type ApodDto struct {
	Copyright      string     `json:"copyright"`
	Date           CustomTime `json:"date"`
	Explanation    string     `json:"explanation"`
	MediaType      string     `json:"media_type"`
	ServiceVersion string     `json:"service_version"`
	Title          string     `json:"title"`
	URL            string     `json:"url"`
	HDURL          string     `json:"hdurl"`
	Data           []byte     `json:"data"`
}

type ApodChanelType struct {
	ApodData *ApodData
	Error    error
}

type GetApodsRequest struct {
	Limit  *int32 `query:"limit"`
	Offset *int32 `query:"offset"`
}

type GetApodByDateRequest struct {
	Date CustomTime
}

type GetApodsResponse struct {
	Apods  []ApodDto `json:"apods"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type GetApodByDateResponse struct {
	Apod ApodDto `json:"apod"`
}
