package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
	//"github.com/bradfitz/gomemcache/memcache"
)

const URL string = "http://nginx-svc:8086"

func MemcachedPut(key, value string){
	payload := strings.NewReader(value)
	url := URL + key 
	req, err := http.NewRequest("PUT", url, payload)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}

func MemcachedGet(key string){
	url := URL + key
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	
	s, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("body conetent: %s", s)
	fmt.Println(resp)
}

func MemcachedDel(key string){
	url := URL + key
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
}

func HelloServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(os.Stdout, "%s %s %s\n", r.Method, r.URL, r.Proto)
	
	for k, v := range r.Header {
		fmt.Fprintf(os.Stdout, "Header[%q] = %q\n", k, v)
	}
	
	fmt.Fprintf(os.Stdout, "Host = %q\n", r.Host)
	fmt.Fprintf(os.Stdout, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(os.Stdout, "Form[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "URL = %q\n", r.URL)
	key := fmt.Sprintf("%s", r.URL)
	value := "EXTRACT_HEADERS\r\nContent-Type: text/plain\r\n\r\n" + key
	MemcachedPut(key, value)
	MemcachedGet(key)
}

type fooHandler struct{
}

func (m *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "%s %s %s\n", r.Method, r.URL, r.Proto)
        
        for k, v := range r.Header {
                fmt.Fprintf(os.Stdout, "Header[%q] = %q\n", k, v)
        }       
        
        fmt.Fprintf(os.Stdout, "Host = %q\n", r.Host)
        fmt.Fprintf(os.Stdout, "RemoteAddr = %q\n", r.RemoteAddr)
        if err := r.ParseForm(); err != nil {
                log.Print(err)
        }       
        
        for k, v := range r.Form {
                fmt.Fprintf(os.Stdout, "Form[%q] = %q\n", k, v)
        }       
        
        fmt.Fprintf(w, "URL = %q\n", r.URL)
        key := fmt.Sprintf("%s", r.URL)
        value := "EXTRACT_HEADERS\r\nContent-Type: text/plain\r\n\r\n" + key
        MemcachedPut(key, value)
        MemcachedGet(key)
}

func main() {
	
	go func() {
		err := http.ListenAndServe(":8081", &fooHandler{})
		if err != nil {
			log.Fatal("ListenAndServe failed, error info: ", err)
		}
	}()
	
	err := http.ListenAndServeTLS(":8080", "/etc/nginx/cert.crt", "/etc/nginx/cert.key", &fooHandler{})
	if err != nil {
		log.Fatal("ListenAndServeTLS failed, error info: ", err)
	}
}
