package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	resp, err := http.Get("https://play.golang.org/p/B27qT17Aeot")
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	rand.Seed(time.Now().UnixNano())
	chars := []rune(
		"abcdefghijklmnopqrstuvwxyz" +
			"0123456789" +
			"0123456789" +
			"ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	length := 16 // taille de la chaine
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))]) // créer une chaine de carractere aléatoire
	}
	str := b.String()
	str += ".html"                                                           // remplacez par ce que vous souhaitez comme extention (.html / .go / .txt)
	file, err := os.OpenFile(str, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600) // nome le fichier créer avec un nom aléatoire
	defer file.Close()                                                       // on ferme automatiquement à la fin de notre programme
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(string(body)) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(body))
}
