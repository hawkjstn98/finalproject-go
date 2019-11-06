package forum_services

import(
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/user_repository"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/object/forum"
	"log"
	"github.com/hawkjstn98/FinalProjectEnv/main/entity/request"
)

func GetThreadPage() string {
	threads := thread_repository.GetThreadPage()

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

func MapThreadToPage(threads []*forum.Thread) (threadsPage []forum.ThreadPage) {
	for i := range threads{
		log.Println("thread: ", threads[i])
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