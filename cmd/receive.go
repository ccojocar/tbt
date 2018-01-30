package cmd

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/cosmincojocar/tbt/tree"
	"github.com/spf13/cobra"
)

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Receive a binary tree from a remote node",
	Long:  "This command starts a server to receive a binary tree from a remote node",
	Run: func(cmd *cobra.Command, args []string) {
		server, err := net.Listen("tcp", address)
		if err != nil {
			fmt.Printf("Failed to start the server. Error: %v\n", err)
			os.Exit(-1)
		}
		defer func() {
			server.Close()
			fmt.Println("Server terminated")
		}()

		fmt.Printf("Server started at %s...\n", address)
		for {
			connection, err := server.Accept()
			if err != nil {
				fmt.Printf("Connection accept error: %v\n", err)
				os.Exit(-1)
			}
			go receiveTree(connection)
		}
	},
}

func receiveTree(conn net.Conn) {
	defer conn.Close()
	timeoutDuration := 5 * time.Second
	bufReader := bufio.NewReader(conn)
	var nodes []string
	for {
		conn.SetReadDeadline(time.Now().Add(timeoutDuration))
		bytes, err := bufReader.ReadBytes(',')
		if err != nil {
			root := reconstructTree(nodes)
			printTree(root)
			return
		}
		nodes = append(nodes, string(bytes[:len(bytes)-1]))
	}
}

func reconstructTree(nodes []string) *tree.Tree {
	return tree.NewFromPreOrderedSeq(nodes)
}

func printTree(root *tree.Tree) {
	quit := make(chan int)
	valueCh := make(chan string)
	defer close(quit)

	go tree.WalkPreOrder(root, valueCh, quit)

	for {
		value := <-valueCh
		if value == "" {
			return
		}
		fmt.Printf("%s,", value)
	}
}

// Address is the address where to listen for incoming binary tree
var address string

func init() {
	rootCmd.AddCommand(receiveCmd)

	receiveCmd.Flags().StringVarP(&address, "address", "a", "", "Address to listen for incoming binary tree")
}
