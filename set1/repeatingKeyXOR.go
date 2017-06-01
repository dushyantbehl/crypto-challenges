package crypto_set1

import (
  "encoding/hex"
)

func RepeatingKeyXOR(key string, text string) (hexCipherText string) {
    lenKey := len(key)
    bytes := make([]byte, len(text))

    for i := range bytes {
      char := key[i%lenKey]
      bytes[i] = char ^ text[i]
    }
    return hex.EncodeToString(bytes)
}
