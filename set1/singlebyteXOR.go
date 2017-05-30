package crypto_set1

import (
  "fmt"
  "encoding/hex"
  "strings"
)

const MaxUint64 = ^uint64(0)

// Should add space and period to this list
// as they are important characters of engligh language as well.
const EnglishCharacters = "abcdefghijklmnopqrstuvwxyz"

// English single letter frequencies (in percent %)
// Taken from - http://practicalcryptography.com/
var english_monogram_freq = map[byte]float64{
  'a':8.55,   'k':0.81,   'u':2.68,
  'b':1.60,   'l':4.21,   'v':1.06,
  'c':3.16,   'm':2.53,   'w':1.83,
  'd':3.87,   'n':7.17,   'x':0.19,
  'e':12.10,  'o':7.47,   'y':1.72,
  'f':2.18,   'p':2.07,   'z':0.11,
  'g':2.09,   'q':0.10,
  'h':4.96,   'r':6.33,
  'i':7.33,   's':6.73,
  'j':0.22,   't':8.94,
}

/*
  The Chi-squared Statistic is a measure of how similar two
  categorical probability distributions are.  If the two distributions
  are identical, the chi-squared statistic is 0, if the distributions
  are very different, some higher number will result
*/
func ChiSquaredStatisticScore(str string) (score uint64) {
  freq := make(map[byte]float64)
  n := len(str)
  score = 0

  for i:=0; i<n; i++ {
    char := str[i]
    if char >= 'A' && char <= 'Z' {  // Convert to lower case
      char = char - 'A' + 'a'
    }
    if v, ok := freq[char]; ok {
      freq[char] = v + 1
    } else {
      freq[char] = 1
    }
  }

  for i:=0; i<len(EnglishCharacters); i++ {
      char := EnglishCharacters[i]
      C, ok := freq[char]
      if !ok {
        C = 0
      }
      E := (english_monogram_freq[char]/float64(100))*float64(n)
      score += uint64((C-E)*(C-E)/E)
  }
  return
}

func DecodeSingleByteXOR(str string) (key byte, text string, err error) {
  str = strings.ToLower(str)  // Convert to lower case
  score := MaxUint64
  msgBytes := make([]byte, len(str))

  inBytes, err := hex.DecodeString(str)
  if err != nil {
    return key, text, fmt.Errorf("Failed to decode input string (%v)\n",str)
  }

  for i:=0; i<len(EnglishCharacters); i++ {
    char := EnglishCharacters[i]
    // Take a character and get the message by reversing XOR.
    for j:=0; j<len(inBytes); j++ {
      msgBytes[j] = inBytes[j] ^ char
    }
    // Convert message to string.
    msg := string(msgBytes)
    chiScore := ChiSquaredStatisticScore(msg)
    if (chiScore < score) {
      key = char
      text = msg
      score = chiScore
    }
  }

  return key, text, nil
}
