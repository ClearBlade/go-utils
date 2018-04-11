package constants

const (
	NOT_YET_IMPLEMENTED              = "Endpoint is not yet implemented"
	INVALID_METHOD                   = "Invalid method"
	INTERNAL_SERVER_ERROR            = "Internal Server Error"
	MISSING_BODY_IN_REQUEST_ERROR    = "Body Required"
	ROUTE_NOT_FOUND                  = "Route not found"
	UNREACHABLE_CODE                 = "Unreachable Code Reached. Incredible."
	INVALID_PERMISSIONS              = "Invalid Permissions"
	NOT_AUTHORIZED_FOR_SYSTEM        = "User does not have access to system"
	FETCH_EVENT_DEFINITIONS_ERROR    = "Unable to fetch event definitions"
	FETCH_ALL_EVENT_HANDLERS         = "Unable to fetch event handlers"
	FETCH_EVENT_HANDLER              = "Unable to fetch event handler"
	CREATE_EVENT_HANDLER_ERROR       = "Unable to create event handler"
	UPDATE_EVENT_HANDLER_ERROR       = "Unable to update event handler"
	REMOVE_EVENT_HANDLER_ERROR       = "Unable to remove event handler"
	NO_TRIGGER_NAME_SPECIFIED        = "No name specified for trigger"
	FETCH_ALL_TIMERS_ERROR           = "Unable to fetch all timers"
	CREATE_TIMER_ERROR               = "Unable to create timer"
	UPDATE_TIMER_ERROR               = "Unable to update timer"
	DELETE_TIMER_ERROR               = "Unable to delete timer"
	FETCH_TIMER_ERROR                = "Unable to fetch timer"
	JSON_MARSHAL_ERROR               = "Unable to marshal JSON"
	JSON_UNMARSHAL_ERROR             = "Unable to unmarshal JSON"
	JSON_DECODE_ERROR                = "Poorly formatted json object"
	NO_ANALYTICS_QUERY_PRESENT       = "No query present"
	UNABLE_TO_PARSE_QUERY            = "Unable to parse query"
	UNABLE_TO_GET_STORAGE_IN_BYTES   = "Unable to fetch storage size"
	BAD_REQUEST_ERROR                = "Bad request"
	DEVICE_ARG_VALIDATION_ERROR      = "Unable to validate arguments for device"
	DEVICE_ROLE_ARG_VALIDATION_ERROR = "Unable to validate arguments for device roles"
	FETCH_ALL_USERS_ERROR            = "Unable to fetch all users"
	USER_MISSING_USER_KEY_ERROR      = "Must contain 'user' key with apropos value"
	USER_MISSING_CHANGES_KEY_ERROR   = "Must contain 'changes' key with apropos data"
	USER_ROLE_ARG_VALIDATION_ERROR   = "Unable to validate arguments for user roles"
	SERVICE_NAME_TOO_SHORT_ERROR     = "Name too short"
	SERVICE_CODE_NOT_FOUND_ERROR     = "No service with that name"
	ENDPOINT_DEPRECATED_ERROR        = "This endpoint is deprecated"

	GITHUB_AUTH_TOKEN_MISSING_ERROR          = "Request header must contain github Oauth access-token"
	GITHUB_REPO_CREATION_ERROR               = "Unable to create github repo"
	NON_FATAL_IPM_ERROR                      = "Some Errors occured, but didn't completely fail"
	EXPORT_TO_FILE_SYSTEM_ERROR              = "Unable to perform selective export to filesystem"
	MISSING_QUERY_TAG_ERROR                  = "queryTag is missing when fetching tags"
	CREATE_ANON_USER_WITH_SYSTEM_CREDS_ERROR = "Unable to create anon user with system creds"
	SYSTEM_TAG_FETCH_ERROR                   = "Unable to fetch system tags"
)