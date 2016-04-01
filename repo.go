package main

import "fmt"

var currentId int

var images Images

func init() {
	RepoCreateImage(Image{Id: 1, Location: "Pikacu Ustanin Tas Firini", Path: "C://Windovz31"})
	RepoCreateImage(Image{Id: 2, Location: "Ultron Ikinci El Elektronik", Path: "/Userz/Mac/L1brari3s/DarwinWasHere"})
}

func RepoFindImage(id int) Image {
	for _, t := range images {
		if t.Id == id {
			return t
		}
	}
	// yoksa geriye bos dondur
	return Image{}
}

func RepoCreateImage(t Image) Image {
	currentId += 1
	t.Id = currentId
	images = append(images, t)
	return t
}

func RepoDestroyImage(id int) error {
	for i, t := range images {
		if t.Id == id {
			images = append(images[:i], images[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Image with id of %d to delete", id)
}
