// Code generated by candi v1.18.0.

package shared

// this file only for example

import (
	"context"

	"github.com/golangid/candi/candishared"
	"github.com/golangid/candi/logger"
)

// DefaultMiddleware for middleware validator example
type DefaultMiddleware struct {
}

// ValidateToken implement TokenValidator
func (DefaultMiddleware) ValidateToken(ctx context.Context, token string) (*candishared.TokenClaim, error) {
	var tokenClaim candishared.TokenClaim
	tokenClaim.Subject = "USER_ID"

	logger.LogI("validate token: allowed")
	return &tokenClaim, nil
}

// CheckPermission implement interfaces.ACLPermissionChecker
func (DefaultMiddleware) CheckPermission(ctx context.Context, userID string, permissionCode string) (role string, err error) {
	/* add check allow permission for user access (is given "userID" can access "permissionCode" ?)
	if !contains(getAllPermissionFromUser(userID), permissionCode) {
		return role, errors.New("Forbidden")
	}
	*/
	logger.LogIf("check permission: users with id '%s' can access resource with permission code '%s' (return role for this user is 'superadmin')", userID, permissionCode)
	return "superadmin", nil
}
