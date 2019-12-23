/*---------------------------------------------------------------------------------------------
 *  Copyright (c) IBAX. All rights reserved.
 *  See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package api

import (
	"net/http"

	"github.com/IBAX-io/go-ibax/packages/smart"
func getTestHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	jsonResponse(w, &getTestResult{
		Value: smart.GetTestValue(params["name"]),
	})
}