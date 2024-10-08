package main

import (
	"fmt"
	bip39 "github.com/tyler-smith/go-bip39"
	"strings"
)

func main() {
	/* // Generate a mnemonic for memorization or user-friendly seeds
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
	seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

	masterKey, _ := bip32.NewMasterKey(seed)
	publicKey := masterKey.PublicKey()

	// Display mnemonic and keys
	fmt.Println("Mnemonic: ", mnemonic)
	fmt.Println("Master private key: ", masterKey)
	fmt.Println("Master public key: ", publicKey) */
	/*words, err := generateMnemonicWords(30)
	if err != nil {
		fmt.Printf("生成的助记词数量为:%d, 助记词为:%v \n", len(words), words)
		fmt.Println(err)
	}*/
	words := gen30Words()
	fmt.Printf("选取的 %d 个单词是：%v", len(words), words)
	// testRemoveDuplicates()
}

func generateMnemonicWords(count int) ([]string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		fmt.Println(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, err
	}

	words := strings.Split(mnemonic, " ")
	if len(words) < count {
		return words, fmt.Errorf("not enough words generated")
	}
	fmt.Printf("生成的助记词数量为:%d, 助记词为:%v \n", len(words), words)
	return words[:count], nil
}

func gen30Words() []string {
	words := make([]string, 0, 48)

	for {
		entropy, err := bip39.NewEntropy(256)
		if err != nil {
			fmt.Println(err)
			return words
		}
		mnemonic1, err := bip39.NewMnemonic(entropy)
		if err != nil {
			fmt.Println(err)
			return words
		}
		words = append(words, strings.Split(mnemonic1, " ")...)
		if len(words) >= 48 {
			break
		}
	}
	fmt.Printf("一共生成的助记词数量 %d, 分别是:%v \n", len(words), words)

	// 去重
	words = RemoveDuplicates(words)
	// 取前 30 个
	return words[:30]
}

func testRemoveDuplicates() {
	mnemonics := []string{"edge", "now", "blur", "luggage", "motor", "wire", "pistol", "camera", "pole", "piece", "genuine", "decorate", "happy", "liberty", "estate", "achieve", "motor", "pistol", "piece", "happy", "estate"}
	fmt.Printf("remove 前 size=%d, mnemonics=%v \n", len(mnemonics), mnemonics)
	mnemonics = RemoveDuplicates(mnemonics)
	fmt.Printf("remove 后 size=%d, mnemonics=%v \n", len(mnemonics), mnemonics)
}

func RemoveDuplicates(strs []string) []string {
	// Use a map to track seen strings.
	seen := make(map[string]struct{}, 48)
	var result []string

	for _, str := range strs {
		// If the string is not in the map, add it to the result and mark it as seen.
		if _, exists := seen[str]; !exists {
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}
	fmt.Printf("result 数量为：%d, result=%v \n", len(result), result)
	return result
}
