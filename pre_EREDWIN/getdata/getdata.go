package getdata

import (
	"fmt"
	"strconv"
	//"time"
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
	dbstring := "SELECT A FROM (SELECT TOP 49 convert(varchar, OUT_DM, 100) AS A,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.DbTime(dbstring)
	//一天只顯示一筆天數
	for i := 0; i <= len(arr1) -1 ; i++ {
		if i != 24 && i != 0 && i != 48{
			nTime := arr1[i]
			str := nTime[10:]
			//寫回arr1取代
			if str != "  1:55AM" {
				arr1[i] = str
			} 
			//fmt.Println(str)
		}
		
		
	}
	return arr1
}
/**標準差**/
func RMSE_value() ([]float64,[]float64)  {
	//時間
	//nTine := time.Now
	//預測值
	//dbstring := "SELECT top 192 RATIO_6H_lstm ,RATIO_12H_lstm,RATIO_24H_lstm,RATIO_48H_lstm,Newpeople_lstm,Total_people_lstm,Total_people1_lstm,Total_people2_lstm,Total_people3_lstm,Total_edwin_lstm FROM [EREDWIN].[dbo].[LSTM_outpred] order by num desc"
	dbstring := "SELECT top 720 isnull(RATIO_6H_lstm,999) ,isnull(RATIO_12H_lstm,999),isnull(RATIO_24H_lstm,999),isnull(RATIO_48H_lstm,999),isnull(Newpeople_lstm,999),isnull(Total_people_lstm,999),isnull(Total_people1_lstm,999),isnull(Total_people2_lstm,999),isnull(Total_people3_lstm,999),isnull(Total_edwin_lstm,999) FROM [EREDWIN].[dbo].[LSTM_outpred] order by num desc"
	arr1 := db.Rmse_1(dbstring)
	dbstring = "SELECT top 720 RATIO_6H_srmx ,RATIO_12H_srmx,RATIO_24H_srmx,RATIO_48H_srmx,Newpeople_srmx,Total_people_srmx,Total_people1_srmx,Total_people2_srmx,Total_people3_srmx,Total_edwin_srmx FROM [EREDWIN].[dbo].[SARIMAX_outpred] order by num desc"
	arr2 := db.Rmse_1(dbstring)
	//實際值
	dbstring = "SELECT top 744 RATIO_6H *100,RATIO_12H *100,RATIO_24H *100,RATIO_48H *100,Newpeople,Total_people,Total_people1,Total_people2,Total_people3,Total_edwin FROM [EREDWIN].[dbo].[UnifyEDWIN] order by num desc"
	arr3 := db.Rmse_1(dbstring)
	//格式化slice
	var arr4 = make([]float64, 0)
	var arr5 = make([]float64, 0)
	//計算值從零開始
	lsmta := 0.0 
	srmxb := 0.0
	//實際值從[0][0]開始 預測值從[24][0]開始
	c := 0
	for j := 0; j < 10; j++ {
		for i := 24 ; i <= len(arr1)- 1; i++ {
			//計算值去掉預測的24筆
			//過濾NULL值 將NULL填999 排除計算
			//目前只過濾LSTM
			if arr1[i][j] == 999 {
				c++
				continue
			} else {
				lsmta += (arr3[c][j] - arr1[i][j]) * (arr3[c][j] - arr1[i][j])
				srmxb += (arr3[c][j] - arr2[i][j]) * (arr3[c][j] - arr2[i][j])
				c++
			}
		}
		c = 0
		lsmta = math.Sqrt(lsmta/float64(len(arr1)))
		srmxb = math.Sqrt(srmxb/float64(len(arr1)))
		//取小數點後兩位
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", lsmta), 64)
		arr4 = append(arr4, value)
		value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", srmxb), 64)
		arr5 = append(arr5, value)
		//歸零值 準備計算下一筆
		lsmta = 0.0
		srmxb = 0.0
	}
	//fmt.Println(arr4)
	return arr4, arr5
}


