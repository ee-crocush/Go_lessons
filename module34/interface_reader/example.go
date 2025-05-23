package interface_reader

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	body := strings.NewReader("hello world")

	req, _ := http.NewRequest("POST", "https://google.com", body)
	client := http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
