package main

import "fmt"

func main(){
	mnemonic:="all hour make first leader extend hole alien behind guard gospel lava path output census museum junior mass reopen famous sing advance salt reform"
	seed,err:=NewSeedWithErrorChecking(mnemonic,"TREZOR")
	if err!=nil{
		println("err=",err.Error())
	}
	fmt.Printf("Seed=%+v\n",seed)
}
