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
  privateKeyPEM := pem.EncodeToMemory(&pem.Block{
    Type: "RSA PRIVATE KEY",
    Bytes: prrivateKeyBytes,
  }) 
  err = os.WriteFile("private.pem", privateKeyPEM, 0644)
  if err != nil {
    panic(err)
  }

  publicKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        panic(err)
    }
    publicKeyPEM := pem.EncodeToMemory(&pem.Block{
        Type:  "RSA PUBLIC KEY",
        Bytes: publicKeyBytes,
    })
    err = os.WriteFile("public.pem", publicKeyPEM, 0644)
    if err != nil {
        panic(err)
  }
}
