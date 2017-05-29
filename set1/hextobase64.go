package crypto_set1

import (
    "fmt"
    "encoding/hex"
    "encoding/binary"
)

const base64Encoding = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func reverseByteSlice(a []byte) {
    for i := len(a)/2-1; i >= 0; i-- {
        opp := len(a)-1-i
        a[i], a[opp] = a[opp], a[i]
    }
}

func HexToBase64(in string) (out string, err error) {
    out = ""       // Initialize return value

    bytes, err := hex.DecodeString(in)
    if err != nil {
        return out, fmt.Errorf("Failed to decode hex string\n")
    }

    n := len(bytes)
    base := make([]byte, 8) // 64 bit wide.

    /*
        We take 6 bytes i.e. 48 bits at a time.
        A byte is 8 bits and a base64 number is 6 bits so
        48 is the LCM of both.
    */
    for i := n; i>0; i = i-6 {
        j := i-6
        if j < 0 {
            j = 0
        }

        for k:=i; k>j; k-- {
            base[i-k] = bytes[k-1]
        }
        val := binary.LittleEndian.Uint64(base)

        curr := ((i-j)*8)/6

        /* Now we extract 6 bits at a time from the 48 bit number
           to generate base64encoding */
        for ; curr>0; curr-- {
            /* Due to little endian order, next value should
               come before the previous value */
            out = string(base64Encoding[(val & 0x3F)]) + out
            val >>= 6
        }
    }
    return out, nil
}
