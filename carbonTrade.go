

/*
used for CodeJam 2017

*/

package main

import (
	"errors"
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// type Mortgage struct {
// 	Owner            string `json:"owner"`
// 	AccountID        string `json:"accountID"`
// 	OriginalAmt      string `json:"originalAmt"`
// 	CurrentBal       string `json:"currentBal"`
// 	PaymentAmt       string `json:"paymentAmt"`
// 	Frequency        string `json:"frequency"`
// 	Status           string `json:"status"`
// }

//==============================================================================================================================
//	Carrier - Defines the structure for a Carrier object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON owner -> Struct owner.
//==============================================================================================================================

type Carrier struct {
  Name            string `json:"name"`
  Delivered       string `json:"delivered"`
}
//==============================================================================================================================
//	customer - Defines the structure for a customer object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON owner -> Struct owner.
//==============================================================================================================================

type Customer struct {
  Name            string `json:"name"`
  Delivered       string `json:"delivered"`
  CurrentBal      string `json:"currentBal"`
}
//==============================================================================================================================
//	XOM - Defines the structure for a XOM object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON owner -> Struct owner.
//==============================================================================================================================
type XOM struct {
  Name            string `json:"name"`
  Delivered       string `json:"delivered"`
  CurrentBal      string `json:"currentBal"`
}

type Order  struct{
  SerialNum   string `json:"serialNum"`
  AckCus      string `json:"ackCus"`     //customer acknowledgement
  AckXom      string `json:"ackXom"`    //customer acknowledgement
  AckCarrier  string `json:"ackCarrier"`//customer acknowledgement
  Cost        string `json:"cost"`
}



// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
return nil, nil
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

  switch function {
  case "createOrder":
    return t.createOrder(stub,args)
  }
return nil,nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
  switch function {
  case "readOrder":
    return t.readOrder(stub,args)
  }

  return nil,nil
}

func (t *SimpleChaincode) createOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args) != 5 {
    return nil, errors.New("Incorrect number of arguments. Expecting 5")
  }
    // SN,err:= args[0]
    // AC,err:=args[1]
    // AX,err:=args[2]
    // ACR,err:=args[3]
    // Cost,err:=args[4]
    return nil

}

func (t *SimpleChaincode) readOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args) != 1 {
      return nil, errors.New("Incorrect number of arguments. Expecting AccountID")
    }

    value, err := stub.GetState(args[0]) //args 0 = serialNum
    if err != nil {
      return nil, errors.New("blank error msg")
    }
    return nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
