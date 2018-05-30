package appCode

const (
	INVALID_METHOD                   = 101
	INTERNAL_SERVER_ERROR            = 102
	ROUTE_NOT_FOUND_ERROR            = 103
	JSON_MARSHAL_ERROR               = 104
	JSON_UNMARSHAL_ERROR             = 105
	JSON_DECODE_ERROR                = 106
	READ_BODY_ERROR                  = 107
	NOT_YET_IMPLEMENTED              = 108
	UNREACHABLE_CODE                 = 109
	MISSING_BODY_IN_REQUEST_ERROR    = 110
	DEVELOPER_ERROR                  = 111
	BAD_SSL_REQUEST_ERROR            = 112
	SSL_LOGIN_ERROR                  = 113
	GET_TRUST_CHAIN_ERROR            = 114
	EXTRACT_CERTS_FROM_CACERTS_ERROR = 115
	VERIFY_CERT_ERROR                = 116
	CERT_SIGNATURE_ERROR             = 117
	REVOKED_CERT_ERROR               = 118
	OPEN_CERTIFICATE_ERROR           = 119
	ENDPOINT_DEPRECATED_ERROR        = 120
	PAYLOAD_PARSE_ERROR              = 121

	PERMISSIONS_ERROR                   = 200
	NOT_AUTHORIZED_FOR_SYSTEM           = 201
	BAD_USER_REGISTRATION_REQUEST_ERROR = 202
	USER_REGISTRATION_ERROR             = 203
	INVALID_TOKEN_ERROR                 = 204

	BAD_USER_REQUEST_ERROR               = 204
	USER_READ_ERROR                      = 205
	USER_CREATE_ERROR                    = 206
	USER_DELETE_ERROR                    = 207
	USER_UPDATE_ERROR                    = 208
	USER_ARG_VALIDATION_ERROR            = 209
	USER_ROLE_READ_ERROR                 = 210
	USER_ROLE_ADD_ERROR                  = 211
	USER_ROLE_UPDATE_ERROR               = 212
	USER_ROLE_DELETE_ERROR               = 213
	USER_ROLE_ARG_VALIDATION_ERROR       = 214
	FETCH_ALL_USERS_ERROR                = 215
	MISSING_USER_KEY_IN_REQUEST_ERROR    = 216
	MISSING_CHANGES_KEY_IN_REQUEST_ERROR = 217
	BAD_USER_COLUMN_REQUEST_ERROR        = 218
	CREATE_USER_COLUMN_ERROR             = 219
	DELETE_USER_COLUMN_ERROR             = 220
	READ_USER_COLUMNS_ERROR              = 221
	NO_USER_FOR_ID                       = 222
	USER_DOESNT_EXIST_ERROR              = 223

	COLLECTION_COUNT_ERROR                    = 300
	BAD_COLLECTION_NAME_COLUMNS_REQUEST_ERROR = 301
	NO_SYSTEM_FOR_COLLECTION_ID               = 302
	COLLECTION_COLUMNS_ERROR                  = 303
	COLLECTION_ID_BY_NAME_ERROR               = 304
	BAD_QUERY_REQUEST_ERROR                   = 305
	GET_COLLECTION_DATA_ERROR                 = 306
	COLLECTION_RETURN_DATA_ERROR              = 307
	COLLECTION_INSERT_ERROR                   = 308
	MISSING_QUERY_ERROR                       = 309
	QUERY_PARSE_ERROR                         = 310
	NO_CHANGES_IN_QUERY_ERROR                 = 311
	TYPE_QUERY_REQUEST_ERROR                  = 312
	TYPES_TO_UPSERT_ERROR                     = 313
	COLLECTION_UPDATE_ERROR                   = 314
	COLLECTION_DELETE_ERROR                   = 315
	NO_COLLECTION_ID_IN_REQUEST_ERROR         = 316
	SELECT_COLLECTIONS_ERROR                  = 317

	LIBRARY_READ_ERROR                     = 400
	LIBRARY_CREATE_ERROR                   = 401
	LIBRARY_UPDATE_ERROR                   = 402
	LIBRARY_DELETE_ERROR                   = 403
	BAD_LIBRARY_REQUEST_ERROR              = 404
	EXECUTE_CODE_ERROR                     = 405
	GET_SERVICE_CODE_ERROR                 = 406
	NAME_TOO_SHORT_ERROR                   = 407
	INVALID_SERVICE_NAME_ERROR             = 408
	SERVICE_CREATION_ERROR                 = 409
	SERVICE_DELETION_ERROR                 = 410
	GET_SERVICE_LIST_ERROR                 = 411
	BAD_GET_LOGS_REQUEST_ERROR             = 412
	TOGGLE_LOGGING_ERROR                   = 413
	GET_LOGGING_ENABLED_ERROR              = 414
	GET_FAILED_SERVICE_FROM_LIST_ERROR     = 415
	RETRY_FAILED_SERVICE_ERROR             = 416
	BAD_RETRY_FAILED_SERVICE_REQUEST_ERROR = 417
	GET_COMPLETED_SERVICE_FROM_LIST_ERROR  = 418
	SERVICE_CODE_NOT_FOUND_ERROR           = 419
	EDIT_SERVICE_CODE_ERROR                = 420
	LIBRARY_NOT_FOUND_ERROR                = 421

	MESSAGE_HISTORY_FETCH_ERROR       = 500
	MESSAGE_HISTORY_DELETE_ERROR      = 501
	BAD_MESSAGE_HISTORY_REQUEST_ERROR = 502
	CURRENT_TOPICS_FETCH_ERROR        = 503
	BAD_PUBLISH_REQUEST_ERROR         = 504
	MESSAGE_INJECTION_ERROR           = 505
	BAD_AVERAGE_PAYLOAD_REQUEST_ERROR = 506
	AVERAGE_PAYLOAD_ERROR             = 507

	FETCH_EVENT_DEFINITIONS_ERROR = 600
	FETCH_ALL_EVENT_HANDLERS      = 601
	CREATE_EVENT_HANDLER_ERROR    = 602
	FETCH_EVENT_HANDLER_ERROR     = 603
	UPDATE_EVENT_HANDLER_ERROR    = 604
	REMOVE_EVENT_HANDLER_ERROR    = 605
	FETCH_ALL_TIMERS_ERROR        = 606
	CREATE_TIMER_ERROR            = 607
	FETCH_TIMER_ERROR             = 608
	UPDATE_TIMER_ERROR            = 609
	DELETE_TIMER_ERROR            = 610
	NO_TRIGGER_NAME_SPECIFIED     = 611

	BAD_ANALYTICS_REQUEST_ERROR               = 900
	STORAGE_ENDPOINT_ERROR                    = 901
	COUNT_ENDPOINT_ERROR                      = 902
	EVENT_TOTALS_ENDPOINT_ERROR               = 903
	EVENT_LIST_ENDPOINT_ERROR                 = 904
	USER_EVENT_ENDPOINT_ERROR                 = 905
	USER_ACTION_REPORT_ERROR                  = 906
	BAD_PLATFORM_STATS_QUERY_REQUEST_ERROR    = 907
	PLATFORM_STATS_QUERY_ERROR                = 908
	PLATFORM_STATS_HISTORY_QUERY_ERROR        = 909
	BAD_PLATFORM_DB_STATS_QUERY_REQUEST_ERROR = 910
	PLATFORM_DB_STATS_QUERY_ERROR             = 911
	BAD_PLATFORM_LOGS_QUERY_REQUEST_ERROR     = 912
	PLATFORM_LOGS_QUERY_ERROR                 = 913
	PLATFORM_LOGS_ALL_ERROR                   = 914
	USER_LOGIN_ERROR                          = 915
	ANON_USER_LOGIN_ERROR                     = 916
	USER_LIST_COUNT_ERROR                     = 917
	AUTH_OR_REG_OVERRIDE_WITH_SERVICE_ERROR   = 918
	USER_IS_AUTHED_ERROR                      = 919
	USER_LOGOUT_ERROR                         = 920
	GET_USER_INFO_ERROR                       = 921
	UPDATE_USER_INFO_ERROR                    = 922
	BAD_UPDATE_USER_INFO_REQUEST_ERROR        = 923
	BAD_CHANGE_PASSWORD_REQUEST_ERROR         = 924
	UPDATE_USER_PASSWORD_ERROR                = 925
	GET_USER_LIST_ERROR                       = 926

	READ_PORTAL_LIST_ERROR       = 1000
	CREATE_PORTAL_ERROR          = 1001
	INVALID_PORTAL_REQUEST_ERROR = 1002
	READ_PORTAL_ERROR            = 1003
	UPDATE_PORTAL_ERROR          = 1004
	DELETE_PORTAL_ERROR          = 1005

	BAD_DEVICES_REQUEST_ERROR        = 1100
	DEVICE_AUTH_ERROR                = 1101
	DEVICE_READ_ERROR                = 1102
	DEVICE_CREATE_ERROR              = 1103
	DEVICE_DELETE_ERROR              = 1104
	DEVICE_UPDATE_ERROR              = 1105
	DEVICE_ARG_VALIDATION_ERROR      = 1106
	DEVICE_ROLE_READ_ERROR           = 1107
	DEVICE_ROLE_CREATE_ERROR         = 1108
	DEVICE_ROLE_UPDATE_ERROR         = 1109
	DEVICE_ROLE_DELETE_ERROR         = 1110
	DEVICE_ROLE_ARG_VALIDATION_ERROR = 1111
	DEVICE_COLUMN_CREATE_ERROR       = 1112
	DEVICE_COLUMN_READ_ERROR         = 1113
	DEVICE_COLUMN_DELETE_ERROR       = 1114
	DEVICE_COUNT_ERROR               = 1115

	BAD_SYSYTEM_SYNC_REQUEST_ERROR = 1200
	SYSTEM_SYNC_ERROR              = 1201
	BAD_EDGE_SYNC_REQUEST_ERROR    = 1202
	EDGE_SYNC_UPDATE_ERROR         = 1203

	METRICS_FETCH_ERROR = 1300

	ADMIN_SYSTEM_USAGE_ERROR                 = 1400
	ADMIN_DEVELOPER_ERROR                    = 1401
	BAD_SYSTEM_ADMIN_REQUEST_ERROR           = 1402
	BAD_TTL_TOKEN_ERROR                      = 1403
	SYSTEM_ADMIN_ERROR                       = 1404
	ALL_SYSTEMS_ERROR                        = 1405
	DISABLED_SYSTEMS_ERROR                   = 1406
	BAD_BIO_NOTES_REQUEST_ERROR              = 1407
	BIO_NOTES_ERROR                          = 1408
	ALL_DEVELOPERS_ERROR                     = 1409
	BAD_COLLECTION_MANAGEMENT_REQUEST_ERROR  = 1410
	COLLECTION_MANAGEMENT_ERROR              = 1411
	BAD_ALL_COLLECTIONS_REQUEST_ERROR        = 1412
	ALL_COLLECTIONS_ERROR                    = 1413
	GET_DEVELOPERS_FOR_SYSTEM_ERROR          = 1414
	UPDATE_DEVELOPERS_FOR_SYSTEM_ERROR       = 1415
	GET_SYSTEMS_FOR_DEVELOPER_ERROR          = 1416
	BAD_ADMIN_REGISTRATION_REQUEST_ERROR     = 1417
	ADMIN_REGISTRATION_ERROR                 = 1418
	GET_DEV_DATA_ERROR                       = 1419
	BAD_ADMIN_AUTH_REQUEST_ERROR             = 1420
	ADMIN_AUTH_ERROR                         = 1421
	ADMIN_LOGOUT_ERROR                       = 1422
	BAD_ADMIN_CHANGE_PASSWORD_REQUEST_ERROR  = 1423
	ADMIN_CHANGE_PASSWORD_ERROR              = 1424
	GET_ROLE_INFO_ERROR                      = 1425
	BAD_CREATE_ROLE_REQUEST_ERROR            = 1426
	CREATE_ROLE_ERROR                        = 1427
	BAD_UPDATE_ROLE_REQUEST_ERROR            = 1428
	UPDATE_ROLE_ERROR                        = 1429
	BAD_DELETE_ROLE_REQUEST_ERROR            = 1430
	DELETE_ROLE_ERROR                        = 1431
	BAD_UPDATE_PASSWORD_AS_DEV_REQUEST_ERROR = 1432
	SYSTEM_USAGE_ERROR                       = 1433
	BAD_SYSTEM_NAME                          = 1434

	CREATE_PLUGIN_ERROR          = 1500
	INVALID_PLUGIN_REQUEST_ERROR = 1501
	READ_PLUGIN_ERROR            = 1502
	READ_PLUGIN_LIST_ERROR       = 1503
	UPDATE_PLUGIN_ERROR          = 1504
	DELETE_PLUGIN_ERROR          = 1505

	BAD_EDGE_COLUMN_REQUEST_ERROR  = 1600
	CREATE_EDGE_COLUMN_ERROR       = 1601
	DELETE_EDGE_COLUMN_ERROR       = 1602
	READ_EDGE_COLUMNS_ERROR        = 1603
	READ_EDGE_LIST_ERROR           = 1604
	READ_EDGE_ERROR                = 1605
	INVALID_EDGE_REQUEST_ERROR     = 1606
	CREATE_EDGE_ERROR              = 1607
	UPDATE_EDGE_ERROR              = 1608
	DELETE_EDGE_ERROR              = 1609
	UPGRADE_EDGE_ERROR             = 1610
	BAD_UPGRADE_EDGE_REQUEST_ERROR = 1611
	EDGE_NOT_CONNECTED_ERROR       = 1612
	EDGE_PROXY_ERROR               = 1613
	EDGES_COUNT_ERROR              = 1614

	BAD_SYSTEM_DEPLOY_REQUEST_ERROR = 1701
	CREATE_DEPLOY_ERROR             = 1702
	DELETE_DEPLOY_ERROR             = 1703
	READ_ALL_DEPLOY_ERROR           = 1704
	UPDATE_DEPLOY_ERROR             = 1705

	BAD_ADAPTORS_REQUEST_ERROR = 1900
	ADAPTOR_READ_ERROR         = 1901
	ADAPTOR_CREATE_ARGS_ERROR  = 1902
	ADAPTOR_UPDATE_ARGS_ERROR  = 1903
	ADAPTOR_CREATE_ERROR       = 1904
	ADAPTOR_UPDATE_ERROR       = 1905
	ADAPTOR_DELETE_ERROR       = 1906
	ADAPTOR_FILE_ARGS_ERROR    = 1907
	ADAPTOR_FILE_DEPLOY_ERROR  = 1908
	ADAPTOR_FILE_CONTROL_ERROR = 1909

	EDGE_GROUPS_READ_ERROR   = 2100
	EDGE_GROUPS_CREATE_ERROR = 2101
	EDGE_GROUPS_UPDATE_ERROR = 2102
	EDGE_GROUPS_DELETE_ERROR = 2103

	DEPLOYMENTS_READ_ERROR   = 4100
	DEPLOYMENTS_CREATE_ERROR = 4101
	DEPLOYMENTS_UPDATE_ERROR = 4102
	DEPLOYMENTS_DELETE_ERROR = 4103

	// IPM Specific Errors: (Blocked Range: 5000-6999)
	NON_FATAL_IPM_ERROR                      = 5001
	QUERY_PARAM_ERROR                        = 5002
	CREATE_ANON_USER_WITH_SYSTEM_CREDS_ERROR = 5003
	EXPORTER_INITIALIZATION_ERROR            = 5004
	SYSTEM_FOLDER_CREATION_ERROR             = 5005
	IMPORTING_USER_INFO_ERROR                = 5006
	FAILED_SYSTEM_IMPORT                     = 5007
	GITHUB_MISSING_AUTH_TOKEN_ERROR          = 5101
	GITHUB_REPO_CREATION_ERROR               = 5102
	GITHUB_REPO_PROCESSING_FOR_IMPORT_ERROR  = 5103
	GITHUB_AUTHCODE_REQUEST_ERROR            = 5103
	GITHUB_TOKEN_REQUEST_ERROR               = 5104
	GITHUB_TOKEN_REVOKE_ERROR                = 5105
	GITHUB_AUTHORIZATION_CHECK_ERROR         = 5107
	GITHUB_GET_SHA_OF_REPO_ERROR             = 5106
	GITHUB_GET_TREE_OF_REPO_ERROR            = 5108
	EXPORT_TO_FILE_SYSTEM_ERROR              = 5201
	SYSTEM_TAG_FETCH_ERROR                   = 5300
	OS_REMOVE_DIR_ERROR                      = 6000
	LIST_FILES_FROM_DIR_ERROR                = 5008
	INVALID_SYSTEM_NAME                      = 5009

	// cb-console Specific Errors: (Blocked Range: 7000-8999)
	PARSE_ZIP_FILE_ERROR = 7000
	MISSING_TOKEN_ERROR  = 7001

	// folder-specific errors: (Blocked Range: 9000-9999)

	FOLDERS_READ_ERROR   = 9000
	FOLDERS_CREATE_ERROR = 9001
	FOLDERS_UPDATE_ERROR = 9002
	FOLDERS_DELETE_ERROR = 9003
)
