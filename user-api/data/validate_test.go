package data

import (
	"testing"
)

func TestCheckSex(t *testing.T) {
	cases := []struct {
		user User
		name string
		want error
	}{
		{
			user: User{
				Sex: 0,
			},
			want: nil,
			name: "sex_0",
		},
		{
			user: User{
				Sex: 1,
			},
			want: nil,
			name: "sex_1",
		},
		{
			user: User{
				Sex: 2,
			},
			want: nil,
			name: "sex_2",
		},
		{
			user: User{
				Sex: 3,
			},
			want: nil,
			name: "sex_3",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.user.CheckSex()
			if tc.want != got {
				t.Errorf("Want '%v', got '%v'", tc.want, got)
			}
		})
	}
}

func TestCheckMail(t *testing.T) {
	cases := []struct {
		user User
		name string
		want interface{}
	}{
		{
			user: User{
				Mail: "abc@gmail.com",
			},
			name: "mail_abc@gmail.com",
			want: nil,
		},
		{
			user: User{
				Mail: "abc",
			},
			name: "mail_abc",
			want: "Mail Invalid",
		},
		{
			user: User{
				Mail: "abc@",
			},
			name: "mail_abc@",
			want: "Mail Invalid",
		},
		{
			user: User{
				Mail: "abc.gmail.com",
			},
			name: "mail_abc.gmail.com",
			want: "Mail Invalid",
		},
		{
			user: User{
				Mail: "_@_",
			},
			name: "mail_@",
			want: "Mail Invalid",
		},
		{
			user: User{
				Mail: "",
			},
			name: "mail_''",
			want: "Empty Mail",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.CheckMail()
			if err == nil && tc.want != err {
				t.Errorf("Want '%v', got '%v'", tc.want, err)
			}
			if err != nil && tc.want != err.Error() {
				t.Errorf("Want '%v', got '%v'", tc.want, err.Error())
			}
		})
	}
}

func TestCheckUserName(t *testing.T) {
	cases := []struct {
		user User
		name string
		want interface{}
	}{
		{
			user: User{
				UserName: "chloeeee",
			},
			name: "n_chloeeee",
			want: nil,
		},
		{
			user: User{
				UserName: "CHLOEEEE",
			},
			name: "n_CHLOEEEE",
			want: nil,
		},
		{
			user: User{
				UserName: "12345_678",
			},
			name: "n_12345_678",
			want: ErrorUesrName,
		},
		{
			user: User{
				UserName: "Aabc-123",
			},
			name: "n__Aabc-123",
			want: ErrorUesrName,
		},
		{
			user: User{
				UserName: "_@_",
			},
			name: "n_@",
			want: ErrorUesrName,
		},
		{
			user: User{
				UserName: "12345",
			},
			name: "n_12345",
			want: ErrorUesrName,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.CheckUserName()
			if err == nil && tc.want != err {
				t.Errorf("Want '%v', got '%v'", tc.want, err)
			}
			if err != nil && tc.want != err.Error() {
				t.Errorf("Want '%v', got '%v'", tc.want, err.Error())
			}
		})
	}
}

func TestCheckPassword(t *testing.T) {
	cases := []struct {
		user User
		name string
		want interface{}
	}{
		{
			user: User{
				Password: "Chloe",
			},
			name: "p_Chloe",
			want: ErrorPassword["len"],
		},
		{
			user: User{
				Password: "CHLOe_!@><(12345678)-_-_-=",
			},
			name: "p_CHLOe_!@><(12345678)-_-_-=",
			want: nil,
		},
		{
			user: User{
				Password: "12345!@678",
			},
			name: "p_12345!@678",
			want: ErrorPassword["lower"],
		},
		{
			user: User{
				Password: "!^abc$_123",
			},
			name: "p_!^abc$_123",
			want: ErrorPassword["upper"],
		},
		{
			user: User{
				Password: "_@_",
			},
			name: "p_@_",
			want: ErrorPassword["len"],
		},
		{
			user: User{
				Password: "12345Abc",
			},
			name: "p_12345",
			want: ErrorPassword["symbol"],
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.CheckPassword()

			if err == nil && tc.want != err {
				t.Errorf("Want '%v', got '%v'", tc.want, err)
			}
			if err != nil && tc.want != err.Error() {
				t.Errorf("Want '%v', got '%v'", tc.want, err.Error())
			}
		})
	}
}
