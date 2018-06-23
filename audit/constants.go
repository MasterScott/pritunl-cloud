package audit

const (
	AdminLogin                 = "admin_login"
	AdminLoginFailed           = "admin_login_failed"
	AdminLogout                = "admin_logout"
	AdminPrimaryApprove        = "admin_primary_approve"
	AdminSecondaryApprove      = "admin_secondary_approve"
	AdminDeviceApprove         = "admin_device_approve"
	AdminDeviceRegisterRequest = "admin_device_register_request"
	AdminDeviceRegister        = "admin_device_register"

	ProxyLogin                 = "proxy_login"
	ProxyLoginFailed           = "proxy_login_failed"
	ProxyLogout                = "proxy_logout"
	ProxyPrimaryApprove        = "proxy_primary_approve"
	ProxySecondaryApprove      = "proxy_secondary_approve"
	ProxyDeviceApprove         = "proxy_device_approve"
	ProxyDeviceRegisterRequest = "proxy_device_register_request"
	ProxyDeviceRegister        = "proxy_device_register"

	UserLogin                 = "user_login"
	UserLoginFailed           = "user_login_failed"
	UserLogout                = "user_logout"
	UserLogoutAll             = "user_logout_all"
	UserPrimaryApprove        = "user_primary_approve"
	UserSecondaryApprove      = "user_secondary_approve"
	UserDeviceApprove         = "user_device_approve"
	UserDeviceRegisterRequest = "user_device_register_request"
	UserDeviceRegister        = "user_device_register"
	UserAccountDisable        = "user_account_disable"

	DuoApprove      = "duo_approve"
	DuoDeny         = "duo_deny"
	OneLoginApprove = "one_login_approve"
	OneLoginDeny    = "one_login_deny"
	OktaApprove     = "okta_approve"
	OktaDeny        = "okta_deny"
)
