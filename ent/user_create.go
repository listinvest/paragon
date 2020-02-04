// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/kcarretto/paragon/ent/event"
	"github.com/kcarretto/paragon/ent/job"
	"github.com/kcarretto/paragon/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	Name         *string
	OAuthID      *string
	PhotoURL     *string
	SessionToken *string
	IsActivated  *bool
	IsAdmin      *bool
	jobs         map[int]struct{}
	events       map[int]struct{}
}

// SetName sets the Name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.Name = &s
	return uc
}

// SetOAuthID sets the OAuthID field.
func (uc *UserCreate) SetOAuthID(s string) *UserCreate {
	uc.OAuthID = &s
	return uc
}

// SetPhotoURL sets the PhotoURL field.
func (uc *UserCreate) SetPhotoURL(s string) *UserCreate {
	uc.PhotoURL = &s
	return uc
}

// SetSessionToken sets the SessionToken field.
func (uc *UserCreate) SetSessionToken(s string) *UserCreate {
	uc.SessionToken = &s
	return uc
}

// SetNillableSessionToken sets the SessionToken field if the given value is not nil.
func (uc *UserCreate) SetNillableSessionToken(s *string) *UserCreate {
	if s != nil {
		uc.SetSessionToken(*s)
	}
	return uc
}

// SetIsActivated sets the IsActivated field.
func (uc *UserCreate) SetIsActivated(b bool) *UserCreate {
	uc.IsActivated = &b
	return uc
}

// SetNillableIsActivated sets the IsActivated field if the given value is not nil.
func (uc *UserCreate) SetNillableIsActivated(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsActivated(*b)
	}
	return uc
}

// SetIsAdmin sets the IsAdmin field.
func (uc *UserCreate) SetIsAdmin(b bool) *UserCreate {
	uc.IsAdmin = &b
	return uc
}

// SetNillableIsAdmin sets the IsAdmin field if the given value is not nil.
func (uc *UserCreate) SetNillableIsAdmin(b *bool) *UserCreate {
	if b != nil {
		uc.SetIsAdmin(*b)
	}
	return uc
}

// AddJobIDs adds the jobs edge to Job by ids.
func (uc *UserCreate) AddJobIDs(ids ...int) *UserCreate {
	if uc.jobs == nil {
		uc.jobs = make(map[int]struct{})
	}
	for i := range ids {
		uc.jobs[ids[i]] = struct{}{}
	}
	return uc
}

// AddJobs adds the jobs edges to Job.
func (uc *UserCreate) AddJobs(j ...*Job) *UserCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return uc.AddJobIDs(ids...)
}

// AddEventIDs adds the events edge to Event by ids.
func (uc *UserCreate) AddEventIDs(ids ...int) *UserCreate {
	if uc.events == nil {
		uc.events = make(map[int]struct{})
	}
	for i := range ids {
		uc.events[ids[i]] = struct{}{}
	}
	return uc
}

// AddEvents adds the events edges to Event.
func (uc *UserCreate) AddEvents(e ...*Event) *UserCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddEventIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.Name == nil {
		return nil, errors.New("ent: missing required field \"Name\"")
	}
	if err := user.NameValidator(*uc.Name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"Name\": %v", err)
	}
	if uc.OAuthID == nil {
		return nil, errors.New("ent: missing required field \"OAuthID\"")
	}
	if uc.PhotoURL == nil {
		return nil, errors.New("ent: missing required field \"PhotoURL\"")
	}
	if uc.SessionToken != nil {
		if err := user.SessionTokenValidator(*uc.SessionToken); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"SessionToken\": %v", err)
		}
	}
	if uc.IsActivated == nil {
		v := user.DefaultIsActivated
		uc.IsActivated = &v
	}
	if uc.IsAdmin == nil {
		v := user.DefaultIsAdmin
		uc.IsAdmin = &v
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value := uc.Name; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldName,
		})
		u.Name = *value
	}
	if value := uc.OAuthID; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldOAuthID,
		})
		u.OAuthID = *value
	}
	if value := uc.PhotoURL; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldPhotoURL,
		})
		u.PhotoURL = *value
	}
	if value := uc.SessionToken; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldSessionToken,
		})
		u.SessionToken = *value
	}
	if value := uc.IsActivated; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: user.FieldIsActivated,
		})
		u.IsActivated = *value
	}
	if value := uc.IsAdmin; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: user.FieldIsAdmin,
		})
		u.IsAdmin = *value
	}
	if nodes := uc.jobs; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.JobsTable,
			Columns: []string{user.JobsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: job.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := uc.events; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.EventsTable,
			Columns: []string{user.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: event.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
