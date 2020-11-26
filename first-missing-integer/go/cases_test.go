package firstmissinginteger

var testCases = []struct {
	description string
	list        []int
	expected    int
}{
	{
		description: "1st case",
		list:        []int{3, 4, -1, 1},
		expected:    2,
	},
	{
		description: "2nd case",
		list:        []int{1, 2, 0},
		expected:    3,
	},
}
