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

### Simulations

## Incubator

### Pitch

## Ads

# Corporate

## SanctionScreening

## Compliance

### Audits

## Treasury

## Cards

### Controls

## Risk

### Fraud

#### Rules

## Anomalies

# Web3

## Wallets

## Transactions

## NFTs

# Payments

## International

## Fx

# Sustainability

## Investments
