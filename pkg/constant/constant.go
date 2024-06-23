package constant

// server
const ServerReadTimeout = 30
const ServerReadHeaderTimeout = 10
const ServerWriteTimeout = 30

// request
const RequestTimeout = 120

// background process timeout
const BackgroundProcessTimeout = 300

const ApiPattern string = "/api"
const V1 = "/v1"
const FilesPattern = "/files"

var AllowedMimeTypes = [2]string{
	"text/",
	"image/",
}

// cors configs
var AllowedMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
var AllowedHeaders = []string{"*"}
var AllowCredentials = false

const InternalServerError = "internal server error"
const BadRequest = "bad request"
