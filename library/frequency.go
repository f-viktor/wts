package wts

import(
  "fmt"
  "sort"
)

// Turning the map into this structure
type frequency struct {
    Key   rune
    Value int
}

func FrequencyMain(message string){
  sayHi("Frequency Analysis", message)

  ordered := CountRunes(message)

  for _, item := range ordered {
    fmt.Printf("%c => %d\n",item.Key, item.Value)
  }
}

// frequency
func CountRunes(s string) []frequency{

  uniq := getUniq(s)
  ordered := orderBy(uniq)


  return ordered
}

func getUniq(s string) map[rune]int{
  freqmap := make(map[rune]int)

  // If the key(values of the slice) is not equal
  // to the already present value in new slice (list)
  // then we map it. else we add one
  for _, entry := range []rune(s) {
      if (freqmap[entry] == 0) {
        freqmap[entry] = 1
      } else {
        freqmap[entry]++
      }
  }

  return freqmap
}

// convert to slice and order by frequency
func orderBy(uniq map[rune]int) []frequency {
  var ss []frequency
  for k, v := range uniq {
      ss = append(ss, frequency{k, v})
  }

  // Then sorting the slice by value, higher first.
  sort.Slice(ss, func(i, j int) bool {
      return ss[i].Value > ss[j].Value
  })

  return ss
}
