package crypto_set1

import (
    "fmt"
    "encoding/hex"
)

func FixedXOR(a string, b string) (out string, err error) {
    out = ""

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

    for i:=0; i<len(a_bytes); i++ {
        a_bytes[i] = a_bytes[i] ^ b_bytes[i]
    }

    out = hex.EncodeToString(a_bytes)

    return out, nil
}
