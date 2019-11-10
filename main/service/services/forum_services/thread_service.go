package forum_services

import (
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
)

func GetThreadPage(page *request.ThreadRequest) string {
	threads, _ := thread_repository.GetThreadPage(page.Page)
	threadsPage := MapThreadToPage(threads)

	end := len(threadsPage)
	start := GetStart(end)

	result, _ := json.Marshal(threadsPage[int(start):end])
	return string(result)
}

func GetThreadCategoryPage(category *request.ThreadCategoryRequest) string{
	threads := thread_repository.GetThreadCategory(category)
	threadsPage := MapThreadToPage(threads)

	end := len(threadsPage)
	start := GetStart(end)

	result, _ := json.Marshal(threadsPage[int(start):end])
	return string(result)
}

func MapThreadToPage(threads []*forum.Thread) (threadsPage []forum.Thread) {
	for i := range threads{
		var currThread forum.Thread
		imageLink := user_repository.GetUserImage(threads[i].MakerUsername)

		currThread.Id = threads[i].Id
		currThread.Timestamp = threads[i].Id.Timestamp()
		currThread.Name = threads[i].Name
		currThread.Category = threads[i].Category
		currThread.MakerUsername = threads[i].MakerUsername
		currThread.MakerImage = imageLink
		currThread.Description = threads[i].Description
		currThread.CommentCount = thread_repository.GetCommentCount(threads[i].Id.Hex())

		threadsPage = append(threadsPage, currThread)
	}
	return threadsPage
}

func GetMaxPage(category *request.ThreadCategoryRequest) int {
	var threads []*forum.Thread

	if "" == category.Category {
		threads, _ = thread_repository.GetThreadPage(category.Page)
	} else{
		threads = thread_repository.GetThreadCategory(category)
	}
	threadsPage := MapThreadToPage(threads)
	if len(threadsPage) % 10 == 0 {
		return len(threadsPage)/10
	} else {
		return (len(threadsPage)/10) + 1
	}
}

func GetStart(end int) (int){
	if int(end) > 10{
		if int(end) % 10 == 0{
			return int(end) - 10
		} else {
			return (int(end) / 10) * 10
		}
	} else {
		return 0
	}
}
