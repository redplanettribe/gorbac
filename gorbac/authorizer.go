package gorbac

import (
	"errors"
	"strings"
)

var (
	ErrEmptyPermission  = errors.New("permission string is empty")
	ErrPermissionDenied = errors.New("permission denied")
)

type Authorizer interface {
	// Authorize checks if the roles have the permission to perform the action on the resource.
	// actionResource is a string in the format "action:resource".
	// Returns true if the roles have the permission, false otherwise.
	Authorize(roles []string, actionResource string) (bool, error)
}
type authorizer struct {
	permissions *Permissions
}

// NewAuthorizer creates a new Authorizer with the given Permissions. The Authorizer checks if the roles have the permission to perform the action on the resource. Generate the permissions with NewPermissions.
func NewAuthorizer(p *Permissions) Authorizer {
	return authorizer{permissions: p}
}

func (a authorizer) Authorize(roles []string, actionResource string) (bool, error) {
	roleMap := NewRoles(roles)
	action, resource := parsePermission(actionResource)
	if action == "" || resource == "" {
		return false, ErrEmptyPermission
	}
	hasPermission := a.permissions.HasPermission(roleMap, action, resource)
	if !hasPermission {
		return false, ErrPermissionDenied
	}
	return true, nil
}

func parsePermission(permission string) (action, resource string) {
	parts := strings.Split(permission, ":")
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}
