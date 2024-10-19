// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

package client

/*
func TestCreateSession(t *testing.T) {
	l := slog.New(slog.NewTextHandler(io.Discard, nil))

	e, err := engine.NewEngine(l)
	if err != nil {
		t.Fatalf("Failed to create a new engine: %v", err)
	}
	defer e.Shutdown()

	c := config.NewConfig()
	if err := config.AcquireConfig("", "config.yml", c); err != nil {
		t.Errorf("AcquireConfig failed: %v", err)
	}

	client := NewClient("http://localhost:4000/graphql")
	if _, err := client.CreateSession(c); err != nil {
		t.Errorf("CreateSession failed: %v", err)
	}
}

func TestCreateAsset(t *testing.T) {
	l := slog.New(slog.NewTextHandler(io.Discard, nil))

	e, err := engine.NewEngine(l)
	if err != nil {
		t.Fatalf("Failed to create a new engine: %v", err)
	}
	defer e.Shutdown()

	c := config.NewConfig()
	if err := config.AcquireConfig("", "config.yml", c); err != nil {
		t.Errorf("AcquireConfig failed: %v", err)
	}

	client := NewClient("http://localhost:4000/graphql")
	token, _ := client.CreateSession(c)

	addr, _ := netip.ParseAddr("192.168.0.1")
	asset := oamnet.IPAddress{Address: addr, Type: "IPv4"}
	data := types.AssetData{
		OAMAsset: asset,
		OAMType:  asset.AssetType(),
	}

	a := types.Asset{Session: token, Name: "Asset#1", Data: data}
	if err := client.CreateAsset(a, token); err != nil {
		t.Errorf("CreateAsset failed: %v", err)
	}
}

func TestSubscribe(t *testing.T) {
	l := slog.New(slog.NewTextHandler(io.Discard, nil))

	e, err := engine.NewEngine(l)
	if err != nil {
		t.Fatalf("Failed to create a new engine: %v", err)
	}
	defer e.Shutdown()

	c := config.NewConfig()
	if err := config.AcquireConfig("", "config.yml", c); err != nil {
		t.Errorf("AcquireConfig failed: %v", err)
	}

	client := NewClient("http://localhost:4000/graphql")
	token, _ := client.CreateSession(c)

	ch, err := client.Subscribe(token)
	if err != nil {
		t.Errorf("Subscribe failed: %v", err)
	}
	time.Sleep(time.Second)

	select {
	case <-ch:
	default:
		t.Error("Failed to receive a message from the channel")
	}
}
*/
