package main

var aa = "G"

func main() {
	nn()
	mm()
}

func nn() {
	print(aa)
}

func mm() {
	aa = "O"
	print(aa)
}