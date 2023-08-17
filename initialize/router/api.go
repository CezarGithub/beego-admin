package router

// Header ["Authorization"] = X-API-Key
// Header [X-API-Key] contains JWT key
// Header [X-Requested-With] = XMLHttpRequest
// POST method only
// JWT key in response to login at admin/api/login with user + pass or MAC address
// Returns myController.ResponseSuccessWithData("My message", data, myController.Ctx) or ResponseWIthError...
type API int

const (
	Version1 API = iota + 1
	Version2
	Version3
)
