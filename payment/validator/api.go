package validator

import (
	"payment/data"
	transactionRepo "payment/repository"
	"sync"
)

var initOnce = sync.Once{}
var instanceOnce = sync.Once{}

type ValidatorService interface {
	Validate(transaction data.Transaction) Valid
	Close()
}

type validator interface {
	Validate(transaction data.Transaction, valid Valid)
}

func Init(in chan data.Transaction, out chan<- data.Transaction) {
	initOnce.Do(func() {
		chanIn = in
		chanOut = out
	})
}

func GetInstance() ValidatorService {
	instanceOnce.Do(func() {
		repository := transactionRepo.GetInstance()
		done := make(chan struct{})
		v := &validatorService{repository, chanIn, chanOut, done}
		instance = v

		go v.init()
	})

	return instance
}
