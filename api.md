# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserLoginResponse">UserLoginResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserRegisterResponse">UserRegisterResponse</a>

Methods:

- <code title="post /users/login">client.Users.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserLoginParams">UserLoginParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserLoginResponse">UserLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/register">client.Users.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserRegisterParams">UserRegisterParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserRegisterResponse">UserRegisterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Me

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeGetResponse">UserMeGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeUpdateResponse">UserMeUpdateResponse</a>

Methods:

- <code title="get /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeGetResponse">UserMeGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeUpdateParams">UserMeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeUpdateResponse">UserMeUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Preferences

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceGetResponse">UserMePreferenceGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceUpdateResponse">UserMePreferenceUpdateResponse</a>

Methods:

- <code title="get /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceGetResponse">UserMePreferenceGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceUpdateParams">UserMePreferenceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMePreferenceUpdateResponse">UserMePreferenceUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceListResponse">UserMeDeviceListResponse</a>

Methods:

- <code title="get /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceListParams">UserMeDeviceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeDeviceListResponse">UserMeDeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Biometrics

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricGetStatusResponse">UserMeBiometricGetStatusResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>

Methods:

- <code title="get /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricGetStatusResponse">UserMeBiometricGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/me/biometrics/verify">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricVerifyParams">UserMeBiometricVerifyParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountGetResponse">AccountGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountListResponse">AccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountLinkResponse">AccountLinkResponse</a>

Methods:

- <code title="get /accounts/{accountId}/details">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountGetResponse">AccountGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/me">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountListParams">AccountListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountListResponse">AccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /accounts/link">client.Accounts.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountService.Link">Link</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountLinkParams">AccountLinkParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountLinkResponse">AccountLinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListPendingResponse">AccountTransactionListPendingResponse</a>

Methods:

- <code title="get /accounts/{accountId}/transactions/pending">client.Accounts.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionService.ListPending">ListPending</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListPendingParams">AccountTransactionListPendingParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountTransactionListPendingResponse">AccountTransactionListPendingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Statements

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementListResponse">AccountStatementListResponse</a>

Methods:

- <code title="get /accounts/{accountId}/statements">client.Accounts.Statements.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementListParams">AccountStatementListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountStatementListResponse">AccountStatementListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Overdraft

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftUpdateResponse">AccountOverdraftUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftGetResponse">AccountOverdraftGetResponse</a>

Methods:

- <code title="put /accounts/{accountId}/overdraft-settings">client.Accounts.Overdraft.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftUpdateParams">AccountOverdraftUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftUpdateResponse">AccountOverdraftUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/overdraft-settings">client.Accounts.Overdraft.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AccountOverdraftGetResponse">AccountOverdraftGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionGetResponse">TransactionGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionListResponse">TransactionListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionAddNotesResponse">TransactionAddNotesResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionCategorizeResponse">TransactionCategorizeResponse</a>

Methods:

- <code title="get /transactions/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionGetResponse">TransactionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionListParams">TransactionListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionListResponse">TransactionListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/notes">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.AddNotes">AddNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionAddNotesParams">TransactionAddNotesParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionAddNotesResponse">TransactionAddNotesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/categorize">client.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionService.Categorize">Categorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionCategorizeParams">TransactionCategorizeParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionCategorizeResponse">TransactionCategorizeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recurring

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringListResponse">TransactionRecurringListResponse</a>

Methods:

- <code title="get /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringListParams">TransactionRecurringListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionRecurringListResponse">TransactionRecurringListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Insights

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightGetTrendsResponse">TransactionInsightGetTrendsResponse</a>

Methods:

- <code title="get /transactions/insights/spending-trends">client.Transactions.Insights.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightService.GetTrends">GetTrends</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#TransactionInsightGetTrendsResponse">TransactionInsightGetTrendsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Budgets

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetGetResponse">BudgetGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetUpdateResponse">BudgetUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetListResponse">BudgetListResponse</a>

Methods:

- <code title="get /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetGetResponse">BudgetGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetUpdateParams">BudgetUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetUpdateResponse">BudgetUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budgets">client.Budgets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetListParams">BudgetListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#BudgetListResponse">BudgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Investments

## Portfolios

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioGetResponse">InvestmentPortfolioGetResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioUpdateResponse">InvestmentPortfolioUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>

Methods:

- <code title="get /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioGetResponse">InvestmentPortfolioGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioUpdateParams">InvestmentPortfolioUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioUpdateResponse">InvestmentPortfolioUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioListParams">InvestmentPortfolioListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /investments/portfolios/{portfolioId}/rebalance">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioService.Rebalance">Rebalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioRebalanceParams">InvestmentPortfolioRebalanceParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>

