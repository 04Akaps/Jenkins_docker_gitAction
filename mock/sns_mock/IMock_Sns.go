package sns_mock

type ISnsMockInterface interface {
	NewPost(interface{}) error
}
