package main

import (
	"fmt"
	"net/http"
	"math"
)

func raizQuadrada() string {
 var x = 0.0001;

 for  i := 0; i < 10; i++ {
	 x += math.Sqrt(x*x)
 } 

 return "Code.education Rocks!";
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, raizQuadrada() )
	});

	fmt.Println(http.ListenAndServe(":80", nil))

}