Methods:

- <code title="get /investments/assets/search">client.Investments.Assets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetSearchParams">InvestmentAssetSearchParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

## Advisor

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorChatResponse">AIAdvisorChatResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorHistoryResponse">AIAdvisorHistoryResponse</a>

Methods:

- <code title="post /ai/advisor/chat">client.AI.Advisor.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorService.Chat">Chat</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorChatParams">AIAdvisorChatParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorChatResponse">AIAdvisorChatResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/advisor/chat/history">client.AI.Advisor.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorService.History">History</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorHistoryParams">AIAdvisorHistoryParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorHistoryResponse">AIAdvisorHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolListResponse">AIAdvisorToolListResponse</a>

Methods:

- <code title="get /ai/advisor/tools">client.AI.Advisor.Tools.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolListParams">AIAdvisorToolListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdvisorToolListResponse">AIAdvisorToolListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Oracle

### Simulate

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunAdvancedResponse">AIOracleSimulateRunAdvancedResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunStandardResponse">AIOracleSimulateRunStandardResponse</a>

Methods:

- <code title="post /ai/oracle/simulate/advanced">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateService.RunAdvanced">RunAdvanced</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunAdvancedParams">AIOracleSimulateRunAdvancedParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunAdvancedResponse">AIOracleSimulateRunAdvancedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/oracle/simulate">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateService.RunStandard">RunStandard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunStandardParams">AIOracleSimulateRunStandardParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulateRunStandardResponse">AIOracleSimulateRunStandardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Simulations

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>

Methods:

- <code title="get /ai/oracle/simulations/{simulationId}">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, simulationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*AIOracleSimulationGetResponse, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/oracle/simulations">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationListParams">AIOracleSimulationListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Incubator

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorGeneratePitchResponse">AIIncubatorGeneratePitchResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>

Methods:

- <code title="post /ai/incubator/pitch">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorService.GeneratePitch">GeneratePitch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorGeneratePitchParams">AIIncubatorGeneratePitchParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorGeneratePitchResponse">AIIncubatorGeneratePitchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/incubator/pitches">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorService.ListPitches">ListPitches</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorListPitchesParams">AIIncubatorListPitchesParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Pitch

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchSubmitFeedbackResponse">AIIncubatorPitchSubmitFeedbackResponse</a>

Methods:

- <code title="get /ai/incubator/pitch/{pitchId}/details">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /ai/incubator/pitch/{pitchId}/feedback">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchService.SubmitFeedback">SubmitFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchSubmitFeedbackParams">AIIncubatorPitchSubmitFeedbackParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIIncubatorPitchSubmitFeedbackResponse">AIIncubatorPitchSubmitFeedbackResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ads

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdListResponse">AIAdListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateResponse">AIAdGenerateResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGetOperationResponse">AIAdGetOperationResponse</a>

Methods:

- <code title="get /ai/ads">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdListParams">AIAdListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdListResponse">AIAdListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/ads/generate">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.Generate">Generate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateParams">AIAdGenerateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGenerateResponse">AIAdGenerateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/ads/operations/{operationId}">client.AI.Ads.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdService.GetOperation">GetOperation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#AIAdGetOperationResponse">AIAdGetOperationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Corporate

## SanctionScreening

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateSanctionScreeningScreenResponse">CorporateSanctionScreeningScreenResponse</a>

Methods:

- <code title="post /corporate/sanction-screening">client.Corporate.SanctionScreening.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateSanctionScreeningService.Screen">Screen</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateSanctionScreeningScreenParams">CorporateSanctionScreeningScreenParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateSanctionScreeningScreenResponse">CorporateSanctionScreeningScreenResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Compliance

### Audits

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>

Methods:

- <code title="post /corporate/compliance/audits">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditRequestParams">CorporateComplianceAuditRequestParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/compliance/audits/{auditId}/report">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditService.GetReport">GetReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, auditID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Treasury

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryForecastCashFlowResponse">CorporateTreasuryForecastCashFlowResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetLiquidityPositionsResponse">CorporateTreasuryGetLiquidityPositionsResponse</a>

Methods:

- <code title="get /corporate/treasury/cash-flow/forecast">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryService.ForecastCashFlow">ForecastCashFlow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryForecastCashFlowParams">CorporateTreasuryForecastCashFlowParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryForecastCashFlowResponse">CorporateTreasuryForecastCashFlowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/treasury/liquidity-positions">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryService.GetLiquidityPositions">GetLiquidityPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateTreasuryGetLiquidityPositionsResponse">CorporateTreasuryGetLiquidityPositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cards

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListResponse">CorporateCardListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardFreezeResponse">CorporateCardFreezeResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssueVirtualResponse">CorporateCardIssueVirtualResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListTransactionsResponse">CorporateCardListTransactionsResponse</a>

