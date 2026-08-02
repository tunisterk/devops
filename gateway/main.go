package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://localhost:8081/employees")

	if err != nil {
		http.Error(w, "Employee Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	fmt.Fprintf(w, `
<html>
<head>
<title>Employee Directory</title>
</head>

<body>

<h1>Employee Directory</h1>

<pre>%s</pre>

</body>
</html>
`, string(body))

}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gateway OK"))
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)

	log.Println("Gateway running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}