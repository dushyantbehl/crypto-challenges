package crypto_set1

import (
    "fmt"
    "encoding/hex"
)

func FixedXOR(a string, b string) (string, error) {
    var out string
    if len(a) != len(b) {
        return out, fmt.Errorf("Input buffers have unequal length\n")
    }
    a_bytes, err := hex.DecodeString(a)
    if err != nil {
        return out, fmt.Errorf("Failed to decode input string a - (%v)\n",a)
    }
    b_bytes, err := hex.DecodeString(b)
    if err != nil {
        return out, fmt.Errorf("Failed to decode input string b - (%v)\n",b)
    }
    for i, v := range a_bytes {
        a_bytes[i] = v ^ b_bytes[i]
    }
    out = hex.EncodeToString(a_bytes)
    return out, nil
}
