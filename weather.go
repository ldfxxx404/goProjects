package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://pro.openweathermap.org/data/2.5/forecast/hourly?lat=55°44′24.00″&lon=37°36′36.00″&appid=e48cbcc68a0436e635a33101817fd625")
	if err != nil {
		fmt.Println("Ошибка создания запроса: ", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка чтения запроса: ", err)
		return
	}
	fmt.Println(string(body))
}
