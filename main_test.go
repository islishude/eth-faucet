package main

import "testing"

func Test_checkEthAddress(t *testing.T) {
	tests := []struct {
		name    string
		address string
		want    bool
	}{
		{name: "UpperCase", address: "0x23a27adea399f2F2Fc69c49aB45acC00378e4EFA", want: true},
		{name: "LowerCase", address: "0x23a27adea399f2f2fc69c49ab45acc00378e4efa", want: true},
		{name: "invalid length", address: "0x23a27adea399f2F2Fc69c49aB45acC00378e4E", want: false},
		{name: "invalid letters", address: "0x23a27adea399f2Fxyz49aB45acC00378e4E", want: false},
		{name: "without 0x prefix", address: "23a27adea399f2F2Fc69c49aB45acC00378e4EFA", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEthAddress(tt.address); got != tt.want {
				t.Errorf("checkEthAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
