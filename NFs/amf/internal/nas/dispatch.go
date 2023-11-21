package nas

import (
	"errors"
	"fmt"

	"github.com/JieranAgileCore/amf/internal/context"
	"github.com/JieranAgileCore/amf/internal/gmm"
	"github.com/JieranAgileCore/amf/internal/logger"
	"github.com/JieranAgileCore/nas"
	"github.com/JieranAgileCore/openapi/models"
	"github.com/JieranAgileCore/util/fsm"
)

func Dispatch(ue *context.AmfUe, accessType models.AccessType, procedureCode int64, msg *nas.Message) error {
	if msg.GmmMessage == nil {
		return errors.New("Gmm Message is nil")
	}

	if msg.GsmMessage != nil {
		return errors.New("GSM Message should include in GMM Message")
	}

	if ue.State[accessType] == nil {
		return fmt.Errorf("UE State is empty (accessType=%q). Can't send GSM Message", accessType)
	}

	return gmm.GmmFSM.SendEvent(ue.State[accessType], gmm.GmmMessageEvent, fsm.ArgsType{
		gmm.ArgAmfUe:         ue,
		gmm.ArgAccessType:    accessType,
		gmm.ArgNASMessage:    msg.GmmMessage,
		gmm.ArgProcedureCode: procedureCode,
	}, logger.GmmLog)
}
