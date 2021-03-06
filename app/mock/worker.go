package mock

import(
	"github.com/bobsar0/AutoTrade/model"
	"strconv"
	"log"
)

//Worker provides Transaction services 
type Worker struct{
	Sess model.Session
	AddOrUpdateDbChan chan model.DbData
	GetDbChan chan model.DbData
}
//Spins out a new Worker
func NewWorker() *Worker{
	return &Worker{
		
	}
}

var _ model.TransactionService = &Worker{} //Enforces that Worker implements TransactionService interface

//A test function to simulate getting ticker price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)FuncThatReturnTicker()float64{
		return 0.002134442
}
//A test function to simulate getting balance price from a trading site
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)FuncThatReturnBalance()float64{
	return 0.026654442
}

//A test function to simulate placing order on a trading site with mock inputs and some outputs determined by the inputs
//In the real world, this will be achieved by communicating with the site's API
func(w *Worker)FuncThatPlacesOrder(in model.OrderInput) model.OrderOutput{
	quant,_ := strconv.ParseFloat(in.Quantity, 64)
	return model.OrderOutput{
		OrderID: 100,
		ClientID: 1,
		Symbol: in.Symbol,
		Ticker: in.Ticker,
		Quantity: quant,
		Balance: 0.02661173,
	}
}
func(w *Worker)AddTransaction(ts model.Transaction) (model.TransactionID, error){
	if user := w.Sess.Authenticate(); user.Username != "Steve" {
		log.Println("Wrong Username")
		return "", model.ErrUnauthorized
	}
	cChan := make(chan model.DbResp)
	w.AddOrUpdateDbChan <-model.DbData{Transaction: ts, CallerChan: cChan}
	res := <-cChan
	return res.TransID, res.Err
}

func(w *Worker)GetTransaction(id model.TransactionID) (model.Transaction, error){
	if user := w.Sess.Authenticate(); user.Username != "Chidi" {
		log.Println("Wrong Name")
		return model.Transaction{}, model.ErrUnauthorized
	}
	cChan := make(chan model.DbResp)
	w.GetDbChan <-model.DbData{TransID: id, CallerChan: cChan}
	res := <-cChan
	return res.Transaction, res.Err
}
