package structs

import (
	"time"
)

// StatusUserDB ...
type StatusUserDB int

// StatusLogin ...
type StatusLogin int

const (
	StatusUserDBNone     StatusUserDB = iota // StatusUserDBNone
	StatusUserDBOK                           // StatusUserDBOK - юзер найден
	StatusUserDBNotFound                     // StatusUserDBNotFound - юзер не найден
	StatusUserDBNotValid                     // StatusUserDBNotValid - на входе невалидные данные
)

const (
	StatusLoginNone         StatusLogin = iota // StatusLoginNone
	StatusLoginOK                              // StatusLoginOK - юзер залогинен
	StatusLoginUserNotFound                    // StatusLoginUserNotFound - юзер не найден
	StatusLoginNotValid                        // StatusLoginNotValid - на входе невалидные данные
)

func (s StatusUserDB) GetDescription() string {
	switch s {
	case StatusUserDBOK:
		return "Успешное получение анкеты пользователя"
	case StatusUserDBNotFound:
		return "Анкета не найдена"
	case StatusUserDBNotValid:
		return "Невалидные данные"
	default:
		return "Непредвиденный статус"
	}
}

func (l StatusLogin) GetDescription() string {
	switch l {
	case StatusLoginOK:
		return "Успешная аутентификация"
	case StatusLoginUserNotFound:
		return "Пользователь не найден"
	case StatusLoginNotValid:
		return "Невалидные данные"
	default:
		return "Непредвиденный статус"
	}
}

type User struct {
	FirstName  string
	SecondName string
	DateBirth  time.Time
	Biography  string
	City       string
	Pass       string
}

type UserLogin struct {
	ID   int64
	Hash string
	Pass string
}
