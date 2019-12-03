package main


import (
	"io/ioutil"
	"os"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
)


func getSha(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		var sha = sha256.Sum256(bytes)
		return hex.EncodeToString(sha[:])
	}
	return ""
}

func main() {
	var args = os.Args[1:]
	var argsLength = len(args)
	if argsLength == 1 {
		var firstArg = args[0]
		var hash = getSha(firstArg)
		fmt.Println(hash)
	} else if argsLength == 2 {
		var firstArg = args[0]
		var secondArg = args[1]
		var firstHash = getSha(firstArg)
		var secondHash = getSha(secondArg)
		if firstHash == secondHash {
			fmt.Println("File hashes match")
			fmt.Println("Hash:", firstHash)
		} else {
			fmt.Println("File hashes DO NOT MATCH")
			fmt.Println("First file:", firstHash)
			fmt.Println("Second file:", secondHash)
		}
	} else {
		fmt.Println("sha256 <file1> <file2>")
	}
}