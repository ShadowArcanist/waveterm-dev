// Copyright 2026, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

// Telemetry has been neutralized in this build. All collection, local
// storage, and cloud transmission have been removed. The public functions
// below are retained as inert no-ops so existing call sites keep compiling,
// but they never record, persist, or upload any data.
package telemetry

import (
	"context"

	"github.com/wavetermdev/waveterm/pkg/telemetry/telemetrydata"
	"github.com/wavetermdev/waveterm/pkg/wshrpc"
)

const MaxTzNameLen = 50
const ActivityEventName = "app:activity"
const WshRunEventName = "wsh:run"

type ActivityType struct {
	Day           string        `json:"day"`
	Uploaded      bool          `json:"-"`
	TData         TelemetryData `json:"tdata"`
	TzName        string        `json:"tzname"`
	TzOffset      int           `json:"tzoffset"`
	ClientVersion string        `json:"clientversion"`
	ClientArch    string        `json:"clientarch"`
	BuildTime     string        `json:"buildtime"`
	OSRelease     string        `json:"osrelease"`
}

type TelemetryData struct {
	ActiveMinutes int                          `json:"activeminutes"`
	FgMinutes     int                          `json:"fgminutes"`
	OpenMinutes   int                          `json:"openminutes"`
	NumTabs       int                          `json:"numtabs"`
	NewTab        int                          `json:"newtab"`
	Displays      []wshrpc.ActivityDisplayType `json:"displays,omitempty"`
	Renderers     map[string]int               `json:"renderers,omitempty"`
	Blocks        map[string]int               `json:"blocks,omitempty"`
	WshCmds       map[string]int               `json:"wshcmds,omitempty"`
	Conn          map[string]int               `json:"conn,omitempty"`
}

func GetTosAgreedTs() int64 {
	return 0
}

func IsTelemetryEnabled() bool {
	return false
}

func IsAutoUpdateEnabled() bool {
	return false
}

func AutoUpdateChannel() string {
	return ""
}

func GoUpdateActivityWrap(update wshrpc.ActivityUpdate, debugStr string) {
}

func GoRecordTEventWrap(tevent *telemetrydata.TEvent) {
}

func RecordTEvent(ctx context.Context, tevent *telemetrydata.TEvent) error {
	return nil
}

func UpdateActivity(ctx context.Context, update wshrpc.ActivityUpdate) error {
	return nil
}

func TruncateActivityTEventForShutdown(ctx context.Context) error {
	return nil
}
