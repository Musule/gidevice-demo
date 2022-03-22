package Api

import(
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func GoNativeHttpClientGet()  {
	url := "http://www.baidu.com"
  	method := "GET"
	payload := strings.NewReader(``)
	client := &http.Client {}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
  	// req.Header.Add("Cookie", "11111")

	res, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Println(string(body))
}

func TestGoHttpClientPost(){
	// 请求地址
	url := "http://"
	// 请求方式
	method := "POST"
	// 请求body 
	payload := strings.NewReader(`{
		"loginName": "",
		"plainPassword": ""
	}`)

  client := &http.Client {
  }
 
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")
  req.Header.Add("Cookie", "")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}