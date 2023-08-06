package message

const (

	// Sber income

	TypeIncomeSberFromSber              = "IncomeSberFromSber"
	TypeIncomeSberFromTinkoffOneStep    = "IncomeSberFromTinkoffOneStep"
	TypeIncomeSberFromTinkoffTwoStep    = "IncomeSberFromTinkoffTwoStep"
	TypeIncomeSberFromTinkoffSBPOneStep = "IncomeSberFromTinkoffSBPOneStep"
	TypeIncomeSberFromTinkoffSBPTwoStep = "IncomeSberFromTinkoffSBPTwoStep"
	TypeIncomeSberFromAlphaOneStep      = "ncomeSberFromAlphaOneStep"
	TypeIncomeSberFromAlphaTwoStep      = "IncomeIncomeSberFromAlphaTwoStep"

	// Sber other

	TypeOtherSberPaid              = "OtherSberPaid"
	TypeOtherSberInsufficientFunds = "OtherSberInsufficientFunds"
	TypeOtherSberCancellation      = "OtherSberCancellation"
	TypeOtherSberPaymentRequest    = "OtherSberPaymentRequest"
	TypeOtherSberReplenishmentATM  = "OtherSberReplenishmentATM"

	// Tinkoff Income

	TypeIncomeTinkoff        = "TypeIncomeTinkoff"
	TypePushIncomeTinkoff    = "TypePushIncomeTinkoff"
	TypePushAllIncomeTinkoff = "TypePushIncomeTinkoff"
	// Tinkoff Other

	TypeOtherTinkoffPaid              = "OtherTinkoffPaid"
	TypeOtherTinkoffPaymentRequest    = "OtherTinkoffPaymentRequest"
	TypeOtherTinkoffPaidSBP           = "OtherTinkoffPaidSBP"
	TypeOtherTinkoffReplenishmentATM  = "OtherTinkoffReplenishmentATM"
	TypeOtherTinkoffInsufficientFunds = "OtherTinkoffInsufficientFunds"

	TypeUnknown = "TypeUnknown"
)
