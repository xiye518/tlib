package snowflake

import (
	"fmt"
	"testing"
)

func Test_snowflake(t *testing.T) {
	// Create a new Node with a Node number of 1
	node, err := NewNode(1)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Generate a snowflake ID.
	id := node.Generate()

	// Print out the ID in a few different ways.
	t.Logf("Int64  ID: %d\n", id)
	t.Logf("String ID: %s\n", id)
	t.Logf("Base2  ID: %s\n", id.Base2())
	t.Logf("Base64 ID: %s\n", id.Base64())

	// Print out the ID's timestamp
	t.Logf("ID Time  : %d\n", id.Time())

	// Print out the ID's node number
	t.Logf("ID Node  : %d\n", id.Node())

	// Print out the ID's sequence number
	t.Logf("ID Step  : %d\n", id.Step())

	// Generate and print, all in one.
	t.Logf("ID       : %d\n", node.Generate().Int64())
}
