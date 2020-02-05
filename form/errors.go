package form

type Errors interface {
	Get(key string) []string
}
