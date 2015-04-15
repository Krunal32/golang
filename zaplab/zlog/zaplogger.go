package zlog

import (
	"fmt"
	"github.com/uis-dat320-fall2014/labs/zaplab/chzap"
	"strings"
	"time"
)

type ZapLogger interface {
	LogZap(z chzap.ChZap)
	Entries() int
	AverageDuration() time.Duration // computes average zap duration
	Top10() ChanViewersList         // returns list of Top10 channels
	Viewers(channelName string) int
	Channels() []string
	ChannelsViewers() ChanViewersList
}
type ChannelViewers struct {
	Channel string
	Viewers int
}

func (cv ChannelViewers) String() string {
	return fmt.Sprintf("%s: %d", cv.Channel, cv.Viewers)
}

//Print  string representation of chanviewerslist
func (cv ChanViewersList) String() string {
	str := make([]string, 0, len(cv))
	if cv != nil {
		for _, ch := range cv {
			str = append(str, ch.String())
		}
	}
	return strings.Join(str, "\n")
}

type ChanViewersList []ChannelViewers

type ByViewers struct{ ChanViewersList } // redundant
// implementing sort interface
func (t ChanViewersList) Len() int { return len(t) }
func (t ChanViewersList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t ChanViewersList) Less(i, j int) bool {
	return t[i].Viewers > t[j].Viewers
}