/*******************                    LSTM                    *******************/
func RATIO_6H_lstm() []float64 {
	dbstring := "SELECT RATIO_6H_lstm FROM (SELECT TOP 49 round(RATIO_6H_lstm,2)RATIO_6H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_12H_lstm() []float64 {
	dbstring := "SELECT RATIO_12H_lstm FROM (SELECT TOP 49 round(RATIO_12H_lstm,2)RATIO_12H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_24H_lstm() []float64 {
	dbstring := "SELECT RATIO_24H_lstm FROM (SELECT TOP 49 round(RATIO_24H_lstm,2)RATIO_24H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_48H_lstm() []float64 {
	dbstring := "SELECT RATIO_48H_lstm FROM (SELECT TOP 49 round(RATIO_48H_lstm,2)RATIO_48H_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Newpeople_lstm() []float64 {
	dbstring := "SELECT Newpeople_lstm FROM (SELECT TOP 49 round(Newpeople_lstm,2)Newpeople_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people_lstm() []float64 {
	dbstring := "SELECT Total_people_lstm FROM (SELECT TOP 49 round(Total_people_lstm,2)Total_people_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people1_lstm() []float64 {
	dbstring := "SELECT Total_people1_lstm FROM (SELECT TOP 49 round(Total_people1_lstm,2)Total_people1_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people2_lstm() []float64 {
	dbstring := "SELECT Total_people2_lstm FROM (SELECT TOP 49 round(Total_people2_lstm,2)Total_people2_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people3_lstm() []float64 {
	dbstring := "SELECT Total_people3_lstm FROM (SELECT TOP 49 round(Total_people3_lstm,2)Total_people3_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_edwin_lstm() []float64 {
	dbstring := "SELECT Total_edwin_lstm FROM (SELECT TOP 49 round(Total_edwin_lstm,2)Total_edwin_lstm,NUM FROM LSTM_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
/*
/*******************                    SRMX                    *******************/
func RATIO_6H_srmx() []float64 {
	dbstring := "SELECT RATIO_6H_srmx FROM (SELECT TOP 49 round(RATIO_6H_srmx,2)RATIO_6H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_12H_srmx() []float64 {
	dbstring := "SELECT RATIO_12H_srmx FROM (SELECT TOP 49 round(RATIO_12H_srmx,2)RATIO_12H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_24H_srmx() []float64 {
	dbstring := "SELECT RATIO_24H_srmx FROM (SELECT TOP 49 round(RATIO_24H_srmx,2)RATIO_24H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_48H_srmx() []float64 {
	dbstring := "SELECT RATIO_48H_srmx FROM (SELECT TOP 49 round(RATIO_48H_srmx,2)RATIO_48H_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Newpeople_srmx() []float64 {
	dbstring := "SELECT Newpeople_srmx FROM (SELECT TOP 49 round(Newpeople_srmx,2)Newpeople_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people_srmx() []float64 {
	dbstring := "SELECT Total_people_srmx FROM (SELECT TOP 49 round(Total_people_srmx,2)Total_people_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people1_srmx() []float64 {
	dbstring := "SELECT Total_people1_srmx FROM (SELECT TOP 49 round(Total_people1_srmx,2)Total_people1_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people2_srmx() []float64 {
	dbstring := "SELECT Total_people2_srmx FROM (SELECT TOP 49 round(Total_people2_srmx,2)Total_people2_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_people3_srmx() []float64 {
	dbstring := "SELECT Total_people3_srmx FROM (SELECT TOP 49 round(Total_people3_srmx,2)Total_people3_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func Total_edwin_srmx() []float64 {
	dbstring := "SELECT Total_edwin_srmx FROM (SELECT TOP 49 round(Total_edwin_srmx,2)Total_edwin_srmx,NUM FROM SARIMAX_outpred ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}

/*******************                    實際值                    *******************/

func RATIO_6H() []float64 {
	dbstring := "SELECT RATIO_6H*100 as RATIO_6H FROM (SELECT TOP 25 RATIO_6H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_12H() []float64 {
	dbstring := "SELECT RATIO_12H*100 as RATIO_12H FROM (SELECT TOP 25 RATIO_12H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_24H() []float64 {
	dbstring := "SELECT RATIO_24H*100 as RATIO_24H FROM (SELECT TOP 25 RATIO_24H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
	arr1 := db.Db(dbstring)
	return arr1
}
func RATIO_48H() []float64 {
	dbstring := "SELECT RATIO_48H*100 as RATIO_48H FROM (SELECT TOP 25 RATIO_48H,NUM FROM UnifyEDWIN ORDER BY NUM DESC) AS T ORDER BY NUM"
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
