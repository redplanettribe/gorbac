package gorbac

import (
	"fmt"
)

// RolePermissions is a map that stores the permissions for each role.
// The structure is: role -> action -> resource -> resource
type RolePermissions map[string]map[string]map[string]string

// Permissions represents the permissions configuration for the application.
type Permissions struct {
	roles       RolePermissions
	currentRole string
}

// NewPermissions creates a new Permissions instance with no roles defined.
func NewPermissions() *Permissions {
	return &Permissions{
		roles:       make(RolePermissions),
		currentRole: "",
	}
}

// Roles is a set of role names for efficient lookups.
type Roles map[string]struct{}

// NewRoles creates a new Roles set from a slice of role names.
func NewRoles(roles []string) *Roles {
	r := make(Roles)
	for _, role := range roles {
		r[role] = struct{}{}
	}
	return &r
}

// AddRole adds a new role to the permissions configuration and sets it as the current role.
// If the role already exists, it only sets it as the current role.
func (p *Permissions) AddRole(role string) *Permissions {
	if _, exists := p.roles[role]; !exists {
		p.roles[role] = make(map[string]map[string]string)
	}
	p.currentRole = role
	return p
}

// CustomAction assigns the permission to perform the specified action on the given resources
// to the current role.
func (p *Permissions) CustomAction(action string, resources ...string) *Permissions {
	if p.currentRole == "" {
		fmt.Printf("Role %s does not exist, ignoring %s action \n", p.currentRole, action)
		return p
	}
	if _, exists := p.roles[p.currentRole]; !exists {
		p.AddRole(p.currentRole)
	}
	if _, exists := p.roles[p.currentRole][action]; !exists {
		p.roles[p.currentRole][action] = make(map[string]string, 0)
	}

	for _, resource := range resources {
		if _, exists := p.roles[p.currentRole][action][resource]; !exists {
			p.roles[p.currentRole][action][resource] = resource
		}
	}
	return p
}

// Write assigns the permission to write to the given resources to the current role.
func (p *Permissions) Write(resources ...string) *Permissions {
	return p.CustomAction("write", resources...)
}

// Read assigns the permission to read from the given resources to the current role.
func (p *Permissions) Read(resources ...string) *Permissions {
	return p.CustomAction("read", resources...)
}

// Delete assigns the permission to delete the given resources to the current role.
func (p *Permissions) Delete(resources ...string) *Permissions {
	return p.CustomAction("delete", resources...)
}

// Inherit copies all permissions from the parent role to the current role.
func (p *Permissions) Inherit(parent string) *Permissions {
	if _, exists := p.roles[parent]; !exists {
		fmt.Printf("Role %s does not exist\n", parent)
		return p
	}
	if _, exists := p.roles[p.currentRole]; !exists {
		p.AddRole(p.currentRole)
	}
	for action, resources := range p.roles[parent] {
		if _, exists := p.roles[p.currentRole][action]; !exists {
			p.roles[p.currentRole][action] = make(map[string]string, 0)
		}
		for resource := range resources {
			if _, exists := p.roles[p.currentRole][action][resource]; !exists {
				p.roles[p.currentRole][action][resource] = resource
			}
		}
	}
	return p
}

// HasPermission checks if any of the provided roles have the permission to perform
// the action on the resource.
func (p *Permissions) HasPermission(roles *Roles, action, resource string) bool {
	for role := range *roles {
		if _, exists := p.roles[role]; exists {
			if _, exists := p.roles[role][action]; exists {
				if _, exists := p.roles[role][action][resource]; exists {
					return true
				}
			}
		}
	}
	return false
}
