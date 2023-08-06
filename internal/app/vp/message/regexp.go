package message

const (

	// Sber regexp

	PatternTypeIncomeSberFromSber              = "^([A-Z]+[0-9]+) ([0-9]{2}:[0-9]{2}) Перевод ([0-9.]+р) от ([0-9а-яА-Я ]+.) Баланс: ([0-9.]+р)$"
	PatternTypeIncomeSberFromTinkoffOneStep    = "^([a-zA-Z-0-9]+) ([0-9]{2}:[0-9]{2}) зачисление ([0-9.]+)р Тинькофф Банк Баланс: ([0-9.]+)р$"
	PatternTypeIncomeSberFromTinkoffTwoStep    = "^[a-zA-Zа-яА-Я0-9 ]+ [0-9]{2}.[0-9]{2}.[0-9]{2} зачислен перевод ([0-9.]+)р из Тинькофф Банк от ([А-Яа-я ]+).$"
	PatternTypeIncomeSberFromTinkoffSBPOneStep = `^([0-9А-Яa-zA-ZА-Я]+) ([0-9]{2}:[0-9]{2}) ([А-Яа-я ]+). перевел\(а\) вам ([0-9]+)р.$`
	PatternTypeIncomeSberFromTinkoffSBPTwoStep = "^[a-zA-Z-0-9]+ [0-9]{2}:[0-9]{2} зачисление ([0-9.]+)р TINKOFF Баланс: ([0-9.]+)р$"
	PatternTypeIncomeSberFromAlphaOneStep      = "^([a-zA-Z-0-9]+) ([0-9]{2}:[0-9]{2}) зачисление ([0-9. ]+)р Альфа Банк Баланс: ([0-9.]+)р$"
	PatternTypeIncomeSberFromAlphaTwoStep      = "^[a-zA-Zа-яА-Я0-9 ]+ [0-9]{2}.[0-9]{2}.[0-9]{2} зачислен перевод ([0-9.]+)р из Альфа Банк от ([А-Яа-я ]+).+$"

	// Sber other

	PatternTypeOtherSberPaid              = "^([0-9А-Яa-zA-ZА-Я-]+) ([0-9]{2}:[0-9]{2}) Покупка ([0-9 .]+)р ([A-Za-z0-9А-Яа-я*. ]+) Баланс: ([0-9. ]+)р+$"
	PatternTypeOtherSberInsufficientFunds = "^([0-9А-Яa-zA-ZА-Я-]+) Мало средств. Покупка ([0-9 .]+)р ([A-Za-z0-9А-Яа-я*. ]+)+$"
	PatternTypeOtherSberCancellation      = "^([0-9А-Яa-zA-ZА-Я-]+) ([0-9]{2}:[0-9]{2}) Отмена покупки ([0-9 .]+)р ([A-Za-z0-9А-Яа-я*. ]+) Баланс: ([0-9.]+)р+$"
	PatternTypeOtherSberPaymentRequest    = "^Никому не сообщайте код: [a-zA-Z0-9]+. После подтверждения произойдет списание ([0-9.]+) RUB ([A-Za-zА-Яа-я0-9. ]+). Комиссия за покупки не взимается.+$"
	PatternTypeOtherSberReplenishmentATM  = "^([a-zA-Z-0-9]+) ([0-9]{2}:[0-9]{2}) зачисление ([0-9. ]+)р ATM [a-zA-Z0-9-]+ Баланс: ([0-9.]+)р+$"

	// Tinkoff regexp

	PatternTypeIncomeTinkoff        = "^Пополнение, счет RUB. ([0-9]+) RUB. ([А-Яа-я .]+)?Доступно ([0-9. ]+) RUB$"
	PatternTypePushIncomeTinkoff    = "^Пополнение на ([0-9]+) ₽, счет RUB. ([А-Яа-я .]+)? Доступно ([0-9. ]+) ₽$"
	PatternTypePushAllIncomeTinkoff = "^Платеж на ([0-9]+) ₽, счет RUBБаланс([0-9. ]+) ₽$"

	// Tinkoff other

	PatternTypeOtherTinkoffPaid              = `Покупка, карта ([\*0-9]+). ([0-9]+) RUB. ([А-Яа-яA-Za-z\* .]+)?Доступно ([0-9. ]+) RUB`
	PatternTypeOtherTinkoffPaymentRequest    = "Никому не говорите код 2120! ([A-Za-zА-Яа-я]+). Сумма ([0-9. ]+) RUB"
	PatternTypeOtherTinkoffPaidSBP           = `Оплата СБП, счет RUB. ([0-9]+) RUB. ([А-Яа-яA-Za-z\* .]+)?Доступно ([0-9. ]+) RUB`
	PatternTypeOtherTinkoffReplenishmentATM  = `Пополнение, счет RUB. ([0-9]+) RUB. ([А-Яа-яA-Za-z\* .]+)?Доступно ([0-9. ]+) RUB`
	PatternTypeOtherTinkoffInsufficientFunds = `^Отказ ([A-Za-zА-Яа-я0-9-.]+). Недостаточно средств. Карта ([\*0-9 ]+)+$`
)
