package main

import "fmt"
import "encoding/base64"
import "crypto/sha1"

func main(){
  var nama = "Andi Muhammad Tamrin"
  var encodeString = base64.StdEncoding.EncodeToString([]byte(nama))
  fmt.Println("Hasil Encoding string : ", encodeString)

  var sha = sha1.New()
  sha.Write([]byte(nama))
  var enskripsi = sha.Sum(nil)
  var stringenkripsi = fmt.Sprintf("%x", enskripsi)

  fmt.Println("Hasil enskripsi string = ",stringenkripsi)
}
