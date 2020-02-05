package tags

type Errors interface {
	Get(key string) []string
}
