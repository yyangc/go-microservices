package data

import (
	"testing"
)

func TestCheckSex(t *testing.T) {
	cases := []struct {
		sex  uint8
		name string
		want bool
	}{
		{
			sex:  0,
			want: true,
			name: "sex_0",
		},
		{
			sex:  1,
			want: true,
			name: "sex_1",
		},
		{
			sex:  2,
			want: true,
			name: "sex_2",
		},
		{
			sex:  3,
			want: false,
			name: "sex_3",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := checkSex(&tc.sex)

			if tc.want != got {
				t.Errorf("Want '%t', got '%t'", tc.want, got)
			}
		})
	}
}

func TestIsEmailValid(t *testing.T) {
	cases := []struct {
		mail string
		name string
		want bool
	}{
		{
			mail: "abc@gmail.com",
			name: "mail_abc@gmail.com",
			want: true,
		},
		{
			mail: "abc",
			name: "mail_abc",
			want: false,
		},
		{
			mail: "abc@",
			name: "mail_abc@",
			want: false,
		},
		{
			mail: "abc.gmail.com",
			name: "mail_abc.gmail.com",
			want: false,
		},
		{
			mail: "_@_",
			name: "mail_@",
			want: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isEmailValid(&tc.mail)

			if tc.want != got {
				t.Errorf("Want '%t', got '%t'", tc.want, got)
			}
		})
	}
}
