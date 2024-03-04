package todo

import (
	"encoding/json"
	"fmt"
)

type Item struct {
    Text string
}

func saveItems(filename string, items []Item) error {
    b, err := json.Marshal(items)
    
    if err != nil {
        return err
    }
    fmt.Println(string(b))
    
    return nil
}
