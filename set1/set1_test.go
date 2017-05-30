package crypto_set1

import (
    "testing"
    "fmt"
)

func TestFixedXOR(t *testing.T) {
    in1 := "1c0111001f010100061a024b53535009181c"
    in2 := "686974207468652062756c6c277320657965"

    want := "746865206b696420646f6e277420706c6179"

    got, err := FixedXOR(in1, in2)

    if err != nil || got != want {
        t.Errorf("TestFixedXOR FAILED\n")
    }
    fmt.Printf("TestFixedXOR - Succeeded\n")
}

func TestHexToBase64(t *testing.T) {
    in := "49276d206b696c6c696e6720796f7572"
    in += "20627261696e206c696b65206120706f"
    in += "69736f6e6f7573206d757368726f6f6d"

    want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiB"
    want += "saWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

    got, err := HexToBase64(in)

    if err != nil || got != want {
        t.Errorf("TestHexToBase64 FAILED\n")
    }
    fmt.Printf("TestHexToBase64 - Succeeded\n")
}

func TestDecodeSingleByteXOR(t *testing.T) {
  in := "1b37373331363f78151b7f2b783431333d"
  in += "78397828372d363c78373e783a393b3736"

  key, _, err := DecodeSingleByteXOR(in)
  if err != nil || key != 'x' {
    t.Errorf("TestDecodeSingleByteXOR FAILED\n")
  }
  fmt.Printf("TestDecodeSingleByteXOR - Succeeded\n")
}
