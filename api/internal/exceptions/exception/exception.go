package exception

type Exception struct {
	err error
}

func New(obj map[string]interface{}) *Exception {

	return &Exception{}
}