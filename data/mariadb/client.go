package mariadb

import (
	"database/sql"
	"errors"
	"github.com/cu2koo/go-ra-data/model"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

const driverName = "mysql"

type MariaDB struct {
	cfg *mysql.Config
}

func New() (*MariaDB, error) {
	vp := viper.New()
	vp.AddConfigPath("./data/mariadb")
	vp.SetConfigName("config")
	vp.SetConfigType("json")

	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &MariaDB{
		cfg: &mysql.Config{
			User:                 vp.GetString("user"),
			Passwd:               vp.GetString("passwd"),
			Addr:                 vp.GetString("addr"),
			DBName:               vp.GetString("dbName"),
			AllowNativePasswords: true,
		},
	}, nil
}

func (db *MariaDB) CreateUser(user *model.User) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	insert, err := udb.Query("INSERT INTO users (mail, password) VALUES (?, ?);", *user.Mail, *user.Password)
	if err != nil {
		return err
	}
	insert.Close()

	rows, err := udb.Query("SELECT * FROM users WHERE mail = ?;", *user.Mail)
	if err != nil {
		return err
	}
	defer rows.Close()

	var userId int
	for rows.Next() {
		var (
			id       int
			mail     string
			password string
		)
		err = rows.Scan(&id, &mail, &password)
		userId = id
	}

	fi, err := udb.Query("INSERT INTO financialInformation (userId) VALUES (?);", userId)
	if err != nil {
		return err
	}
	fi.Close()

	goal, err := udb.Query("INSERT INTO goals (userId) VALUES (?);", userId)
	if err != nil {
		return err
	}
	goal.Close()

	pd, err := udb.Query("INSERT INTO periodicalDeposits (userId) VALUES (?);", userId)
	if err != nil {
		return err
	}
	pd.Close()

	ra, err := udb.Query("INSERT INTO riskAssessments (userId) VALUES (?);", userId)
	if err != nil {
		return err
	}
	ra.Close()

	return nil
}

func (db *MariaDB) GetUser(userId *int) (*model.User, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM users WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.User
	for rows.Next() {
		var (
			id       int
			mail     string
			password string
		)
		err = rows.Scan(&id, &mail, &password)
		result.UserId = &id
		result.Mail = &mail
		result.Password = &password
	}
	if result.UserId == nil {
		return nil, errors.New("user do not exist")
	}

	return &result, nil
}

func (db *MariaDB) GetUsers() (*model.Users, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.Users
	for rows.Next() {
		var (
			id       int
			mail     string
			password string
		)
		err = rows.Scan(&id, &mail, &password)
		user := &model.User{
			UserId:   &id,
			Mail:     &mail,
			Password: &password,
		}
		result = append(result, user)
	}

	if len(result) <= 0 {
		return &model.Users{}, nil
	}

	return &result, nil
}

