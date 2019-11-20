package request

type SearchThreadRequest struct {
	Username 	string 	`json:"username" bson:"username"`
	SearchKey 	string	`json:"searchKey" bson:"searchKey"`
	Page 		int 	`json:"page" bson:"page"`
}
