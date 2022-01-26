package wallet

const (
	privateKey    string = "3077020101042049319f52e92f6ba86a384a882f1e9cb3837518cab7441ac5542efcd6b0962bcba00a06082a8648ce3d030107a144034200045522f84857f582bd3ac4ef609487ac0a692baf56de5d70ab7b028d1a709b2dabb3079c95ca23d675a682bfac76dc04ed7ce792dc9950d2f2d243638a24d404c6"
	signature     string = "7e193ff6bbe34ba6249dc2fe2c320fcd8ea1e1522b08c9dec43198a49547bb0728d903021bb70b83462612328de63ae135f8f2c9419868a5353ce9d5f0cbe532"
	hashedMessage string = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
)

func Start() {
	// 위 privateKey, signature, hashedMessage 를  만드는 과정
	// 실제 blockchain, wallet에선 사용되질 않을 코드지만
	// 이런 흐름으로 private key, signature, hashedMessage가 만들어지고 있다는 것을 알면 좋을듯

	// privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)

	// fmt.Printf("%x\n\n", keyAsBytes)

	// hashAsBytes, err := hex.DecodeString(hashedMessage)

	// utils.HandleErr(err)

	// r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	// signature := append(r.Bytes(), s.Bytes()...)

	// fmt.Printf("signature => %x\n\n", signature)

	// utils.HandleErr(err)

	// fmt.Printf("R:%d\nS:%d\n", r, s)

	// ok := ecdsa.Verify(&privateKey.PublicKey, hashAsBytes, r, s)

	// fmt.Println(ok)
}
