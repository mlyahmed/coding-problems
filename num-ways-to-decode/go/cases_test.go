package numwaystodecode

var testCases = []struct {
	description string
	code        string
	expected    int
}{
	{
		description: "1st case",
		code:        "111",
		expected:    3,
	},
	{
		description: "2nd case",
		code:        "257",
		expected:    2,
	},
	{
		description: "3th case",
		code:        "3214",
		expected:    3,
	},
}
