package adduptok

var addUpToKCases = []struct {
	description string
	list        []int
	k           int
	expected    bool
}{
	{
		description: "1st case",
		list:        []int{10, 15, 3, 7},
		k:           17,
		expected:    true,
	},
	{
		description: "2nd case",
		list:        []int{100, 150, 43, 17},
		k:           17,
		expected:    false,
	},
	{
		description: "3nd case",
		list:        []int{99, 350, 43, 11},
		k:           110,
		expected:    true,
	},
}
