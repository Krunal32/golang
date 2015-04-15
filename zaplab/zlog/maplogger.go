// +build !solution
// Leave an empty line above this comment.
package zlog

import (
	"fmt"
	"github.com/uis-dat320-fall2014/labs/zaplab/chzap"
	"github.com/uis-dat320-fall2014/labs/zaplab/util"
	"sort"
    //"strings"
	"time"
)

type ZapMap struct {
	Zaps  map[string]chzap.ChZap
	Chans map[string]ChannelViewers
    Durations []time.Duration
}

func NewViewersZapLogger() ZapLogger {
	zs := ZapMap{Zaps: make(map[string]chzap.ChZap), Chans: make(map[string]ChannelViewers),Durations:make([]time.Duration,0)}
	return &zs
}
func (zs *ZapMap) LogZap(z chzap.ChZap) {
       durs:=&(*zs).Durations
	   ip:=z.IP.String()
	   if  z.IsZap() {
	    if ch,exists:=((*zs).Zaps)[ip];exists{
         dur:=ch.Duration(z) 
         if dur >0 {
       
         *durs=append(*durs,dur)
         }
        ((*zs).Zaps)[ip]=z
   	    } else {
          // z.ZDuration=time.Duration(0)
	    ((*zs).Zaps)[ip]=z
	    }
	     zs.logViewer(z)
	    }
	
}
// log viewer counts
func (zs *ZapMap) logViewer(z chzap.ChZap) {
	ch := &zs.Chans
	_, exists := ((*zs).Zaps)[z.IP.String()] // checks for existence of ip
	to, fr := z.ToChan, z.FromChan 
	if v, ok := (*ch)[fr]; ok && exists { // channel exists
		(*ch)[fr] = ChannelViewers{Channel: fr, Viewers: v.Viewers - 1} // zapped from channel. Lost 1 viewer
	} else {
		(*ch)[fr] = ChannelViewers{Channel: fr, Viewers: 0} // channel doesen't exist in map
	}
	if v, ok := (*ch)[to]; ok {
		(*ch)[to] = ChannelViewers{Channel: to, Viewers: v.Viewers + 1} // zapped to channel. Gained 1 viewer
	} else {
		(*ch)[to] = ChannelViewers{Channel: to, Viewers: 0} //

	}
}
// number of zaps in map and also the number of unique IPs
func (zs *ZapMap) Entries() int {
	return len((*zs).Zaps)
}
func (zs *ZapMap) String() string {
	return fmt.Sprintf("SS: %d", len((*zs).Zaps))
}
// conversion from map of channelviewers to a string array representation
func (zs *ZapMap) Channels() []string {
	defer util.TimeElapsed(time.Now(), "simple.Channels")
	str := make([]string, 0)
	//copy map values to a slice
	zap := (*zs).Chans
	for _, v := range zap {
		str = append(str, v.Channel)
	}
	return str
}
// conversion from default map to slice
// this is done in order to sort the data later
func (zs *ZapMap) ChannelsViewers() ChanViewersList {
	defer util.TimeElapsed(time.Now(), "simple.ChannelsViewers")
	zap := (*zs).Chans
	chslice := make(ChanViewersList, 0, len(zap))
	for _, v := range zap {
		chslice = append(chslice, v)
	}
	return chslice
}
// obtains list of channelviewers 
// sorts list and returns top10
func (zm *ZapMap) Top10() ChanViewersList {
	ch := zm.ChannelsViewers()
	if len(ch) >= 10 { 
		sort.Sort(ch)
		s := (ch)[:10]
		return s
	}
	return nil
}
// compute viewers 
// a lookup in the viewers map 
func (zs *ZapMap) Viewers(chName string) int {
	if v, ok := ((*zs).Chans)[chName]; ok {
		return v.Viewers
	} else {
		return 0
	}
}
// compute average duration between zaps
func (zs *ZapMap) AverageDuration() time.Duration {
   var totalDur,n int64
   for _, dur := range (*zs).Durations {
          totalDur+=int64(dur) // convert to int64 for computation
   }
   if n=int64(len((*zs).Durations));n>0 { 
	return time.Duration((totalDur/n)) // convert back to a time.Duration
   }
   return 0
}
