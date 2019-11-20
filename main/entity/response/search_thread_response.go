package response

import "github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"

type SearchThreadResponse struct {
	Response BasicResponse   `json:"response" bson:"response"`
	Thread   []*forum.Thread `json:"thread" bson:"thread"`
	MaxPage  int             `json:"maxPage" bson:"maxPage"`
}