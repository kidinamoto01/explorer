package stake

import (
	"fmt"
	"github.com/tendermint/tendermint/rpc/client"
	"log"
	"testing"
	//"github.com/tendermint/go-wire/data"
)

//
//func TestSearchTransaction(t *testing.T) {
//
//		c :=client.NewHTTP("tcp://gaia-2-node0.testnets.interblock.io:46657", "/websocket")
//		log.Printf("get client ")
//		// now we query for the tx.
//		// since there's only one tx, we know index=0.
//		hash:= "630DC4121836E35CCCDF804C71E49E745CE2B035"
//
//
//		results, err := c.TxSearch(fmt.Sprintf("tx.hash='%v'", hash),false)
//
//		if err != nil{
//			fmt.Println("Got error",err)
//		}else{
//			fmt.Println(len(results))
//			ptx := results[0]
//			fmt.Println(ptx.TxResult.Tags)
//		}
//
//
//}

func TestSearchStakeTransaction(t *testing.T) {

	c := client.NewHTTP("tcp://gaia-2-node0.testnets.interblock.io:46657", "/websocket")
	log.Printf("get client ")
	// now we query for the tx.
	// since there's only one tx, we know index=0.
	hash := "DE390ABEB8A7626DDA11C3CDB056D3217B0A2434"

	results, err := c.TxSearch(fmt.Sprintf("tx.hash='%v'", hash), false)
	//results, err := c.TxSearch("tx.hash=630DC4121836E35CCCDF804C71E49E745CE2B035",false)

	//result,_ :=c.Status()
	//fmt.Println(result)

	if err != nil {
		fmt.Println("Got error", err)
	} else {
		fmt.Println(len(results))
		ptx := results[0]
		fmt.Println(ptx.TxResult.Tags)
		fmt.Println(ptx.TxResult)

	}

}
