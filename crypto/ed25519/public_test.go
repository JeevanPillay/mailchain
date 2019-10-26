package ed25519

import (
	"reflect"
	"testing"
)

func TestPublicKey_Bytes(t *testing.T) {
	tests := []struct {
		name string
		pk   PublicKey
		want []byte
	}{
		{
			"sofia",
			sofiaPublicKey,
			sofiaPublicKeyBytes,
		},
		{
			"charlotte",
			charlottePublicKey,
			charlottePublicKeyBytes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pk.Bytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicKey.Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPublicKeyFromBytes(t *testing.T) {
	type args struct {
		keyBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *PublicKey
		wantErr bool
	}{
		{
			"sofia",
			args{
				sofiaPublicKeyBytes,
			},
			&sofiaPublicKey,
			false,
		},
		{
			"err-too-short",
			args{
				[]byte{0x72, 0x3c, 0xaa, 0x23},
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublicKeyFromBytes(tt.args.keyBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicKeyFromBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicKeyFromBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
