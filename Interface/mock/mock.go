package mock

type Reliever struct {
	Context string
}

func (r Reliever) Get(url string) string {
	r.Context = url
	return r.Context
}
