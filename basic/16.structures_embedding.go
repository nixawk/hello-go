package main

import "fmt"

type Request struct {
	Resource string
}

type ResourceFormatter struct{}

func (r *ResourceFormatter) FormatHTTP(resource string) string {
	return fmt.Sprintf("http://%s", resource)
}

func (r *ResourceFormatter) FormatHTTPS(resource string) string {
	return fmt.Sprintf("https://%s", resource)
}

type AuthenticatedRequest struct {
	Request
	Username, Password string
	ResourceFormatter
}

func main() {
	ar := new(AuthenticatedRequest)
	ar.Resource = "www.example.com/request"
	ar.Username = "sam"
	ar.Password = "P@ssword"

	fmt.Println(ar.FormatHTTP(ar.Resource))
	fmt.Println(ar.FormatHTTPS(ar.Resource))

	fmt.Printf("%#v\n", ar)
}

/* https://www.programming-books.io/essential/go/5c99711b5d2b467d85fe082b1bef3268-composition-and-embedding */
