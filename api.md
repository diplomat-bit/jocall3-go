# Shared Params Types

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared#AddressParam">AddressParam</a>

# Shared Response Types

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared#Transaction">Transaction</a>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserLoginResponse">UserLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserRegisterResponse">UserRegisterResponse</a>

Methods:

- <code title="post /users/login">client.Users.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserLoginParams">UserLoginParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserLoginResponse">UserLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/logout">client.Users.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserService.Logout">Logout</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/register">client.Users.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserRegisterParams">UserRegisterParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserRegisterResponse">UserRegisterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PasswordReset

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetConfirmResponse">UserPasswordResetConfirmResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetInitiateResponse">UserPasswordResetInitiateResponse</a>

Methods:

- <code title="post /users/password-reset/confirm">client.Users.PasswordReset.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetConfirmParams">UserPasswordResetConfirmParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetConfirmResponse">UserPasswordResetConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/password-reset/initiate">client.Users.PasswordReset.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetInitiateParams">UserPasswordResetInitiateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserPasswordResetInitiateResponse">UserPasswordResetInitiateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Me

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeGetResponse">UserMeGetResponse</a>

Methods:

- <code title="get /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeGetResponse">UserMeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="delete /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Preferences

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceGetResponse">UserMePreferenceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceUpdateResponse">UserMePreferenceUpdateResponse</a>

Methods:

- <code title="get /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceGetResponse">UserMePreferenceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceUpdateParams">UserMePreferenceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceUpdateResponse">UserMePreferenceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Security

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityGetLogResponse">UserMeSecurityGetLogResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityRotateKeysResponse">UserMeSecurityRotateKeysResponse</a>

Methods:

- <code title="get /users/me/security/log">client.Users.Me.Security.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityService.GetLog">GetLog</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityGetLogParams">UserMeSecurityGetLogParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityGetLogResponse">UserMeSecurityGetLogResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/me/security/rotate-keys">client.Users.Me.Security.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityService.RotateKeys">RotateKeys</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeSecurityRotateKeysResponse">UserMeSecurityRotateKeysResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceListResponse">UserMeDeviceListResponse</a>

Methods:

- <code title="get /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceListResponse">UserMeDeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/me/devices/{deviceId}">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceService.Deregister">Deregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceRegisterParams">UserMeDeviceRegisterParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Biometrics

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricGetStatusResponse">UserMeBiometricGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>

Methods:

- <code title="delete /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/me/biometrics/enroll">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricService.Enroll">Enroll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricEnrollParams">UserMeBiometricEnrollParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricGetStatusResponse">UserMeBiometricGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/me/biometrics/verify">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricVerifyParams">UserMeBiometricVerifyParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountListResponse">AccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountLinkResponse">AccountLinkResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOpenResponse">AccountOpenResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountGetDetailsResponse">AccountGetDetailsResponse</a>

Methods:

- <code title="get /accounts/me">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /accounts/{accountId}">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.Close">Close</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /accounts/link">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.Link">Link</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountLinkParams">AccountLinkParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountLinkResponse">AccountLinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/open">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.Open">Open</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOpenParams">AccountOpenParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOpenResponse">AccountOpenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/details">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountGetDetailsResponse">AccountGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListArchivedResponse">AccountTransactionListArchivedResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListPendingResponse">AccountTransactionListPendingResponse</a>

Methods:

- <code title="get /accounts/{accountId}/transactions/archived">client.Accounts.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionService.ListArchived">ListArchived</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListArchivedParams">AccountTransactionListArchivedParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListArchivedResponse">AccountTransactionListArchivedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/transactions/pending">client.Accounts.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionService.ListPending">ListPending</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListPendingResponse">AccountTransactionListPendingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BalanceHistory

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountBalanceHistoryGetResponse">AccountBalanceHistoryGetResponse</a>

Methods:

- <code title="get /accounts/{accountId}/balance-history">client.Accounts.BalanceHistory.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountBalanceHistoryService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountBalanceHistoryGetParams">AccountBalanceHistoryGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountBalanceHistoryGetResponse">AccountBalanceHistoryGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Statements

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementListResponse">AccountStatementListResponse</a>

Methods:

