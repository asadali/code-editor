package code

const (
	//keys
	SECRET_BIN_KEY 			= "#"
	EXISTS_KEY				= "EXIST"
	TRIBS_KEY				= "Tribs"
	FOLLOWING_KEY			= "Following"
	FOLLOWERS_KEY			= "Followers"

	PRIMARY_USERS_KEY		= "PRIMARY"
	REPLICA_ONE_USERS_KEY	= "REPLICA_ONE"
	REPLICA_TWO_USERS_KEY	= "REPLICA_TWO"

	//Strings
	USER_LIST				= "UserList"
	DOC_LIST				= "DocList"

	//numbers
	MIN_DOCS				= 20
)
var USER_LIST_KEYS = [...]string{
								"Tribs",
								"Following",
								"Followers",
								"UserList",
}
var SPECIAL_KEYS = [...]string{
						PRIMARY_USERS_KEY,
						REPLICA_ONE_USERS_KEY,
						REPLICA_TWO_USERS_KEY,
					}
