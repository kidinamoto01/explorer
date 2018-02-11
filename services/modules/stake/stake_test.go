package stake

import (
	"testing"
	"log"
	"fmt"
	"github.com/tendermint/tendermint/rpc/client"
	"github.com/tendermint/go-wire/data"

)

func TestSearchStakeTransaction(t *testing.T) {

	c :=client.NewHTTP("tcp://gaia-2-node0.testnets.interblock.io:46657", "/websocket")

	log.Printf("get client ")
	// now we query for the tx.
	// since there's only one tx, we know index=0.
	hash:= data.Bytes("630DC4121836E35CCCDF804C71E49E745CE2B035")
	results, err := c.TxSearch(fmt.Sprintf("tx.hash='%v'", hash),false)


	if err != nil{
		fmt.Println("Got error",err)
	}else{
		fmt.Println(results)
	}


}