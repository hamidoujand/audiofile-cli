package command

import (
	"flag"
	"testing"

	"github.com/hamidoujand/audiofile/internal/interfaces"
)

func TestGetCommand_Run(t *testing.T) {
	type fields struct {
		client interfaces.Client
		fs     *flag.FlagSet
		id     string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc := &GetCommand{
				client: tt.fields.client,
				fs:     tt.fields.fs,
				id:     tt.fields.id,
			}
			if err := gc.Run(); (err != nil) != tt.wantErr {
				t.Errorf("GetCommand.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
