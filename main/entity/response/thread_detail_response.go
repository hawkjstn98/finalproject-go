package response

import "github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"

type ThreadDetailResponse struct {
	Response    BasicResponse          `json:"response" bson:"response"`
	Thread      *forum.Thread          `json:"thread" bson:"thread"`
	CommentList []*forum.ObjectComment `json:"commentList" bson:"commentList"`
	MaxPage     int64                  `json:"maxPage" bson:"maxPage"`
}

type CreateThreadCommentResponse struct {
	Response BasicResponse `json:"response" bson:"response"`
}