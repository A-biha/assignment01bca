package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var hashArr = [] string {}
var updatedHashArr = [] string {}
type block struct {
		nonce int
		transaction string
		previousHash string
		currHash string
	}

type TransList struct {
	list[]*block
}
	
func CalculateHash(final string) string{
	new := ""
	hash := sha256.Sum256([]byte(final))
	//fmt.Printf("%x", hash)

	new = fmt.Sprintf("%x", hash)
	hashArr = append(hashArr,new)

	return new
}
	func NewBlock (transaction string, nonce int, previousHash string)*block {
		b := new(block)
		b.transaction = transaction
		b.nonce = nonce
		b.previousHash = previousHash

		str := strconv.Itoa(nonce)
		final := transaction + previousHash + str
		b.currHash=CalculateHash(final)
		return b
	}

	func (ls *TransList) Createblock (transaction string, nonce int)*block{
		length:=len(ls.list)
		
		ph:=""
		if length==0{
			ph =""
		}
		if length>0{
			ph =ls.list[length-1].currHash		
		}

		st := NewBlock (transaction, nonce, ph)
		ls.list = append (ls.list,st)

		return st
	}
	
		//fmt.Println("In String: ", temp)
		func (ls *TransList) ListBlocks(){
			length:=len(ls.list)
			i:=0
			//temp:=" "
			for(i<length){
				fmt.Println("%s List %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
				fmt.Println("")
				fmt.Println("Transaction: ", ls.list[i].transaction)
				fmt.Println("Nonce: ", ls.list[i].nonce)
				fmt.Println("Current Hash: ", ls.list[i].currHash)
				fmt.Println("Previous Hash: ", ls.list[i].previousHash)
				i++
			}
	}
	func (ls *TransList) ChangeBlock(id int, newTrans string) string {
			ls.list[id].transaction=newTrans
			/*fmt.Println("------------------------UPDATED BLOCKS----------------------------")
			ListBlocks()*/
			//b := new(block)
			var nonce int 
			var previousHash string
			
			ls.list[id].transaction = newTrans
			nonce = ls.list[id].nonce
			previousHash = ls.list[id].previousHash 

			str := strconv.Itoa(nonce)
			final := newTrans + previousHash + str
			//b.currHash=CalculateHash(final)

			upnew := ""
			hash := sha256.Sum256([]byte(final))
			//fmt.Printf("%x", hash)

			upnew = fmt.Sprintf("%x", hash)
			ls.list[id].currHash = upnew
			updatedHashArr = append(updatedHashArr,upnew)
			return upnew
	}
	func (ls *TransList) VerifyChain() {
		length:=len(ls.list)
		i := 0
		//change := " "
		
		
		for(i<length-1){
			if ls.list[i].currHash != ls.list[i+1].previousHash {
				fmt.Println("Hash Changed in the blockchain in the block with id ", i)
			}
			
			//fmt.Println("Current Hash1: ", ls.list[i].currHash)
			//fmt.Println("Current Hash2: ", updatedHashArr[i])
			i++
		}
		fmt.Println("No change detected in any other block.")
	}
	func main() {
		block := new(TransList)

		block.Createblock("Alice to Bob", 123)
		block.Createblock("Lily to Marshal", 123)
		block.Createblock("Frankenstein to Sheldon", 123)
		block.Createblock("Sheldon to Leonard", 123)
		block.ListBlocks()

		var id int
		var newTrans string 
		fmt.Print("Enter ID of the block to be changed: ")
		fmt.Scan(&id)

		fmt.Print("Change the transaction to: ")
		fmt.Scan(&newTrans)

		block.ChangeBlock(id,newTrans)
	
		fmt.Println(" ")
		fmt.Println(" ")
		fmt.Println("------------------------UPDATED BLOCKS----------------------------")
		fmt.Println(" ")
		block.ListBlocks()
		fmt.Println(" ")
		fmt.Println("------------------------VERIFYING BLOCKS----------------------------")
		fmt.Println(" ")
		block.VerifyChain()
		
	}
