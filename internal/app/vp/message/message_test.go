package message

import (
	"context"
	"regexp"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestSaveMessageWithDB(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Skip()
	}

	db := client.Database("database_test")
	logger := logrus.NewEntry(logrus.StandardLogger())

	s := NewService(db, logger)
	err = s.Save(ctx, "test", "test2", "sms", "MIR-8105 10:06 зачисление 100р Тинькофф Банк Баланс: 558.02р")

	if err != nil {
		t.Error(err)
	}

	err = s.Save(ctx, "test", "test2", "sms", "МИР Сберкарта8105 26.07.23 зачислен перевод 100р из Тинькофф Банк от МАКСИМ ВИКТОРОВИЧ П.")

	if err != nil {
		t.Error(err)
	}
}

func Test(t *testing.T) {
	rxp, err := regexp.Compile(PatternTypeIncomeSberFromSber)
	if err != nil {
		panic(err)
	}

	if ok := rxp.Match([]byte(`VISA3200 22:08 Перевод 4580р от Екатерина И. Баланс: 20018.96р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffOneStep)

	if ok := rxp.Match([]byte(`MIR-8105 10:06 зачисление 100р Тинькофф Банк Баланс: 558.02р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffTwoStep)

	if ok := rxp.Match([]byte(`МИР Сберкарта8105 26.07.23 зачислен перевод 100р из Тинькофф Банк от МАКСИМ ВИКТОРОВИЧ П.`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffSBPOneStep)

	if ok := rxp.Match([]byte(`МИР8105 10:59 Любовь П. перевел(а) вам 104р.`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromTinkoffSBPTwoStep)

	if ok := rxp.Match([]byte(`MIR-8105 10:58 зачисление 104р TINKOFF Баланс: 1068.02р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromAlphaOneStep)

	if ok := rxp.Match([]byte(`MIR-8105 10:58 зачисление 104р Альфа Банк Баланс: 1068.02р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromAlphaOneStep)

	if ok := rxp.Match([]byte(`MIR-8105 10:58 зачисление 104р Альфа Банк Баланс: 1068.02р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeIncomeSberFromAlphaTwoStep)

	if ok := rxp.Match([]byte(`МИР Сберкарта8105 26.07.23 зачислен перевод 100р из Альфа Банк от Максим Викторович П. Сообщение: Перевод денежных средств.`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherSberPaid)

	if ok := rxp.Match([]byte(`MIR-8105 16:09 Покупка 164р YANDEX*4215*DOSTAVKA Баланс: 458.02р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherSberInsufficientFunds)

	if ok := rxp.Match([]byte(`MIR-8105 Мало средств. Покупка 1145р YANDEX*4121*TAXI`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherSberCancellation)

	if ok := rxp.Match([]byte(`MIR-8105 08:45 Отмена покупки 300р WHOOSH BIKE Баланс: 1102.02р`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherSberPaymentRequest)

	if ok := rxp.Match([]byte(`Никому не сообщайте код: 319405. После подтверждения произойдет списание 500.00 RUB VOXIMPLANT.COM. Комиссия за покупки не взимается. За пополнение карт предусмотрена комиссия sberbank.ru/sms/df/`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherSberReplenishmentATM)

	if ok := rxp.Match([]byte(`VISA3200 17:04 зачисление 26000р ATM 60031224 Баланс: 77275.63р`)); !ok {
		t.Error(ok)
	}

	// Tinkoff

	rxp, err = regexp.Compile(PatternTypeIncomeTinkoff)

	if ok := rxp.Match([]byte(`Пополнение, счет RUB. 100 RUB. Максим П. Доступно 82005.91 RUB`)); !ok {
		t.Error(ok)
	}

	if ok := rxp.Match([]byte(`Пополнение, счет RUB. 100 RUB. Доступно 82105.91 RUB`)); !ok {
		t.Error(ok)
	}
	rxp, err = regexp.Compile(PatternTypePushIncomeTinkoff)
	msg := `Пополнение на 100 ₽, счет RUB. Максим П.  \nДоступно 87 859.89 ₽`
	msg = strings.ReplaceAll(msg, `\n`, "")
	msg = strings.ReplaceAll(msg, "\u00a0", " ")
	if ok := rxp.Match([]byte(msg)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypePushAllIncomeTinkoff)

	msg = `Платеж на 500 ₽, счет RUB\nБаланс 87 409.89 ₽`
	msg = strings.ReplaceAll(msg, `\n`, "")
	msg = strings.ReplaceAll(msg, "\u00a0", " ")
	t.Log(msg)
	if ok := rxp.Match([]byte(msg)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypePushAllIncomeTinkoff)
	msg = `МИР8105 10:59 Любовь П. перевел(а) вам 104р.`

	// Tinkoff other

	rxp, err = regexp.Compile(PatternTypeOtherTinkoffPaid)

	if ok := rxp.Match([]byte(`Покупка, карта *8367. 500 RUB. YM*freelance. Доступно 81505.91 RUB`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherTinkoffPaymentRequest)

	if ok := rxp.Match([]byte(`Никому не говорите код 2120! freelance. Сумма 500.00 RUB`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherTinkoffPaidSBP)

	if ok := rxp.Match([]byte(`Оплата СБП, счет RUB. 2597 RUB. GAZPROMNEF. Доступно 172005.91 RUB`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherTinkoffReplenishmentATM)

	if ok := rxp.Match([]byte(`Пополнение, счет RUB. 210000 RUB. Банкомат. Доступно 233759.71 RUB`)); !ok {
		t.Error(ok)
	}

	rxp, err = regexp.Compile(PatternTypeOtherTinkoffInsufficientFunds)

	if ok := rxp.Match([]byte(`Отказ YandexGo. Недостаточно средств. Карта *8367`)); !ok {
		t.Error(ok)
	}
}

func TestFields(t *testing.T) {
	rxp, _ := regexp.Compile(PatternTypeIncomeSberFromTinkoffSBPOneStep)

	r := rxp.FindAllStringSubmatch(`МИР8105 10:59 Любовь П. перевел(а) вам 104р.`, -1)

	t.Log(r[0])
	if len(r) != 1 {
		t.Error("match should be 1")
	}

	if len(r[0]) != 5 {
		t.Error("match should be 3", len(r[0]))
	}
}
