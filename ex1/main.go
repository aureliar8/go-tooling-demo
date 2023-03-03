package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()
	part1()
	log.Println("Part 1 done")
	part2()
	log.Println("Part 2 done")
	part3()
	log.Println("Part 3 done")
	log.Println("Finished sucessfully in", time.Now().Sub(start))
}
