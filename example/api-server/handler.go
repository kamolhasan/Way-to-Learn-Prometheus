package server

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/macaron.v1"
)

func GetBooks(ctx *macaron.Context) {
	log.Println("received Get(all) request from: " + ctx.Req.RemoteAddr)

	var books []Book
	if err := Engine.Find(&books); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	bookList := BookList{
		Items: books,
	}
	ctx.JSON(http.StatusOK, bookList)

	// prometheus
	prom_httpRequestTotal.With(prometheus.Labels{"method":"GET","code":strconv.Itoa(ctx.Resp.Status())}).Inc()
}

func demo(w http.ResponseWriter, r *http.Request)   {

}
func GetBook(ctx *macaron.Context) {
	log.Println("received Get(single) request from: " + ctx.Req.RemoteAddr)
	book := Book{
		ID: ctx.Params("id"),
	}
	if exist, err := Engine.Get(&book); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else if exist {
		ctx.JSON(http.StatusFound, book)
	} else {
		ctx.JSON(http.StatusNotFound, nil)
	}

	// prometheus
	prom_httpRequestTotal.With(prometheus.Labels{"method":"GET","code":strconv.Itoa(ctx.Resp.Status())}).Inc()
}

func PostBook(ctx *macaron.Context, list BookList) {
	log.Println("received Post request from: " + ctx.Req.RemoteAddr)
	for _, val := range list.Items {
		book := Book{
			ID: val.ID,
		}
		has, err := Engine.Get(&book)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		}
		if !has {
			_, err := Engine.Insert(val)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
			}
		}
	}

	var books []Book
	if err := Engine.Find(&books); err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	bookList := BookList{
		Items: books,
	}
	ctx.JSON(http.StatusOK, bookList)

	// prometheus
	prom_httpRequestTotal.With(prometheus.Labels{"method":"POST","code":strconv.Itoa(ctx.Resp.Status())}).Inc()

}

func UpdateBook(ctx *macaron.Context, book Book) {
	log.Println("received Update request from: " + ctx.Req.RemoteAddr)

	_, err := Engine.ID(ctx.Params("id")).Update(book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, book)

	// prometheus
	prom_httpRequestTotal.With(prometheus.Labels{"method":"POST","code":strconv.Itoa(ctx.Resp.Status())}).Inc()

}

func NotFoundFunc(ctx *macaron.Context) {
	ctx.JSON(http.StatusBadRequest, nil)
	fmt.Println("lsdkajfl;fddddddddddddddddddd")
	prom_notFoundTotal.With(prometheus.Labels{"method":ctx.Req.Method, "URL": ctx.Req.RequestURI}).Inc()
}
