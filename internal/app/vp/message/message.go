package message

import (
	"regexp"
	"sync"
)

const (

	// Sber income

	TypeIncomeSberFromSber               = "IncomeSberFromSber"
	TypeIncomeSberFromTinkoffOneStep     = "IncomeSberFromTinkoffOneStep"
	TypeIncomeSberFromTinkoffTwoStep     = "IncomeSberFromTinkoffTwoStep"
	TypeIncomeSberFromAlphaOneStep       = "ncomeSberFromAlphaOneStep"
	TypeIncomeIncomeSberFromAlphaTwoStep = "IncomeIncomeSberFromAlphaTwoStep"

	// Sber other

	TypeOtherSberPaid              = "OtherSberPaid"
	TypeOtherSberInsufficientFunds = "OtherSberInsufficientFunds"
	TypeOtherSberCancellation      = "OtherSberCancellation"
	TypeOtherSberPaymentRequest    = "OtherSberPaymentRequest"
	TypeOtherSberReplenishmentATM  = "OtherSberReplenishmentATM"

	// Tinkoff Income

	TypeIncomeTinkoffFromSber    = "IncomeTinkoffFromSber"
	TypeIncomeTinkoffFromTinkoff = "IncomeNinkoffFromTinkoff"
	TypeIncomeTinkoffFromAlpha   = "IncomeNinkoffFromAlpha"

	// Tinkoff Other

	TypeOtherTinkoffPaid              = "OtherTinkoffPaid"
	TypeOtherTinkoffPaymentRequest    = "OtherTinkoffPaymentRequest"
	TypeOtherTinkoffPaidSBP           = "OtherTinkoffPaidSBP"
	TypeOtherTinkoffReplenishmentATM  = "OtherTinkoffReplenishmentATM"
	TypeOtherTinkoffInsufficientFunds = "OtherTinkoffInsufficientFunds"
)

var mapTypesRegExp map[string]*regexp.Regexp

func init() {
	var so sync.Once
	so.Do(func() {

		mapTypesRegExp = make(map[string]*regexp.Regexp)

		rxp, err := regexp.Compile("([A-Z]+)([0-9]+) ([0-9]{2}):([0-9]{2}) Перевод ([0-9.]+р) от ([0-9а-яА-Я ]+.) Баланс: ([0-9.]+р)")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeSberFromSber] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeSberFromTinkoffOneStep] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeSberFromTinkoffTwoStep] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeSberFromAlphaOneStep] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeIncomeSberFromAlphaTwoStep] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherSberPaid] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherSberInsufficientFunds] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherSberCancellation] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherSberPaymentRequest] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherSberReplenishmentATM] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeTinkoffFromSber] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeTinkoffFromTinkoff] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeIncomeTinkoffFromAlpha] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherTinkoffPaid] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherTinkoffPaymentRequest] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherTinkoffPaidSBP] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherTinkoffReplenishmentATM] = rxp
		rxp, err = regexp.Compile("")

		if err != nil {
			panic(err)
		}

		mapTypesRegExp[TypeOtherTinkoffInsufficientFunds] = rxp
	})
}

func getRegExp(typeKind string) *regexp.Regexp {
	if r, ok := mapTypesRegExp[typeKind]; ok {
		return r
	}
	return nil
}

type Message interface {
	GetType() string
	GetFrom() string
	GetCardType() string
	GetCardNumber() string
	GetCard() string
	GetDate() string
	GetBalance() string
	GetAmount() string
}

func ParseMessage(msg string) Message {
	return nil
}

func parseMessageType(msg string) string {
	return ""
}
