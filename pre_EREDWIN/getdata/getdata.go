package getdata

import (
	//"fmt"
	"math"
	"pre_EREDWIN/db"
	_ "github.com/denisenkom/go-mssqldb"
	//"github.com/go-echarts/go-echarts/opts"

)
type parameter struct {
	Day1 string
	Day2 string
	
}

/**取X軸 後24筆預測值**/
func Time() []string {
	dbstring := "SELECT A FROM (SELECT TOP 48 convert(varchar, OUT_DM, 100) AS A,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.DbTime(dbstring)
	return arr1
}
/**標準差**/
func RMSE_value() ([]float64,[]float64)  {
	dbstring := "SELECT top 720 RATIO_6H_lstm,RATIO_12H_lstm,RATIO_24H_lstm,RATIO_48H_lstm,Newpeople_lstm,Total_people_lstm,Total_people2_lstm,Total_people3_lstm,Total_edwin_lstm FROM [EREDWIN].[dbo].[LSTM_outpred]"
	arr1 := db.Rmse(dbstring)
	dbstring = "SELECT top 720  RATIO_6H_srmx,RATIO_12H_srmx,RATIO_24H_srmx,RATIO_48H_srmx,Newpeople_srmx,Total_people_srmx,Total_people2_srmx,Total_people3_srmx,Total_edwin_srmx FROM [EREDWIN].[dbo].[SARIMAX_outpred]"
	arr2 := db.Rmse(dbstring)
	//實際值
	dbstring = "SELECT top 720  RATIO_6H,RATIO_12H,RATIO_24H,RATIO_48H,Newpeople,Total_people,Total_people2,Total_people3,Total_edwin FROM [EREDWIN].[dbo].[UnifyEDWIN]"
	arr3 := db.Rmse(dbstring)
	var arr4 = make([]float64, 0)
	var arr5 = make([]float64, 0)
	arr4Value := 0.0 
	arr5Value := 0.0
	for j := 0; j < 9; j++ {
		for i := 24 ; i <= len(arr1) - 1; i++ {
			arr4Value += (arr3[i][j] - arr1[i][j]) * (arr3[i][j] - arr1[i][j])
			arr5Value += (arr3[i][j] - arr2[i][j]) * (arr3[i][j] - arr2[i][j])
		}
		arr4Value = math.Sqrt(arr4Value / float64(len(arr1)))
		arr5Value = math.Sqrt(arr5Value / float64(len(arr1)))
		arr4 = append(arr4, arr4Value)
		arr5 = append(arr5, arr5Value)
		arr4Value = 0.0
		arr5Value = 0.0
	}
	return arr4, arr5
}


/*******************                    LSTM                    *******************/
func RATIO_6H_lstm() []float64 {
	dbstring := "SELECT RATIO_6H_lstm FROM (SELECT TOP 48 RATIO_6H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_12H_lstm() []float64 {
	dbstring := "SELECT RATIO_12H_lstm FROM (SELECT TOP 48 RATIO_12H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_24H_lstm() []float64 {
	dbstring := "SELECT RATIO_24H_lstm FROM (SELECT TOP 48 RATIO_24H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_48H_lstm() []float64 {
	dbstring := "SELECT RATIO_48H_lstm FROM (SELECT TOP 48 RATIO_48H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Newpeople_lstm() []float64 {
	dbstring := "SELECT Newpeople_lstm FROM (SELECT TOP 48 Newpeople_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people_lstm() []float64 {
	dbstring := "SELECT Total_people_lstm FROM (SELECT TOP 48 Total_people_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people1_lstm() []float64 {
	dbstring := "SELECT Total_people1_lstm FROM (SELECT TOP 48 Total_people1_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people2_lstm() []float64 {
	dbstring := "SELECT Total_people2_lstm FROM (SELECT TOP 48 Total_people2_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people3_lstm() []float64 {
	dbstring := "SELECT Total_people3_lstm FROM (SELECT TOP 48 Total_people3_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_edwin_lstm() []float64 {
	dbstring := "SELECT Total_edwin_lstm FROM (SELECT TOP 48 Total_edwin_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
/*
/*******************                    SRMX                    *******************/
func RATIO_6H_srmx() []float64 {
	dbstring := "SELECT RATIO_6H_srmx FROM (SELECT TOP 48 RATIO_6H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_12H_srmx() []float64 {
	dbstring := "SELECT RATIO_12H_srmx FROM (SELECT TOP 48 RATIO_12H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_24H_srmx() []float64 {
	dbstring := "SELECT RATIO_24H_srmx FROM (SELECT TOP 48 RATIO_24H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_48H_srmx() []float64 {
	dbstring := "SELECT RATIO_48H_srmx FROM (SELECT TOP 48 RATIO_48H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Newpeople_srmx() []float64 {
	dbstring := "SELECT Newpeople_srmx FROM (SELECT TOP 48 Newpeople_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people_srmx() []float64 {
	dbstring := "SELECT Total_people_srmx FROM (SELECT TOP 48 Total_people_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people1_srmx() []float64 {
	dbstring := "SELECT Total_people1_srmx FROM (SELECT TOP 48 Total_people1_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people2_srmx() []float64 {
	dbstring := "SELECT Total_people2_srmx FROM (SELECT TOP 48 Total_people2_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people3_srmx() []float64 {
	dbstring := "SELECT Total_people3_srmx FROM (SELECT TOP 48 Total_people3_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_edwin_srmx() []float64 {
	dbstring := "SELECT Total_edwin_srmx FROM (SELECT TOP 48 Total_edwin_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}

/*******************                    實際值                    *******************/

func RATIO_6H() []float64 {
	dbstring := "SELECT RATIO_6H FROM (SELECT TOP 25 RATIO_6H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_12H() []float64 {
	dbstring := "SELECT RATIO_12H FROM (SELECT TOP 25 RATIO_12H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_24H() []float64 {
	dbstring := "SELECT RATIO_24H FROM (SELECT TOP 25 RATIO_24H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_48H() []float64 {
	dbstring := "SELECT RATIO_48H FROM (SELECT TOP 25 RATIO_48H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Newpeople() []float64 {
	dbstring := "SELECT Newpeople FROM (SELECT TOP 25 Newpeople,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people() []float64 {
	dbstring := "SELECT Total_people FROM (SELECT TOP 25 Total_people,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people1() []float64 {
	dbstring := "SELECT Total_people1 FROM (SELECT TOP 25 Total_people1,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people2() []float64 {
	dbstring := "SELECT Total_people2 FROM (SELECT TOP 25 Total_people2,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people3() []float64 {
	dbstring := "SELECT Total_people3 FROM (SELECT TOP 25 Total_people3,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_edwin() []float64 {
	dbstring := "SELECT Total_edwin FROM (SELECT TOP 25 Total_edwin,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
