package user

import "time"

type User struct {
	Id           int
	Name         string
	DateComplete time.Time
	Status       bool
}

func (this *User) CompleteUser(id int, name string, dateComplete time.Time, status bool) {
	this.Id = id
	this.Name = name
	this.DateComplete = dateComplete
	this.Status = status

}
