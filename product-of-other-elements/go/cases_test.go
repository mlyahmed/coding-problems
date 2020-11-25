package productofotherlements

var productOfOtherElementsCases = []struct {
	description string
	list        []int
	expected    []int
}{
	{
		description: "1st case",
		list:        []int{1, 2, 3, 4, 5},
		expected:    []int{120, 60, 40, 30, 24},
	},
	{
		description: "2nd case",
		list:        []int{3, 2, 1},
		expected:    []int{2, 3, 6},
	},
	{
		description: "3th case",
		list:        []int{13, 15, 8, 24, 41},
		expected:    []int{118080, 102336, 191880, 63960, 37440},
	},
}
