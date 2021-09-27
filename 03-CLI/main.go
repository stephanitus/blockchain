package main

func main() {
	//Read in blockchain
	bc := NewBlockchain()
	defer bc.db.Close()

	//Handle CLI
	cli := CLI{bc}
	cli.Run()
}
