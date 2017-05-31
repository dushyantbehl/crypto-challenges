package crypto_set1

import (
  "fmt"
  "encoding/hex"
  "os"
  "bufio"
)

// Very big number
const MaxFloat64 = float64(^uint64(0))

const charStart byte = 32
const charEnd byte = 126

// English single letter frequencies (in percent %)
// Taken from - http://practicalcryptography.com/
var english_monogram_freq = map[byte]float64{
  'a':8.167,   'k':0.772,   'u':2.758,
  'b':1.492,   'l':4.025,   'v':0.978,
  'c':2.782,   'm':2.406,   'w':2.360,
  'd':4.253,   'n':6.749,   'x':0.150,
  'e':12.702,  'o':7.507,   'y':1.974,
  'f':2.228,   'p':1.929,   'z':0.074,
  'g':2.015,   'q':0.095,
  'h':6.094,   'r':5.987,
  'i':6.966,   's':6.327,
  'j':0.153,   't':9.056,
}

var ascii_ignored_characters = map[byte]int {
  // Null, HT, LF, VT, CR, Space
  0:1, 9:1, 10:1, 11:1, 13:1, 32:1,
  //  '!', ',', '.', '?', ''', '"'
  33:1, 44:1, 46:1, 63:1, 39:1, 34:1,
}

/*
  The Chi-squared Statistic is a measure of how similar two
  categorical probability distributions are.  If the two distributions
  are identical, the chi-squared statistic is 0, if the distributions
  are very different, some higher number will result
*/
func ChiSquaredStatisticScore(str string) (float64) {
  var char byte
  var n int = 0
  var score float64 = 0
  freq := make(map[byte]float64)

  for i:=0; i<len(str); i++ {
    char = str[i]
    if _, ok := ascii_ignored_characters[char]; ok {
      continue;
    }
    if char >= 48 && char <= 57 { //Ignore numbers 0-9
      continue;
    }
    if char < 'A' || (char > 'Z' && char < 'a') || char > 'z' {
      // Ignore non alphabetic characters
      return MaxFloat64
    }
    if char >= 'A' && char <= 'Z' {  // Convert to lower case
      char = char - 'A' + 'a'
    }
    if v, ok := freq[char]; ok {
      freq[char] = v + 1
    } else {
      freq[char] = 1
    }
    n++
  }

  for char = 'a'; char <= 'z'; char++ {
      C, ok := freq[char]
      if !ok {
        C = 0
      }
      E := (english_monogram_freq[char]/float64(100))*float64(n)
      score += (C-E)*(C-E)/E
  }
  return score
}

func DecodeSingleByteXOR(str string) (score float64, key byte, text string, err error) {
  score = MaxFloat64
  var char byte

  inBytes, err := hex.DecodeString(str)
  if err != nil {
    err = fmt.Errorf("Failed to decode input string (%v)\n",str)
    return
  }
  msgBytes := make([]byte, len(inBytes))

  for char = charStart; char < charEnd; char++ {
    // Take a character and get the message by reversing XOR.
    for i, v := range inBytes {
      msgBytes[i] = v ^ char
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
  err = nil
  return
}

func DetectSingleCharacterXOR(filename string) (key byte, text string, err error) {

  score := MaxFloat64
  f, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer f.Close() // Discard file close errors

  var lines []string
  scanner := bufio.NewScanner(f)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  for _,line := range lines {
    s, k, t, e := DecodeSingleByteXOR(line)
    if e != nil {
      err = e
      return
    }
    if (s < score) {
      score = s
      key = k
      text = t
    }
  }
  fmt.Printf("Possible Outcome - Encryption key - %v, Decoded text - %v\n",key,text)
  return key, text, nil
}
