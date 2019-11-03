package forum_services

import(
	"encoding/json"
	"github.com/hawkjstn98/FinalProjectEnv/main/repository/thread_repository"
)

func GetThreadPage() string {
	thread := thread_repository.GetThreadPage()
	result, _ := json.Marshal(thread)
	return string(result)
}