package message

import (
	"regexp"
	"sync"
)

func init() {
	var so sync.Once
	so.Do(func() {
		mapTypesRegExp = make(map[string]*regexp.Regexp)

		rxp, err := regexp.Compile(PatternTypeIncomeSberFromSber)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromSber] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffOneStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromTinkoffOneStep] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffTwoStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromTinkoffTwoStep] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffTwoStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromTinkoffTwoStep] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffSBPOneStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromTinkoffSBPOneStep] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffSBPTwoStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromTinkoffSBPTwoStep] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromAlphaOneStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromAlphaOneStep] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeSberFromAlphaTwoStep)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeSberFromAlphaTwoStep] = rxp

		// Other sber
		rxp, err = regexp.Compile(PatternTypeOtherSberPaid)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherSberPaid] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherSberInsufficientFunds)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherSberInsufficientFunds] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherSberCancellation)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherSberCancellation] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherSberPaymentRequest)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherSberPaymentRequest] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherSberReplenishmentATM)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherSberReplenishmentATM] = rxp

		rxp, err = regexp.Compile(PatternTypeIncomeTinkoff)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeIncomeTinkoff] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherTinkoffPaid)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherTinkoffPaid] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherTinkoffPaymentRequest)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherTinkoffPaymentRequest] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherTinkoffPaidSBP)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherTinkoffPaidSBP] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherTinkoffReplenishmentATM)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherTinkoffReplenishmentATM] = rxp

		rxp, err = regexp.Compile(PatternTypeOtherTinkoffInsufficientFunds)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypeOtherTinkoffInsufficientFunds] = rxp

		rxp, err = regexp.Compile(PatternTypePushIncomeTinkoff)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypePushIncomeTinkoff] = rxp

		rxp, err = regexp.Compile(PatternTypePushAllIncomeTinkoff)
		if err != nil {
			panic(err)
		}
		mapTypesRegExp[TypePushAllIncomeTinkoff] = rxp

		steps = make(map[string]string)
		steps[TypeIncomeSberFromTinkoffTwoStep] = TypeIncomeSberFromTinkoffOneStep
		steps[TypeIncomeSberFromTinkoffSBPTwoStep] = TypeIncomeSberFromTinkoffSBPOneStep
		steps[TypeIncomeSberFromAlphaTwoStep] = TypeIncomeSberFromAlphaOneStep

		mapFields = make(map[string][]string)

		mapFields[TypeIncomeSberFromSber] = []string{"body", "card", "time", "amount", "from", "balance"}
		mapFields[TypeIncomeSberFromTinkoffOneStep] = []string{"body", "card", "time", "amount", "balance"}
		mapFields[TypeIncomeSberFromTinkoffTwoStep] = []string{"body", "amount", "from"}
		mapFields[TypeIncomeSberFromTinkoffSBPOneStep] = []string{"body", "card", "time", "from", "amount"}
		mapFields[TypeIncomeSberFromTinkoffSBPTwoStep] = []string{"body", "amount", "balance"}
		mapFields[TypeIncomeSberFromAlphaOneStep] = []string{"body", "card", "time", "amount", "balance"}
		mapFields[TypeIncomeSberFromAlphaTwoStep] = []string{"body", "amount", "from"}

		mapFields[TypeOtherSberPaid] = []string{"body", "card", "time", "amount", "from", "balance"}
		mapFields[TypeOtherSberInsufficientFunds] = []string{"body", "card", "amount", "from"}
		mapFields[TypeOtherSberCancellation] = []string{"body", "card", "time", "amount", "from", "balance"}
		mapFields[TypeOtherSberPaymentRequest] = []string{"body", "amount"}
		mapFields[TypeOtherSberReplenishmentATM] = []string{"body", "card", "time", "amount", "atm", "balance"}

		mapFields[TypeIncomeTinkoff] = []string{"body", "amount", "from", "balance"}

		mapFields[TypeOtherTinkoffPaid] = []string{"body", "card", "amount", "from", "balance"}
		mapFields[TypeOtherTinkoffPaymentRequest] = []string{"body", "card", "amount"}
		mapFields[TypeOtherTinkoffPaidSBP] = []string{"body", "amount", "balance"}
		mapFields[TypeOtherTinkoffReplenishmentATM] = []string{"body", "amount", "balance"}
		mapFields[TypeOtherTinkoffInsufficientFunds] = []string{"body", "from", "card"}
		mapFields[TypePushIncomeTinkoff] = []string{"body", "amount", "from", "balance"}
		mapFields[TypePushAllIncomeTinkoff] = []string{"body", "amount", "balance"}
	})
}
