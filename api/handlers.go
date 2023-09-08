package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tapansaikia/slot-book-server/types"
)

func (ctx *HandlerContext) bookSlot(w http.ResponseWriter, r *http.Request) {
	//...use ctx.store to query the database...

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleErr(w, err, http.StatusInternalServerError)
		return
	}

	// fmt.Println(string(body))

	payload := types.Booking{}
	err = json.NewDecoder(bytes.NewReader(body)).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Do something with payload
	fmt.Println(payload)
}
