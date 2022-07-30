package service

import (
	"errors"
	"github.com/cu2koo/go-ra-data/model"
	"github.com/cu2koo/go-ra-data/model/data"
	"github.com/cu2koo/go-ra-data/model/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserService struct {
	data *data.User
}

func NewUserService(data data.User) (*UserService, error) {
	if data == nil {
		return nil, errors.New("no mariadb added to service")
	}
	us := &UserService{
		data: &data,
	}

	return us, nil
}

func (us *UserService) CreateUser(c *gin.Context) {
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.CreateUser(&user)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetUser(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetUser(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) GetUsers(c *gin.Context) {
	userDatabase := *us.data
	result, err := userDatabase.GetUsers()
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) UpdateUser(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var user model.User
	err = c.ShouldBind(&user)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateUser(&userId, &user)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) DeleteUser(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.DeleteUser(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetFinancialInformation(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetFinancialInformation(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) UpdateFinancialInformation(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var fi model.FinancialInformation
	err = c.ShouldBind(&fi)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateFinancialInformation(&userId, &fi)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetGoal(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetGoal(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) UpdateGoal(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var goal model.Goal
	err = c.ShouldBind(&goal)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateGoal(&userId, &goal)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetPeriodicalDeposits(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetPeriodicalDeposits(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) UpdatePeriodicalDeposits(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var pd model.PeriodicalDeposits
	err = c.ShouldBind(&pd)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdatePeriodicalDeposits(&userId, &pd)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetRiskAssessment(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetRiskAssessment(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) UpdateRiskAssessment(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var ra model.RiskAssessments
	err = c.ShouldBind(&ra)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateRiskAssessment(&userId, &ra)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) AddExperience(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var experience model.Experience
	err = c.ShouldBind(&experience)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.AddExperience(&userId, &experience)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) UpdateExperience(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var experience model.Experience
	err = c.ShouldBind(&experience)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateExperience(&id, &experience)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) DeleteExperience(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.DeleteExperience(&id)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetExperiences(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetExperiences(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) AddDeposit(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var deposit model.Deposit
	err = c.ShouldBind(&deposit)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.AddDeposit(&userId, &deposit)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) UpdateDeposit(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var deposit model.Deposit
	err = c.ShouldBind(&deposit)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateDeposit(&id, &deposit)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) DeleteDeposit(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.DeleteDeposit(&id)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetDeposits(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetDeposits(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (us *UserService) AddSecurity(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var security model.Security
	err = c.ShouldBind(&security)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.AddSecurity(&userId, &security)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) UpdateSecurity(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}
	var security model.Security
	err = c.ShouldBind(&security)
	if err != nil {
		response.FormBindingFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.UpdateSecurity(&id, &security)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) DeleteSecurity(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	err = userDatabase.DeleteSecurity(&id)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	response.ActionSuccessful(c)
}

func (us *UserService) GetSecurities(c *gin.Context) {
	userIdString := c.Params.ByName("userId")
	userId, err := strconv.Atoi(userIdString)
	if err != nil {
		response.IdConversionFailed(c)
		return
	}

	userDatabase := *us.data
	result, err := userDatabase.GetSecurities(&userId)
	if err != nil {
		response.ActionFailed(c)
		return
	}

	c.JSON(http.StatusOK, result)
}
