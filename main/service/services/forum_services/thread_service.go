package forum_services

import (
	"encoding/json"
	"fmt"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	response2 "github.com/hawkjstn98/FinalProjectEnv/main/entity/response"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"log"
)

func GetThreadPage(page *request.ThreadRequest) string {
	threads := thread_repository.GetThreadPage(page.Page)
	threadsPage := MapThreadToPage(threads)

	end := len(threadsPage)
	start := GetStart(end)

	result, _ := json.Marshal(threadsPage[int(start):end])
	return string(result)
}

func GetThreadCategoryPage(category *request.ThreadCategoryRequest) string {
	threads := thread_repository.GetThreadCategory(category.Category)
	threadsPage := MapThreadToPage(threads)
	result, _ := json.Marshal(threadsPage)
	log.Println(category)
	return string(result)
}

func MapThreadToPage(threads []*forum.Thread) (threadsPage []forum.ThreadPage) {
	for i := range threads {
		var currThread forum.ThreadPage
		imageLink := user_repository.GetUserImage(threads[i].MakerUsername)

		currThread.Id = threads[i].Id
		currThread.Timestamp = threads[i].Timestamp
		currThread.Name = threads[i].Name
		currThread.Category = threads[i].Category
		currThread.MakerUsername = threads[i].MakerUsername
		currThread.MakerImage = imageLink
		currThread.Description = threads[i].Description
		currThread.CommentNumber = len(threads[i].CommentList)

		threadsPage = append(threadsPage, currThread)
	}
	return threadsPage
}

func GetMaxPage(category *request.ThreadMaxPageRequest) int {
	var threads []*forum.Thread

	if "" == category.Category {
		threads = thread_repository.GetThreadPage(category.Page)
	} else {
		threads = thread_repository.GetThreadCategory(category.Category)
	}
	threadsPage := MapThreadToPage(threads)
	//log.Println("Category: ", len(threadsPage))
	if len(threadsPage)%10 == 0 {
		return len(threadsPage) / 10
	} else {
		return (len(threadsPage) / 10) + 1
	}
}

func GetStart(end int) (int) {
	if int(end) > 10 {
		if int(end)%10 == 0 {
			return int(end) - 10
		} else {
			return (int(end) / 10) * 10
		}
	} else {
		return 0
	}
}

func CreateThread(threadRequest *request.CreateThreadRequest) string {
	response := new(response2.CreateThreadResponse)

	fmt.Println("KONTOL ", threadRequest)

	if "" == threadRequest.MakerUsername || "" == threadRequest.Name || "" == threadRequest.Category || "" == threadRequest.Description || threadRequest.Timestamp.Time().IsZero() {
		response.Response.Message = "Invalid Request Format"
		response.Response.ResponseCode = "Failed To Add Or Update PhoneNumber"
	}

	var thread = new(forum.Thread)
	thread.Timestamp = threadRequest.Timestamp
	thread.Description = threadRequest.Description
	thread.Category = threadRequest.Category
	thread.MakerUsername = threadRequest.MakerUsername
	thread.Name = threadRequest.Name

	res, msg := thread_repository.CreateThread(thread)

	if res {
		response.Response.Message = msg
		response.Response.ResponseCode = "Create Thread Success"
	} else {
		response.Response.Message = "Create Thread failed, "+msg
		response.Response.ResponseCode = "Create Thread Failed"
	}

	result, _ := json.Marshal(response)
	return string(result)

}
