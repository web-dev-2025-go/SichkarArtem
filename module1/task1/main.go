package main

import (
	"SichkarArtem/module1/task1/models"
)

func main() {
	me := models.NewMe("Artem", 1, "Duikt", "draw")
	// use copy age login
	me.NotChangeAge(100) // return 100 but in struct age dosent chnage
	me.Describe()        // age was 1
	// now chagne age without copy struct
	me.ChangeAge(100) // return
	// user with new age
	me.Describe()
	songs := []string{"first-song", "second-song"}
	friend := models.NewFriend("Friend", 2, "Duit", songs...)

	friend.GetMusic()
	music := []string{"third-song"}
	friend.AddMusic(music...)
	friend.Describe()
}
