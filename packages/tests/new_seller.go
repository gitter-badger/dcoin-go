package main

import (
	"fmt"
	"github.com/democratic-coin/dcoin-go/packages/utils"
	"tests_utils"
)

func main() {

	f:=tests_utils.InitLog()
	defer f.Close()

	txType := "NewSeller";
	txTime := "1427383713";
	userId := []byte("2")
	var blockId int64 = 128008

	var txSlice [][]byte
	// hash
	txSlice = append(txSlice, []byte("22cb812e53e22ee539af4a1d39b4596d"))
	// type
	txSlice = append(txSlice,  utils.Int64ToByte(utils.TypeInt(txType)))
	// time
	txSlice = append(txSlice, []byte(txTime))
	// user_id
	txSlice = append(txSlice, []byte("91573"))
	//arbitration_trust_list \[[0-9]{1,10}(,[0-9]{1,10}){0,100}\]
	txSlice = append(txSlice, []byte("[1,2,3,6,4,91573]"))
	//arbitration_days_refund
	txSlice = append(txSlice, []byte("30"))
	//holdback_pct_array \[\"[0-9]{1,3}(\.[0-9]{2})?\"(,\"[0-9]{1,3}(\.[0-9]{2}\")?){3}\]
	txSlice = append(txSlice, []byte(`["100.00", "90.55", "0.00"]`))
	// sign
	txSlice = append(txSlice, []byte("11111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"))

	blockData := new(utils.BlockData)
	blockData.BlockId = blockId
	blockData.Time = utils.StrToInt64(txTime)
	blockData.UserId = utils.BytesToInt64(userId)

	err := tests_utils.MakeTest(txSlice, blockData, txType, "work_and_rollback");
	if err != nil {
		fmt.Println(err)
	}

}
