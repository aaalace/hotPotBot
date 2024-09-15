package messages

import (
	"fmt"
	"hotPotBot/internal/consts"
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

// Messages

const SuccessfulRandomCardDropTitle = "Поздравляем, тебе выпала карта -"

const SuccessfulCraftTitle = "Поздравляем, ты скрафтил -"
