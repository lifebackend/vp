package message

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

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

	TypeIncomeTinkoff     = "TypeIncomeTinkoff"
	TypePushIncomeTinkoff = "TypePushIncomeTinkoff"
	// Tinkoff Other

	TypeOtherTinkoffPaid              = "OtherTinkoffPaid"
	TypeOtherTinkoffPaymentRequest    = "OtherTinkoffPaymentRequest"
	TypeOtherTinkoffPaidSBP           = "OtherTinkoffPaidSBP"
	TypeOtherTinkoffReplenishmentATM  = "OtherTinkoffReplenishmentATM"
	TypeOtherTinkoffInsufficientFunds = "OtherTinkoffInsufficientFunds"

	TypeUnknown = "TypeUnknown"

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

	PatternTypeIncomeTinkoff     = "Пополнение, счет RUB. ([0-9]+) RUB. ([А-Яа-я .]+)?Доступно ([0-9. ]+) RUB"
	PatternTypePushIncomeTinkoff = "Пополнение на ([0-9]+) ₽, счет RUB. ([А-Яа-я .]+)? Доступно ([0-9. ]+) ₽"

	// Tinkoff other

	PatternTypeOtherTinkoffPaid              = `Покупка, карта ([\*0-9]+). ([0-9]+) RUB. ([А-Яа-яA-Za-z\* .]+)?Доступно ([0-9. ]+) RUB`
	PatternTypeOtherTinkoffPaymentRequest    = "Никому не говорите код 2120! ([A-Za-zА-Яа-я]+). Сумма ([0-9. ]+) RUB"
	PatternTypeOtherTinkoffPaidSBP           = `Оплата СБП, счет RUB. ([0-9]+) RUB. ([А-Яа-яA-Za-z\* .]+)?Доступно ([0-9. ]+) RUB`
	PatternTypeOtherTinkoffReplenishmentATM  = `Пополнение, счет RUB. ([0-9]+) RUB. ([А-Яа-яA-Za-z\* .]+)?Доступно ([0-9. ]+) RUB`
	PatternTypeOtherTinkoffInsufficientFunds = `^Отказ ([A-Za-zА-Яа-я0-9-.]+). Недостаточно средств. Карта ([\*0-9 ]+)+$`
)

var (
	mapTypesRegExp map[string]*regexp.Regexp
	steps          map[string]string
	mapFields      map[string][]string
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

		steps = make(map[string]string)
		steps[TypeIncomeSberFromTinkoffTwoStep] = PatternTypeIncomeSberFromTinkoffOneStep
		steps[TypeIncomeSberFromTinkoffSBPTwoStep] = PatternTypeIncomeSberFromTinkoffSBPOneStep
		steps[TypeIncomeSberFromAlphaTwoStep] = PatternTypeIncomeSberFromAlphaOneStep

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

	})

}

type Service struct {
	collection *mongo.Collection
}

func NewService(client *mongo.Client) *Service {
	collection := client.Database("database").Collection("messages")
	return &Service{collection: collection}
}

func (s *Service) Save(ctx context.Context, deviceID string, from string, typeMsg string, msg string) error {
	t, m := parseMessage(msg)

	if ok, t := isTwoStep(t); ok {
		return s.collection.
			FindOneAndUpdate(
				ctx,
				bson.D{
					{"deviceID", deviceID},
					{"typeMsg", typeMsg},
					{"type", t},
					{"msg", bson.D{{"amount", m["amount"]}}},
				},
				bson.D{{"", ""}}).
			Err()
	}

	b := bson.D{
		{"deviceID", deviceID},
		{"from", from},
		{"typeMsg", typeMsg},
		{"msg", m},
		{"type", t},
		{"datetime", time.Now().Unix()},
	}

	_, err := s.collection.InsertOne(ctx, b)

	return err
}

func isTwoStep(tp string) (bool, string) {
	stp, ok := steps[tp]
	return ok, stp
}

func mapDataToField(m map[string]string, data []string, fields []string) {
	var mx sync.Mutex

	mx.Lock()
	fmt.Println(fields)
	fmt.Println(data)

	for i, f := range fields {
		fmt.Println(f, "->>>>>", data[i], "---->>", i)
		m[f] = data[i]
	}
	mx.Unlock()
}

func getFieldsByType(tp string) []string {
	if r, ok := mapFields[tp]; ok {
		return r
	}
	return nil
}

func parseMessage(msg string) (string, map[string]string) {
	msg = strings.ReplaceAll(msg, "\n", "")
	msg = strings.ReplaceAll(msg, "\u00a0", " ")
	var mx sync.Mutex
	m := make(map[string]string)
	for k, r := range mapTypesRegExp {
		if ok := r.MatchString(msg); ok {
			data := r.FindAllStringSubmatch(msg, -1)
			fields := getFieldsByType(k)

			mapDataToField(m, data[0], fields)

			return k, m
		}
	}
	mx.Lock()
	m["body"] = msg
	mx.Unlock()
	return TypeUnknown, m
}
