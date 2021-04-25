package models

import (
	"database/sql"
	"fmt"
	"time"
)

// GetGenTokenBalanceByID ...
func (tkn *TokenBalanceModel) GetGenTokenBalanceByID(ID string) error {
	sp, _ := RetriveSP("GET_Balance_BY_ID")
	err := Db.QueryRow(sp, ID).Scan(&tkn.BonusTokenBalance, &tkn.PaidTokenBalance, &tkn.BalanceDate, &tkn.LastUpdatedDate)
	if err != nil && err != sql.ErrNoRows {

		fmt.Println("error while Fetching token Balance " + err.Error())

		return fmt.Errorf("Error While Getting token Balance %v ", err.Error())
	}
	if err == sql.ErrNoRows {
		tkn.BonusTokenBalance = 0
		tkn.PaidTokenBalance = 0
		tkn.BalanceDate = time.Now()
		tkn.LastUpdatedDate = time.Now()
	}

	return nil
}

// AllocateTokensToID ...
func (tkn *TokenAllocationModel) AllocateTokensToID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}
	sp, _ := RetriveSP("Token_Allocation_ins")
	tknBalSp, _ := RetriveSP("Token_Balance_ins")
	stmt, err := Db.Prepare(sp)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S6TKN001", fmt.Errorf("Error While WhilePreparing Allocating tokens %v ", err.Error()), successResp,
		}
		return Job
	}
	tbStmt, err := Db.Prepare(tknBalSp)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S6TKN001", fmt.Errorf("Error While Preparing Balance tokens %v ", err.Error()), successResp,
		}
		return Job
	}
	var bonusTokens float64
	var paidTokens float64
	if tkn.ModeOfTokenissue == "Bonus" {
		bonusTokens = tkn.AllocatedTokens
		paidTokens = 0
	} else if tkn.ModeOfTokenissue == "Paid" {
		paidTokens = tkn.AllocatedTokens
		bonusTokens = 0
	} else {
		fmt.Println("Error while Allocating Tokens")
		Job <- DbModelError{
			"500", "S6TKN002", fmt.Errorf("Error While inseting Token Allocation Table , Invalid Mode of Issue %s", tkn.ModeOfTokenissue), successResp,
		}
		return Job
	}
	defer stmt.Close()
	defer tbStmt.Close()
	currentTime := time.Now().Format(time.RFC3339)

	_, err = stmt.Exec(tkn.StakeholderID, tkn.AllocatedTokens, currentTime, tkn.PaymentID, tkn.ModeOfTokenissue)

	if err != nil {

		fmt.Println("error while inserting " + err.Error())
		Job <- DbModelError{
			"500", "S6TKN002", fmt.Errorf("Error While inseting Token Allocation Table %v ", err.Error()), successResp,
		}
		return Job
	}
	_, err = tbStmt.Exec(tkn.StakeholderID, bonusTokens, paidTokens, currentTime, bonusTokens, paidTokens, currentTime)
	if err != nil {

		fmt.Println("error while inserting " + err.Error())
		Job <- DbModelError{
			"500", "S6TKN002", fmt.Errorf("Error While inseting Token Balance Table %v ", err.Error()), successResp,
		}
		return Job
	}

	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError

	return Job
}

// GetAllocateTokensToID ...
func (tkn *AllocatedTokens) GetAllocateTokensToID(ID string) error {
	sp, _ := RetriveSP("GET_Token_Allocation_BY_ID")
	rows, err := Db.Query(sp, ID)
	if err != nil && err != sql.ErrNoRows {

		fmt.Println("error while Fetching token Allocations " + err.Error())

		return fmt.Errorf("Error While Getting token Allocations %v ", err.Error())
	}

	if err == sql.ErrNoRows {
		tkn.AllocatedTokens = append(tkn.AllocatedTokens, TokenAllocationModel{})
	}

	defer rows.Close()

	for rows.Next() {
		var newAlloc TokenAllocationModel
		err = rows.Scan(&newAlloc.AllocatedTokens, &newAlloc.AllocatedDate, &newAlloc.PaymentID, &newAlloc.ModeOfTokenissue)
		if err != nil {
			return fmt.Errorf("Cannot read the Allocation  Rows %v", err.Error())
		}
		tkn.AllocatedTokens = append(tkn.AllocatedTokens, newAlloc)
	}

	return nil
}

