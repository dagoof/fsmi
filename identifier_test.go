package fsmi

import (
	"fmt"
)

type StringIdentifier string

const (
	NilSI uint64 = iota
	PendingSI
	ActiveSI
	FinishedSI
	Pending  StringIdentifier = "nothing"
	Active   StringIdentifier = "pending"
	Finished StringIdentifier = "fininshed"
)

func (s StringIdentifier) Identity() uint64 {
	switch s {
	case Pending:
		return PendingSI
	case Active:
		return ActiveSI
	case Finished:
		return FinishedSI
	}
	return NilSI
}

func ExampleIdentifier() {
	printIdentity := func(identifier Identifier) {
		fmt.Println(identifier.Identity())
	}

	printIdentity(Pending)
	printIdentity(Active)
	printIdentity(Finished)
	printIdentity(StringIdentifier("invalid"))
	// Output:
	// 1
	// 2
	// 3
	// 0
}
