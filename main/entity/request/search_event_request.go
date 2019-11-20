package request

type SearchEventRequest struct {
	Username 	string 	`json:"username" bson:"username"`
	SearchKey 	string	`json:"searchKey" bson:"searchKey"`
	Page 		int 	`json:"page" bson:"page"`
	Latitude    string  `json:"latitude" bson:"latitude"`
	Longitude   string  `json:"longitude" bson:"longitude"`
}
