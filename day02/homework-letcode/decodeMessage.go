package main

import (
	"fmt"
	"strings"
)

func decodeMessage(key string, message string) string {
	alphabets := "abcdefghijklmnopqrstuvwxyz"
	hashMap := make(map[string]string)
	key = strings.ReplaceAll(key, " ", "")
	fmt.Println(key)
	for i := 0; i < len(alphabets); i++ {
		hashMap[string(key[i])] = string(alphabets[i])
		fmt.Println(string(key[i]), string(alphabets[i]))
	}
	// key = qwertyui
	// hashMap[q] = a
	// hashMap[w] = b
	//....

	var answer string
	for i := 0; i < len(message); i++ {
		if string(message[i]) == " " {
			answer += string(message[i])
			fmt.Println(answer)
			continue
		}
		if value, ok := hashMap[string(message[i])]; ok {
			answer += value
		}
	}
	return answer
}

func main() {
	fmt.Println(decodeMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb"))
}
