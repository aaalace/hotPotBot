package consts

import "time"

// Card Types
const (
	TypeAll = iota
	TypeSingle
	TypeAlbum
)

// Craft Prices
const (
	Free            = 0
	CraftAlbumPrice = 10
)

const LogsFlushInterval = time.Hour * 24
const AppLocation = "Europe/Moscow"
const AdminNick = "@aaalace"
const TechSupportNick = "@aaalace"
const FixedCooldown = 6 * time.Hour
const StartCarouselIndex = 1
const InlineDataDelimiter = "&"