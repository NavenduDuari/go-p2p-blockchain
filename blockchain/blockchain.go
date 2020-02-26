package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

func main() {

}

type Block struct {
	Index     int
	Data      int
	Timestamp string
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func CalculateHash(block Block) string {
	record := strconv.Itoa(block.Index) + strconv.Itoa(block.Data) + block.Timestamp + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

func IsBlockValid(currentBlock, prevBlock Block) bool {
	if prevBlock.Index+1 != currentBlock.Index {
		return false
	}

	if prevBlock.Hash != currentBlock.PrevHash {
		return false
	}

	if CalculateHash(currentBlock) != currentBlock.Hash {
		return false
	}

	return true
}

func GenerateBlock(prevBlock Block, data int) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Data = data
	newBlock.Timestamp = t.String()
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}