- <code title="get /accounts/{accountId}/statements">client.Accounts.Statements.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementListResponse">AccountStatementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/statements/{statementId}/pdf">client.Accounts.Statements.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementService.DownloadPdf">DownloadPdf</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, statementID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Overdraft

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftGetSettingsResponse">AccountOverdraftGetSettingsResponse</a>

Methods:

- <code title="get /accounts/{accountId}/overdraft-settings">client.Accounts.Overdraft.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftService.GetSettings">GetSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftGetSettingsResponse">AccountOverdraftGetSettingsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{accountId}/overdraft-settings">client.Accounts.Overdraft.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftService.UpdateSettings">UpdateSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftUpdateSettingsParams">AccountOverdraftUpdateSettingsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionListResponse">TransactionListResponse</a>

Methods:

- <code title="get /transactions/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/notes">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.AddNotes">AddNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionAddNotesParams">TransactionAddNotesParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /transactions/{transactionId}/categorize">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.Categorize">Categorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionCategorizeParams">TransactionCategorizeParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go/shared#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/{transactionId}/dispute">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.Dispute">Dispute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionDisputeParams">TransactionDisputeParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /transactions/{transactionId}/split">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.Split">Split</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionSplitParams">TransactionSplitParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Recurring

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringListResponse">TransactionRecurringListResponse</a>

Methods:

- <code title="post /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringNewParams">TransactionRecurringNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringListResponse">TransactionRecurringListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /transactions/recurring/{recurringId}">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringService.Cancel">Cancel</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, recurringID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Insights

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightGetFutureFlowResponse">TransactionInsightGetFutureFlowResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightGetSpendingTrendsResponse">TransactionInsightGetSpendingTrendsResponse</a>

Methods:

- <code title="get /transactions/insights/future-flow">client.Transactions.Insights.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightService.GetFutureFlow">GetFutureFlow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightGetFutureFlowResponse">TransactionInsightGetFutureFlowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/insights/spending-trends">client.Transactions.Insights.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightService.GetSpendingTrends">GetSpendingTrends</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightGetSpendingTrendsResponse">TransactionInsightGetSpendingTrendsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

## Advisor

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorChatResponse">AIAdvisorChatResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorGetHistoryResponse">AIAdvisorGetHistoryResponse</a>

Methods:

- <code title="post /ai/advisor/chat">client.AI.Advisor.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorService.Chat">Chat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorChatParams">AIAdvisorChatParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorChatResponse">AIAdvisorChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/advisor/chat/history">client.AI.Advisor.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorService.GetHistory">GetHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorGetHistoryResponse">AIAdvisorGetHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolListResponse">AIAdvisorToolListResponse</a>

Methods:

- <code title="get /ai/advisor/tools">client.AI.Advisor.Tools.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolListResponse">AIAdvisorToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/advisor/tools/{toolId}/enable">client.AI.Advisor.Tools.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolService.Enable">Enable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, toolID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Oracle

### Simulate

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunAdvancedResponse">AIOracleSimulateRunAdvancedResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunStandardResponse">AIOracleSimulateRunStandardResponse</a>

Methods:

- <code title="post /ai/oracle/simulate/advanced">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateService.RunAdvanced">RunAdvanced</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunAdvancedParams">AIOracleSimulateRunAdvancedParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunAdvancedResponse">AIOracleSimulateRunAdvancedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/oracle/simulate/monte-carlo">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateService.RunMonteCarlo">RunMonteCarlo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunMonteCarloParams">AIOracleSimulateRunMonteCarloParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /ai/oracle/simulate">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateService.RunStandard">RunStandard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunStandardParams">AIOracleSimulateRunStandardParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunStandardResponse">AIOracleSimulateRunStandardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Predictions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionGetInflationForecastResponse">AIOraclePredictionGetInflationForecastResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionGetMarketCrashProbabilityResponse">AIOraclePredictionGetMarketCrashProbabilityResponse</a>

Methods:

- <code title="get /ai/oracle/predictions/inflation">client.AI.Oracle.Predictions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionService.GetInflationForecast">GetInflationForecast</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionGetInflationForecastParams">AIOraclePredictionGetInflationForecastParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionGetInflationForecastResponse">AIOraclePredictionGetInflationForecastResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/oracle/predictions/market-crash-probability">client.AI.Oracle.Predictions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionService.GetMarketCrashProbability">GetMarketCrashProbability</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOraclePredictionGetMarketCrashProbabilityResponse">AIOraclePredictionGetMarketCrashProbabilityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Simulations

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationGetResponse">AIOracleSimulationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>

