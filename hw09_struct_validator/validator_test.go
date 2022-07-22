package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code     int    `validate:"in:200,404,500"`
		Body     string `json:"omitempty"`
		UserInfo User   `validate:"nested"`
	}

	Md5sum struct {
		Sum string `validate:"len:32symbol"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:     "012345678901234567890123456789012345",
				Age:    22,
				Email:  "email@address.com",
				Role:   "admin",
				Phones: []string{"89991112233", "89992233111"},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:     "0123456789012345678901234567890123451",
				Age:    222,
				Email:  "emailaddress.com",
				Role:   "adminl",
				Phones: []string{"819991112233", "89992233111"},
			},
			expectedErr: ValidationErrors{
				{
					Field: "ID is 0123456789012345678901234567890123451",
					Err:   ErrorStringLengthIsNotEqual,
				},
				{
					Field: "Age is 222",
					Err:   ErrorIntMoreThanMax,
				},
				{
					Field: "Email is emailaddress.com",
					Err:   ErrorStringRegexpNotMatch,
				},
				{
					Field: "Role is adminl",
					Err:   ErrorStringNotIncludedInSet,
				},
				{
					Field: "Phones is [819991112233 89992233111]",
					Err:   ErrorStringLengthIsNotEqual,
				},
			},
		},
		{
			in: App{
				Version: "0.2.7",
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "0.2.73",
			},
			expectedErr: ValidationErrors{
				{
					Field: "Version is 0.2.73",
					Err:   ErrorStringLengthIsNotEqual,
				},
			},
		},
		{
			in:          Token{},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 200,
				UserInfo: User{
					ID:     "012345678901234567890123456789012345",
					Age:    22,
					Email:  "email@address.com",
					Role:   "admin",
					Phones: []string{"89991112233", "89992233111"},
				},
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 210,
				UserInfo: User{
					ID:     "012345678901234111123567890123456789012345",
					Age:    222,
					Email:  "emailaddress.com",
					Role:   "aewfdmin",
					Phones: []string{"819991112233", "89992233111"},
				},
			},
			expectedErr: ValidationErrors{
				{
					Field: "Code is 210",
					Err:   ErrorIntNotIncludedInSet,
				},
				{
					Field: "UserInfo is {012345678901234111123567890123456789012345  222 emailaddress.com aewfdmin [819991112233 89992233111] []}", //nolint
					Err: ValidationErrors{
						{
							Field: "ID is 012345678901234111123567890123456789012345",
							Err:   ErrorStringLengthIsNotEqual,
						},
						{
							Field: "Age is 222",
							Err:   ErrorIntMoreThanMax,
						},
						{
							Field: "Email is emailaddress.com",
							Err:   ErrorStringRegexpNotMatch,
						},
						{
							Field: "Role is aewfdmin",
							Err:   ErrorStringNotIncludedInSet,
						},
						{
							Field: "Phones is [819991112233 89992233111]",
							Err:   ErrorStringLengthIsNotEqual,
						},
					},
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
			_ = tt
		})
	}
}

func TestValidateProgramError(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: Md5sum{
				Sum: "123",
			},
			expectedErr: strconv.ErrSyntax,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.ErrorIs(t, err, tt.expectedErr)
			_ = tt
		})
	}
}
