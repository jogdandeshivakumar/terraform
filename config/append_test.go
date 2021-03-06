package config

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestAppend(t *testing.T) {
	cases := []struct {
		c1, c2, result *Config
		err            bool
	}{
		{
			&Config{
				Atlas: &AtlasConfig{
					Name: "foo",
				},
				Modules: []*Module{
					&Module{Name: "foo"},
				},
				Outputs: []*Output{
					&Output{Name: "foo"},
				},
				ProviderConfigs: []*ProviderConfig{
					&ProviderConfig{Name: "foo"},
				},
				Resources: []*Resource{
					&Resource{Name: "foo"},
				},
				Variables: []*Variable{
					&Variable{Name: "foo"},
				},
				Locals: []*Local{
					&Local{Name: "foo"},
				},

				unknownKeys: []string{"foo"},
			},

			&Config{
				Atlas: &AtlasConfig{
					Name: "bar",
				},
				Modules: []*Module{
					&Module{Name: "bar"},
				},
				Outputs: []*Output{
					&Output{Name: "bar"},
				},
				ProviderConfigs: []*ProviderConfig{
					&ProviderConfig{Name: "bar"},
				},
				Resources: []*Resource{
					&Resource{Name: "bar"},
				},
				Variables: []*Variable{
					&Variable{Name: "bar"},
				},
				Locals: []*Local{
					&Local{Name: "bar"},
				},

				unknownKeys: []string{"bar"},
			},

			&Config{
				Atlas: &AtlasConfig{
					Name: "bar",
				},
				Modules: []*Module{
					&Module{Name: "foo"},
					&Module{Name: "bar"},
				},
				Outputs: []*Output{
					&Output{Name: "foo"},
					&Output{Name: "bar"},
				},
				ProviderConfigs: []*ProviderConfig{
					&ProviderConfig{Name: "foo"},
					&ProviderConfig{Name: "bar"},
				},
				Resources: []*Resource{
					&Resource{Name: "foo"},
					&Resource{Name: "bar"},
				},
				Variables: []*Variable{
					&Variable{Name: "foo"},
					&Variable{Name: "bar"},
				},
				Locals: []*Local{
					&Local{Name: "foo"},
					&Local{Name: "bar"},
				},

				unknownKeys: []string{"foo", "bar"},
			},

			false,
		},

		// Terraform block
		{
			&Config{
				Terraform: &Terraform{
					RequiredVersion: "A",
				},
			},
			&Config{},
			&Config{
				Terraform: &Terraform{
					RequiredVersion: "A",
				},
			},
			false,
		},

		{
			&Config{},
			&Config{
				Terraform: &Terraform{
					RequiredVersion: "A",
				},
			},
			&Config{
				Terraform: &Terraform{
					RequiredVersion: "A",
				},
			},
			false,
		},

		// appending configs merges terraform blocks
		{
			&Config{
				Terraform: &Terraform{
					RequiredVersion: "A",
				},
			},
			&Config{
				Terraform: &Terraform{
					Backend: &Backend{
						Type: "test",
					},
				},
			},
			&Config{
				Terraform: &Terraform{
					RequiredVersion: "A",
					Backend: &Backend{
						Type: "test",
					},
				},
			},
			false,
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%02d", i), func(t *testing.T) {
			actual, err := Append(tc.c1, tc.c2)
			if err != nil != tc.err {
				t.Errorf("unexpected error: %s", err)
			}

			if !reflect.DeepEqual(actual, tc.result) {
				t.Errorf("wrong result\ngot: %swant: %s", spew.Sdump(actual), spew.Sdump(tc.result))
			}
		})
	}
}
