package main

import (
	"embed"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/* TODO: To handle multiple form types
var formTypes = map[string]FormType{
	"tagged": {
		DBTable:      "tagged_offer_contact_forms",
		TeamTemplate: "tagged_offer_contact_form.html",
		UserTemplate: "offer_contact_form_sender_reply.html",
		TeamSubject:  "{{.Tag}} | Offer Contact Form | Resourceinn | Ad",
		UserSubject:  "Resourceinn (HRM) | Thank You!",
	},
	// Add other form types here as needed
}
*/

//go:embed templates/*
var templatesFS embed.FS

func Main(args map[string]interface{}) map[string]interface{} {

	response := make(map[string]interface{})

	form, err := extractFormFromArgs(args)
	if err != nil {
		response["error"] = err.Error()
		return response
	}

	if err := insertFormIntoDB(form); err != nil {
		response["error"] = "Failed to insert into database"
		return response
	}

	if err := sendTeamEmail(form); err != nil {
		response["error"] = "Failed to send email to the team"
		return response
	}

	if err := sendUserEmail(form); err != nil {
		response["error"] = "Failed to send reply email to the user"
		return response
	}

	response["message"] = "Thank you for contacting"
	return response
}

func sendTeamEmail(form TaggedOfferContactForm) error {
	teamSubject := fmt.Sprintf("%s | Offer Contact Form | Resourceinn | Ad", form.Tag)
	cc := []string{"geekinntech@gmail.com", "gogglebakers@gmail.com", "noman@resourceinn.com", "ahsan@resourceinn.com", "marketing@resourceinn.com"}
	to := "sales@resourceinn.com"
	teamTemplatePath := getTemplatePath("tagged_offer_contact_form.html")
	return sendEmail(teamTemplatePath, teamSubject, to, cc, form)
}

func sendUserEmail(form TaggedOfferContactForm) error {
	userSubject := "Resourceinn (HRM) | Thank You!"
	userTemplatePath := getTemplatePath("offer_contact_form_sender_reply.html")
	return sendEmail(userTemplatePath, userSubject, form.Email, nil, form)
}

func getTemplatePath(filename string) string {
	return "templates/" + filename
}
