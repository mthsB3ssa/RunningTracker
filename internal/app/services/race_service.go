package services

type RaceService interface {
	CreateRace(runnerId int, distance float64, duration float64, typeOfRunning float64, date time)
}
