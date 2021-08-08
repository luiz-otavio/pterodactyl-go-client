package pterogo

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/luiz-otavio/ptero-go/bootstrap"
	"github.com/luiz-otavio/ptero-go/pterodactyl"
)

func main() {
	client := pterodactyl.NewConnection(
		os.Args[1],
		os.Args[2],
		bootstrap.HTTPOption{
			Timeout:     time.Second.Milliseconds() * 15,
			RequestSize: 30,
		},
	)

	value, err := client.ServerById(1)

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(value.GetObject("attributes").Get("name"))
}
