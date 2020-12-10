package main

import (
	"fmt"
	"log"

	"gopkg.in/ldap.v3"
)

func main() {

	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", "nakagusuku.ie.u-ryukyu.ac.jp", 389))
	if err != nil {
		log.Fatal(err)
	}

	searchRequest := ldap.NewSearchRequest(
		"cn=iesudoers,ou=Group,ou=ie,o=u-ryukyu,c=jp", // base info
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=organaizationalPerson)(uid=%s))", "hoge"),
		[]string{"dn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range sr.Entries {
		fmt.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("uid"))
	}
}
