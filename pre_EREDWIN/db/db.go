package db

import (
	"database/sql"
	"fmt"
	"log"
	//"time"
	_ "github.com/denisenkom/go-mssqldb"
	//"github.com/go-echarts/go-echarts/opts"

)


func Db(n1 string) []float64 {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		"----", "----", "----", "----")

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}

	defer conn.Close()


	//建立連接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	//通過連接對象執行查詢
	rows, err := db.Query(n1)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()
	
	//var rowsData []AccessRegion 
	
	//遍歷
	var arr1 = []float64{}
	var test1 float64

	for rows.Next() {
		//RATIO_6H_lstm := make([]opts.BarData, 0)
       //var row := new(arr1)
        rows.Scan(&test1)
		//fmt.Println(RATIO_6H_lstm)
        arr1 = append(arr1, test1)
		//fmt.Println(arr1)
    }
	
	return arr1
}

func DbTime(n1 string) []string {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		"----", "----", "----", "----")

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}

	defer conn.Close()


	//建立連接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	//通過連接對象執行查詢
	rows, err := db.Query(n1)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()
	
	
	//遍歷
	var arr1 = []string{}
	var test1 string
	for rows.Next() {
		//RATIO_6H_lstm := make([]opts.BarData, 0)
       //var row := new(arr1)
        rows.Scan(&test1)
		//fmt.Println(RATIO_6H_lstm)
        arr1 = append(arr1, test1)
		//fmt.Println(arr1)
    }
	
	return arr1
}

func Rmse(n1 string) []float64 {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		"----", "----", "----", "----")

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}

	defer conn.Close()


	//建立連接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	//通過連接對象執行查詢
	rows, err := db.Query(n1)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()
	
	
	//遍歷
	var arr1 = []float64{}
	var t01,t02,t03,t04,t05,t06,t07,t08,t09 float64
	for rows.Next() {
		//RATIO_6H_lstm := make([]opts.BarData, 0)
       //var row := new(arr1)
        rows.Scan(&t01,&t02,&t03,&t04,&t05,&t06,&t07,&t08,&t09)
		//fmt.Println(RATIO_6H_lstm)
        arr1 = append(arr1, t01,t02,t03,t04,t05,t06,t07,t08,t09)
		//fmt.Println(arr1)
    }
	return arr1
}
func Rmse_1(n1 string) [][]float64 {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		"----", "----", "----", "----")

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}

	defer conn.Close()


	//建立連接
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open Connection failed:", err.Error())
	}
	defer db.Close()

	//通過連接對象執行查詢
	rows, err := db.Query(n1)
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()
	
	
	//遍歷
	rmseRow := make([][]float64, 0)
	var t01,t02,t03,t04,t05,t06,t07,t08,t09,t10 float64
	for rows.Next() {
		//RATIO_6H_lstm := make([]opts.BarData, 0)
       //var row := new(arr1)
        rows.Scan(&t01,&t02,&t03,&t04,&t05,&t06,&t07,&t08,&t09,&t10)
		//fmt.Println(RATIO_6H_lstm)
        rmseValue := make([]float64, 0)
		rmseValue = append(rmseValue, t01, t02, t03, t04, t05, t06, t07, t08, t09, t10)
		rmseRow = append(rmseRow, [][]float64{rmseValue}...)
	}

	return rmseRow
}
