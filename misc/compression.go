package main

import (
	"compress/gzip"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"
)

//Kode for funksjoner og for programmet.
//for å kjøre programmet må du skrive "go run compression.go -filename"
//filen må inneholde hex verdi uten noen form for mellomrom eller andre symboler som kan gi den en utf-8 invalid byte hvis ikke vil programmet panic og close.

func main() {
	args := os.Args
	file := args[1]

	d := returnHexASCII(file)
	a := returnBase64(d)
	compressBase64(a)

}

func readFile(file string) string {
	b, err := ioutil.ReadFile(file)
	fileValue := len(b)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	if fileValue < 125 {
		fmt.Println("Hex string:", str)
	} else {
		fmt.Println("Hex stringen er på", fileValue, "tegn")
	}
	return str
}

// Returnere en ascii/utf8 representasjon
func returnHexASCII(hex1 string) string {
	fileRead := readFile(hex1)

	ascii, err := hex.DecodeString(fileRead)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", ascii)

	fmt.Println("Fra hex til ASCII:", str)
	fmt.Printf("Størrelse i byte: %T, %d \n", str, unsafe.Sizeof(str))
	fmt.Println("Lengde: ", len(str))
	fmt.Println("")

	return str
}

// Returnere en base64 representasjon
func returnBase64(s string) string {

	// ASCII til base64
	e := base64.StdEncoding.EncodeToString([]byte(s))
	// Lengden av base64 strengen
	r := base64.StdEncoding.EncodedLen(len(e))

	if r < 100 {
		fmt.Println("Fra ASCII til base64:", e)
	} else {
		fmt.Println("Base64 er på", r, "tegn")
	}

	fmt.Printf("Størrelse i byte for base64: %T, %d \n", e, unsafe.Sizeof(e))
	fmt.Println("Lengden på stringen i base64:", r)

	return e
}

// Komprimerer til .gz
func compressBase64(b64String string) {

	newFile, err := os.Create("compression.gz")
	if err != nil {
		fmt.Print(err)
	}
	w := gzip.NewWriter(newFile)

	fmt.Println("Komprimerer nå til .gz")

	w.Write([]byte(b64String))

	w.Close()
}
