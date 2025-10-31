package acl

import (
	"encoding/json"
	"fmt"
)

// ACL — Access Control List

func ParseACL(ACLData string) (map[string][]string, error) {

	aclData := make(map[string][]string)
	err := json.Unmarshal([]byte(ACLData), &aclData)
	if err != nil {
		return nil, err
	}

	isThisUserExistsInACL(aclData, "logger1")

	return aclData, nil
}

func isThisUserExistsInACL(acl map[string][]string, primaryKey string) bool {

	if val, ok := acl[primaryKey]; ok {
		fmt.Print("Users found: ", val)
		return true
	} else {
		fmt.Print("User not found")
		return false

	}
}
