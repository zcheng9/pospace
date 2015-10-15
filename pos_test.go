package pos

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var prover *Prover = nil
var size int = 4
var graphDir string = "graph"

var expHashes [][hashSize]byte = nil
var expMerkle [][hashSize]byte = nil

func TestPoS(t *testing.T) {

}

//Sanity check using simple graph
//[0 0 0 0]
//[1 0 0 0]
//[1 0 0 0]
//[1 0 1 0]
func TestComputeHash(t *testing.T) {
	hashes := make([][]byte, size)
	for i := range hashes {
		f, _ := os.Open(fmt.Sprintf("%s/%d/hash", graphDir, i))
		hashes[i] = make([]byte, hashSize)
		f.Read(hashes[i])
	}

	var result   [hashSize]byte

	for i := range expHashes {
		copy(result[:], hashes[i])
		if expHashes[i] != result {
			log.Fatal("Hash mismatch:", expHashes[i], result)
		}

	}
}

func TestMerkleTree(t *testing.T) {
	result := make([][hashSize]byte, 2*size)
	for i := 1; i < size; i++ {
		f, _ := os.Open(fmt.Sprintf("%s/merkle/%d", graphDir, i))
		buf := make([]byte, hashSize)
		f.Read(buf)
		copy(result[i][:], buf)
	}
	for i := 0; i < size; i++ {
		f, _ := os.Open(fmt.Sprintf("%s/%d/hash", graphDir, i))
		buf := make([]byte, hashSize)
		f.Read(buf)
		copy(result[i+size][:], buf)
	}

	for i := 2*size-1; i > 0; i-- {
		if expMerkle[i] != result[i] {
			log.Fatal("Merkle node mismatch:", i, expMerkle[i], result[i])
		}
	}

}

func TestMain(m *testing.M) {
	prover = setup(size, graphDir)
	prover.InitGraph()
	os.Exit(m.Run())
}