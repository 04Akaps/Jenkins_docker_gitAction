package sns_mock

type SnsMock struct {
	IMock ISnsMockInterface
}

func (s *SnsMock) Use(d interface{}) error {
	return s.IMock.NewPost(d)
}
