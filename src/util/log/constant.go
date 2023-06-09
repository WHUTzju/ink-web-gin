package log

const (
	LogCategoryLogrus   = "logrus"
	LogCategoryZap      = "zap"
	LogLineNumKey       = "LineNum"
	LogErrorKey         = "Err"
	LogSkipHelperCtxKey = "LogSkipHelper"
	LogHiddenSqlCtxKey  = "LogHiddenSql"
)

const (
	MiddlewareUrlPrefix                      = "api"
	MiddlewareIdempotencePrefix              = "idempotence"
	MiddlewareIdempotenceExpire              = 24
	MiddlewareIdempotenceTokenName           = "api-idempotence-token"
	MiddlewareOperationLogNotLogin           = "not login"
	MiddlewareOperationLogApiCacheKey        = "operation_log_api"
	MiddlewareOperationLogSkipPathDict       = "OperationLogSkipPath"
	MiddlewareOperationLogMaxCountBeforeSave = 100
	MiddlewareRequestIdCtxKey                = "RequestId"
	MiddlewareTraceIdCtxKey                  = "TraceId"
	MiddlewareSpanIdCtxKey                   = "SpanId"
	MiddlewareTransactionTxCtxKey            = "tx"
	MiddlewareTransactionForceCommitCtxKey   = "ForceCommitTx"
	MiddlewareJwtUserCtxKey                  = "user"
	MiddlewareSignSeparator                  = "|"
	MiddlewareSignTokenHeaderKey             = "X-Sign-Token"
	MiddlewareSignAppIdHeaderKey             = "appid"
	MiddlewareSignTimestampHeaderKey         = "timestamp"
	MiddlewareSignSignatureHeaderKey         = "signature"
	MiddlewareAccessLogIpLogKey              = "Ip"
	MiddlewareParamsQueryCtxKey              = "ParamsQuery"
	MiddlewareParamsBodyCtxKey               = "ParamsBody"
	MiddlewareParamsNullBody                 = "{}"
	MiddlewareParamsQueryLogKey              = "Query"
	MiddlewareParamsBodyLogKey               = "Body"
	MiddlewareParamsRespCtxKey               = "ResponseBody"
	MiddlewareParamsRespLogKey               = "Resp"
	MiddlewareCorsOrigin                     = "*"
	MiddlewareCorsHeaders                    = "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Sign-Token,api-idempotence-token"
	MiddlewareCorsMethods                    = "OPTIONS,GET,POST,PUT,PATCH,DELETE"
	MiddlewareCorsExpose                     = "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type"
	MiddlewareCorsCredentials                = "true"
)
