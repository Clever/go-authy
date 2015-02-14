package authygo

import(
    "testing"
)

func Test_RequestSms(t *testing.T) {
    api := NewAuthyApi("bf12974d70818a08199d17d5e2bae630")
    api.ApiUrl = "http://sandbox-api.authy.com"

    userResponse, err := api.RegisterUser(UserOpts{
		Email: "foo@example.com",
		PhoneNumber: "432-123-1111",
		CountryCode: 1,
	})
    verification, err := api.RequestSms(userResponse.User.Id, true)

    if err != nil {
        t.Error("External error found", err)
    }

    if verification.Valid != true {
        t.Error("Verification should be valid.")
    }
}