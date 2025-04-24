package config

import "os"

type AuthenticationClient struct {
	ClientId     string
	ClientSecret string
	RedirectUrl  string
}

var VERSION string = "dev"

const (
	SIGNER_S3_URL         = "https://t1ey25e0i2.execute-api.us-east-1.amazonaws.com"
	DOWNLOAD_OPERATION    = "download"
	UPLOAD_OPERATION      = "upload"
	EPHIMERAL_SERVER_PORT = "8888"
	SHORTENER_URL         = "https://short-service.corvux.co"
)

var OLX_CONFIG_PATH string = os.Getenv("HOME") + "/.olx/auth.json"
var AuthenticationClientConfig = AuthenticationClient{
	ClientId:     "6e75vs6eqq7ghql61o8cm21ig5",
	ClientSecret: "c7dp96t3t0rfcclb96p6ng4odpiojqp3f4fppj51ev5nt5l7sf3",
	RedirectUrl:  "http://localhost:" + EPHIMERAL_SERVER_PORT + "/code",
}
var S3SignerServiceAllowOperations = []string{"upload", "download"}
