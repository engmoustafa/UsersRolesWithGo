package authorisation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func rolesSetup()[]Role{
	return []Role{
		Role{id: 1, description: "System Administrator", parentID: 0},
		Role{id: 2, description: "Location Manasger", parentID: 1},
		Role{id: 3, description: "Supervisor", parentID: 2},
		Role{id: 4, description: "Employee", parentID: 3},
		Role{id: 5, description: "Trainer", parentID: 3},
	}
}
func usersSetup()[]User{
	return []User {
		User{id:1, name: "Adam Admin", roleID: 1},
		User{id:2, name: "Emily Employee", roleID: 4},
		User{id:3, name: "Sam Supervisor", roleID: 3},
		User{id:4, name: "Mary Manager", roleID: 2},
		User{id:5, name: "Steve Trainer", roleID: 5},
	}
}

func TestRoleHierarchyRepo_getSubOrdinates_ShouldReturnTwoUsers(t *testing.T) {
	//Arrange
	rolesList := rolesSetup()
	usersList := usersSetup()
	usersRepo := RoleHierarchyRepo{}

	//Act
	usersRepo.setRolesList(rolesList)
	usersRepo.setUsersList(usersList)

	ordinatesUsersList,error := usersRepo.getSubOrdinates(3)

	//Assert
	assert.Nil(t, error)
	assert.Equal(t, 2,len(ordinatesUsersList))
	assert.Contains(t, ordinatesUsersList,&usersList[1])
	assert.Contains(t, ordinatesUsersList,&usersList[4])
}

func TestRoleHierarchyRepo_getSubOrdinates_ShouldReturnFourUsers(t *testing.T) {
	//Arrange
	rolesList := rolesSetup()
	usersList := usersSetup()
	usersRepo := RoleHierarchyRepo{}

	//Act
	usersRepo.setRolesList(rolesList)
	usersRepo.setUsersList(usersList)

	ordinatesUsersList,error := usersRepo.getSubOrdinates(1)

	//Assert
	assert.Nil(t, error)
	assert.Equal(t, 4,len(ordinatesUsersList))
	assert.Contains(t, ordinatesUsersList,&usersList[1])
	assert.Contains(t, ordinatesUsersList,&usersList[2])
	assert.Contains(t, ordinatesUsersList,&usersList[3])
	assert.Contains(t, ordinatesUsersList,&usersList[4])
}

func TestRoleHierarchyRepo_getSubOrdinates_ShouldNotReturnAny(t *testing.T) {
	//Arrange
	rolesList := rolesSetup()
	usersList := usersSetup()
	usersRepo := RoleHierarchyRepo{}

	//Act
	usersRepo.setRolesList(rolesList)
	usersRepo.setUsersList(usersList)

	ordinatesUsersList,error := usersRepo.getSubOrdinates(5)

	//Assert
	assert.Nil(t, error)
	assert.Equal(t, 0,len(ordinatesUsersList))
}

func TestRoleHierarchyRepo_getBosses_ShouldReturnThreeUsers(t *testing.T) {
	//Arrange
	rolesList := rolesSetup()
	usersList := usersSetup()
	usersRepo := RoleHierarchyRepo{}

	//Act
	usersRepo.setRolesList(rolesList)
	usersRepo.setUsersList(usersList)

	ordinatesUsersList,error := usersRepo.getBosses(5)

	//Assert
	assert.Nil(t, error)
	assert.Equal(t, 3,len(ordinatesUsersList))
	assert.Contains(t, ordinatesUsersList,&usersList[0])
	assert.Contains(t, ordinatesUsersList,&usersList[2])
	assert.Contains(t, ordinatesUsersList,&usersList[3])
}

func TestRoleHierarchyRepo_getBosses_ShouldNotReturnAny(t *testing.T) {
	//Arrange
	rolesList := rolesSetup()
	usersList := usersSetup()
	usersRepo := RoleHierarchyRepo{}

	//Act
	usersRepo.setRolesList(rolesList)
	usersRepo.setUsersList(usersList)

	ordinatesUsersList,error := usersRepo.getBosses(1)

	//Assert
	assert.Nil(t, error)
	assert.Equal(t, 0,len(ordinatesUsersList))
}