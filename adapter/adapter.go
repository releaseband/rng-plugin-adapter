package adapter

type RNG interface {
	Random(min, max uint32) uint32
}

type Random interface {
	RNG
	SetSeed(seed int64)
	RandomData(n uint) []byte
}