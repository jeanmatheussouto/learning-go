package main

import(
  "fmt"
  "os"
  "strings"
)

func main()  {
  words := os.Args[1:]

  stats := buildStats(words)

  showStats(stats)
}

func buildStats(words []string) map[string]int {
  stats := make(map[string]int)

  for _, word := range words {
    letter := strings.ToUpper(string(word[0]))

    count, found := stats[letter]

    if found {
      stats[letter] = count + 1
    } else {
      stats[letter] = 1
    }
  }
  return stats
}

func showStats(stats map[string]int) {
  fmt.Println("Count of words grouped by initial letter")
  for letter, count := range stats {
    fmt.Printf("%s = %d\n", letter, count)
  }
}
