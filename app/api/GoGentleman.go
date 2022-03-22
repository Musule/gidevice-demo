package Api


import (
	"fmt"
  
	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/body"
	"gopkg.in/h2non/gentleman.v2/mux"
	"gopkg.in/h2non/gentleman.v2/plugins/url"
  
  )


  /*
 	[API]
	 https://pkg.go.dev/gopkg.in/h2non/gentleman.v2 
  */

  func GentlemanSampleRequest(){
	  // Create a new client
  cli := gentleman.New()

  // Define the Base URL
  cli.URL("http://httpbin.org/post")

  // Create a new request based on the current client
  req := cli.Request()

  // Method to be used
  req.Method("POST")

  // Define the JSON payload via body plugin
  data := map[string]string{"foo": "bar"}
  req.Use(body.JSON(data))

  // Perform the request
  res, err := req.Send()
  // 
  if err != nil {
    fmt.Printf("Request error: %s\n", err)
    return
  }
  // 
  if !res.Ok {
    fmt.Printf("Invalid server response: %d\n", res.StatusCode)
    return
  }

  // 响应状态码
  fmt.Printf("Status: %d\n", res.StatusCode)
  // 响应内容
  fmt.Printf("Body: %s", res.String())
  }


  func GentlemanSendJsonBody()  {
	  // Create a new client
  cli := gentleman.New()

  // Define the Base URL
  cli.URL("http://httpbin.org/post")

  // Create a new request based on the current client
  req := cli.Request()

  // Method to be used
  req.Method("POST")

  // Define the JSON payload via body plugin
  data := map[string]string{"foo": "bar"}
  req.Use(body.JSON(data))

  // Perform the request
  res, err := req.Send()
  if err != nil {
    fmt.Printf("Request error: %s\n", err)
    return
  }
  if !res.Ok {
    fmt.Printf("Invalid server response: %d\n", res.StatusCode)
    return
  }

  fmt.Printf("Status: %d\n", res.StatusCode)
  fmt.Printf("Body: %s", res.String())
	  
  }

func GentlemanCompositionViaMultiplexer()  {
	 // Create a new client
	 cli := gentleman.New()

	 // Define the server url (must be first)
	 cli.Use(url.URL("http://httpbin.org"))
   
	 // Create a new multiplexer based on multiple matchers
	 mx := mux.If(mux.Method("GET"), mux.Host("httpbin.org"))
   
	 // Attach a custom plugin on the multiplexer that will be executed if the matchers passes
	 mx.Use(url.Path("/headers"))
   
	 // Attach the multiplexer on the main client
	 cli.Use(mx)
   
	 // Perform the request
	 res, err := cli.Request().Send()
	 if err != nil {
	   fmt.Printf("Request error: %s\n", err)
	   return
	 }
	 if !res.Ok {
	   fmt.Printf("Invalid server response: %d\n", res.StatusCode)
	   return
	 }
   
	 fmt.Printf("Status: %d\n", res.StatusCode)
	 fmt.Printf("Body: %s", res.String())
}