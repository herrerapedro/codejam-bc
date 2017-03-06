

/*
used for CodeJam 2017

*/

package main

import (
	"errors"
	"fmt"
	"strconv"
  "encoding/json"
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
  CurrentBal      string `json:"currentBal"`
}
//==============================================================================================================================
//	XOM - Defines the structure for a XOM object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON owner -> Struct owner.
//==============================================================================================================================
type XOM struct {
  Name            string `json:"name"`
  CurrentBal      string `json:"currentBal"`
}


//==============================================================================================================================
//	Order - Defines the structure for a Order object. JSON on right tells it what JSON fields to map to
//			  that element when reading a JSON object into the struct e.g. JSON owner -> Struct owner.
//==============================================================================================================================
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
  case "createCustomer":
    return t.createCustomer(stub,args)
  case "createXOM":
    return t.createXOM(stub,args)
  case "tradeFunds":
    return t.tradeFunds(stub,args)
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

//==============================================================================================================================
//  Transfer funds
//==============================================================================================================================
func (t *SimpleChaincode) tradeFunds(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
A:=args[0]
B:=args[1]
C:=args[2]

amount:=strconv.Atoi(C)

//get the balances
var provider XOM
valueA, err := stub.GetState(A)
err=json.Unmarshal(valueA, &provider)
CBX,err:=strconv.Atoi(provider.CurrentBal)

var customer Customer
valueB, err := stub.GetState(B)
err=json.Unmarshal(valueB, &customer)
CBC,err:=strconv.Atoi(customer.CurrentBal)

//now the trade
CBC=CBC-amount
CBX=CBX+amount

a1,_:=json.Marshal(provider)
err = stub.PutState(A,[]byte(a1))

b1,_:=json.Marshal(customer)
err = stub.PutState(B,[]byte(b1))


}
//==============================================================================================================================
//  XOM Creation
//==============================================================================================================================
func (t *SimpleChaincode) createXOM(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 3")
  }

  X:=args[0]//serial number
  valueC, err := stub.GetState(X)//get state of previous Order
  var carbon XOM//create order object
  err = json.Unmarshal(valueC, &carbon)

  X2 := &XOM{
    Name: args[0],
    currentBal:args[1],
  }

  x1,_:=json.Marshal(X2)
  serialnum:= args[0]

  err = stub.PutState(serialnum, []byte(x1)) // passes serialNum as the key value for searching blockchain
  if err != nil {
    return nil, err

  }
  return nil,nil
}
//==============================================================================================================================
//  Customer Creation
//==============================================================================================================================
func (t *SimpleChaincode) createCustomer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args) != 2 {
    return nil, errors.New("Incorrect number of arguments. Expecting 3")
  }

  C:=args[0]//serial number
  valueC, err := stub.GetState(C)//get state of previous Order
  var customer Customer//create order object
  err = json.Unmarshal(valueC, &customer)

  C2 := &Customer{
    Name: args[0],
    currentBal:args[1],
  }

  c1,_:=json.Marshal(C2)
  serialnum:= args[0]

  err = stub.PutState(serialnum, []byte(c1)) // passes serialNum as the key value for searching blockchain
  if err != nil {
    return nil, err

  }
  return nil,nil
}

//==============================================================================================================================
//  Read Customer
//==============================================================================================================================
func (t *SimpleChaincode) readCustomer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

  if len(args) != 1 {
      return nil, errors.New("Incorrect number of arguments. Expecting Serial Number")
    }

    value, err := stub.GetState(args[0]) //args 0 = serialNum
    if err != nil {
      return nil, errors.New("blank error msg")
    }
    return value,nil

}

//==============================================================================================================================
//  Order Creation
//==============================================================================================================================
func (t *SimpleChaincode) createOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args) != 5 {
    return nil, errors.New("Incorrect number of arguments. Expecting 5")
  }

  O:=args[0]//serial number
  valueO, err := stub.GetState(O)//get state of previous Order
  var order Order//create order object
  err = json.Unmarshal(valueO, &order)

  O2 := &Order{
    SerialNum: args[0],
    AckCus: args[1],
    AckXom: args[2],
    AckCarrier: args[3],
    Cost: args[4],
  }

  o1,_:=json.Marshal(O2)
  serialnum:= args[0]

  err = stub.PutState(serialnum, []byte(o1)) // passes serialNum as the key value for searching blockchain
  if err != nil {
    return nil, err

  }
  return nil,nil
}

//==============================================================================================================================
//  Read Order
//==============================================================================================================================
func (t *SimpleChaincode) readOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
  if len(args) != 1 {
      return nil, errors.New("Incorrect number of arguments. Expecting Serial Number")
    }

    value, err := stub.GetState(args[0]) //args 0 = serialNum
    if err != nil {
      return nil, errors.New("blank error msg")
    }
    return value,nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
