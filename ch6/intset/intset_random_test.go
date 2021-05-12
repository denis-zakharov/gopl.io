package intset

import (
	"math/rand"
	"testing"
	"time"
)

// Compare IntSet operations with random input against IntMapSet implementation.

func randomInts(rng *rand.Rand) []int {
	n := rng.Intn(10) // a random length up to 49
	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = rng.Intn(1000)
	}

	return res
}

func createIntSet(ints []int) IntSet {
	var intSet IntSet
	for _, i := range ints {
		intSet.Add(i)
	}
	return intSet
}

func createIntMapSet(ints []int) IntMapSet {
	var intMapSet = CreateIntMapSet()
	for _, i := range ints {
		intMapSet.Add(i)
	}
	return intMapSet
}

func randomSets(rndSet1, rndSet2 []int) (is1, is2 IntSet, ims1, ims2 IntMapSet) {
	is1 = createIntSet(rndSet1)
	is2 = createIntSet(rndSet2)
	ims1 = createIntMapSet(rndSet1)
	ims2 = createIntMapSet(rndSet2)
	return
}

func TestRandomIntSet(t *testing.T) {
	// Initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("PRNG is initialized: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		// inputs
		rndSet1 := randomInts(rng)
		rndSet2 := randomInts(rng)

		// IntersectWith
		is1, is2, ims1, ims2 := randomSets(rndSet1, rndSet2)
		is1Str := is1.String()
		is2Str := is2.String()
		is1.IntersectWith(&is2)
		ims1.IntersectWith(&ims2)
		if !ims1.EqualToIntSet(&is1) {
			t.Errorf("%v IntersectWith %v = %v, wanted %v", is1Str, is2Str, is1.String(), ims1.Elems())
		}

		// UnionWith
		is1, is2, ims1, ims2 = randomSets(rndSet1, rndSet2)
		is1Str = is1.String()
		is2Str = is2.String()
		is1.UnionWith(&is2)
		ims1.UnionWith(&ims2)
		if !ims1.EqualToIntSet(&is1) {
			t.Errorf("%v UnionWith %v = %v, wanted %v", is1Str, is2Str, is1.String(), ims1.Elems())
		}

		// DifferenceWith
		is1, is2, ims1, ims2 = randomSets(rndSet1, rndSet2)
		is1Str = is1.String()
		is2Str = is2.String()
		is1.DifferenceWith(&is2)
		ims1.DifferenceWith(&ims2)
		if !ims1.EqualToIntSet(&is1) {
			t.Errorf("%v DifferenceWith %v = %v, wanted %v", is1Str, is2Str, is1.String(), ims1.Elems())
		}

		// SymmetricDifference
		is1, is2, ims1, ims2 = randomSets(rndSet1, rndSet2)
		is1Str = is1.String()
		is2Str = is2.String()
		is1.SymmetricDifference(&is2)
		ims1.SymmetricDifference(&ims2)
		if !ims1.EqualToIntSet(&is1) {
			t.Errorf("%v SymmetricDiff %v = %v, wanted %v", is1Str, is2Str, is1.String(), ims1.Elems())
		}
	}
}
