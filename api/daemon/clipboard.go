// Copyright 2020-2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPL-3.0
// license that can be found in the LICENSE file.

package daemon

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"changkun.de/x/midgard/internal/clipboard"
	"changkun.de/x/midgard/internal/config"
	"changkun.de/x/midgard/internal/hotkey"
	"changkun.de/x/midgard/internal/types"
	"changkun.de/x/midgard/internal/utils"
)

func (m *Daemon) watchLocalClipboard(ctx context.Context) {
	last := time.Now()
	hotkey.Handle(ctx, func() {
		if time.Since(last) < time.Second*5 {
			log.Println("pressing hotkey too fast, ignoring")
			return
		}
		last = time.Now()

		var msg string
		defer func() {
			log.Print(msg)
			clipboard.Local.Write(
				types.MIMEPlainText, utils.StringToBytes(msg))
		}()

		log.Println("hotkey triggered")
		res, err := utils.Request(
			http.MethodPut,
			types.EndpointAllocateURL,
			&types.AllocateURLInput{
				Source: types.SourceUniversalClipboard,
			})
		if err != nil {
			msg = fmt.Sprintf("cannot perform allocate request, err: %v", err)
			return
		}
		var out types.AllocateURLOutput
		err = json.Unmarshal(res, &out)
		if err != nil {
			msg = fmt.Sprintf("cannot parse requested URL, err: %v", err)
			return
		}
		if out.URL == "" {
			msg = out.Message
		} else {
			msg = config.Get().Domain + out.URL
		}
	})

	textCh := clipboard.Local.Watch(ctx, types.MIMEPlainText)
	imagCh := clipboard.Local.Watch(ctx, types.MIMEImagePNG)
	for {
		select {
		case <-ctx.Done():
			return
		case text, ok := <-textCh:
			if !ok {
				return
			}

			d := &types.PutToUniversalClipboardInput{}
			d.Type = types.MIMEPlainText
			d.Data = utils.BytesToString(text)
			d.DaemonID = m.ID
			b, _ := json.Marshal(d)
			log.Println("local clipboard has changed as text, sync to server...")
			m.writeCh <- &types.WebsocketMessage{
				Action:  types.ActionClipboardPut,
				UserID:  m.ID,
				Message: "local clipboard has changed",
				Data:    b,
			}
		case img, ok := <-imagCh:
			if !ok {
				return
			}
			d := &types.PutToUniversalClipboardInput{}
			d.Type = types.MIMEImagePNG
			d.Data = base64.StdEncoding.EncodeToString(img)
			d.DaemonID = m.ID
			b, _ := json.Marshal(d)
			log.Println("local clipboard has changed as image, sync to server...")
			m.writeCh <- &types.WebsocketMessage{
				Action:  types.ActionClipboardPut,
				UserID:  m.ID,
				Message: "local clipboard has changed",
				Data:    b,
			}
		}
	}
}