// TokenTransactionsToID ...
func (tkn *TokenTransactionsModel) TokenTransactionsToID() <-chan DbModelError {
	Job := make(chan DbModelError, 1)
	fmt.Printf("\n%+v\n", tkn)
	successResp := map[string]string{}
	var customError DbModelError
	if CheckPing(&customError); customError.Err != nil {
		Job <- customError
		return Job
	}

	var tknBal TokenBalanceModel

	tksp, _ := RetriveSP("GET_Balance_BY_ID")
	fmt.Println(tkn.StakeholderID, "====UD19940000000002===================------------------")
	err := Db.QueryRow(tksp, tkn.StakeholderID).Scan(&tknBal.BonusTokenBalance, &tknBal.PaidTokenBalance, &tknBal.BalanceDate, &tknBal.LastUpdatedDate)
	if err != nil && err != sql.ErrNoRows {

		fmt.Println("error while Fetching token Balance " + err.Error())
		Job <- DbModelError{
			"500", "S6TKN003", fmt.Errorf("Error While Getting token Balance %v ", err.Error()), successResp,
		}
		return Job
	}
	if err == sql.ErrNoRows {
		Job <- DbModelError{
			"500", "S6TKN003", fmt.Errorf("Insufficient Token Balance (0) required (%s) ", fmt.Sprintf("%.2f", tkn.BonusTokensTransacted+tkn.PaidTokensTransacted)), successResp,
		}
		return Job
	}
	if tknBal.BonusTokenBalance < tkn.BonusTokensTransacted {
		fmt.Println("Insufficient Bonus Token balance" + err.Error())
		Job <- DbModelError{
			"500", "S6TKN001", fmt.Errorf("Transaction failed due to Insufficient Token Balance"), successResp,
		}
		return Job
	}

	if tknBal.PaidTokenBalance < tkn.PaidTokensTransacted {
		fmt.Println("Insufficient Paid Token balance")
		Job <- DbModelError{
			"500", "S6TKN001", fmt.Errorf("Transaction failed due to Insufficient Token Balance"), successResp,
		}
		return Job
	}

	sp, _ := RetriveSP("Token_TX_ins")
	tknBalSp, _ := RetriveSP("Token_Balance_ins")
	stmt, err := Db.Prepare(sp)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S6TKN001", fmt.Errorf("Error While Adding token Transaction %v ", err.Error()), successResp,
		}
		return Job
	}
	tbStmt, err := Db.Prepare(tknBalSp)
	if err != nil {

		fmt.Println("error while inserting" + err.Error())
		Job <- DbModelError{
			"500", "S6TKN001", fmt.Errorf("Error While Preparing Balance tokens %v ", err.Error()), successResp,
		}
		return Job
	}
	defer stmt.Close()
	defer tbStmt.Close()
	currentTime := time.Now()

	stmt.Exec(tkn.StakeholderID, tkn.BonusTokensTransacted, tkn.PaidTokensTransacted, tkn.TransactionID, currentTime)

	if err != nil {

		fmt.Println("error while inserting " + err.Error())
		Job <- DbModelError{
			"500", "S6TKN002", fmt.Errorf("Error While inseting Token Transaction Table %v ", err.Error()), successResp,
		}
		return Job
	}
	tbStmt.Exec(tkn.StakeholderID, tkn.BonusTokensTransacted-(tkn.BonusTokensTransacted*2), tkn.PaidTokensTransacted-(tkn.PaidTokensTransacted*2), currentTime, tkn.BonusTokensTransacted-(tkn.BonusTokensTransacted*2), tkn.PaidTokensTransacted-(tkn.PaidTokensTransacted*2), currentTime)
	if err != nil {

		fmt.Println("error while inserting " + err.Error())
		Job <- DbModelError{
			"500", "S6TKN002", fmt.Errorf("Error While inseting Token Balance Table %v ", err.Error()), successResp,
		}
		return Job
	}

	customError.ErrTyp = "000"
	customError.SuccessResp = successResp

	Job <- customError

	return Job
}

// GetTokenTransactionsForID ...
func (tkn *TokenTransactions) GetTokenTransactionsForID(ID string) error {

	sp, _ := RetriveSP("GET_Token_TX_BY_ID")
	rows, err := Db.Query(sp, ID)
	if err != nil && err != sql.ErrNoRows {

		fmt.Println("error while Fetching token Transactions " + err.Error())

		return fmt.Errorf("Error While Getting token Transactions %v ", err.Error())
	}
	defer rows.Close()
	if err == sql.ErrNoRows {
		tkn.Transactions = append(tkn.Transactions, TokenTransactionsModel{})

	} else {

		for rows.Next() {
			var newTx TokenTransactionsModel
			err = rows.Scan(&newTx.BonusTokensTransacted, &newTx.PaidTokensTransacted, &newTx.TransactionID, &newTx.TransactionDate, &newTx.LastUpdatedDate)
			if err != nil {
				return fmt.Errorf("Cannot read the Transactions  Rows %v", err.Error())
			}
			tkn.Transactions = append(tkn.Transactions, newTx)
		}
	}

	return nil
}
