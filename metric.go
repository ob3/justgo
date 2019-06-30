package justgo

type Metric interface {
	Increment(key string)
}
