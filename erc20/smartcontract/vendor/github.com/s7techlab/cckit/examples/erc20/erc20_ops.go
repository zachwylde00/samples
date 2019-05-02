package erc20

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"github.com/pkg/errors"
	"github.com/s7techlab/cckit/identity"
	r "github.com/s7techlab/cckit/router"
)

const (
	BalancePrefix   = `BALANCE`
	AllowancePrefix = `APPROVE`
)

var (
	ErrNotEnoughFunds                   = errors.New(`not enough funds`)
	ErrForbiddenToTransferToSameAccount = errors.New(`forbidden to transfer to same account`)
	ErrSpenderNotHaveAllowance          = errors.New(`spender not have allowance for amount`)
)

type (
	Transfer struct {
		From   identity.PublicKeyID
		To     identity.PublicKeyID
		Amount int
	}

	Approve struct {
		From    identity.PublicKeyID
		Spender identity.PublicKeyID
		Amount  int
	}
)

func querySymbol(c r.Context) (interface{}, error) {
	return c.State().Get(SymbolKey)
}

func queryName(c r.Context) (interface{}, error) {
	return c.State().Get(NameKey)
}

func queryTotalSupply(c r.Context) (interface{}, error) {
	return c.State().Get(TotalSupplyKey)
}

func queryBalanceOf(c r.Context) (interface{}, error) {
	return getBalance(c, c.ArgString(`publicKey`))
}

func invokeTransfer(c r.Context) (interface{}, error) {
	// transfer target
	// toMspId := c.ParamString(`toMspId`)
	// toCertId := c.ParamString(`toCertId`)
	toPublicKey := c.ParamString(`toPublicKey`)

	//transfer amount
	amount := c.ParamInt(`amount`)

	// get information about tx creator
	//invoker, err := identity.FromStub(c.Stub())

	pKey, keyErr := cid.GetID(c.Stub())

	if keyErr != nil {
		return nil, keyErr
	}

	fmt.Println("Invoker public key: ", pKey)

	// Disallow to transfer token to same account
	if pKey == toPublicKey {
		return nil, ErrForbiddenToTransferToSameAccount
	}

	// get information about invoker balance from state
	invokerBalance, err := getBalance(c, pKey)
	if err != nil {
		return nil, err
	}

	// Check the funds sufficiency
	if invokerBalance-amount < 0 {
		return nil, ErrNotEnoughFunds
	}

	// Get information about recipient balance from state
	recipientBalance, err := getBalance(c, toPublicKey)
	if err != nil {
		return nil, err
	}

	// Update payer and recipient balance
	if err = setBalance(c, pKey, invokerBalance-amount); err != nil {
		return nil, err
	}

	if err = setBalance(c, toPublicKey, recipientBalance+amount); err != nil {
		return nil, err
	}

	// Trigger event with name "transfer" and payload - serialized to json Transfer structure
	if err = c.SetEvent(`transfer`, &Transfer{
		To: identity.PublicKeyID{
			PublicKey: toPublicKey},
		Amount: amount,
	}); err != nil {
		return nil, err
	}

	// return current invoker balance
	return invokerBalance - amount, nil
}

func queryAllowance(c r.Context) (interface{}, error) {
	return getAllowance(c, c.ParamString(`ownerPublicKey`), c.ParamString(`spenderPublicKey`))
}

func invokeApprove(c r.Context) (interface{}, error) {
	// spenderMspId := c.ParamString(`spenderMspId`)
	// spenderCertId := c.ParamString(`spenderCertId`)
	spenderPublicKey := c.ParamString(`spenderPublicKey`)
	amount := c.ParamInt(`amount`)

	pKey, keyErr := cid.GetID(c.Stub())

	if keyErr != nil {
		return nil, keyErr
	}

	if err := setAllowance(c, pKey, spenderPublicKey, amount); err != nil {
		return nil, err
	}

	if err := c.SetEvent(`approve`, &Approve{
		From: identity.PublicKeyID{
			PublicKey: pKey},
		Spender: identity.PublicKeyID{
			PublicKey: spenderPublicKey},
		Amount: amount,
	}); err != nil {
		return nil, err
	}

	return true, nil
}

func invokeTransferFrom(c r.Context) (interface{}, error) {

	fromPublicKey := c.ParamString(`fromPublicKey`)
	toPublicKey := c.ParamString(`toPublicKey`)
	amount := c.ParamInt(`amount`)

	pKey, keyErr := cid.GetID(c.Stub())

	if keyErr != nil {
		return nil, keyErr
	}

	// check method invoker has allowances
	allowance, err := getAllowance(c, fromPublicKey, pKey)

	if err != nil {
		return nil, err
	}

	if allowance < amount {
		return nil, ErrSpenderNotHaveAllowance
	}

	balance, err := getBalance(c, fromPublicKey)
	if err != nil {
		return nil, err
	}

	if balance-amount < 0 {
		return nil, ErrNotEnoughFunds
	}

	recipientBalance, err := getBalance(c, toPublicKey)
	if err != nil {
		return nil, err
	}

	if err = setBalance(c, fromPublicKey, balance-amount); err != nil {
		return nil, err
	}
	if err = setBalance(c, toPublicKey, recipientBalance+amount); err != nil {
		return nil, err
	}
	if err = setAllowance(c, fromPublicKey, pKey, allowance-amount); err != nil {
		return nil, err
	}

	if err = c.SetEvent(`transfer`, &Transfer{
		From: identity.PublicKeyID{
			PublicKey: fromPublicKey},
		To: identity.PublicKeyID{
			PublicKey: toPublicKey},
		Amount: amount,
	}); err != nil {
		return nil, err
	}

	// return current invoker balance
	return balance - amount, nil
}

// === internal functions, not "public" chaincode functions

// setBalance puts balance value to state
func balanceKey(ownerPublicKey string) []string {
	return []string{BalancePrefix, ownerPublicKey}
}

func allowanceKey(ownerPublicKey, spenderPublicKey string) []string {
	return []string{AllowancePrefix, ownerPublicKey, spenderPublicKey}
}

func getBalance(c r.Context, publicKey string) (int, error) {
	return c.State().GetInt(balanceKey(publicKey), 0)
}

// setBalance puts balance value to state
func setBalance(c r.Context, publicKey string, balance int) error {
	return c.State().Put(balanceKey(publicKey), balance)
}

func getAllowance(c r.Context, ownerPublicKey, spenderPublicKey string) (int, error) {
	return c.State().GetInt(allowanceKey(ownerPublicKey, spenderPublicKey), 0)
}

func setAllowance(c r.Context, ownerPublicKey, spenderPublicKey string, amount int) error {
	return c.State().Put(allowanceKey(ownerPublicKey, spenderPublicKey), amount)
}
