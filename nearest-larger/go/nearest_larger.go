package nearestLarger

func NearestLarger(input []int, element int) int {
  if element >= len(input) {
    return -1
  }

  x := 1
  xnext, xprevious := true, true
  for ; xnext || xprevious; x++  {

    xnext, xprevious = element + x < len(input), element - x >= 0
    nearest := element
    
    if xnext && input[element + x] > input[nearest] {
      nearest = element + x
    }

    if xprevious && input[element - x] > input[nearest] {
      nearest = element -x
    }

    if nearest != element {
      return nearest
    }

  }

  return -1
}
