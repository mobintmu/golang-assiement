package rearrange

func Rearrange(input string) string {
	FinalString := ""

	myMap := MakeMap(input)

	for counter := 0; len(FinalString) < len(input); counter++ {

		max := MaxMap(myMap)
		FinalString += max.Key
		myMap[max.Key]--
		max = MaxMapExceptBefore(myMap, max.Key)
		if max.Value == 0 {
			break
		}
		FinalString += max.Key
		myMap[max.Key]--

	}

	if len(FinalString) == len(input) {
		return FinalString
	}

	return ""

}

func MakeMap(input string) map[string]int {

	m := make(map[string]int)
	for _, v := range input {
		m[string(v)]++
	}
	return m
}

type Pair struct {
	Key   string
	Value int
}

func MaxMap(data map[string]int) Pair {

	var max Pair
	for char, value := range data {
		if value > max.Value {
			max = Pair{char, value}
		}
	}
	return max
}
func MaxMapExceptBefore(data map[string]int, inputChar string) Pair {

	var max Pair
	for char, value := range data {
		if value > max.Value && char != inputChar {
			max = Pair{char, value}
		}
	}
	return max
}
