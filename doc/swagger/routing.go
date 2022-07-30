package swagger

// swagger:route POST /user/ User createUser
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/ User getUsers
//
//	Responses:
//		200: userResponses
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId} User getUser
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: userResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId} User updateUser
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route DELETE /user/{userId} User deleteUser
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/fi User getFinancialInformation
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: financialInformationResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/fi User updateFinancialInformation
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/goal User getGoal
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: goalResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/goal User updateGoal
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/pd User getPeriodicalDeposits
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: periodicalDepositsResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/pd User updatePeriodicalDeposits
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/ra User getRiskAssessment
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: riskAssessmentResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/ra User updateRiskAssessment
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route POST /user/{userId}/experience User addExperience
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/experience User getExperiences
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: experienceResponses
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/experience/{id} User updateExperience
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//		+ 	name: id
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route DELETE /user/{userId}/experience/{id} User deleteExperience
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//		+ 	name: id
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route POST /user/{userId}/deposit User addDeposit
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/deposit User getDeposits
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: depositResponses
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/deposit/{id} User updateDeposit
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//		+ 	name: id
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route DELETE /user/{userId}/deposit/{id} User deleteDeposit
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//		+ 	name: id
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route POST /user/{userId}/security User addSecurity
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /user/{userId}/security User getSecurity
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: securityResponses
//		400: errorResponse
//		501: errorResponse

// swagger:route PUT /user/{userId}/security/{id} User updateSecurity
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//		+ 	name: id
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route DELETE /user/{userId}/security/{id} User deleteSecurity
//
// 	Parameters:
//		+ 	name: userId
//			in: path
//			required: true
//			type: integer
//		+ 	name: id
//			in: path
//			required: true
//			type: integer
//
//	Responses:
//		200: successResponse
//		400: errorResponse
//		501: errorResponse

// swagger:route GET /market/prices/ Market getPrices
//
// 	Parameters:
//		+ 	name: symbol
//			in: query
//			required: true
//			type: string
//		+ 	name: period
//			in: query
//			type: string
//
//	Responses:
//		200: priceDataResponses
//		400: errorResponse
//		501: errorResponse
