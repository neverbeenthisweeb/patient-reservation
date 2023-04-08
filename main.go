package main

import (
	"patientreservation/app"

	"github.com/gofiber/fiber/v2"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup repo
	repo := app.NewRepo()
	repo.SetSlot(app.NewRepoSlotImpl())
	repo.SetPatient(app.NewRepoPatientImpl())
	repo.SetDoctor(app.NewRepoDoctorImpl())
	repo.SetReservation(app.NewRepoReservationImpl())

	// Setup infra
	infra := app.NewInfra()
	infra.SetQueueCounter(app.NewQueueCounterImpl())

	// Setup use case
	usecase := app.NewUsecase(repo, infra)

	// Setup controller
	fiberApp := fiber.New(
		fiber.Config{
			ErrorHandler: app.ControllerErrHandler,
		},
	)
	controller := app.NewController(usecase)
	controller.Start(fiberApp)
}