Methods:

- <code title="get /ai/oracle/simulations/{simulationId}">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, simulationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationGetResponse">AIOracleSimulationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/oracle/simulations">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Incubator

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorSubmitPitchResponse">AIIncubatorSubmitPitchResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorValidateIdeaResponse">AIIncubatorValidateIdeaResponse</a>

Methods:

- <code title="get /ai/incubator/pitches">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorService.ListPitches">ListPitches</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/incubator/pitch">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorService.SubmitPitch">SubmitPitch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorSubmitPitchParams">AIIncubatorSubmitPitchParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorSubmitPitchResponse">AIIncubatorSubmitPitchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/incubator/validate">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorService.ValidateIdea">ValidateIdea</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorValidateIdeaParams">AIIncubatorValidateIdeaParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorValidateIdeaResponse">AIIncubatorValidateIdeaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Analysis

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisGenerateSwotResponse">AIIncubatorAnalysisGenerateSwotResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisScanCompetitorsResponse">AIIncubatorAnalysisScanCompetitorsResponse</a>

Methods:

- <code title="post /ai/incubator/analysis/swot">client.AI.Incubator.Analysis.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisService.GenerateSwot">GenerateSwot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisGenerateSwotParams">AIIncubatorAnalysisGenerateSwotParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisGenerateSwotResponse">AIIncubatorAnalysisGenerateSwotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/incubator/analysis/competitors">client.AI.Incubator.Analysis.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisService.ScanCompetitors">ScanCompetitors</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisScanCompetitorsParams">AIIncubatorAnalysisScanCompetitorsParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorAnalysisScanCompetitorsResponse">AIIncubatorAnalysisScanCompetitorsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Pitch

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>

Methods:

- <code title="get /ai/incubator/pitch/{pitchId}/details">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /ai/incubator/pitch/{pitchId}/feedback">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchService.SubmitFeedback">SubmitFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchSubmitFeedbackParams">AIIncubatorPitchSubmitFeedbackParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Ads

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdListResponse">AIAdListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateCopyResponse">AIAdGenerateCopyResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateVideoResponse">AIAdGenerateVideoResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdOptimizeCampaignResponse">AIAdOptimizeCampaignResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGetOperationStatusResponse">AIAdGetOperationStatusResponse</a>

Methods:

- <code title="get /ai/ads">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdListResponse">AIAdListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/ads/generate/copy">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.GenerateCopy">GenerateCopy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateCopyParams">AIAdGenerateCopyParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateCopyResponse">AIAdGenerateCopyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/ads/generate/video">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.GenerateVideo">GenerateVideo</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateVideoParams">AIAdGenerateVideoParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateVideoResponse">AIAdGenerateVideoResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/ads/optimize">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.OptimizeCampaign">OptimizeCampaign</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdOptimizeCampaignParams">AIAdOptimizeCampaignParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdOptimizeCampaignResponse">AIAdOptimizeCampaignResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/ads/operations/{operationId}">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.GetOperationStatus">GetOperationStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGetOperationStatusResponse">AIAdGetOperationStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Agent

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentGetCapabilitiesResponse">AIAgentGetCapabilitiesResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentGetPromptsResponse">AIAgentGetPromptsResponse</a>

Methods:

- <code title="get /ai/agent/capabilities">client.AI.Agent.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentService.GetCapabilities">GetCapabilities</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentGetCapabilitiesResponse">AIAgentGetCapabilitiesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/agent/prompts">client.AI.Agent.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentService.GetPrompts">GetPrompts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentGetPromptsResponse">AIAgentGetPromptsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /ai/agent/prompts">client.AI.Agent.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentService.UpdatePrompts">UpdatePrompts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAgentUpdatePromptsParams">AIAgentUpdatePromptsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Models

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelFineTuneResponse">AIModelFineTuneResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelListVersionsResponse">AIModelListVersionsResponse</a>

Methods:

