package main

import (
	"task1/model"
)

func main() {
	me := model.NewMe("Artem", 1, "Duikt", "draw")
	// use copy age login
	me.NotChangeAge(100) // return 100 but in struct age dosent chnage
	me.Describe()        // age was 1
	// now chagne age without copy struct
	me.ChangeAge(100) // return
	// user with new age
	me.Describe()
	songs := []string{"first-song", "second-song"}
	friend := model.NewFriend("Friend", 2, "Duit", songs...)

	friend.GetMusic()
	music := []string{"third-song"}
	friend.AddMusic(music...)
	friend.Describe()
}
