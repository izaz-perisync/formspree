package service

import (
	"context"
	"fmt"
	"time"
)

func (s *Service) ProjectSetUp(ctx context.Context, p Project) error {
	if p.Name == "" {
		return ErrMissingName
	}

	if p.Type == "" {
		return ErrMissingProjectType
	}

	t := time.Now()

	if p.Id == 0 {

		if err := s.db.QueryRow(`INSERT INTO public.projects (project_name, project_type, created_at, updated_at, created_by) VALUES($1, $2, $3, $4,1) returning id;`, p.Name, p.Type, t, t).Scan(&p.Id); err != nil {
			fmt.Println("insert err", err)
			return ErrNotInserted
		}
	} else {

		r, err := s.db.Exec(`UPDATE public.projects SET  project_name=$1, updated_at=$2,timezone=$3,restricted_domain=$4 where id=$5;`, p.Name, t, p.TimeZone, p.Domain, p.Id)
		if err != nil {
			fmt.Println("update err", err)
			return ErrUpdate
		}

		c, _ := r.RowsAffected()
		if c == 0 {
			return ErrUpdate
		}
	}

	return nil

}

func (s *Service) DeleteProject(ctx context.Context, f Filter) error {
	if f.Id == 0 {
		return ErrMissingId
	}
	r, err := s.db.Exec(`DELETE FROM public.projects WHERE id=$1 and created_by=$2`, f.Id, 1)
	if err != nil {
		return ErrDelete
	}
	c, _ := r.RowsAffected()
	if c == 0 {
		return ErrDelete
	}

	return nil
}