- <code title="post /ai/models/fine-tune">client.AI.Models.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelService.FineTune">FineTune</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelFineTuneParams">AIModelFineTuneParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelFineTuneResponse">AIModelFineTuneResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/models/versions">client.AI.Models.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelService.ListVersions">ListVersions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIModelListVersionsResponse">AIModelListVersionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Corporate

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateOnboardEntityResponse">CorporateOnboardEntityResponse</a>

Methods:

- <code title="post /corporate/onboard">client.Corporate.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateService.OnboardEntity">OnboardEntity</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateOnboardEntityParams">CorporateOnboardEntityParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateOnboardEntityResponse">CorporateOnboardEntityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Compliance

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenAdverseMediaResponse">CorporateComplianceScreenAdverseMediaResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenPepResponse">CorporateComplianceScreenPepResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenSanctionsResponse">CorporateComplianceScreenSanctionsResponse</a>

Methods:

- <code title="post /corporate/compliance/media">client.Corporate.Compliance.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceService.ScreenAdverseMedia">ScreenAdverseMedia</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenAdverseMediaParams">CorporateComplianceScreenAdverseMediaParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenAdverseMediaResponse">CorporateComplianceScreenAdverseMediaResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/compliance/pep">client.Corporate.Compliance.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceService.ScreenPep">ScreenPep</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenPepParams">CorporateComplianceScreenPepParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenPepResponse">CorporateComplianceScreenPepResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/compliance/sanctions">client.Corporate.Compliance.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceService.ScreenSanctions">ScreenSanctions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenSanctionsParams">CorporateComplianceScreenSanctionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceScreenSanctionsResponse">CorporateComplianceScreenSanctionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Audits

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>

Methods:

- <code title="post /corporate/compliance/audits">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditRequestParams">CorporateComplianceAuditRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/compliance/audits/{auditId}/report">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditService.GetReport">GetReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, auditID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Treasury

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryOptimizeLiquidityResponse">CorporateTreasuryOptimizeLiquidityResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetCashFlowForecastResponse">CorporateTreasuryGetCashFlowForecastResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetLiquidityPositionsResponse">CorporateTreasuryGetLiquidityPositionsResponse</a>

Methods:

- <code title="post /corporate/treasury/bulk-payouts">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryService.ExecuteBulkPayouts">ExecuteBulkPayouts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryExecuteBulkPayoutsParams">CorporateTreasuryExecuteBulkPayoutsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /corporate/treasury/liquidity/optimize">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryService.OptimizeLiquidity">OptimizeLiquidity</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryOptimizeLiquidityParams">CorporateTreasuryOptimizeLiquidityParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryOptimizeLiquidityResponse">CorporateTreasuryOptimizeLiquidityResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/treasury/cash-flow/forecast">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryService.GetCashFlowForecast">GetCashFlowForecast</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetCashFlowForecastParams">CorporateTreasuryGetCashFlowForecastParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetCashFlowForecastResponse">CorporateTreasuryGetCashFlowForecastResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/treasury/liquidity-positions">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryService.GetLiquidityPositions">GetLiquidityPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetLiquidityPositionsResponse">CorporateTreasuryGetLiquidityPositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Sweeping

Methods:

- <code title="post /corporate/treasury/sweeping/rules">client.Corporate.Treasury.Sweeping.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasurySweepingService.ConfigureRules">ConfigureRules</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasurySweepingConfigureRulesParams">CorporateTreasurySweepingConfigureRulesParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /corporate/treasury/sweeping/execute">client.Corporate.Treasury.Sweeping.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasurySweepingService.Execute">Execute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasurySweepingExecuteParams">CorporateTreasurySweepingExecuteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Pooling

Methods:

- <code title="post /corporate/treasury/liquidity/pooling">client.Corporate.Treasury.Pooling.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryPoolingService.Configure">Configure</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryPoolingConfigureParams">CorporateTreasuryPoolingConfigureParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListResponse">CorporateCardListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssuePhysicalResponse">CorporateCardIssuePhysicalResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssueVirtualResponse">CorporateCardIssueVirtualResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListTransactionsResponse">CorporateCardListTransactionsResponse</a>

Methods:

