package forum_services

import(
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"log"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
)

func GetThreadPage(page *request.ThreadRequest) string {
	threads := thread_repository.GetThreadPage(page.Page)
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
	log.Println(category)
	return string(result)
}

func MapThreadToPage(threads []*forum.Thread) (threadsPage []forum.ThreadPage) {
	for i := range threads{
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

func GetMaxPage(category *request.ThreadCategoryRequest) int {
	var threads []*forum.Thread

	if "" == category.Category {
		threads = thread_repository.GetThreadPage(category.Page)
	} else{
		threads = thread_repository.GetThreadCategory(category)
	}
	threadsPage := MapThreadToPage(threads)
	//log.Println("Category: ", len(threadsPage))
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