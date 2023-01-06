package controllers

import (
	"gallery/views"
	"html/template"
	"net/http"
)

type StaticController struct {
}

func (sc StaticController) Home(w http.ResponseWriter, r *http.Request) {
	views.Render(w, nil, "master.html", "home.html")
}

func (sc StaticController) Contact(w http.ResponseWriter, r *http.Request) {
	views.Render(w, nil, "master.html", "contact.html")
}

func (sc StaticController) Faq(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "I changed my log-in email but my notifications are still coming from the old address. Why?",
			Answer:   "Calendar Invitations within your event types will come from the email address that is associated with your connected calendar. You can switch your connected calendar, or update your event type notifications to Email Confirmations so notifications come from <a href='mailto:notifications@calendly.com'>notifications@calendly.com.</a>",
		},
		{
			Question: "Trouble signing in",
			Answer:   "How to regain access to your Stripe account if you can no longer sign in or have forgotten some of your log in information.",
		},
		{
			Question: "Where can I download my invoices from?",
			Answer:   "Our invoices are sent from our payment provider to the email address that we have on file for your account. Contact us and we’ll be happy to locate any invoices that you’re missing. ",
		},
	}
	views.Render(w, data, "master.html", "faq.html")
}