- <code title="get /corporate/cards">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListParams">CorporateCardListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListResponse">CorporateCardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/{cardId}/freeze">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.Freeze">Freeze</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardFreezeParams">CorporateCardFreezeParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /corporate/cards/physical">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.IssuePhysical">IssuePhysical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssuePhysicalParams">CorporateCardIssuePhysicalParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssuePhysicalResponse">CorporateCardIssuePhysicalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/virtual">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.IssueVirtual">IssueVirtual</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssueVirtualParams">CorporateCardIssueVirtualParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssueVirtualResponse">CorporateCardIssueVirtualResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/cards/{cardId}/transactions">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.ListTransactions">ListTransactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListTransactionsResponse">CorporateCardListTransactionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Controls

Methods:

- <code title="put /corporate/cards/{cardId}/controls">client.Corporate.Cards.Controls.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardControlService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardControlUpdateParams">CorporateCardControlUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Risk

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskGetExposureResponse">CorporateRiskGetExposureResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskRunStressTestResponse">CorporateRiskRunStressTestResponse</a>

Methods:

- <code title="get /corporate/risk/exposure">client.Corporate.Risk.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskService.GetExposure">GetExposure</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskGetExposureResponse">CorporateRiskGetExposureResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/risk/stress-test">client.Corporate.Risk.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskService.RunStressTest">RunStressTest</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskRunStressTestParams">CorporateRiskRunStressTestParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskRunStressTestResponse">CorporateRiskRunStressTestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Fraud

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudAnalyzeResponse">CorporateRiskFraudAnalyzeResponse</a>

Methods:

- <code title="post /corporate/risk/fraud/analyze">client.Corporate.Risk.Fraud.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudService.Analyze">Analyze</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudAnalyzeParams">CorporateRiskFraudAnalyzeParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudAnalyzeResponse">CorporateRiskFraudAnalyzeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleListResponse">CorporateRiskFraudRuleListResponse</a>

Methods:

- <code title="post /corporate/risk/fraud/rules">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleNewParams">CorporateRiskFraudRuleNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /corporate/risk/fraud/rules/{ruleId}">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleUpdateParams">CorporateRiskFraudRuleUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /corporate/risk/fraud/rules">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleListResponse">CorporateRiskFraudRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Governance

### Proposals

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalListResponse">CorporateGovernanceProposalListResponse</a>

Methods:

- <code title="post /corporate/governance/proposals">client.Corporate.Governance.Proposals.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalNewParams">CorporateGovernanceProposalNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /corporate/governance/proposals">client.Corporate.Governance.Proposals.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalListResponse">CorporateGovernanceProposalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/governance/proposals/{proposalId}/vote">client.Corporate.Governance.Proposals.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalService.Vote">Vote</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, proposalID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateGovernanceProposalVoteParams">CorporateGovernanceProposalVoteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Anomalies

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>

Methods:

- <code title="get /corporate/anomalies">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/anomalies/{anomalyId}/status">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyService.UpdateStatus">UpdateStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, anomalyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyUpdateStatusParams">CorporateAnomalyUpdateStatusParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Web3

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3GetNetworkStatusResponse">Web3GetNetworkStatusResponse</a>

Methods:

- <code title="get /web3/network/status">client.Web3.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3Service.GetNetworkStatus">GetNetworkStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3GetNetworkStatusResponse">Web3GetNetworkStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Wallets

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletNewResponse">Web3WalletNewResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletListResponse">Web3WalletListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletGetBalancesResponse">Web3WalletGetBalancesResponse</a>

Methods:

- <code title="post /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletNewParams">Web3WalletNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletNewResponse">Web3WalletNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletListResponse">Web3WalletListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /web3/wallets/connect">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletConnectParams">Web3WalletConnectParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /web3/wallets/{walletId}/balances">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.GetBalances">GetBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, walletID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletGetBalancesResponse">Web3WalletGetBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionSendResponse">Web3TransactionSendResponse</a>

Methods:

- <code title="post /web3/transactions/bridge">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionService.BridgeChain">BridgeChain</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionBridgeChainParams">Web3TransactionBridgeChainParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /web3/transactions/initiate">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionInitiateParams">Web3TransactionInitiateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /web3/transactions/send">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionService.Send">Send</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionSendParams">Web3TransactionSendParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionSendResponse">Web3TransactionSendResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /web3/transactions/swap">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionService.SwapTokens">SwapTokens</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionSwapTokensParams">Web3TransactionSwapTokensParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## NFTs

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTListResponse">Web3NFTListResponse</a>

