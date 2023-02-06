package main

type Implementation interface {
	implementOperation() string
}

type ImplementationA struct{}

func (imp *ImplementationA) implementOperation() string {
	return "Result of implementation A"
}

type ImplementationB struct{}

func (imp *ImplementationB) implementOperation() string {
	return "Result of implementation B"
}
