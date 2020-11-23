package nearestLarger

var nearestLargerCases = []struct {
  description string
  input []int
  element int
  expected int
}{
  {
    description: "1st case",
    input : []int{4, 1, 3, 5, 6},
    element : 0,
    expected : 3,
  },
  {
    description: "2nd case",
    input : []int{1, 3, 5, 4, 12, 20, 6},
    element : 2,
    expected : 4,
  },
  {
    description: "3th case",
    input : []int{1, 17, 33, 50, 6, 7, 13, 4, 9},
    element : 6,
    expected : 3,
  },
  {
    description: "4th case",
    input : []int{1, 17, 33, 50, 6, 7, 13, 4, 9},
    element : 3,
    expected : -1,
  },
  {
    description: "5th case",
    input : []int{13, 7, 5, 9, 4, 3, 20, 50, 1},
    element : 2,
    expected : 3,
  },
  {
    description: "6th case",
    input : []int{8, 3, 5, 4, 12, 20, 6},
    element : 2,
    expected : 4,
  },
  {
    description: "7th case",
    input : []int{6},
    element : 0,
    expected : -1,
  },
  {
    description: "8th case",
    input : []int{13, 7, 5, 9, 4, 3, 20, 50, 1},
    element : 20,
    expected : -1,
  },
  {
    description: "9th case",
    input : []int{},
    element : 0,
    expected : -1,
  },
}