Methods:

- <code title="get /corporate/cards">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListParams">CorporateCardListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListResponse">CorporateCardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/{cardId}/freeze">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.Freeze">Freeze</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardFreezeParams">CorporateCardFreezeParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardFreezeResponse">CorporateCardFreezeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/virtual">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.IssueVirtual">IssueVirtual</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssueVirtualParams">CorporateCardIssueVirtualParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardIssueVirtualResponse">CorporateCardIssueVirtualResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/cards/{cardId}/transactions">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardService.ListTransactions">ListTransactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListTransactionsParams">CorporateCardListTransactionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardListTransactionsResponse">CorporateCardListTransactionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Controls

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardControlUpdateResponse">CorporateCardControlUpdateResponse</a>

Methods:

- <code title="put /corporate/cards/{cardId}/controls">client.Corporate.Cards.Controls.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardControlService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardControlUpdateParams">CorporateCardControlUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateCardControlUpdateResponse">CorporateCardControlUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Risk

### Fraud

#### Rules

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleUpdateResponse">CorporateRiskFraudRuleUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleListResponse">CorporateRiskFraudRuleListResponse</a>

Methods:

- <code title="put /corporate/risk/fraud/rules/{ruleId}">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleUpdateParams">CorporateRiskFraudRuleUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleUpdateResponse">CorporateRiskFraudRuleUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/risk/fraud/rules">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleListParams">CorporateRiskFraudRuleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateRiskFraudRuleListResponse">CorporateRiskFraudRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Anomalies

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyUpdateStatusResponse">CorporateAnomalyUpdateStatusResponse</a>

Methods:

- <code title="get /corporate/anomalies">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyListParams">CorporateAnomalyListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/anomalies/{anomalyId}/status">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyService.UpdateStatus">UpdateStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, anomalyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyUpdateStatusParams">CorporateAnomalyUpdateStatusParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#CorporateAnomalyUpdateStatusResponse">CorporateAnomalyUpdateStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Web3

## Wallets

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletListResponse">Web3WalletListResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletConnectResponse">Web3WalletConnectResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletGetBalanceResponse">Web3WalletGetBalanceResponse</a>

Methods:

- <code title="get /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletListParams">Web3WalletListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletListResponse">Web3WalletListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletConnectParams">Web3WalletConnectParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletConnectResponse">Web3WalletConnectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /web3/wallets/{walletId}/balances">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletService.GetBalance">GetBalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, walletID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletGetBalanceParams">Web3WalletGetBalanceParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3WalletGetBalanceResponse">Web3WalletGetBalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionInitiateResponse">Web3TransactionInitiateResponse</a>

Methods:

- <code title="post /web3/transactions/initiate">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionInitiateParams">Web3TransactionInitiateParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3TransactionInitiateResponse">Web3TransactionInitiateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## NFTs

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTListResponse">Web3NFTListResponse</a>

Methods:

- <code title="get /web3/nfts">client.Web3.NFTs.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTListParams">Web3NFTListParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#Web3NFTListResponse">Web3NFTListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

## International

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalGetStatusResponse">PaymentInternationalGetStatusResponse</a>

Methods:

- <code title="get /payments/international/{paymentId}/status">client.Payments.International.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentInternationalGetStatusResponse">PaymentInternationalGetStatusResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Fx

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxConvertResponse">PaymentFxConvertResponse</a>
- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>

Methods:

- <code title="post /payments/fx/convert">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxService.Convert">Convert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxConvertParams">PaymentFxConvertParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxConvertResponse">PaymentFxConvertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/fx/rates">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxService.GetRates">GetRates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxGetRatesParams">PaymentFxGetRatesParams</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sustainability

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityGetFootprintResponse">SustainabilityGetFootprintResponse</a>

Methods:

- <code title="get /sustainability/carbon-footprint">client.Sustainability.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityService.GetFootprint">GetFootprint</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityGetFootprintResponse">SustainabilityGetFootprintResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Investments

Response Types:

- <a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityInvestmentAnalyzeImpactResponse">SustainabilityInvestmentAnalyzeImpactResponse</a>

Methods:

- <code title="get /sustainability/investments/impact">client.Sustainability.Investments.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityInvestmentService.AnalyzeImpact">AnalyzeImpact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go">githubcomjocall3go</a>.<a href="https://pkg.go.dev/github.com/diplomat-bit/jocall3-go#SustainabilityInvestmentAnalyzeImpactResponse">SustainabilityInvestmentAnalyzeImpactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
