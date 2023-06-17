# noiserand

The `noiserand` package implements a random number generator that is based on multidimensional noise, using one of the dimensions as a seed, which allows slight alterations to the seed to slightly alter the generated sequence.

## What is this good for?

Well, for example it can be used for procedural generation and if you have found a seed that produces an interesting result, you can explore neighboring seeds which will produce slight variations of the original output.

## Usage

To use the `noiserand` package, first import it into your Go code:

```go
import "github.com/Flokey82/noiserand"
```

Then, create a new `NoiseRand` object with a random seed and a noise function:

```go
r := noiserand.New(12345.0, func(x, y float64) float64 {
    // Replace this with your own noise function.
    return 0.0
})
```

You can then use the `Intn`, `Int63`, and `Float64` methods to generate random numbers:

```go
// Generate a random integer in the range [0, 10).
n := r.Intn(10)

// Generate a random 63-bit integer.
n64 := r.Int63()

// Generate a random float in the range [0.0, 1.0).
f := r.Float64()

// Generate a pseudo-random permutation of the integers [0,n).
p := r.Perm(10)
```

## Example output

```go
package main

import (
	"log"

	"github.com/Flokey82/noiserand"
	"github.com/ojrac/opensimplex-go"
)

func main() {
    // Get a noise function that returns values in the range [0, 1).
	noise := opensimplex.NewNormalized(0)

    // Generate a random sequence of numbers using the given seed.
	seed := 12345.0
	log.Printf("rand with seed %f", seed)
	rnd := noiserand.New(seed, noise.Eval2)
	for i := 0; i < 10; i++ {
		log.Printf("%d: Float64() = %f", i, rnd.Float64())
	}

    // Slightly modify the seed and generate a new sequence of numbers.
	seed += 0.01
	log.Printf("rand with seed %f", seed)
	rnd = noiserand.New(seed, noise.Eval2)
	for i := 0; i < 10; i++ {
		log.Printf("%d: Float64() = %f", i, rnd.Float64())
	}
}
```

Output:

```
 rand with seed 12345.000000
 0: Float64() = 0.416808
 1: Float64() = 0.579680
 2: Float64() = 0.485875
 3: Float64() = 0.456083
 4: Float64() = 0.651193
 5: Float64() = 0.556411
 6: Float64() = 0.464984
 7: Float64() = 0.278727
 8: Float64() = 0.166394
 9: Float64() = 0.252161
 rand with seed 12345.010000
 0: Float64() = 0.414928
 1: Float64() = 0.581802
 2: Float64() = 0.495654
 3: Float64() = 0.462045
 4: Float64() = 0.646941
 5: Float64() = 0.550813
 6: Float64() = 0.458039
 7: Float64() = 0.282221
 8: Float64() = 0.168412
 9: Float64() = 0.253999
```

Note how the second sequence is very similar to the first one, but not identical. Imagine if you have a seed that generates a great procedural world, but you'd like it to be just a little bit different... that's what this package is for.

## Contributions

We welcome contributions to the noiserand package! If you have ideas for improvements or new features, please feel free to submit a pull request on GitHub. We also appreciate bug reports and feedback on how the package is being used. By working together, we can make the noiserand package even better for everyone. Thank you for your support!

## License

The `noiserand` package is licensed under the MIT License. See the LICENSE file for more information.

