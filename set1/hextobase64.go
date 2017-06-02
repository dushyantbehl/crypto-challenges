package crypto_set1

import (
    "fmt"
    "encoding/hex"
)

const base64Encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func reverseByteSlice(a []byte) {
    for i := len(a)/2-1; i >= 0; i-- {
        opp := len(a)-1-i
        a[i], a[opp] = a[opp], a[i]
    }
}

func HexToBase64(in string) (string, error) {
    out := ""
    padding := ""
    extra := len(in)%3

    bytes, err := hex.DecodeString(in)
    if err != nil {
        return out, fmt.Errorf("Failed to decode hex string\n")
    }

    if extra == 2 {
      bytes = append(bytes, 0)
      padding = "="
    } else if extra == 1 {
      bytes = append(bytes, 0, 0)
      padding = "=="
    }

    for i := len(bytes); i>0; i-=3 {
        val := (uint32(bytes[i-1])<<0 |
                uint32(bytes[i-2])<<8 |
                uint32(bytes[i-3])<<16)

        /* Now we extract 6 bits at a time from the 24 bit number
           to generate base64encoding */
        for j:=0; j<4; j++ {
            out = string(base64Encoding[val&0x3f]) + out
            val >>= 6
        }
    }

    n := len(out) - len(padding)
    out = out[0:n]+padding
    return out, nil
}
