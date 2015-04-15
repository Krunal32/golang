// +build !solution
// Leave an empty line above this comment.
package zlog

import (
	"fmt"
	"github.com/uis-dat320-fall2014/labs/zaplab/chzap"
	"github.com/uis-dat320-fall2014/labs/zaplab/util"
	"sort"
	"time"
)

type Zaps []chzap.ChZap

func NewSimpleZapLogger() ZapLogger {
	zs := make(Zaps, 0)
	return &zs
}
func (zs *Zaps) LogZap(z chzap.ChZap) {
	if z.IsZap() == true {
		*zs = append(*zs, z)
	}
}
func (zs *Zaps) Entries() int {
	return len(*zs)
}
func (zs *Zaps) String() string {
	return fmt.Sprintf("SS: %d", len(*zs))
}
func (zs *Zaps) Viewers(chName string) int {
	defer util.TimeElapsed(time.Now(), "simple.Viewers")
	viewers := 0
	for _, v := range *zs {
		if v.ToChan == chName {
			viewers++
		}
		if v.FromChan == chName {
			viewers--
		}
	}
	return viewers
}

//return list of channels (with no duplicates) by going through the zaps
//using channelnames as keys with empty struct as value to save memory
// then converting to a string slice
func (zs *Zaps) Channels() []string {
	defer util.TimeElapsed(time.Now(), "simple.Channels")
	if len(*zs) < 1 {
		return nil
	}
	cvmap := make(map[string]struct{})
	for _, v := range *zs {
		cvmap[v.ToChan] = struct{}{}
		cvmap[v.FromChan] = struct{}{}
	}
	strarr := make([]string, 0, len(cvmap))
	for key := range cvmap {
		strarr = append(strarr, key)
	}
	return strarr
}
func (zs *Zaps) ChannelsViewers() ChanViewersList {
	defer util.TimeElapsed(time.Now(), "simple.ChannelsViewers")
	strarr := zs.Channels()
	if strarr == nil || len(strarr) == 0 {
		return nil
	}
	cvarr := make(ChanViewersList, len(strarr))
	for _, str := range strarr {
		i := zs.Viewers(str)
		cvarr = append(cvarr, ChannelViewers{Channel: str, Viewers: i})
	}
	return cvarr
}

func (zs *Zaps) Top10() ChanViewersList { // working implementation
	ch := zs.ChannelsViewers()
	if len(ch) >= 10 {
		sort.Sort(ch)
		s := (ch)[:10]
		return s
	}
	return nil
}
func (zs *Zaps) AverageDuration() time.Duration { // not implemented on this structure
	return 0
}
