package cdrConvert

import (
	"encoding/hex"
	"strings"
	"time"

	"github.com/JieranAgileCore/chf/cdr/asn"
	"github.com/JieranAgileCore/chf/cdr/cdrType"
	"github.com/JieranAgileCore/openapi/models"
)

func MultiUnitUsageToCdr(multiUnitUsageList []models.MultipleUnitUsage) []cdrType.MultipleUnitUsage {
	cdrMultiUnitUsageList := make([]cdrType.MultipleUnitUsage, 0, len(multiUnitUsageList))

	for _, multiUnitUsage := range multiUnitUsageList {
		usedUnitContainer := UsedUnitContainerToCdr(multiUnitUsage.UsedUnitContainer)
		cdrMultiUnitUsage := cdrType.MultipleUnitUsage{
			cdrType.RatingGroupId{
				int64(multiUnitUsage.RatingGroup),
			},
			usedUnitContainer,
			&cdrType.NetworkFunctionName{
				asn.IA5String(multiUnitUsage.UPFID),
			},
			// TODO convert PDUAddress, not exist in current spec
			nil,
		}
		cdrMultiUnitUsageList = append(cdrMultiUnitUsageList, cdrMultiUnitUsage)
	}

	return cdrMultiUnitUsageList
}

// TODO: Only convert Local Sequence Number, Uplink, Downlink, Total Volumn,
//       Service Specific Units currently.
func UsedUnitContainerToCdr(usedUnitContainerList []models.UsedUnitContainer) []cdrType.UsedUnitContainer {
	cdrUsedUnitContainerList := make([]cdrType.UsedUnitContainer, 0, len(usedUnitContainerList))

	for _, usedUnitContainer := range usedUnitContainerList {
		serviceSpecificUnits := int64(usedUnitContainer.ServiceSpecificUnits)
		cdrUsedUnitContainer := cdrType.UsedUnitContainer{
			LocalSequenceNumber: &cdrType.LocalSequenceNumber{
				int64(usedUnitContainer.LocalSequenceNumber),
			},
			DataVolumeUplink: &cdrType.DataVolumeOctets{
				int64(usedUnitContainer.UplinkVolume),
			},
			DataVolumeDownlink: &cdrType.DataVolumeOctets{
				int64(usedUnitContainer.DownlinkVolume),
			},
			DataTotalVolume: &cdrType.DataVolumeOctets{
				int64(usedUnitContainer.TotalVolume),
			},
			ServiceSpecificUnits: &serviceSpecificUnits,
		}
		cdrUsedUnitContainerList = append(cdrUsedUnitContainerList, cdrUsedUnitContainer)
	}

	return cdrUsedUnitContainerList
}

// TODO convert type Trigger
func TriggersToCdr(triggers []models.Trigger) []cdrType.Trigger {
	cdrTriggers := make([]cdrType.Trigger, 0, len(triggers))

	return cdrTriggers
}

// format: YYMMDDhhmmssShhmm
// BCD encoded
func TimeStampToCdr(t *time.Time) cdrType.TimeStamp {
	ts := make(asn.OctetString, 9)

	_, tz := t.Zone()
	ts[0] = (byte(t.Year()%100/10) << 4) | (byte(t.Year() % 10))
	ts[1] = (byte(t.Month()/10) << 4) | (byte(t.Month() % 10))
	ts[2] = (byte(t.Day()/10) << 4) | (byte(t.Day() % 10))
	ts[3] = (byte(t.Hour()/10) << 4) | (byte(t.Hour() % 10))
	ts[4] = (byte(t.Minute()/10) << 4) | (byte(t.Minute() % 10))
	ts[5] = (byte(t.Second()/10) << 4) | (byte(t.Second() % 10))
	if tz >= 0 {
		ts[6] = byte('+')
	} else {
		ts[6] = byte('-')
	}
	ts[7] = (byte(tz/3600/10) << 4) | (byte(tz / 3600 % 10))
	ts[8] = (byte(tz%3600/10) << 4) | (byte(tz % 3600 % 10))
	cdrTimeStamp := cdrType.TimeStamp{
		Value: ts,
	}

	return cdrTimeStamp
}

func PlmnIdToCdr(modelsPlmnid models.PlmnId) cdrType.PLMNId {
	var hexString string
	mcc := strings.Split(modelsPlmnid.Mcc, "")
	mnc := strings.Split(modelsPlmnid.Mnc, "")
	if len(modelsPlmnid.Mnc) == 2 {
		hexString = mcc[1] + mcc[0] + "f" + mcc[2] + mnc[1] + mnc[0]
	} else {
		hexString = mcc[1] + mcc[0] + mnc[0] + mcc[2] + mnc[2] + mnc[1]
	}

	var cdrPlmnId cdrType.PLMNId
	if plmnId, err := hex.DecodeString(hexString); err == nil {
		cdrPlmnId.Value = plmnId
	}
	return cdrPlmnId
}
