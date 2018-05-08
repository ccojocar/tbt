package cmd

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/ccojocar/tbt/tree"
	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a binary tree over to a remote node",
	Long:  "This command sends a binary tree over to a remote node",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(`Please provide the binary tree as a pre-ordered, sequence in
                        (e.g 1,3,#,#,4,#,#). Where # indicates a node termination.`)
		}

		connection, err := net.Dial("tcp", addressReceiver)
		if err != nil {
			fmt.Printf("Failed to connect to the remote node. Error: %v", err)
			os.Exit(-1)
		}
		defer connection.Close()
		nodes := strings.Split(args[0], ",")
		root := tree.NewFromPreOrderedSeq(nodes)

		quit := make(chan int)
		valuesCh := make(chan string)
		defer close(quit)

		go tree.WalkPreOrder(root, valuesCh, quit)

		for {
			value := <-valuesCh
			if value == "" {
				return
			}
			value = fmt.Sprintf("%s,", value)
			connection.Write([]byte(value))
		}
	},
}

// Address is the address of the remote node
var addressReceiver string

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringVarP(&addressReceiver, "address", "a", "", "Address of the remote node where the binary tree will be sent")
}
