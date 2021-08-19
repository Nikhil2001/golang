package main

import "fmt"

type Album struct {
	name string
	songs []string
}

func mapExamples() {

	
var singerAlbumMap map[string]Album
fmt.Println(singerAlbumMap)

singers :=make(map[string]Album)
fmt.Println(singers)

singers["duaLipa"] = Album{"future Nostalgia",[]string{"levitating","don't start now"}}
singers["Hemachandra"] = Album{"rrr",[]string{"dara dum dum"}}
fmt.Println(singers,"number of singers: ",len(singers))
value,present := singers["Hemachandra"]
fmt.Println(value,present)

value1,present1 := singers["otherSingerName"]
fmt.Printf("%q,%v \n",value1,present1)

delete(singers,"otherSingerName")
fmt.Println(singers)
delete(singers,"Hemachandra")
fmt.Println(singers)
}