package main

import (
	"errors"
	"fmt"
)


func main() {

	hossam := &human{}
	cat := &animal{}
	steps := []point{
		{1,2,3},
		{2,4,7},
	}

	err := move(hossam,steps)
	if err != nil {
		fmt.Println(err)
	}

	steps = []point{
		{1,2,2},
	}
	err = move(cat,steps)

	if err != nil {
		fmt.Println(err)
	}

}

// accepts type walker which is considered a parent for both animal and human cause the implement its 2 functions
func move(walker walker, point []point) error  {
	for _,point:= range point {
		err := walker.walk(point)
		if err != nil {
			return err
		}
	}
	return nil
}

type object struct {
	position point
}
type human struct {
	object
}

type point struct {
	x float32
	y float32
	z float32
}


type animal struct {
	object
}

type walker interface {
	walk(p point) error
	getPosition() point
}

type talker interface {
	talk(statement string)
}

func (human human) walk(point point) error{
	if point.x < 0 || point.y < 0 || point.z < 0 {
		return errors.New("invalid point")
	}

	human.position = point
	fmt.Println("Human Walked To Position",human.position)
	return nil
}


func (human human) getPosition() point{
	return human.position
}


func (animal animal) walk(point point) error{
	if point.x < 0 || point.y < 0 || point.z < 0 {
		return errors.New("invalid point")
	}

	animal.position = point
	fmt.Println("animal Walked To Position",animal.position)
	return nil
}


func (animal animal) getPosition() point{
	return animal.position
}
