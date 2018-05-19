package handler

import
(
	"crypto/rand"
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/Jeffail/gabs"
	"log"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"github.com/gorilla/mux"
)


type Meta struct {
	Title string `json:"title"`
	Price string `json:"price"`
	Image string `json:"image"`
}

type Product struct {
	Url string `json:"url"`
	Meta Meta `json:"meta"`
}

type Request struct {
	ID string `json:"id"`
	Products []Product `json:"products"`
}
var requests []Request

func responseJSON(data interface{}, response http.ResponseWriter, request *http.Request) {
	js, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.Header().Set("Content-Type", "application/json")
	response.Write(js)
}

func randReqID() string {
	b := make([]byte, 4)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func ClearImageUrl(s string) string {
	s = s[2:]
	split := strings.Split(s, "\"")
	return split[0]
}

func ClearTitle(s string) string {
	s = strings.Replace(s,"  ", "", -1)
	s = strings.Replace(s,"\n", "", -1)
	return s
}

func ClearPrice(s string) string {
	s = strings.Replace(s,"  ", "", -1)
	s = strings.Replace(s,"\n", "", -1)
	s = strings.Replace(s,"£", "", -1)
	return s
}
func Parse(url string) Product {
	var product Product
	product.Url = url
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	product.Meta.Title = ClearTitle(doc.Find("#productTitle").Text())
	//product.Meta.Title = doc.Find(".a-size-large").First().Text()
	product.Meta.Price = ClearPrice(doc.Find(".a-color-price").First().Text())

	img, _ := doc.Find(".a-dynamic-image").First().Attr("data-a-dynamic-image")
	product.Meta.Image = ClearImageUrl(img)


	return product
}

func ProcessRequest(urls []string, id string) {
	var req Request
	req.ID = id
	for _, url := range urls {
		product := Parse(url)
		req.Products = append(req.Products, product)
	}
	requests = append(requests, req)
}

func FindRequest(id string) (bool, Request) {
	for _, req := range requests {
		if req.ID == id {
			return true, req
		}
	}
	return false, Request{}
}

var UrlHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var urls []string

	body, _ := ioutil.ReadAll(request.Body)
	json, _ := gabs.ParseJSON(body)
	children, _ := json.S("urls").Children()

	for _, child := range children {
		urls = append(urls, child.Data().(string))
	}

	id := randReqID()
	go ProcessRequest(urls, id)
	res := "Your request is processing, requestID=" + id
	fmt.Fprintf(response, res)
})

var AllRequestsHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	responseJSON(requests, response, request)
})

var GetRequestHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	id := strings.ToLower(mux.Vars(request)["id"])
	exist, req := FindRequest(id)
	if (!exist) {
		fmt.Fprintf(response, "Request doesn't exist or in processing")
	} else {
		responseJSON(req, response, request)
	}
})

var InfoHandler = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Send urls from amazon to api/sendurls")
})