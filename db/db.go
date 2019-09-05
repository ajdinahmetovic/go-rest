package db

// Item struct
type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//DATA array
var DATA []Item

//AddItem function
func AddItem(item Item) {
	DATA = append(DATA, item)
}

//FindItem func
func FindItem(item Item) interface{} {
	for _, i := range DATA {
		if i == item {
			return item
		}
	}
	return nil
}
