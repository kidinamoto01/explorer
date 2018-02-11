package handlers

import (
  "net/http"

  "github.com/gorilla/mux"

  "github.com/tendermint/tmlibs/common"
  "github.com/cosmos/cosmos-sdk/client/commands"
  "github.com/spf13/cast"
  "fmt"
)

// queryStatus is to query block chain status
func queryStatus(w http.ResponseWriter, r *http.Request) {
  c := commands.GetNode()
  status, err := c.Status()
  if err != nil {
    common.WriteError(w, err)
    return
  }
  if err := printResult(w, status); err != nil {
    common.WriteError(w, err)
  }
}


//query the block height
// queryBlock is to query a block by height
func queryBlockHeight(w http.ResponseWriter, r *http.Request) {

  c := commands.GetNode()
  status, err := c.Status()

  if err != nil {
    common.WriteError(w, err)
    return
  }
  if err := printResult(w, status.LatestBlockHeight); err != nil {
    common.WriteError(w, err)
  }
}

//query the block height
// queryBlock is to query a block by height
func queryTotalBond(w http.ResponseWriter, r *http.Request) {

  //c,_ := commands.GetNode().Validators(1)
  c := commands.GetNode()

  status, err := c.Status()

  height := status.LatestBlockHeight

  height64 := cast.ToInt64(height)

  block, err := c.Validators(&height64)
  if err != nil {
    common.WriteError(w, err)
    return
  }

  //loop validators
  sum := int64(0)
  for i:= 0; i < len(block.Validators);i++{
    fmt.Println(block.Validators[i])
    sum += block.Validators[i].VotingPower
    fmt.Println(sum)
  }

  if err := printResult(w, sum); err != nil {
    common.WriteError(w, err)
  }

}


// mux.Router registrars

func RegisterQueryStatus(r *mux.Router) error {
  r.HandleFunc("/status", queryStatus).Methods("GET")
  return nil
}

//register height
func RegisterQueryBlockHeight(r *mux.Router) error {
  r.HandleFunc("/status/height", queryBlockHeight).Methods("GET")
  return nil
}

//register bond
func RegisterQueryTotalBond(r *mux.Router) error {
  r.HandleFunc("/status/bond", queryTotalBond).Methods("GET")
  return nil
}


func RegisterStatus(r *mux.Router) error {
  funcs := []func(*mux.Router) error{
    RegisterQueryStatus,
    RegisterQueryBlockHeight,
    RegisterQueryTotalBond,
  }

  for _, fn := range funcs {
    if err := fn(r); err != nil {
      return err
    }
  }
  return nil
}

// End of mux.Router registrars
