package pcm

import "github.com/go-msvc/utils/results"

var (
	pcmResults, ResultSuccess = results.New()
	FailedGetProfileA         = pcmResults.Add(1, "FAILED_GET_PROFILE_A", "Failed to get subscriber profile for the sender")
	FailedSetProfileA         = pcmResults.Add(2, "FAILED_SET_PROFILE_A", "Failed to set subscriber profile for the sender")
	FailedGetProfileB         = pcmResults.Add(3, "FAILED_GET_PROFILE_B", "Failed to get subscriber profile for the recipient")
	LimitedA                  = pcmResults.Add(4, "LIMITED_A", "Sending too fast")
	LimitedA2B                = pcmResults.Add(5, "LIMITED_A2B", "Sending too fast to recipient")
	FailedSmsRenderA          = pcmResults.Add(6, "FAILED_SMS_RENDER_A", "Failed to execute template for SMS to sender")
	FailedSmsRenderB          = pcmResults.Add(7, "FAILED_SMS_RENDER_B", "Failed to execute template for SMS to recipient")
	FailedSriSmA              = pcmResults.Add(8, "FAILED_SRISM_A", "Failed to do SRI-SM for sender")
	FailedSriSmB              = pcmResults.Add(9, "FAILED_SRISM_B", "Failed to do SRI-SM for recipient")
	FailedFwdSmA              = pcmResults.Add(10, "FAILED_FWDSM_A", "Failed to send FWD-SM to sender")
	FailedFwdSmB              = pcmResults.Add(11, "FAILED_FWDSM_B", "Failed to send FWD-SM to recipient")
	FailedIcapA               = pcmResults.Add(12, "FAILED_ICAP_A", "Failed to get sender status from ICAP")
	IcapANotActive            = pcmResults.Add(13, "ICAP_A_NOT_ACTIVE", "ICAP says MSISDN of sender is not active")
	BlockedByB                = pcmResults.Add(14, "BLOCKED_BY_B", "Recipient blocked receiving of any pcm")
	NotAllowedToBImsiNetwork  = pcmResults.Add(15, "NOT_ALLOW_B_IMSI", "Not allowed to network that B IMSI belongs to")
	NotAllowedToBVlrNetwork   = pcmResults.Add(16, "NOT_ALLOW_B_VLR", "Not allowed to network where B currently roams")
)
