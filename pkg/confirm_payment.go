package pkg

import "fmt"

type ConfirmPaymentRequest struct {
	// application trying to use bpc hack, for information purpose only
	// to identify each user's request one from another
	MDOrder         string `json:"md_order"`
	ACSRequestId    string `json:"acs-request-id"`
	ACSSessionUrl   string `json:"acs-session-url,omitempty"`
	OneTimePassword string `json:"one-time-password"`
	TerminateUrl    string `json:"terminate-url"`
}

type ConfirmPaymentResponse struct {
	Status         HackResponseStatus `json:"status"`
	CurrentAttempt int                `json:"current-attempt,omitempty"`
	TotalAttempts  int                `json:"total-attempts,omitempty"`
	FinalUrl       string             `json:"final-url,omitempty"`
}

func (s *ConfirmPaymentResponse) String() string {
	return fmt.Sprintf("ConfirmPaymentResponse {status: %v, cur: %d, tot: %d, finalUrl: %v}",
		s.Status, s.CurrentAttempt, s.TotalAttempts, s.FinalUrl)
}