Methods:

- <code title="get /web3/nfts">client.Web3.NFTs.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTListResponse">Web3NFTListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /web3/nfts/mint">client.Web3.NFTs.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTService.Mint">Mint</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTMintParams">Web3NFTMintParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Contracts

Methods:

- <code title="post /web3/contracts/deploy">client.Web3.Contracts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3ContractService.Deploy">Deploy</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3ContractDeployParams">Web3ContractDeployParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Payments

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentListResponse">PaymentListResponse</a>

Methods:

- <code title="get /payments/{paymentId}">client.Payments.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /payments">client.Payments.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentListResponse">PaymentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Domestic

Methods:

- <code title="post /payments/domestic/ach">client.Payments.Domestic.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentDomesticService.SendACH">SendACH</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentDomesticSendACHParams">PaymentDomesticSendACHParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /payments/domestic/rtp">client.Payments.Domestic.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentDomesticService.SendRtp">SendRtp</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentDomesticSendRtpParams">PaymentDomesticSendRtpParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /payments/domestic/wire">client.Payments.Domestic.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentDomesticService.SendWire">SendWire</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentDomesticSendWireParams">PaymentDomesticSendWireParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## International

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalGetStatusResponse">PaymentInternationalGetStatusResponse</a>

Methods:

- <code title="get /payments/international/{paymentId}/status">client.Payments.International.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalGetStatusResponse">PaymentInternationalGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /payments/international/sepa">client.Payments.International.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalService.SendSepa">SendSepa</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalSendSepaParams">PaymentInternationalSendSepaParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /payments/international/swift">client.Payments.International.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalService.SendSwift">SendSwift</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalSendSwiftParams">PaymentInternationalSendSwiftParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Fx

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>

Methods:

- <code title="post /payments/fx/deals">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxService.BookDeal">BookDeal</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxBookDealParams">PaymentFxBookDealParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /payments/fx/convert">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxService.ConvertCurrency">ConvertCurrency</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxConvertCurrencyParams">PaymentFxConvertCurrencyParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /payments/fx/rates">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxService.GetRates">GetRates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxGetRatesParams">PaymentFxGetRatesParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sustainability

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityGetCarbonFootprintResponse">SustainabilityGetCarbonFootprintResponse</a>

Methods:

- <code title="get /sustainability/carbon-footprint">client.Sustainability.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityService.GetCarbonFootprint">GetCarbonFootprint</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityGetCarbonFootprintResponse">SustainabilityGetCarbonFootprintResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Offsets

Methods:

- <code title="post /sustainability/offsets/purchase">client.Sustainability.Offsets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityOffsetService.Purchase">Purchase</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityOffsetPurchaseParams">SustainabilityOffsetPurchaseParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /sustainability/offsets/retire">client.Sustainability.Offsets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityOffsetService.Retire">Retire</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityOffsetRetireParams">SustainabilityOffsetRetireParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Impact

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactListGreenProjectsResponse">SustainabilityImpactListGreenProjectsResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactGetPortfolioAnalysisResponse">SustainabilityImpactGetPortfolioAnalysisResponse</a>

Methods:

- <code title="get /sustainability/impact/projects">client.Sustainability.Impact.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactService.ListGreenProjects">ListGreenProjects</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactListGreenProjectsParams">SustainabilityImpactListGreenProjectsParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactListGreenProjectsResponse">SustainabilityImpactListGreenProjectsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sustainability/impact/portfolio">client.Sustainability.Impact.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactService.GetPortfolioAnalysis">GetPortfolioAnalysis</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityImpactGetPortfolioAnalysisResponse">SustainabilityImpactGetPortfolioAnalysisResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Marketplace

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceListProductsResponse">MarketplaceListProductsResponse</a>

Methods:

- <code title="get /marketplace/products">client.Marketplace.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceService.ListProducts">ListProducts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceListProductsResponse">MarketplaceListProductsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Offers

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceOfferListResponse">MarketplaceOfferListResponse</a>

Methods:

