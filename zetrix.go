package main

import (

	"fmt"

	"github.com/armmarov/zetrix-sdk-go-fork/src/model"
	"github.com/armmarov/zetrix-sdk-go-fork/src/sdk"
)


func main() {
	var amount int64 = 100
	// The account to receive
	var destAddress string = "FILL_IN_DEST_ADDRESS_HERE"
	var url string = "https://node.zetrix.com"
	// The account that Gas
	var sourceAddress string = "FILL_IN_SOURCE_ADDRESS_HERE"
	// The fixed write 1000L, the unit is UGas
	var gasPrice int64 = 10
	// Set up the maximum cost 0.01Gas
	var feeLimit int64 = 50000
	//Building SDK objects
	var testSdk sdk.Sdk
	var reqDataInit model.SDKInitRequest
	reqDataInit.SetUrl(url)
	resDataInit := testSdk.Init(reqDataInit)
	if resDataInit.ErrorCode != 0 {
		fmt.Printf(resDataInit.ErrorDesc)
	}
	//Gets the latest Nonce
	var reqDataNonce model.AccountGetNonceRequest
	reqDataNonce.SetAddress(sourceAddress)
	resDataNonce := testSdk.Account.GetNonce(reqDataNonce)
	if resDataNonce.ErrorCode != 0 {
		fmt.Printf(resDataNonce.ErrorDesc)
	}
	//Building Operation
	var reqDataOperation model.GasSendOperation
	reqDataOperation.Init()

	reqDataOperation.SetAmount(amount)
	reqDataOperation.SetDestAddress(destAddress)
	//Building Blob
	var reqDataBlob model.TransactionBuildBlobRequest
	reqDataBlob.SetSourceAddress(sourceAddress)
	reqDataBlob.SetFeeLimit(feeLimit)
	reqDataBlob.SetGasPrice(gasPrice)
	var nonce int64 = resDataNonce.Result.Nonce + 1
	reqDataBlob.SetNonce(nonce)
	reqDataBlob.SetOperation(reqDataOperation)
	resDataBlob := testSdk.Transaction.BuildBlob(reqDataBlob)
	if resDataBlob.ErrorCode != 0 {
		fmt.Printf(resDataBlob.ErrorDesc)
	} else {
		//Sign
		PrivateKey := []string{"FILL_IN_PRIVATE_KEY_HERE"}
		var reqData model.TransactionSignRequest
		reqData.SetBlob(resDataBlob.Result.Blob)
		reqData.SetPrivateKeys(PrivateKey)
		resDataSign := testSdk.Transaction.Sign(reqData)
		if resDataSign.ErrorCode != 0 {
			fmt.Printf(resDataSign.ErrorDesc)
		} else {
			//Submit
			var reqData model.TransactionSubmitRequest
			reqData.SetBlob(resDataBlob.Result.Blob)
			reqData.SetSignatures(resDataSign.Result.Signatures)
			resDataSubmit := testSdk.Transaction.Submit(reqData)
			if resDataSubmit.ErrorCode != 0 {
				fmt.Printf(resDataSubmit.ErrorDesc)
			} else {
				fmt.Printf("Test_Transaction_BuildBlob_Sign_Submit succeed, Hash:", resDataSubmit.Result.Hash)
			}
		}
	}
}