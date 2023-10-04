package http

import (
	"github.com/samarec1812/airline-manager/internal/app/entity"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"

	"github.com/samarec1812/airline-manager/internal/app/service"
)

func CreateAirline(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.airline.CreateAirline"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody createAirlineRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, AirlineErrorResponse(err)) // AirlineErrorResponse(err))

			return
		}

		err := a.CreateAirline(r.Context(), entity.Airline{Code: reqBody.Code, Name: reqBody.Name})
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, AirlineErrorResponse(err)) // UserErrorResponse(err))

			return
		}

		render.JSON(w, r, AirlineSuccessResponse("CODE")) // UserSuccessResponse("OK"))
	}
}

func DeleteAirline(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.airline.DeleteAirline"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody deleteAirlineRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, AirlineErrorResponse(err))

			return
		}

		err := a.RemoveAirline(r.Context(), reqBody.Code)
		if err != nil {
			log.Error("error with delete", err)
			render.JSON(w, r, AirlineErrorResponse(err))

			return
		}

		render.JSON(w, r, AirlineSuccessResponse("CODE")) // UserSuccessResponse("OK"))
	}
}

func CreateAccount(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.account.CreateAccount"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody createAccountRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, AirlineErrorResponse(err)) // AirlineErrorResponse(err))

			return
		}

		err := a.CreateAccount(r.Context(), entity.Account{SchemaID: reqBody.SchemaID})
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, AirlineErrorResponse(err)) // UserErrorResponse(err))

			return
		}

		render.JSON(w, r, AirlineSuccessResponse("CODE")) // UserSuccessResponse("OK"))
	}
}

func DeleteAccount(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.account.DeleteAccount"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody deleteAccountRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, AirlineErrorResponse(err))

			return
		}

		err := a.RemoveAccount(r.Context(), reqBody.AccountID)
		if err != nil {
			log.Error("error with delete", err)
			render.JSON(w, r, AirlineErrorResponse(err))

			return
		}

		render.JSON(w, r, AirlineSuccessResponse("CODE")) // UserSuccessResponse("OK"))
	}
}

func CreateSchema(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.schema.CreateSchema"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody createSchemaRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, AirlineErrorResponse(err)) // AirlineErrorResponse(err))

			return
		}

		err := a.CreateSchema(r.Context(), entity.Schema{Name: reqBody.Name})
		if err != nil {
			log.Error("error with create", err)
			render.JSON(w, r, AirlineErrorResponse(err)) // UserErrorResponse(err))

			return
		}

		render.JSON(w, r, AirlineSuccessResponse("CODE")) // UserSuccessResponse("OK"))
	}
}

func DeleteSchema(log *slog.Logger, a service.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.account.DeleteAccount"

		log = log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var reqBody deleteSchemaRequest
		if err := render.DecodeJSON(r.Body, &reqBody); err != nil {
			log.Error("failed to decode request body", err)
			render.JSON(w, r, AirlineErrorResponse(err))

			return
		}

		err := a.RemoveSchema(r.Context(), reqBody.SchemaID)
		if err != nil {
			log.Error("error with delete", err)
			render.JSON(w, r, AirlineErrorResponse(err))

			return
		}

		render.JSON(w, r, AirlineSuccessResponse("CODE")) // UserSuccessResponse("OK"))
	}
}
