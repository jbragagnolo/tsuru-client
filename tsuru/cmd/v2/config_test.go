// Copyright 2026 tsuru-client authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/fatih/color"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/tsuru/tablecli"
	"github.com/tsuru/tsuru-client/tsuru/cmd/standards"
)

func TestAppInfoViewFromClientConfig(t *testing.T) {
	originalViper, originalDir := defaultViper, TsuruConfigDir
	originalTable, originalColor, originalPadding := tablecli.TableConfig, color.NoColor, standards.SubTableWriterPadding
	defer func() {
		defaultViper, TsuruConfigDir = originalViper, originalDir
		tablecli.TableConfig, color.NoColor, standards.SubTableWriterPadding = originalTable, originalColor, originalPadding
	}()

	TsuruConfigDir = t.TempDir()
	t.Setenv("TSURU_APP_INFO_VIEW", "")
	if err := os.WriteFile(filepath.Join(TsuruConfigDir, "client.yaml"), []byte("app-info-view: detailed\n"), 0600); err != nil {
		t.Fatal(err)
	}
	defaultViper = preSetupViper(nil)
	view, found := AppInfoView()
	assert.True(t, found)
	assert.Equal(t, "detailed", view)

	t.Setenv("TSURU_APP_INFO_VIEW", "simplified")
	view, found = AppInfoView()
	assert.True(t, found)
	assert.Equal(t, "simplified", view)
}

func TestPager(t *testing.T) {
	originalViper := defaultViper
	defer func() { defaultViper = originalViper }()

	tests := []struct {
		name          string
		setup         func(v *viper.Viper)
		expectedPager string
		expectedFound bool
	}{
		{
			name:          "pager key not set",
			setup:         func(v *viper.Viper) {},
			expectedPager: "",
			expectedFound: false,
		},
		{
			name: "pager set to boolean true",
			setup: func(v *viper.Viper) {
				v.Set("pager", true)
			},
			expectedPager: "",
			expectedFound: false,
		},
		{
			name: "pager set to string true",
			setup: func(v *viper.Viper) {
				v.Set("pager", "true")
			},
			expectedPager: "",
			expectedFound: false,
		},
		{
			name: "pager set to boolean false",
			setup: func(v *viper.Viper) {
				v.Set("pager", false)
			},
			expectedPager: "",
			expectedFound: true,
		},
		{
			name: "pager set to string false",
			setup: func(v *viper.Viper) {
				v.Set("pager", "false")
			},
			expectedPager: "",
			expectedFound: true,
		},
		{
			name: "pager set to custom pager command",
			setup: func(v *viper.Viper) {
				v.Set("pager", "less -R")
			},
			expectedPager: "less -R",
			expectedFound: true,
		},
		{
			name: "pager set to another custom pager",
			setup: func(v *viper.Viper) {
				v.Set("pager", "more")
			},
			expectedPager: "more",
			expectedFound: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := viper.New()
			tt.setup(v)
			defaultViper = v

			pager, found := Pager()

			assert.Equal(t, tt.expectedPager, pager)
			assert.Equal(t, tt.expectedFound, found)
		})
	}
}
