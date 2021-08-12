package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	//log 包
	/*log.Print("Hey, I'm a log!")
	fmt.Print("Can you see me?")

	//log.Panic("Hey, I'm an error log!")
	fmt.Print("Can you see me?")

	log.SetPrefix("main(): ")
	log.Print("Hey, I'm a log!")*/
	//log.Fatal("Hey, I'm an error log!")

	//记录到文件
	/*file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Hey, I'm a log!")*/


	//记录框架
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey! I'm a log message!!!!")


	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
}
