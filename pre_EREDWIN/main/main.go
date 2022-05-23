package main

import (
	"pre_EREDWIN/getdata"
    "net/http"
    "log"
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
	Total_people1_lstm []float64
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
	Total_people1_srmx []float64
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
	Total_people1 []float64
	Total_people2 []float64
	Total_people3 []float64
	Total_edwin []float64
	/**標準差**/
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
	data.Total_people1_lstm = getdata.Total_people1_lstm()
	data.Total_people2_lstm = getdata.Total_people2_lstm()
	data.Total_people3_lstm = getdata.Total_people3_lstm()
	data.Total_edwin_lstm = getdata.Total_edwin_lstm()


	data.RMSE_lstm, data.RMSE_srmx = getdata.RMSE_value()
	//data.RMSE_srmx = getdata.RMSE_srmx()

	
	data.RATIO_6H_srmx = getdata.RATIO_6H_srmx()
	data.RATIO_12H_srmx = getdata.RATIO_12H_srmx()
	data.RATIO_24H_srmx = getdata.RATIO_24H_srmx()
	data.RATIO_48H_srmx = getdata.RATIO_48H_srmx()
	data.Newpeople_srmx = getdata.Newpeople_srmx()
	data.Total_people_srmx 	= getdata.Total_people_srmx()
	data.Total_people1_srmx	= getdata.Total_people1_srmx()
	data.Total_people2_srmx = getdata.Total_people2_srmx()
	data.Total_people3_srmx = getdata.Total_people3_srmx()
	data.Total_edwin_srmx = getdata.Total_edwin_srmx()
	
	data.RATIO_6H = getdata.RATIO_6H()
	data.RATIO_12H = getdata.RATIO_12H()
	data.RATIO_24H = getdata.RATIO_24H()
	data.RATIO_48H = getdata.RATIO_48H()
	data.Newpeople = getdata.Newpeople()
	data.Total_people  = getdata.Total_people()
	data.Total_people1  = getdata.Total_people1()
	data.Total_people2 = getdata.Total_people2()
	data.Total_people3 = getdata.Total_people3()
	data.Total_edwin = getdata.Total_edwin()

	tmpl.Execute(w, data)
}

func main() {
	//192.168.226.170
	http.HandleFunc("/", index) //設定存取的路由
	http.HandleFunc("/index", index)
	http.HandleFunc("/total_LSTM_data.csv", downloadFile1)
	http.HandleFunc("/total_SRMX_data.csv", downloadFile2)
	http.HandleFunc("/total_Unify_data.csv", downloadFile3)
    err := http.ListenAndServe("192.168.226.170:9090", nil) //設定監聽的埠
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
	
}

func downloadFile1(w http.ResponseWriter, r *http.Request) {
	file := "total_LSTM_data.csv"

	// 設定此 Header 告訴瀏覽器下載檔案。 如果沒設定則會在新的 tab 開啟檔案。
	// w.Header().Set("Content-Disposition", "attachment; filename="+file)

	http.ServeFile(w, r, file)
}
func downloadFile2(w http.ResponseWriter, r *http.Request) {
	file := "total_SRMX_data.csv"
	http.ServeFile(w, r, file)
}
func downloadFile3(w http.ResponseWriter, r *http.Request) {
	file := "total_Unify_data.csv"
	http.ServeFile(w, r, file)
}