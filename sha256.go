package main


import (
	"io/ioutil"
	"os"
	"fmt"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"unicode"
)


func getSha(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err == nil {
		var sha = sha256.Sum256(bytes)
		return hex.EncodeToString(sha[:])
	}
	return ""
}

func printNoFileFound(fileName string) {
	fmt.Println("No file found", fileName)
}

func isHash(hash string) bool {
	var strLen = len(hash) == 64
	if(!strLen) {
		return false
	}
	
	for _, ch := range hash {
		if !unicode.IsLetter(ch) && !unicode.IsNumber(ch) {
			return false
		}
	}
	
	return true
}

func main() {
	var args = os.Args[1:]
	var argsLength = len(args)
	if argsLength == 1 {
		var firstArg = args[0]
		var hash = getSha(firstArg)
		if hash == "" {
			printNoFileFound(firstArg)
		} else {
			fmt.Println(hash)
		}
	} else if argsLength == 2 {
		var firstArg = args[0]
		var secondArg = args[1]
		var firstHash = getSha(firstArg)
		var secondHash string
		var validHash = isHash(secondArg)
		
		if(validHash) {
			secondHash = strings.ToLower(secondArg)
		} else {
			secondHash = getSha(secondArg)
		}
		
		if firstHash == "" {
			printNoFileFound(firstArg)
		} else if secondHash == "" && !validHash {
			fmt.Println("Invalid hash provided, sha256 hashes should be 64 characters long and alphanumeric.")
		} else if secondHash == "" {
			printNoFileFound(secondArg)
		} else if firstHash == secondHash {
			fmt.Println("Hashes match")
			fmt.Println("Hash:", firstHash)
		} else {
			fmt.Println("Hashes DO NOT MATCH")
			fmt.Println("First:", firstHash)
			fmt.Println("Second:", secondHash)
		}
	} else {
		fmt.Println("sha256 <file1> <file2/hash>")
	}
}