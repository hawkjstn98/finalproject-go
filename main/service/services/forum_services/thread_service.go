package forum_services

import(
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"log"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
	"strconv"
)

func GetThreadPage() string {
	threads, _ := thread_repository.GetThreadPage()

	threadsPage := MapThreadToPage(threads)

	result, _ := json.Marshal(threadsPage)
	return string(result)
}

func GetThreadCategoryPage(category *request.ThreadCategoryRequest) string{
	threads := thread_repository.GetThreadCategory(category.Category)
	threadsPage := MapThreadToPage(threads)
	result, _ := json.Marshal(threadsPage)
	log.Println(category)
	return string(result)
}

func MapThreadToPage(threads []*forum.Thread) (threadsPage []forum.Thread) {
	for i, thread := range threads{
		log.Println("thread: ", threads[i])
		var currThread forum.Thread
		imageLink := user_repository.GetUserImage(threads[i].MakerUsername)

		currThread.Id = threads[i].Id
		currThread.Timestamp = threads[i].Timestamp
		currThread.Name = threads[i].Name
		currThread.Category = threads[i].Category
		currThread.MakerUsername = threads[i].MakerUsername
		currThread.MakerImage = imageLink
		currThread.Description = threads[i].Description
		id, _ := strconv.Atoi(thread.Id.String())
		count, _ := thread_repository.GetThreadDetail(id)
		currThread.CommentCount = count

		threadsPage = append(threadsPage, currThread)
	}
	return threadsPage
}

func GetMaxPage(category *request.ThreadMaxPageRequest) int {
	var threads []*forum.Thread

	if "" == category.Category {
		threads, _ = thread_repository.GetThreadPage()
	} else{
		threads = thread_repository.GetThreadCategory(category.Category)
	}
	threadsPage := MapThreadToPage(threads)
	log.Println("Category: ", len(threadsPage))
	if len(threadsPage) % 10 == 0 {
		return len(threadsPage)/10
	} else {
		return (len(threadsPage)/10) + 1
	}
}