package data

import (
	"testing"
)

func TestCheckSex(t *testing.T) {
	cases := []struct {
		user User
		name string
		want bool
	}{
		{
			user: User{
				Sex: 0,
			},
			want: true,
			name: "sex_0",
		},
		{
			user: User{
				Sex: 1,
			},
			want: true,
			name: "sex_1",
		},
		{
			user: User{
				Sex: 2,
			},
			want: true,
			name: "sex_2",
		},
		{
			user: User{
				Sex: 3,
			},
			want: false,
			name: "sex_3",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.CheckSex()

			if tc.want != got {
				t.Errorf("Want '%t', got '%t'", tc.want, got)
			}
		})
	}
}

func TestCheckMail(t *testing.T) {
	cases := []struct {
		user User
		name string
		want bool
	}{
		{
			user: User{
				Mail: "abc@gmail.com",
			},
			name: "mail_abc@gmail.com",
			want: true,
		},
		{
			user: User{
				Mail: "abc",
			},
			name: "mail_abc",
			want: false,
		},
		{
			user: User{
				Mail: "abc@",
			},
			name: "mail_abc@",
			want: false,
		},
		{
			user: User{
				Mail: "abc.gmail.com",
			},
			name: "mail_abc.gmail.com",
			want: false,
		},
		{
			user: User{
				Mail: "_@_",
			},
			name: "mail_@",
			want: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.CheckMail()

			if tc.want != got {
				t.Errorf("Want '%t', got '%t'", tc.want, got)
			}
		})
	}
}

func TestCheckUserName(t *testing.T) {
	cases := []struct {
		user User
		name string
		want bool
	}{
		{
			user: User{
				UserName: "chloeeee",
			},
			name: "n_chloeeee",
			want: true,
		},
		{
			user: User{
				UserName: "CHLOEEEE",
			},
			name: "n_CHLOEEEE",
			want: true,
		},
		{
			user: User{
				UserName: "12345!@678",
			},
			name: "n_12345!@678",
			want: true,
		},
		{
			user: User{
				UserName: "!^abc$_123",
			},
			name: "n_!^abc$_123",
			want: true,
		},
		{
			user: User{
				UserName: "_@_",
			},
			name: "n_@",
			want: false,
		},
		{
			user: User{
				UserName: "12345",
			},
			name: "n_12345",
			want: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.CheckUserName()

			if tc.want != got {
				t.Errorf("Want '%t', got '%t'", tc.want, got)
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	cases := []struct {
		user User
		name string
		want bool
	}{
		{
			user: User{
				Password: "chloeeee",
			},
			name: "p_chloeeee",
			want: true,
		},
		{
			user: User{
				Password: "CHLOE_!@><(12345678)-_-_-=",
			},
			name: "p_CHLOE_!@><(12345678)-_-_-=",
			want: true,
		},
		{
			user: User{
				Password: "12345!@678",
			},
			name: "p_12345!@678",
			want: true,
		},
		{
			user: User{
				Password: "!^abc$_123",
			},
			name: "p_!^abc$_123",
			want: true,
		},
		{
			user: User{
				Password: "_@_",
			},
			name: "p_@",
			want: false,
		},
		{
			user: User{
				Password: "12345",
			},
			name: "p_12345",
			want: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.CheckPassword()

			if tc.want != got {
				t.Errorf("Want '%t', got '%t'", tc.want, got)
			}
		})
	}
}
