package crawler

type Browser interface {
	Get(url string) ([]Link, error)
}
