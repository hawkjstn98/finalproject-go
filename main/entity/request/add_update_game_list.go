package request

type AddOrUpdateGameListRequest struct {
	Username string `json:"username" bson:"username"`
	GameList []string `json:"gameList" bson:"gameList"`
}
