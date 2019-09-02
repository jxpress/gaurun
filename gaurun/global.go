package gaurun

import (
	"net/http"

	"github.com/jxpress/gaurun/gcm"
	"github.com/sideshow/apns2"

	"go.uber.org/zap"
)

var (
	// Toml configuration for Gaurun
	ConfGaurun ConfToml
	// push notification Queue
	QueueNotification chan RequestGaurunNotification
	// TLS certificate and key for APNs
	CertificatePemIos CertificatePem
	// APNs Auth Key
	AuthKeyIos []byte
	// Stat for Gaurun
	StatGaurun StatApp
	// http client for APNs and GCM/FCM
	APNSClient      *http.Client
	APNSClientToken *apns2.Client
	GCMClient       *gcm.Client
	// access and error logger
	LogAccess *zap.Logger
	LogError  *zap.Logger
	// sequence ID for numbering push
	SeqID uint64
)
