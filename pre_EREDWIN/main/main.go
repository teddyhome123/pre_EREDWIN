package main

import (
	"pre_EREDWIN/getdata"
	"log"
    //"math/rand"
    //"os"
    //"time"
	//"fmt"
    "net/http"
	"html/template"
)

type Sarimax_Lstm struct {
	//X軸長度
	Time []string
	//LSTM
	RATIO_6H_lstm []float64
	RATIO_12H_lstm []float64
	RATIO_24H_lstm []float64
	RATIO_48H_lstm []float64
	Newpeople_lstm []float64
	Total_people_lstm []float64
	Total_people2_lstm []float64
	Total_people3_lstm []float64
	Total_edwin_lstm []float64
	//SRMX
	RATIO_6H_srmx []float64
	RATIO_12H_srmx []float64
	RATIO_24H_srmx []float64
	RATIO_48H_srmx []float64
	Newpeople_srmx []float64
	Total_people_srmx []float64
	Total_people2_srmx []float64
	Total_people3_srmx []float64
	Total_edwin_srmx []float64
	//Actual value
	RATIO_6H []float64
	RATIO_12H []float64
	RATIO_24H []float64
	RATIO_48H []float64
	Newpeople []float64
	Total_people []float64
	Total_people2 []float64
	Total_people3 []float64
	Total_edwin []float64

	RMSE_lstm []float64
	RMSE_srmx []float64

}
 
func index(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("./index.html"))
	data := new(Sarimax_Lstm)
	data.Time = getdata.Time()
	data.RATIO_6H_lstm = getdata.RATIO_6H_lstm()
	data.RATIO_12H_lstm = getdata.RATIO_12H_lstm()
	data.RATIO_24H_lstm = getdata.RATIO_24H_lstm()
	data.RATIO_48H_lstm = getdata.RATIO_48H_lstm()
	data.Newpeople_lstm = getdata.Newpeople_lstm()
	data.Total_people_lstm = getdata.Total_people_lstm()
	data.Total_people2_lstm = getdata.Total_people2_lstm()
	data.Total_people3_lstm = getdata.Total_people3_lstm()
	data.Total_edwin_lstm = getdata.Total_edwin_lstm()

	data.RMSE_lstm, data.RMSE_srmx = getdata.RMSE_value()

	
	data.RATIO_6H_srmx = getdata.RATIO_6H_srmx()
	data.RATIO_12H_srmx = getdata.RATIO_12H_srmx()
	data.RATIO_24H_srmx = getdata.RATIO_24H_srmx()
	data.RATIO_48H_srmx = getdata.RATIO_48H_srmx()
	data.Newpeople_srmx = getdata.Newpeople_srmx()
	data.Total_people_srmx 	= getdata.Total_people_srmx()
	data.Total_people2_srmx = getdata.Total_people2_srmx()
	data.Total_people3_srmx = getdata.Total_people3_srmx()
	data.Total_edwin_srmx = getdata.Total_edwin_srmx()
	
	data.RATIO_6H = getdata.RATIO_6H()
	data.RATIO_12H = getdata.RATIO_12H()
	data.RATIO_24H = getdata.RATIO_24H()
	data.RATIO_48H = getdata.RATIO_48H()
	data.Newpeople = getdata.Newpeople()
	data.Total_people  = getdata.Total_people()
	data.Total_people2 = getdata.Total_people2()
	data.Total_people3 = getdata.Total_people3()
	data.Total_edwin = getdata.Total_edwin()

	tmpl.Execute(w, data)
}

func main() {

	http.HandleFunc("/", index) //設定存取的路由
	http.HandleFunc("/index", index)
    err := http.ListenAndServe(":9090", nil) //設定監聽的埠
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
	//getdata.RMSE_lstm()
}