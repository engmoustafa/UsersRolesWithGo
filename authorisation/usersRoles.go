package authorisation

import (
	"errors"
	"fmt"
)

const (
	userRoleNotFoundMsg = "could not find associated role"
	userDoesNotExistMsg = "user does not exist"
	roleDoesNotExistMsg = "role does not exist"
)

// RoleHierarchyRepo define main repo that manages the UserRole relationships
type RoleHierarchyRepo struct {
	rolesMap map[int]*Role
	usersMap map[int64]*User
	rolesUsersMap map[int][]*User
}

func (r *RoleHierarchyRepo) setUsersList(users []User) (p *[]ProcessingResult){
	processingResult := make([]ProcessingResult,0)
	//Clear the current users Map
	r.usersMap = make(map[int64]*User)

	var recordResult ProcessingResult
	for i:= 0; i< len(users); i++{
		//If user's role is defined, then add the users to the role's list
		if currentRole, ok := r.rolesMap[users[i].roleID];ok {
			//Add user to user's map
			r.usersMap[users[i].id] = &users[i]
			//Retrieve the user's list associated with the role
			if usersList,listExists := r.rolesUsersMap[currentRole.id]; listExists {
				//Append the current user to the list
				usersList = append(usersList,&users[i])
				//TODO reset users list to role
				r.rolesUsersMap[currentRole.id] = usersList
			} else {
				usersList := make([]*User,1)
				usersList[0] = &users[i]
				r.rolesUsersMap[currentRole.id] = usersList
			}

		} else {
			recordResult = ProcessingResult{itemIndex: i, message: String(userRoleNotFoundMsg)}
			processingResult = append(processingResult,recordResult)
		}


	}
	return &processingResult
}

func (r *RoleHierarchyRepo) setRolesList(roles []Role) (p *[]ProcessingResult){
	processingResult := make([]ProcessingResult,0)
	r.rolesMap = make(map[int]*Role)
	r.rolesUsersMap = make(map[int][]*User)
	var parentRole *Role
	for i:= 0; i< len(roles); i++{
		r.rolesMap[roles[i].id] = &roles[i]

		if roles[i].parentID > 0 {
			parentRole = r.rolesMap[roles[i].parentID]
			if !contains(parentRole.childRoles, roles[i].id) {
				parentRole.childRoles = append(parentRole.childRoles, roles[i].id)
			}
		}
	}
	return &processingResult
}

func (r *RoleHierarchyRepo) getSubOrdinates(userID int64) ([]*User, error) {
	ordinatesUsers := make([]*User,0)
	//Retrieve the user for usersMap
	currentUser, userOk := r.usersMap[userID]
	if userOk{
		//Using the User's Role, get the Parent Role
		currentRole, roleOk := r.rolesMap[currentUser.roleID]
		if roleOk {
			// Traverse up the heirarcy levels to the top role, in each level, collect the users for rolesUsersMao
			ordinatesUsers,_ = r.loadSubOrdinatesForRole(currentRole)

		} else {
			return nil, errors.New(roleDoesNotExistMsg + " " + fmt.Sprint(currentUser.roleID))
		}
	} else {
		return nil, errors.New(userDoesNotExistMsg+ " " + fmt.Sprint(userID))
	}

	return ordinatesUsers, nil
}

func (r *RoleHierarchyRepo) loadSubOrdinatesForRole(currentRole *Role) ([]*User, error) {
	childRoles := currentRole.childRoles
	var childRoleUsers []*User = make([]*User,0)
	var subOrdinatesForRole []*User
	for i:= 0; i< len(childRoles); i++ {
		childRoleUsers =  append(childRoleUsers, r.rolesUsersMap[ childRoles[i]]...)
		subOrdinatesForRole,_ = r.loadSubOrdinatesForRole(r.rolesMap[childRoles[i]])
		childRoleUsers = append(childRoleUsers, subOrdinatesForRole...)
	}
	return childRoleUsers, nil
}

func (r *RoleHierarchyRepo) getBosses(userID int64) ([]*User, error){
	bossesUsers := make([]*User,0)
	//Retrieve the user for usersMap
	currentUser, userOk := r.usersMap[userID]
	if userOk{
		//Using the User's Role, get the Parent Role
		currentRole, roleOk := r.rolesMap[currentUser.roleID]
		if roleOk {
			// Traverse up the heirarcy levels to the top role, in each level, collect the users for rolesUsersMao
			parentRoleID := currentRole.parentID
			for {
				if parentRoleID == 0 {
					break
				}
				currentRole = r.rolesMap[parentRoleID]
				//Add all users in next level
				bossesUsers = append(bossesUsers,r.rolesUsersMap[parentRoleID]...)
				parentRoleID = currentRole.parentID
			}
		} else {
			return nil, errors.New(roleDoesNotExistMsg + " " + fmt.Sprint(currentUser.roleID))
		}
	} else {
		return nil, errors.New(userDoesNotExistMsg+ " " + fmt.Sprint(userID))
	}
	return bossesUsers, nil
}
