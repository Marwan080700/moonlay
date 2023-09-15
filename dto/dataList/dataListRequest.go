package datalist

type DataListRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	File        string `json:"file"`
}