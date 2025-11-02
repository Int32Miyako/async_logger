package acl

import (
	"encoding/json"
	"fmt"
)

// ACL — Access Control List
/*
	ACLData string = `{
	"logger1":          ["/main.Admin/Logging"],
	"logger2":          ["/main.Admin/Logging"],
	"stat1":            ["/main.Admin/Statistics"],
	"stat2":            ["/main.Admin/Statistics"],
	"biz_user":         ["/main.Biz/Check", "/main.Biz/Add"],
	"biz_admin":        ["/main.Biz/*"],
	"after_disconnect": ["/main.Biz/Add"]
}`
*/
// допустим я хочу получить мапу из user -> Admin/Logging

func ParseACL(ACLData string) (map[string][]string, error) {
	aclData := make(map[string][]string)
	err := json.Unmarshal([]byte(ACLData), &aclData)
	if err != nil {
		return nil, err
	}

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

func GetMethodsForUser(acl map[string][]string, user string) []string {
	if isThisUserExistsInACL(acl, user) {
		return acl[user]
	}
	return []string{}
}

func IsUserAllowedForMethod(acl map[string][]string, user string, method string) bool {
	methods := GetMethodsForUser(acl, user)
	for _, m := range methods {
		if m == method {
			return true
		}
	}

	return false
}
