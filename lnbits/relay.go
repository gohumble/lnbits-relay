package lnbits

import (
	log "github.com/sirupsen/logrus"
	"strconv"

	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type LNBitsRelay struct {
	httpServer *http.Server
	client     *Client
}

func NewDonationRelay(address string, client *Client) *LNBitsRelay {
	srv := &http.Server{
		Addr: address,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	apiServer := &LNBitsRelay{
		httpServer: srv,
		client:     client,
	}
	apiServer.httpServer.Handler = apiServer.newRouter()
	return apiServer
}

func (w *LNBitsRelay) Start() {
	log.Fatal(w.httpServer.ListenAndServe())
}

func (w *LNBitsRelay) newRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/donate/{amount}", w.invoice).Methods(http.MethodGet)
	return router
}

func (w LNBitsRelay) invoice(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	amount, err := strconv.Atoi(vars["amount"])
	if err != nil {
		writer.WriteHeader(http.StatusPreconditionFailed)
	}
	invoice, err := w.client.Invoice(InvoiceParams{Amount: int64(amount), Out: false, Memo: "Donation"})
	if err != nil {
		return
	}
	writer.Write([]byte(invoice.PaymentRequest))
}
