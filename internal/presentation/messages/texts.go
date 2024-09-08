package messages

import (
	"fmt"
	"hotPotBot/consts"
)

var ForPurchaseWrite = fmt.Sprintf("Для покупки пиши - %s", consts.AdminNick)

var SupportContactText = fmt.Sprintf("При возникновении проблем в работе - %s", consts.TechSupportNick)

// Page titles

const SuccessfulRandomCardDropTitle = "Поздравляем, тебе выпала карта -"

const TutorialTitle = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, in culpa qui officia deserunt mollit anim id est laborum."

const StartPageTitle = "Привет!\nЕсли ты тут впервые - ознакомься с туториалом по кнопке в нижней части бота"

const CardsStoragePageTitle = "💼 Здесь ты можешь ознакомиться со своей коллекцией карт"

const HotPotStudioPageTitle = "🎸 Ты в Hot Pot Studio, тут можно найти много интересного"

const MyAccountPageTitle = "👤 Твой аккаунт:"

const OtherAccountPageTitle = "🔍 Введи ник другого пользователя - @username"

var ShopTitle = fmt.Sprintf("🏦 Добро пожаловать в Hot Pot Shop\nЗдесь ты можешь приобрести все желанные карты\n%s",
	ForPurchaseWrite)

const CraftTitle = "🔮 Добро пожаловать в Hot Pot Сraft\nЗдесь ты можешь обменять свои дубликаты на более крутые карты"

const CraftAgreementTitle = "Подтверди крафт"

const SuccessfulCraftMessage = "Поздравляем, ты скрафтил -"
