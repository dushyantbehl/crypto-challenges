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
        return
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
        return
    }
    fmt.Printf("TestHexToBase64 - Succeeded\n")
}

func TestDecodeSingleByteXOR(t *testing.T) {
  in := "1b37373331363f78151b7f2b783431333d"
  in += "78397828372D363C78373E783a393b3736"

  _, key, text, err := DecodeSingleByteXOR(in)
  if err != nil || key != 'X' {
    t.Errorf("TestDecodeSingleByteXOR FAILED\n")
    return
  }
  fmt.Printf("TestDecodeSingleByteXOR - Succeeded\n")
  fmt.Printf("Possible Outcome - Encryption key - %v, Decoded text - %v\n",key,text)
}

func TestDetectSingleCharacterXOR(t *testing.T) {
  file := "4.txt"

  key, text, err := DetectSingleCharacterXOR(file)
  if err != nil {
    t.Errorf("TestDetectSingleCharacterXOR FAILED\n")
    return
  }
  fmt.Printf("TestDetectSingleCharacterXOR - Succeeded\n")
  fmt.Printf("Possible Outcome - Encryption key - %v, Decoded text - %v\n",key,text)
}

func TestRepeatingKeyXOR(t *testing.T) {
  in := "Burning 'em, if you ain't quick and nimble\n"
  in += "I go crazy when I hear a cymbal"

  want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d62"
  want += "3d63343c2a26226324272765272a282b2f20430a652e2c65"
  want += "2a3124333a653e2b2027630c692b20283165286326302e27282f"

  got := RepeatingKeyXOR("ICE", in)

  if got != want {
    t.Errorf("TestRepeatingKeyXOR FAILED\n")
    return
  }
  fmt.Printf("TestRepeatingKeyXOR - Succeeded\n")
}
