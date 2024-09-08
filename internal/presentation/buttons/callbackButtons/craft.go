package buttons

// craft menu

var DoCraftInlineButton = CallbackButton{
	Title: "Подтверждаю ✅",
	Data:  "do_craft_button_data",
}

var CancelCraftInlineButton = CallbackButton{
	Title: "Отмена ❌️",
	Data:  "", // not used
}

var CraftAlbumInlineButton = CallbackButton{
	Title: "Single ➡️ Album",
	Data:  "craft_album_button_data",
}
