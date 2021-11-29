package adapter

type RNG interface {
	Random(min, max uint32) uint32
	RandomData(n uint) []byte
}

type Random interface {
	RNG
	SetSeed(seed int64)
}