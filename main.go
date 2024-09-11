package main

immport (
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "enconding/pem"
  "os"
)

func main() {
  privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
  if err != nil {
    panic(err)
  }

  publicKey := &privateKey.PublicKey

  prrivateKeyBytes := x509.MarshalPKCS1PrivateKey(prrivateKey)
  privateKeyPEM := 
}
