package main

type Item struct {
	title string
	body  string
}

var database []Item

func GetByName(title string) Item {
	var getItem Item
	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}

	return getItem
}

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

func EditItem(title string, edit Item) Item {
	var changed Item
	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}
	return changed
}

func DelItem(item Item) {
	var del Item
	for idx, val := range database {
		if val.title == edit.title {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	return del
}

func main() {

}
