package core

import (
	"encoding/json"
	"fmt"
	"strings"
)

const (
	SET    = "SET"
	GET    = "GET"
	DELETE = "DELETE"
)

/*
 Parses a query string into a token array

 This function takes a query string as input (e.g. "SET myKey myValue EXPIRE 3600")
 and returns an array containing the tokens.

 The array currently contains 3 elements:

   - tokens[0] - The operation (e.g. "SET", "GET", "DELETE")

   - tokens[1] - The key being operated on (e.g. "myKey")

   - tokens[2] - The value (if setting) or empty (if getting/deleting)

 A future enhancement may include adding a 4th element for the expiration
   - tokens[3] - The expiration time in seconds (e.g. 3600)

 Input:
   - queryString - The query as a string (e.g. "SET foo bar")

 Output:
   - []string - Array of tokens
*/

var kv = NewKeyValueStore()

func QueryProcesser(query string) ([]string, error) {
	// Split the query string into tokens
	tokens := strings.Fields(query)
	if len(tokens) < 2 {
		return nil, fmt.Errorf("invalid operation")
	}

	fmt.Print(len(tokens))
	switch tokens[0] {
	case SET:
		if len(tokens) < 3 {
			return nil, fmt.Errorf("invalid operation")
		}
		_, _ = kv.Set(tokens[1], tokens[2])
		return []string{tokens[1], string(tokens[2])}, nil
	case GET:
		value, ok := kv.Get(tokens[1])
		if !ok {
			return nil, fmt.Errorf("GET: key not found")
		}
		switch value.(type) {
		case string:
			return []string{value.(string)}, nil
		case int, float32, float64:
			return []string{fmt.Sprintf("%d", value.(int))}, nil
		case map[string]interface{}:
			jsonData, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			return []string{string(jsonData)}, nil
		case []interface{}:
			jsonData, err := json.Marshal(value)
			if err != nil {
				return nil, err
			}
			return []string{string(jsonData)}, nil
		}
	case DELETE:
		key, ok := kv.Delete(tokens[1])
		if !ok {
			return nil, fmt.Errorf("DELETE: key not found")
		}
		return []string{key}, nil
	}
	return nil, fmt.Errorf("invalid operation")
}