func (db *MariaDB) UpdateUser(userId *int, user *model.User) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE users SET mail = ?, password = ? WHERE userId = ?;", *user.Mail, *user.Password, *userId)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) DeleteUser(userId *int) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	fi, err := udb.Query("DELETE FROM financialInformation WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	fi.Close()

	goal, err := udb.Query("DELETE FROM goals WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	goal.Close()

	pi, err := udb.Query("DELETE FROM periodicalDeposits WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	pi.Close()

	ra, err := udb.Query("DELETE FROM riskAssessments WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	ra.Close()

	exp, err := udb.Query("DELETE FROM experiences WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	exp.Close()

	d, err := udb.Query("DELETE FROM deposits WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	d.Close()

	s, err := udb.Query("DELETE FROM securities WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	s.Close()

	rows, err := udb.Query("DELETE FROM users WHERE userId = ?;", *userId)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetFinancialInformation(userId *int) (*model.FinancialInformation, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM financialInformation WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.FinancialInformation
	var validationId *int
	for rows.Next() {
		var (
			id             int
			liquidAssets   float64
			illiquidAssets float64
			income         float64
			expenses       float64
			liabilities    float64
			resultUserId   int
		)
		err = rows.Scan(&id, &liquidAssets, &illiquidAssets, &income, &expenses, &liabilities, &resultUserId)
		result.Assets = &model.Assets{}
		result.Assets.Liquid = &liquidAssets
		result.Assets.Illiquid = &illiquidAssets
		result.Income = &income
		result.Expenses = &expenses
		result.Liabilities = &liabilities
		validationId = &resultUserId
	}
	if validationId == nil {
		return nil, errors.New("data do not exist")
	}

	return &result, nil
}

func (db *MariaDB) UpdateFinancialInformation(userId *int, financialInformation *model.FinancialInformation) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE financialInformation SET liquidAssets = ?, illiquidAssets = ?, income = ?, expenses = ?, liabilities = ? WHERE userId = ?;", *financialInformation.Assets.Liquid, *financialInformation.Assets.Illiquid, *financialInformation.Income, *financialInformation.Expenses, *financialInformation.Liabilities, *userId)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetGoal(userId *int) (*model.Goal, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM goals WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.Goal
	var validationId *int
	for rows.Next() {
		var (
			id           int
			startDate    string
			endDate      string
			goal         string
			resultUserId int
		)
		err = rows.Scan(&id, &startDate, &endDate, &goal, &resultUserId)
		result.StartDate = &startDate
		result.EndDate = &endDate
		result.Goal = &goal
		validationId = &resultUserId
	}
	if validationId == nil {
		return nil, errors.New("data do not exist")
	}

	return &result, nil
}

func (db *MariaDB) UpdateGoal(userId *int, goal *model.Goal) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE goals SET startDate = ?, endDate = ?, goal = ? WHERE userId = ?;", *goal.StartDate, *goal.EndDate, *goal.Goal, *userId)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetPeriodicalDeposits(userId *int) (*model.PeriodicalDeposits, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM periodicalDeposits WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.PeriodicalDeposits
	var validationId *int
	for rows.Next() {
		var (
			id           int
			amount       float64
			currency     string
			period       string
			resultUserId int
		)
		err = rows.Scan(&id, &amount, &currency, &period, &resultUserId)
		result.Amount = &amount
		result.Currency = &currency
		result.Period = &period
		validationId = &resultUserId
	}
	if validationId == nil {
		return nil, errors.New("data do not exist")
	}

	return &result, nil
}

func (db *MariaDB) UpdatePeriodicalDeposits(userId *int, periodicalDeposits *model.PeriodicalDeposits) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE periodicalDeposits SET amount = ?, currency = ?, period = ? WHERE userId = ?;", *periodicalDeposits.Amount, *periodicalDeposits.Currency, *periodicalDeposits.Period, *userId)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetRiskAssessment(userId *int) (*model.RiskAssessments, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM riskAssessments WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.RiskAssessments
	var validationId *int
	for rows.Next() {
		var (
			id            int
			riskAppetite  int
			maxVolatility float64
			resultUserId  int
		)
		err = rows.Scan(&id, &riskAppetite, &maxVolatility, &resultUserId)
		result.RiskAppetite = &riskAppetite
		result.MaxVolatility = &maxVolatility
		validationId = &resultUserId
	}
	if validationId == nil {
		return nil, errors.New("data do not exist")
	}

	return &result, nil
}

func (db *MariaDB) UpdateRiskAssessment(userId *int, riskAssessments *model.RiskAssessments) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE riskAssessments SET riskAppetite = ?, maxVolatility = ? WHERE userId = ?;", *riskAssessments.RiskAppetite, *riskAssessments.MaxVolatility, *userId)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) AddExperience(userId *int, experience *model.Experience) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	insert, err := udb.Query("INSERT INTO experiences (type, transactions, timePeriod, userId) VALUES (?, ?, ?, ?);", *experience.Type, *experience.Transactions, *experience.TimePeriod, *userId)
	if err != nil {
		return err
	}
	insert.Close()

	return nil
}

func (db *MariaDB) UpdateExperience(id *int, experience *model.Experience) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE experiences SET type = ?, transactions = ?, timePeriod = ? WHERE id = ?;", *experience.Type, *experience.Transactions, *experience.TimePeriod, *id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) DeleteExperience(id *int) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("DELETE FROM experiences WHERE id = ?;", *id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetExperiences(userId *int) (*model.Experiences, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM experiences WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.Experiences
	for rows.Next() {
		var experience model.Experience
		var (
			id           int
			typ          string
			transactions int
			timePeriod   string
			resultUserId int
		)
		err = rows.Scan(&id, &typ, &transactions, &timePeriod, &resultUserId)
		experience.Id = &id
		experience.Type = &typ
		experience.Transactions = &transactions
		experience.TimePeriod = &timePeriod
		result = append(result, &experience)
	}

	if len(result) <= 0 {
		return &model.Experiences{}, nil
	}

	return &result, nil
}

func (db *MariaDB) AddDeposit(userId *int, deposit *model.Deposit) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	insert, err := udb.Query("INSERT INTO deposits (datetime, amount, currency, target, userId) VALUES (?, ?, ?, ?, ?);", *deposit.Date, *deposit.Amount, *deposit.Currency, *deposit.Target, *userId)
	if err != nil {
		return err
	}
	insert.Close()

	return nil
}

