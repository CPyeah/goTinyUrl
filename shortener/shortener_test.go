package shortener

import (
	"reflect"
	"testing"
)

func TestGenerateShortLink(t *testing.T) {
	type args struct {
		longUrl string
		uuId    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"https://github.com/CPyeah/goTinyUrl", "11212"}, "G5G2hkvi"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateShortLink(tt.args.longUrl, tt.args.uuId); got != tt.want {
				t.Errorf("GenerateShortLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_base58Encoded(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base58Encoded(tt.args.bytes); got != tt.want {
				t.Errorf("base58Encoded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sha256Of(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sha256Of(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sha256Of() = %v, want %v", got, tt.want)
			}
		})
	}
}
