package main

import (
	"fmt"

	"github.com/haobird/learngo/helloecdsa/signature"
)

var (
	pubker = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7mU/3m2URGywg1uYL9rUosMb/i54
4kzJ5xHg/cZ9oM9wC3V3+OMKw+4Tkxs8lMznwMHxxAbVcZUOeRpc8jgrjQ==
-----END PUBLIC KEY-----`
	priker = `-----BEGIN PRIVATE KEY-----
MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCDL9833IX0JVALjoeB1
StEOUtTnLhHsQ3eds2Y47IEE+A==
-----END PRIVATE KEY-----`
	publicKey  = "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE7mU_3m2URGywg1uYL9rUosMb_i544kzJ5xHg_cZ9oM9wC3V3-OMKw-4Tkxs8lMznwMHxxAbVcZUOeRpc8jgrjQ"
	privateKey = "MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCDL9833IX0JVALjoeB1StEOUtTnLhHsQ3eds2Y47IEE-A"
	data       = "USER0001202004151958010871292app0001202004161020152918451test8888123456"
)

func main() {
	fmt.Println(data)
	// signature.TransPem2Base64(priker)
	privateKey = "MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCC70S8JiFtDqOl1_fBSiZLEgpIiQID2HZPoHhUrT3BF3g"
	signature.TransBase642Pem(privateKey, "PRIVATE KEY")
	// privateKey = "MEECAQAwEwYHKoZIzj0CAQYIKoZIzj0DAQcEJzAlAgEBBCDL9833IX0JVALjoeB1StEOUtTnLhHsQ3eds2Y47IEE-A"
	// sign, err1 := signature.Sign(data, privateKey)
	// fmt.Println(sign)
	// fmt.Println(err1)

	// fmt.Println(signature.Hash(data))
	// fmt.Println(signature.GetSHA256HASH(data))

}
