// +build !solution
// Leave an empty line above this comment.
package chzap

import (
	"net"
	"strings"
	"time"
)

const timeFormat = "2006/01/02, 15:04:05"
const dateFormat = "2006/01/02"
const timeOnly = "15:04:05"
const vol, mute, hdmi = "Volume", "Mute", "HDMI"
const timeLen = len(timeFormat)

type StatusChange struct {
	Volume string
	Mute   string
	HDMI   string
}

type ChZap struct {
	Time     time.Time
	IP       net.IP
	Status   StatusChange //contains status change field (for cleaner printouts and consistent data)
	FromChan string
	ToChan   string
}

// modified to return onstat StatusChangely Chzap. Since Chzap struct contains a StatusChange struct.
func NewSTBEvent(event string) (*ChZap, error) {
	chzp := new(ChZap)
	newzap := strings.Split(strings.TrimSpace(event), ", ")              // parse each field into array
	time, err := time.Parse(timeFormat, strings.Join(newzap[0:2], ", ")) // string to time.Time
	*chzp = ChZap{Time: time, IP: net.ParseIP(newzap[2])}                // string to net.IP
	if len(newzap) == 5 {                                                //has toChan and from Chan field
		(*chzp).FromChan = newzap[3]
		(*chzp).ToChan = newzap[4]
		return chzp, err
	}
	if len(newzap) == 4 { //has status field
		if strings.Contains(newzap[3], vol) {
			(*chzp).Status.Volume = newzap[3]
		} else if strings.Contains(newzap[3], mute) {
			(*chzp).Status.Mute = newzap[3]
		} else if strings.Contains(newzap[3], hdmi) {
			(*chzp).Status.HDMI = newzap[3]
		}
		return chzp, err
	}
	return nil, err
}

// returns string representation of Chzaps
// also prints StatusChanges if contained
func (zap ChZap) String() string {
	tmp := []string{""}
	if st := zap.Status.String(); st != "" {
       // zap with statusfield
		tmp = []string{zap.Date(), (zap.IP).String(), st} 
	} else if to, fr := zap.FromChan, zap.ToChan; fr != "" && to != "" { 
        // zap without statusfield
		tmp = []string{zap.Date(), (zap.IP).String(), fr, to}
	}
	return strings.Join(tmp, ", ") // seperate with comma
}
func (schg StatusChange) String() string { 
    // return string representation. Only 1 field should be non-nempty
	return schg.Volume + schg.Mute + schg.HDMI
}

// The duration between receiving (this) zap event and the provided event
func (zap ChZap) Duration(provided ChZap) time.Duration {
	d := time.Duration((zap.Time).Sub(provided.Time)) // difference
	if d < 0 {
		return time.Duration((provided.Time).Sub(zap.Time))
	}
	return d
}
//false if the ChZap is a statuschange, true if it's a zap event
func (zap ChZap) IsZap() bool { 
	if str := zap.Status.Volume + zap.Status.HDMI + zap.Status.Mute; zap.ToChan != "" && str == "" {
		return true
	}
	return false
}
func (zap ChZap) Date() string {
	return zap.Time.Format(timeFormat)
}
