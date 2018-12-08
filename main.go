package main

func main() {
	bc := NewBlockchain()
	bc.AddBlock("A Send B 1 BTC")
	bc.AddBlock("B Send C 1 BTC")

	/*
	for _, block := range bc.blocks {
		fmt.Printf("Version: %d\n", block.Version)
		fmt.Printf("PrevBlochHash: %x\n", block.PrevBlockHash)
		fmt.Printf("CurrBlockHash: %x\n", block.Hash)
		fmt.Printf("TimeStamp: %d\n", block.TimeStamp)
		fmt.Printf("Bits: %d\n", block.Bits)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Isvalid: %v\n", NewProofWork(block).IsValid())
		fmt.Println()
	}
	*/
}
