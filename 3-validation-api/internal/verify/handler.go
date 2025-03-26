package verify

import (
	"encoding/json"
	"go/validation/configs"
	"go/validation/pkg/request"
	"go/validation/pkg/response"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type ValidHandlerDeps struct {
	*configs.Config
}

type ValidHandler struct {
	*configs.Config
}

type Verification struct {
	Verif []ToJson `json:"verif"`
}

var verification Verification

func NewValidHandler(router *http.ServeMux, deps ValidHandlerDeps) {
	handler := &ValidHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *ValidHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		emailRequest, err := request.SendEmail[Mail](&w, r)
		if err != nil {
			return
		}
		hash := request.HashIt()
		e := email.NewEmail()
		e.From = handler.Config.Email.Email
		e.To = []string{emailRequest.Address}
		e.Subject = "Email validation"
		e.HTML = []byte("Verify your email address: http://localhost:8081/verify/{" + hash + "}")
		err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", handler.Config.Email.Email, handler.Config.Password.Password, "smtp.gmail.com"))
		if err != nil {
			response.Json(w, err.Error(), 402)
			return
		}
		toJson := ToJson{
			Address: emailRequest.Address,
			Hash:    hash,
		}
		verification.Verif = append(verification.Verif, toJson)
		jsonData, err := json.Marshal(verification)
		if err != nil {
			panic(err)
		}
		toJson.Write(jsonData)
		response.Json(w, "The verification link was sent on your email", 200)
	}
}

func (handler *ValidHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		hash = strings.Trim(hash, "{}")
		for _, v := range verification.Verif {
			if v.Hash == hash {
				response.Json(w, "Email verified", 200)
				verification.Delete(hash)
				return
			}
		}
		response.Json(w, "Invalid link", 400)
	}
}

func (verif *Verification) Delete(hash string) bool {
	var emails []ToJson
	isDeleted := false
	for _, s := range verif.Verif {
		isMatched := strings.Contains(s.Hash, hash)
		if !isMatched {
			emails = append(emails, s)
			continue
		}
		isDeleted = true
	}
	verif.Verif = emails
	file, err := json.Marshal(verif.Verif)
	if err != nil {
		return isDeleted
	}
	toJson := ToJson{}
	toJson.Write(file)
	return isDeleted
}
