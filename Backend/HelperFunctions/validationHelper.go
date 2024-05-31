package helperfunctions

import "github.com/vishnumenon/budgetapplication/models"

func validateRole(user models.User) bool {
	if user.Role == "Admin" || user.Role == "User" {
		return true
	}
	return false
}
