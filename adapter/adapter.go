package adapter

type RNG interface {
	Random(min, max uint32) uint32
}

type Random interface {
	RNG
	SetSeed(seed int64)
}

type RngPool interface {
	Get() (RNG, int64, error)
	Put(rng RNG)
}
