package app

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const (
	defaultPort = ":4040"
)

var (
	validate = validator.New()
)

type controller struct {
	uc usecase
}

func NewController(uc usecase) *controller {
	return &controller{
		uc: uc,
	}
}

func (ct *controller) GetSlots(c *fiber.Ctx) error {
	slots, err := ct.uc.GetSlots(c.Context())
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(slots)
}

func (ct *controller) CreateReservation(c *fiber.Ctx) error {
	req := CreateReservationRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	err = validate.Struct(req)
	if err != nil {
		return err
	}

	rv, err := ct.uc.CreateReservation(c.Context(), req.PatientID, req.DoctorID, req.SlotID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(rv)
}

func (ct *controller) GetReservations(c *fiber.Ctx) error {
	showCancelled := c.QueryBool("show_cancelled", true)

	rvs, err := ct.uc.GetReservations(c.Context(), showCancelled)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(rvs)
}

func (ct *controller) CancelReservation(c *fiber.Ctx) error {
	req := CancelReservationRequest{}

	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	rv, err := ct.uc.CancelReservation(c.Context(), req.ReservationID)
	if err != nil {
		return err
	}

	return c.Status(http.StatusOK).JSON(rv)
}

func (ct *controller) Start(app *fiber.App) {
	// Middleware
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${locals:requestid} ${status} - ${latency} ${method} ${path}​\n",
	}))

	// Route
	app.Get("/reservations/slots", ct.GetSlots)
	app.Post("/reservations", ct.CreateReservation)
	app.Get("/reservations", ct.GetReservations)
	app.Put("/reservations", ct.CancelReservation)

	// Listen
	err := app.Listen(defaultPort)
	if err != nil {
		panic(err)
	}
}

func ControllerErrHandler(c *fiber.Ctx, err error) error {
	code := http.StatusInternalServerError

	switch err {
	case errPatientNotFound,
		errDoctorNotFound,
		errReservationNotFound,
		errSlotNotFound:
		code = http.StatusNotFound
	case errReservationExists,
		errReservationCancelled:
		code = http.StatusBadRequest
	}
	if _, ok := err.(validator.ValidationErrors); ok {
		code = http.StatusBadRequest
	}

	// Send error response
	errResp := errorResponse{
		ErrorMessage: err.Error(),
	}

	b, errMarshal := json.Marshal(errResp)
	if errMarshal != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error: " + errMarshal.Error())
	}

	errSend := c.Status(code).Send(b)
	if errSend != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error: " + errSend.Error())
	}

	return nil
}

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

type CreateReservationRequest struct {
	PatientID int `json:"patient_id" validate:"required,gte=0"`
	DoctorID  int `json:"doctor_id" validate:"required,gte=0"`
	SlotID    int `json:"slot_id" validate:"required,gte=0"`
}

type CancelReservationRequest struct {
	ReservationID int `json:"reservation_id" validate:"required,gte=0"`
}
