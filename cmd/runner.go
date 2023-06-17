package main

import (
	"log"

	"github.com/Flokey82/noiserand"
	"github.com/ojrac/opensimplex-go"
)

func main() {
	noise := opensimplex.NewNormalized(0)
	seed := 12345.0
	log.Printf("rand with seed %f", seed)
	rnd := noiserand.New(seed, noise.Eval2)
	for i := 0; i < 10; i++ {
		log.Printf("%d: Float64() = %f", i, rnd.Float64())
	}
	seed += 0.01
	log.Printf("rand with seed %f", seed)
	rnd = noiserand.New(seed, noise.Eval2)
	for i := 0; i < 10; i++ {
		log.Printf("%d: Float64() = %f", i, rnd.Float64())
	}
}
