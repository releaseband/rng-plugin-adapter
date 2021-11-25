package adapter

type SeedGenerator interface {
	Seed() int64
}

type RNG interface {
	Random(min, max uint32) uint32
	YesNo() bool
	SetSeed(seed int64)
}


type POOL interface {
	Get() (RNG, int64, error)
}

type PoolMaker interface {
	Make(generator SeedGenerator) POOL
}