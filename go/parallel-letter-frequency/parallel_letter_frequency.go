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

func FrequencyChannelled(s string, c chan FreqMap) {
	c <- Frequency(s)
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(l []string) FreqMap {
	ch := make(chan FreqMap)
	for _, t := range l {
		FrequencyChannelled(t, ch)
	}
	m := FreqMap{}
	for f := range ch {
		for a, b := range f {
			m[a] += b
		}
	}
	return m
}
