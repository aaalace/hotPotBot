package messages

import (
	"fmt"
	"hotPotBot/internal/consts"
	"hotPotBot/internal/db/models"
)

var ForPurchaseWrite = fmt.Sprintf("Для покупки пиши - %s", consts.AdminNick)

var SupportContactText = fmt.Sprintf("При возникновении проблем в работе - %s", consts.TechSupportNick)

// Page titles

const TutorialTitle = "Тут будет туториал"

const StartPageTitle = "Привет!\nЕсли ты тут впервые - ознакомься с туториалом по кнопке в нижней части бота"

const CardsStoragePageTitle = "💼 Здесь ты можешь ознакомиться со своей коллекцией карт"

const HotPotStudioPageTitle = "🎸 Ты в Hot Pot Studio, тут можно найти много интересного"

const MyAccountPageTitle = "👤 Твой аккаунт:"

const OtherAccountPageTitle = "🔍 Введи ник другого пользователя - @username"

var ShopTitle = fmt.Sprintf("🏦 Добро пожаловать в Hot Pot Shop\nЗдесь ты можешь приобрести все желанные карты\n%s",
	ForPurchaseWrite)

const CraftTitle = "🔮 Добро пожаловать в Hot Pot Сraft\nЗдесь ты можешь обменять свои дубликаты на более крутые карты"

const CraftAgreementTitle = "Подтверди крафт"

const WriteUsernameToExchangeTitle = "Введи ник пользователя с которым хочешь обменяться - @username"

// Messages

const SuccessfulRandomCardDropTitle = "Поздравляем, тебе выпала карта -"

const SuccessfulCraftTitle = "Поздравляем, ты скрафтил -"

func SuccessfulExchangeInit(username string) string {
	return fmt.Sprintf(
		"Вы успешно запросили обмен\nСкоро придет уведомление и вы узнаете, на какую карту %s готов обменяться",
		username)
}

func WantToContinueExchange(username, cardName string) string {
	return fmt.Sprintf(
		"Вам пришло предложение обмена от %s\nОн готов отдать карту - %s\nВыберите в разделе Мои Карты ту, которую готовы отдать ему и введите его ник, или откажитесь от обмена",
		username,
		cardName)
}

func ExchangeAgreement(partnerUsername string, partnerCard, myCard *models.Card) string {
	return fmt.Sprintf(
		"У тебя с @%s есть открытый обмен!\n\n📈 Ты получишь - %s [Fame %v]\n📉 Он получит - %s [Fame %v]",
		partnerUsername,
		partnerCard.Name,
		partnerCard.Weight,
		myCard.Name,
		myCard.Weight)
}

func WaitPartnerToAcceptExchange(partnerUsername string) string {
	return fmt.Sprintf(
		"Ждем @%s... Вы получите уведомление, как только он примет свое решение",
		partnerUsername)
}

func SuccessExchange(partnerUsername string, partnerCard *models.Card) string {
	return fmt.Sprintf(
		"Поздравляем, обмен с @%s состоялся!\nВы получили карту:",
		partnerUsername,
		partnerCard.Name,
		partnerCard.Weight)
}

const DeclinedExchange = "Обмен отменен, можете создать новый в разделе своих карт"
