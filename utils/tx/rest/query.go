package rest

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// EncodeTxRequestHandlerFn returns the decode tx REST handler. In particular,
// it takes a base64-encoded bytes, decodes it using the Amino wire protocol,
// and responds with JSON-encoded transaction.
func DecodeTxRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		var req DecodeTxsRequest

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		resp := DecodeTxsResponse{
			Txs: []auth.StdTx{},
		}

		for _, base64str := range req.Txs {
			tx, err := decodeTx(cliCtx.Codec, base64str)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			resp.Txs = append(resp.Txs, tx)
		}

		rest.PostProcessResponseBare(w, cliCtx, &resp)
	}
}

func decodeTx(cdc *codec.Codec, base64str string) (tx auth.StdTx, err error) {
	var res types.StdTx

	bytes, err := base64.StdEncoding.DecodeString(base64str)
	if err != nil {
		return auth.StdTx{}, err
	}

	err = cdc.UnmarshalBinaryLengthPrefixed(bytes, &res)
	if err != nil {
		return auth.StdTx{}, err
	}

	return res, nil
}

func SignTxHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))
			return
		}

		name, passphrase := r.FormValue("name"), r.FormValue("passphrase")

		var signMsg types.StdSignMsg

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &signMsg) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		txBldr := auth.NewTxBuilderFromCLI().
			WithTxEncoder(utils.GetTxEncoder(cliCtx.Codec)).
			WithAccountNumber(signMsg.AccountNumber).
			WithSequence(signMsg.Sequence).
			WithChainID(signMsg.ChainID)

		stdTx := auth.StdTx{
			Msgs:       signMsg.Msgs,
			Fee:        signMsg.Fee,
			Signatures: nil,
			Memo:       signMsg.Memo,
		}

		signedStdTx, err := txBldr.SignStdTx(name, passphrase, stdTx, false)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, signedStdTx)
	}
}

func BroadcastTxHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var stdTx types.StdTx

		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &stdTx) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		txBytes, err := cliCtx.Codec.MarshalBinaryLengthPrefixed(stdTx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		res, err := BroadcastMessage(cliCtx, txBytes)

		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func SignMessage(cliCtx context.CLIContext, chainId string, signer sdk.AccAddress,
	name string, passphrase string, msg []sdk.Msg) ([]byte, error) {

	acc, err := auth.NewAccountRetriever(cliCtx).GetAccount(signer)
	if err != nil {
		return nil, err
	}

	txBldr := auth.NewTxBuilderFromCLI().
		WithTxEncoder(utils.GetTxEncoder(cliCtx.Codec)).
		WithAccountNumber(acc.GetAccountNumber()).
		WithSequence(acc.GetSequence()).
		WithChainID(chainId)

	return txBldr.BuildAndSign(name, passphrase, msg)
}

func BroadcastMessage(cliCtx context.CLIContext, message []byte) ([]byte, error) {
	res, err := cliCtx.BroadcastTxCommit(message)
	if err != nil {
		return nil, err
	}

	txBytes, err := cliCtx.Codec.MarshalJSON(res)
	if err != nil {
		return nil, err
	}

	return txBytes, nil
}

func SignAndBroadcastMessage(cliCtx context.CLIContext, chainId string, signer sdk.AccAddress, account string, passphrase string, msg []sdk.Msg) ([]byte, error) {
	signedMsg, err := SignMessage(cliCtx, chainId, signer, account, passphrase, msg)
	if err != nil {
		return nil, err
	}

	return BroadcastMessage(cliCtx, signedMsg)
}
