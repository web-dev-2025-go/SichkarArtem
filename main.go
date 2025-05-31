package main

import (
	"go-lesson/model"
)

func main() {
	me := model.NewMe("Artem", 1, "Duikt", "draw")
	// use copy age login
	me.Login(10000)
	// return old age
	me.Describe()
	// now chagne age without copy struct
	me.ChangeAge(100)
	// user with new age
	me.Describe()
	songs := []string{"first-song", "second-song"}
	friend := model.NewFriend("Friend", 2, "Duit", songs)

	friend.GetMusic()
	friend.AddMusic([]string{"third-song"})
	friend.Describe()
}
