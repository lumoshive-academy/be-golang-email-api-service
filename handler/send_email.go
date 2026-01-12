package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lumoshive-academy/be-golang-email-api-service/dto"
	"github.com/lumoshive-academy/be-golang-email-api-service/utils"
	"gopkg.in/gomail.v2"
)

type HandlerEmail struct {
	Config utils.Configuration
}

func NewHandlerEmail(config utils.Configuration) *HandlerEmail {
	return &HandlerEmail{Config: config}
}

func (handlerEmail *HandlerEmail) SendEmail(w http.ResponseWriter, r *http.Request) {
	var req dto.EmailRequest

	// Bind JSON body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	// Input validation logic (Manual check to match Node.js explicitly)
	if req.To == "" || req.Subject == "" || req.Text == "" || req.Name == "" {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Field to, subject, and text are required", nil)
		return
	}

	if !utils.ValidateEmail(req.To) {
		utils.ResponseBadRequest(w, http.StatusBadRequest, "Invalid email address", nil)
		return
	}

	// Email HTML Content construction
	htmlContent := fmt.Sprintf(`
			<html>
			  <head>
				<style>
				  .email-container {
					font-family: Arial, sans-serif;
					background-color: #15294e;
					padding: 20px;
					border-radius: 10px;
					max-width: 600px;
					margin: auto;
					font-size: 16px;
					line-height: 1.5;
				  }
				  .email-body { padding: 20px; background-color: white; }
				  .email-footer { text-align: center; padding-top: 20px; color: #fff; font-size: 12px; }
				  .email-title { font-size: 1.1rem; font-weight: bold; margin-bottom: 10px; }
				  .email-sub-title { font-size: 1rem; font-weight: bold; margin-bottom: 8px; }
				  .email-text { font-size: 0.95rem; margin: 0 auto; }
				  .email-footer p { margin: 5px 0; }
				</style>
			  </head>
			  <body>
				<div class="email-container">
				  <div class="email-body">
					<h1 class="email-title">%s</h1>
					<h2 class="email-sub-title">%s</h2>
					<p class="email-text">%s</p>
					<hr>
					<p class="email-text"><strong>Email Receiver:</strong> %s</p>
				  </div>
				  <div class="email-footer">
					<p><span style="color: red; font-weight: bold;">PERHATIAN!</span> Email ini hanya diperuntukan untuk keperluan pembelajaran di Lumoshive Academy</p>
					<hr>
					<p>© 2024 Lumoshive Academy. All rights reserved.</p>
				  </div>
				</div>
			  </body>
			</html>
		`, req.Name, req.Subject, req.Text, req.To)

	// Setup Gomail
	hostUser := handlerEmail.Config.EmailHostUser
	hostPassword := handlerEmail.Config.EmailHostPassword

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("Lumoshive Academy <%s>", hostUser))
	m.SetHeader("To", req.To)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/plain", req.Text) // Fallback plain text
	m.SetBody("text/html", htmlContent)

	// Gmail SMTP config: Host: smtp.gmail.com, Port: 587
	d := gomail.NewDialer("smtp.gmail.com", 587, hostUser, hostPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("Error sending email:", err)
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "Failed to send email", nil)
		return
	}

	fmt.Println("Email sent successfully to:", req.To)
	utils.ResponseSuccess(w, http.StatusOK, "Email sent successfully", nil)
}
