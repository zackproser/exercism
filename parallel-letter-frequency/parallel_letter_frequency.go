package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {

	// Create a channel to receive FreqMap results
	res := make(chan FreqMap)

	// Range over input strings, feeding each into the Frequency function, whose
	// results are then sent into the results channel
	for _, s := range l {
		go func(s string) { res <- Frequency(s) }(s)
	}

	finalCounts := FreqMap{}

	// Iterate over all results, incrementing counts in the final FreqMap
	for range l {
		for c, count := range <-res {
			finalCounts[c] += count
		}
	}
	return finalCounts
}