func (db *MariaDB) UpdateDeposit(id *int, deposit *model.Deposit) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE deposits SET datetime = ?, amount = ?, currency = ?, target = ? WHERE id = ?;", *deposit.Date, *deposit.Amount, *deposit.Currency, *deposit.Target, *id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) DeleteDeposit(id *int) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("DELETE FROM deposits WHERE id = ?;", *id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetDeposits(userId *int) (*model.Deposits, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM deposits WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.Deposits
	for rows.Next() {
		var deposit model.Deposit
		var (
			id           int
			datetime     string
			amount       float64
			currency     string
			target       string
			resultUserId int
		)
		err = rows.Scan(&id, &datetime, &amount, &currency, &target, &resultUserId)
		deposit.Id = &id
		deposit.Date = &datetime
		deposit.Amount = &amount
		deposit.Currency = &currency
		deposit.Target = &target
		result = append(result, &deposit)
	}

	if len(result) <= 0 {
		return &model.Deposits{}, nil
	}

	return &result, nil
}

func (db *MariaDB) AddSecurity(userId *int, security *model.Security) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	insert, err := udb.Query("INSERT INTO securities (symbol, shares, averageCostPerShare, userId) VALUES (?, ?, ?, ?);", *security.Symbol, *security.Shares, *security.AverageCostPerShare, *userId)
	if err != nil {
		return err
	}
	insert.Close()

	return nil
}

func (db *MariaDB) UpdateSecurity(id *int, security *model.Security) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("UPDATE securities SET symbol = ?, shares = ?, averageCostPerShare = ? WHERE id = ?;", *security.Symbol, *security.Shares, *security.AverageCostPerShare, *id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) DeleteSecurity(id *int) error {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer udb.Close()

	rows, err := udb.Query("DELETE FROM securities WHERE id = ?;", *id)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *MariaDB) GetSecurities(userId *int) (*model.Securities, error) {
	udb, err := sql.Open(driverName, db.cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	defer udb.Close()

	rows, err := udb.Query("SELECT * FROM securities WHERE userId = ?;", *userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result model.Securities
	for rows.Next() {
		var security model.Security
		var (
			id                  int
			symbol              string
			shares              float64
			averageCostPerShare float64
			resultUserId        int
		)
		err = rows.Scan(&id, &symbol, &shares, &averageCostPerShare, &resultUserId)
		security.Id = &id
		security.Symbol = &symbol
		security.Shares = &shares
		security.AverageCostPerShare = &averageCostPerShare
		result = append(result, &security)
	}

	if len(result) <= 0 {
		return &model.Securities{}, nil
	}

	return &result, nil
}
