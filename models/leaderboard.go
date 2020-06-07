package models

// LeaderboardQuery holds the query parameters for getting the leaderboard
type LeaderboardQuery struct {
	Draw   int64
	Length int64
	Start  int64
}

// LeaderboardSuccessResponse represents http success response structure required by datatable api
type LeaderboardSuccessResponse struct {
	Draw            int64       `json:"draw"`            // The draw counter that this object is a response to - from the draw parameter sent as part of the data request.
	RecordsTotal    int64       `json:"recordsTotal"`    // Total records in db, before filtering
	RecordsFiltered int64       `json:"recordsFiltered"` // Total records in db, after filtering
	Data            []*GameUser `json:"data"`            // The data to be displayed in the table //TODO: check if interface works
}

// NewLeaderboardSuccessResponse creates a new success response to be consumed by the databale api
func NewLeaderboardSuccessResponse(draw int64, gameUsers []*GameUser, totalCount, totalFiltered int64) LeaderboardSuccessResponse {
	return LeaderboardSuccessResponse{
		Draw:            draw,
		RecordsTotal:    totalCount,
		RecordsFiltered: totalFiltered,
		Data:            gameUsers,
	}
}
