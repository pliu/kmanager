package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Kafka Manager API")
	})

	ldapURL := "ldap://localhost:3890"
	l, err := ldap.DialURL(ldapURL)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	err = l.Bind("cn=admin,ou=people,dc=example,dc=com", "'password'")
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest("dc=example,dc=com", ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false, "(&(objectClass=person))", []string{"cn", "uid", "memberOf"}, nil)
	searchResp, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range searchResp.Entries {
		fmt.Printf("User: %s (UID: %s)\n", entry.GetAttributeValue("cn"), entry.GetAttributeValue("uid"))
		groups := entry.GetAttributeValues("memberOf")
		if len(groups) > 0 {
			fmt.Println("  Groups:")
			for _, groupDN := range groups {
				// Extract group name from the DN
				groupName := extractCNFromDN(groupDN)
				fmt.Printf("    - %s\n", groupName)
			}
		} else {
			fmt.Println("  No group memberships found.")
		}
		fmt.Println()
	}

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// Helper function to extract CN from a Distinguished Name (DN)
func extractCNFromDN(dn string) string {
	parsedDN, err := ldap.ParseDN(dn)
	if err != nil {
		return dn // Return original DN if parsing fails
	}
	for _, rdn := range parsedDN.RDNs {
		for _, attr := range rdn.Attributes {
			if attr.Type == "cn" {
				return attr.Value
			}
		}
	}
	return dn // Return original DN if CN not found
}
