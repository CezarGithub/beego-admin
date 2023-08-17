package global

import "quince/conf"

// URL_CURRENT
const URL_CURRENT = "url://current"

// URL_RELOAD refresh page
const URL_RELOAD = "url://reload"

// URL_BACK Return to the previous page
const URL_BACK = "url://back"

// URL_CLOSE_LAYER Close the current layer pop-up window
const URL_CLOSE_LAYER = "url://close-layer"

// URL_CLOSE_REFRESH Close the current pop-up window and refresh the parent
const URL_CLOSE_REFRESH = "url://close-refresh"

// LOGIN_USER Login user key
const LOGIN_USER = "loginUser"

// LOGIN_USER_ID Login user id
const LOGIN_USER_ID = "LoginUserId"

// LOGIN_USER_ID_SIGN Sign in user signature
const LOGIN_USER_ID_SIGN = "loginUserIdSign"



var (
	// BA_CONFIG conf.Server
	BA_CONFIG conf.Server
)
