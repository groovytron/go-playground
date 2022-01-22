package schemas

type ApiPaginationSchema struct {
	TotalItems int64       `json:"totalItems"`
	Next       *int64      `json:"next"`
	Previous   *int64      `json:"previous"`
	Last       int64       `json:"last"`
	Current    int64       `json:"current"`
	Items      interface{} `json:"items"`
}
