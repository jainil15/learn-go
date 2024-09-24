package models

import (
	"log"
	"slices"
	"testing"
)

func TestModels_Validation(t *testing.T) {
	t.Run("Testing CreatePropertyPayload", func(t *testing.T) {
		payload := CreatePropertyPayload{
			Name:        "Jainil",
			Email:       "jainil@gmail.com",
			PhoneNumber: "081234567890",
			Address:     "Jakarta",
			About:       nil,
		}
		err := payload.Validate()
		if err != nil {
			log.Println(err)
			t.Error("CreatePropertyPayload failed")
		}
	})
	t.Run("Testing CreatePropertyPayload Fail", func(t *testing.T) {
		about := ""
		payload := CreatePropertyPayload{
			Name:        "",
			Email:       "@gmail.com",
			PhoneNumber: "567890",
			Address:     "Jakarta",
			About:       &about,
		}
		err := payload.Validate()
		if err == nil {
			log.Println(err)
			t.Error("CreatePropertyPayload failed")
		}
		if !slices.Equal(err.Get("phone_number"), []string{"Phone number is not valid"}) {
			t.Error("CreatePropertyPayload failed", err.Get("phone_number"))
		}
		if len(err.Get("email")) != 1 {
			t.Error("CreatePropertyPayload failed", err.Get("email"))
		}
		if len(err.Get("name")) != 1 {
			t.Error("CreatePropertyPayload failed", err.Get("name"))
		}
		if len(err.Get("about")) != 1 {
			t.Error("CreatePropertyPayload failed", err.Get("about"))
		}
		log.Printf("Validation error for failing: %v\n", err)
	})
}
