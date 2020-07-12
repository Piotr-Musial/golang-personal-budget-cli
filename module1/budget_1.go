package module1

type Item struct {
	Description string
	Price float32
}

// Budget stores budget information
type Budget struct {
	Max float32
	Items []Item
}

// Item stores item information