- <code title="get /marketplace/offers">client.Marketplace.Offers.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceOfferService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceOfferListResponse">MarketplaceOfferListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /marketplace/offers/{offerId}/redeem">client.Marketplace.Offers.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#MarketplaceOfferService.Redeem">Redeem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, offerID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Lending

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingSubmitApplicationResponse">LendingSubmitApplicationResponse</a>

Methods:

- <code title="post /lending/applications">client.Lending.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingService.SubmitApplication">SubmitApplication</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingSubmitApplicationParams">LendingSubmitApplicationParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingSubmitApplicationResponse">LendingSubmitApplicationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Applications

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingApplicationGetStatusResponse">LendingApplicationGetStatusResponse</a>

Methods:

- <code title="get /lending/applications/{appId}/status">client.Lending.Applications.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingApplicationService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, appID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingApplicationGetStatusResponse">LendingApplicationGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Decisions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingDecisionGetRationaleResponse">LendingDecisionGetRationaleResponse</a>

Methods:

- <code title="get /lending/decisions/{decisionId}/rationale">client.Lending.Decisions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingDecisionService.GetRationale">GetRationale</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, decisionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#LendingDecisionGetRationaleResponse">LendingDecisionGetRationaleResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Investments

## Portfolios

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>

Methods:

- <code title="post /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioNewParams">InvestmentPortfolioNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioUpdateParams">InvestmentPortfolioUpdateParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioListParams">InvestmentPortfolioListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /investments/portfolios/{portfolioId}/rebalance">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.Rebalance">Rebalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioRebalanceParams">InvestmentPortfolioRebalanceParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>

Methods:

- <code title="get /investments/assets/search">client.Investments.Assets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetSearchParams">InvestmentAssetSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Performance

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPerformanceGetHistoricalResponse">InvestmentPerformanceGetHistoricalResponse</a>

Methods:

- <code title="get /investments/performance/historical">client.Investments.Performance.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPerformanceService.GetHistorical">GetHistorical</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPerformanceGetHistoricalParams">InvestmentPerformanceGetHistoricalParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPerformanceGetHistoricalResponse">InvestmentPerformanceGetHistoricalResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# System

## Status

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemStatusGetResponse">SystemStatusGetResponse</a>

Methods:

- <code title="get /system/status">client.System.Status.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemStatusService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemStatusGetResponse">SystemStatusGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemWebhookListResponse">SystemWebhookListResponse</a>

Methods:

- <code title="post /system/webhooks">client.System.Webhooks.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemWebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemWebhookNewParams">SystemWebhookNewParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /system/webhooks">client.System.Webhooks.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemWebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemWebhookListResponse">SystemWebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /system/webhooks/{webhookId}">client.System.Webhooks.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemWebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## AuditLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemAuditLogListResponse">SystemAuditLogListResponse</a>

Methods:

- <code title="get /system/audit-logs">client.System.AuditLogs.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemAuditLogService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemAuditLogListParams">SystemAuditLogListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemAuditLogListResponse">SystemAuditLogListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Sandbox

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemSandboxForceErrorResponse">SystemSandboxForceErrorResponse</a>

Methods:

- <code title="post /system/sandbox/simulate-error">client.System.Sandbox.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemSandboxService.ForceError">ForceError</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemSandboxForceErrorParams">SystemSandboxForceErrorParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemSandboxForceErrorResponse">SystemSandboxForceErrorResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /system/sandbox/reset">client.System.Sandbox.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemSandboxService.Reset">Reset</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Verification

Methods:

- <code title="post /system/verification/biometric-comparison">client.System.Verification.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemVerificationService.CompareBiometrics">CompareBiometrics</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemVerificationCompareBiometricsParams">SystemVerificationCompareBiometricsParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /system/verification/document">client.System.Verification.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemVerificationService.VerifyDocument">VerifyDocument</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemVerificationVerifyDocumentParams">SystemVerificationVerifyDocumentParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Notifications

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemNotificationListTemplatesResponse">SystemNotificationListTemplatesResponse</a>

Methods:

- <code title="get /system/notifications/templates">client.System.Notifications.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemNotificationService.ListTemplates">ListTemplates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemNotificationListTemplatesResponse">SystemNotificationListTemplatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /system/notifications/push">client.System.Notifications.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemNotificationService.SendPush">SendPush</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SystemNotificationSendPushParams">SystemNotificationSendPushParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
