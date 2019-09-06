package db

// Item struct
type Item struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

//data array
var data []Item

//AddItem function
func AddItem(item *Item) {
	data = append(data, *item)
}

//FindItem func
func FindItem(id int) *Item {
	for _, i := range data {
		if i.ID == id {
			return &i
		}
	}
	return nil
}

//GetAllItems func
func GetAllItems() *[]Item {
	return &data
}

//GetLenght func
func GetLenght() int {
	return len(data)
}
