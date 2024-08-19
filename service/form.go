package service

import (
	"context"
	"fmt"
	"math/rand"
	"net/mail"
	"time"
)

func (s *Service) CreatedForm(ctx context.Context, f Form) error {
	if f.Name == "" {
		return ErrMissingName
	}

	if f.Project == 0 {
		return ErrMissingId
	}

	_, err := mail.ParseAddress(string(f.LinkedEmail))
	if err != nil {
		return ErrInvalidEmail
	}
	t := time.Now()
	f.Key = RandomString(8, false)
	r, err := s.db.Exec(`INSERT INTO public.forms (project_id, created_at, updated_at, user_id, "name", linked_email_id, form_key) VALUES($1, $2, $3, 1, $4, $5, $6);`, f.Project, t, t, f.Name, f.LinkedEmail, f.Key)
	if err != nil {
		fmt.Println("update err", err)
		return ErrUpdate
	}

	c, _ := r.RowsAffected()
	if c == 0 {
		return ErrUpdate
	}

	return nil
}

func (s *Service) ToggleForm(ctx context.Context, f Filter, reqType string) error {
	if f.Id == 0 {
		return ErrMissingId
	}

	query := ""

	switch reqType {
	case "active":
		query = `update public.forms  set active=$1 where id=$2`
	case "submission":
		query = `update public.forms  set store_submissions=$1 where id=$2`
	case "reCAPTCHA":
		query = `update public.forms  set reCAPTCHA=$1 where id=$2`
	case "formName":

		query = `update public.forms  set name=$1 where id=$2`

	default:
		return Generic{Msg: "invalid req type"}
	}
	r, err := s.db.Exec(query, f)
	if err != nil {
		fmt.Println("update err", err)
		return ErrUpdate
	}

	c, _ := r.RowsAffected()
	if c == 0 {
		return ErrUpdate
	}
	return nil
}

func RandomString(length int, withSymbols bool) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	var charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789"

	if withSymbols {
		charset += "@!-_^#$%"
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
