package models

// SuccessResponse represents http success response structure required by datatable api
type SuccessResponse struct {
	Draw            int64   `json:"draw"`            // The draw counter that this object is a response to - from the draw parameter sent as part of the data request.
	RecordsTotal    int64   `json:"recordsTotal"`    // Total records in db, before filtering
	RecordsFiltered int64   `json:"recordsFiltered"` // Total records in db, after filtering
	Data            []*Team `json:"data"`            // The data to be displayed in the table //TODO: check if interface works
}

// NewSuccessResponse creates a new success response to be consumed by the databale api
func NewSuccessResponse(draw int64, teams []*Team, totalCount, totalFiltered int64) SuccessResponse {
	return SuccessResponse{
		Draw:            draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalFiltered,
		Data:            teams,
	}
}

// ErrorResponse represents http error response structure required by datatable api
type ErrorResponse struct {
	Error string // Optional: If an error occurs during the running of the server-side processing
}

// NewErrorResponse creates a new error response to be consumed by the databale api
func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{Error: err.Error()}
}
