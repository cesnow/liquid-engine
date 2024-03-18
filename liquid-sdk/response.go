package LiquidSDK

func ResponseError(error string) CmdErrorResponse {
	return CmdErrorResponse{
		Error: error,
	}
}
