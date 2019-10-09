package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

type opts func(con *controller) *controller

func main() {
	//mux := http.NewServeMux()
	//mux.Handle("/raj", http.HandlerFunc(raj))
	//http.ListenAndServe(":8080", mux)

	route := mux.NewRouter()
	//var opt1 opts = func(con *controller) *controller {
	//	if con != nil {
	//		con.thatconfig = "somethavalue"
	//	}
	//	return con
	//}
	//controllers := newcontroller(opt1, opts(func(con *controller) *controller {
	//	if con != nil {
	//		con.thisconfig = "somethisvalue"
	//	}
	//	return con
	//}))
	route.HandleFunc("/raj", timeMiddleWare(loggingMiddleware(raj))).Methods("GET")
	logrus.Debugf("hello %s", "raj")
	//route.HandleFunc("/rajcontroller", controllers.rajcont).Methods("GET")
	http.ListenAndServe(":8080", route)

}

//func rajcont(writer http.ResponseWriter, request *http.Request) {
//
//}
func timeMiddleWare(fun http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		////fmt.Println("time started")
		////t := time.Now()
		fun(writer, request)
		////fmt.Println("time ended ", time.Since(t))
	}
}
func loggingMiddleware(fun http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		//using fmt instead of log
		fmt.Println("this is before calling")
		fun(writer, request)
		fmt.Println("this is after calling ")
	}
}
func (cont *controller) rajcont(w http.ResponseWriter, r *http.Request) {
	time.Sleep(10 * time.Second)
	w.Write([]byte(" new hello raj" + cont.thatconfig + " " + cont.thisconfig))

}

func raj(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(callSomeLogic())
	////longRunningCalculation := func(ctx context.Context) {
	//for i := 0; ; i++ {
	//	fmt.Println("this will be keep on coming")
	//	select {
	//	case <-r.Context().Done():
	//		fmt.Println("done called")
	//		return
	//	default:
	//		time.Sleep(1 * time.Second)
	//		fmt.Printf("Worker %d \n", i)
	//	}
	//}
	//}

	// the context is canceled when the ServeHTTP method returns
	//go longRunningCalculation(r.Context())

	// give some time for longRunningCalculation to do some work
	//time.Sleep(2 * time.Second)
	fmt.Println("is this priting")
	io.WriteString(w, "bazinga!")
}

type controller struct {
	thisconfig string
	thatconfig string
}

func callSomeLogic() int32 {
	somevalue := rand.Int31n(19) * 8923
	return somevalue
	//fmt.Println(somevalue)
}

func newcontroller(options ...opts) *controller {
	cont := &controller{}
	for _, val := range options {
		val(cont)
	}
	return cont
}
