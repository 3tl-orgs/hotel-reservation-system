package core

import sctx "github.com/ngleanhvu/go-booking/shared/srvctx"

func Recover() {
	if r := recover(); r != nil {
		sctx.GlobalLogger().GetLogger("recovered").Errorln(r)
	}
}
