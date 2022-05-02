package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	nonce    int
}

type Blockchain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.nonce = nonce

	return block
}

func GenesisBlock() *Block {
	return CreateBlock("Genesis", []byte{})
}

func (chain *Blockchain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	newBlock := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, newBlock)
}

func InitBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}
