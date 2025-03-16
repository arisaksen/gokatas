// Print statistics about katas you've done.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const katasMd = "katas.md"

var sortByColumn = flag.Int("c", 1, "sort katas by `column`")
var gokatasRepo = flag.String("r", ".", "path to gokatas repository")
var gokatasRandom = flag.Int("random", 0, "pick random number of katas")

func main() {
	flag.Parse()

	log.SetPrefix("gokatas: ")
	log.SetFlags(0)

	if *gokatasRepo != "." {
		if err := os.Chdir(*gokatasRepo); err != nil {
			log.Fatal(err)
		}
	}

	katas, err := Get()
	if err != nil {
		log.Fatalf("getting katas: %v", err)
	}

	Print(katas, *sortByColumn)

	random := rand.Intn(len(katas) + 1)
	if *gokatasRandom == 1 {
		log.Printf("random katas: %v", katas[random].Name)
		timestamp := time.Now().Format(time.DateOnly)
		f, err := os.OpenFile(katasMd, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err = f.WriteString(fmt.Sprintf("* %v: %s\n", timestamp, katas[random].Name)); err != nil {
			panic(err)
		}
	}
}
