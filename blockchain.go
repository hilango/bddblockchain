package main

import (
	//"github.com/joho/godotenv"
	//"github.com/davecgh/go-spew/spew"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"log"
	"encoding/json"
	"io"
	"crypto/sha256"
	"encoding/hex"
)

var Blockchain []Block

func run() error{
	addr:=":40015"
	mux := makMuxRouter()
	log.Println("Listening on ",addr)
	s := &http.Server{
		Addr: addr,
		Handler: mux,
		ReadTimeout:10*time.Second,
		WriteTimeout:10*time.Second,
		MaxHeaderBytes:1<<20,
	}
	if err:=s.ListenAndServe();err!=nil{
		log.Println(err)
		return err
	}
	return nil
}



func makMuxRouter() http.Handler {
	muxRouter:=mux.NewRouter()
	muxRouter.HandleFunc("/get",handleGetBlockchain).Methods("GET")
	muxRouter.HandleFunc("/",handleWriteBlock).Methods("POST")
	return muxRouter
}

type Message struct{
	BPM int
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var m Message

	decoder := json.NewDecoder(r.Body)
	if err:= decoder.Decode(&m); err!=nil{
		respondeWidthJson(w,r,http.StatusBadRequest,r.Body)
		return
	}
	defer r.Body.Close()

	newBlock,err := generateBlock(Blockchain[len(Blockchain)-1],m.BPM)

	if err!=nil{
		respondeWidthJson(w,r,http.StatusCreated,newBlock)

	}

	if isBlockValid(newBlock,Blockchain[len(Blockchain)-1]){
		newBlockchain := append(Blockchain,newBlock)
		replaceChain(newBlockchain)
	}

	respondeWidthJson(w,r,http.StatusCreated,newBlock)
}

func replaceChain(newBlocks []Block) {
		if len(newBlocks) >len(Blockchain){
			Blockchain = newBlocks
		}
}
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index || oldBlock.Hash != newBlock.PreHash ||calculateHash(newBlock)!=newBlock.Hash {
		return false
	}
	return  true
}

func generateBlock(oldBlock Block, BPM int) (Block, error) {
	var newBlock	Block
	t:=time.Now()
	newBlock.Index = oldBlock.Index+1
	newBlock.Timestamp = t.String()
	newBlock.BPM = BPM
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock,nil

}
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PreHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func respondeWidthJson(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response,err := json.MarshalIndent(payload,"","  ")
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Http 500 ： Internal Server Error"))
	}
	w.WriteHeader(code)
	w.Write(response)
}

func handleGetBlockchain(writer http.ResponseWriter, request *http.Request) {
	bytes,err := json.MarshalIndent(Blockchain,""," ")
	if err!=nil{
		http.Error(writer,err.Error(),http.StatusInternalServerError)
		return
	}
	io.WriteString(writer,string(bytes))
}

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PreHash   string
}

func main() {

	genesisBlock:=Block{Index:0,Timestamp:time.Now().String(),BPM:0,Hash:"",PreHash:""}
	Blockchain = append(Blockchain,genesisBlock)

	log.Fatal(run())
}