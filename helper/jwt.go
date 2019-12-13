package helper

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"github.com/anaabdi/todo-go/config"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	jwtSigningKey          = getPrivateKey()
	JWTVerifyingKey        = getPublicKey()
	AccessTokenExpiration  = config.GetInt64("ACCESS_TOKEN_LIFETIME_IN_SEC", 300)
	RefreshTokenExpiration = config.GetInt64("REFRESH_TOKEN_LIFETIME_IN_SEC", 86400)
)

type JWTClaims struct {
	Username   string `json:"uname,omitempty"`
	UserID     string `json:"uid,omitempty"`
	Permission string `json:"permission,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	jwt.StandardClaims
}

func GenerateJWT(claimsSet JWTClaims, tokenLifetimeInSec int64) (string, error) {
	expirationTime := TimeNext(tokenLifetimeInSec)
	log.Print(expirationTime)
	claimsSet.ExpiresAt = expirationTime.UTC().Unix()
	claimsSet.Issuer = "mapsvc"

	t := jwt.NewWithClaims(jwt.SigningMethodRS512, claimsSet)
	return t.SignedString(jwtSigningKey)
}

/*
func ValidateJWT(token string) (claims map[string]interface{}, err error) {
	var jwtObj jwt.JWT
	if jwtObj, err = jws.ParseJWT([]byte(token)); err != nil {
		return
	}
	validator := jws.NewValidator(nil, 0, 0, func(c jws.Claims) error {
		// Custom claims validation, none for now
		claims = c
		return nil
	})
	err = jwtObj.Validate(&jwtPrivateKey.PublicKey, crypto.SigningMethodRS256, validator)
	return
} */

/* func GenerateJWTKey() ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	x509PublicKey, err := x509.MarshalPKIXPublicKey(privateKey.Public())
	if err != nil {
		return nil, err
	}

	fmt.Println("Generated JWT public key (base64 encoded):")
	fmt.Println(base64.StdEncoding.EncodeToString(x509PublicKey))

	return x509.MarshalPKCS1PrivateKey(privateKey), nil
}

func InitJWT(privateKeyBytes []byte) (err error) {
	if jwtPrivateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBytes); err != nil {
		return
	}
	jwtPrivateKey.Precompute()
	return
}

func GetJWTPublicKeyBase64() (string, error) {
	x509PublicKey, err := x509.MarshalPKIXPublicKey(jwtPrivateKey.Public())
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(x509PublicKey), nil
} */

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open("/var/conf/jwt/jwtRS256.key")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open("/var/conf/jwt/jwtRS256.key.pub")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		panic(err)
	}

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
