package casbin

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
)

var cb *casbin.Enforcer

func InitFromSQLLite(db *gorm.DB, rbac_path_str string) {

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		fmt.Println(err.Error())
	}

	cb, err = casbin.NewEnforcer(rbac_path_str, a)
	//cb.AddFunction("keyMatch2", util.KeyMatch2Func)

	if err != nil {
		fmt.Println(err.Error())
	}
	err = cb.LoadPolicy()
	cb.EnableAutoSave(true)

	PreDefineRoleAndPolicy()
}

func PreDefineRoleAndPolicy() {
	_, _ = cb.AddPolicy("MEMBER_PERSONA_USER", "/v1/tasks", "/task.v1.TaskService/GetTasks")
	_, _ = cb.AddPolicy("MEMBER_PERSONA_USER", "/v1/tasks/{id}", "/task.v1.TaskService/GetTaskByID")
	_, _ = cb.AddPolicy("MEMBER_PERSONA_USER", "/v1/tasks", "/task.v1.TaskService/CreateTask")
	_, _ = cb.AddPolicy("MEMBER_PERSONA_USER", "/v1/tasks/{id}", "/task.v1.TaskService/UpdateTask")
	_, _ = cb.AddPolicy("MEMBER_PERSONA_USER", "/v1/tasks/{id}", "/task.v1.TaskService/DeleteTaskByID")
}

func GetCasbin() *casbin.Enforcer {
	return cb
}

func AddRole(role string, policy string) (bool, error) {
	return cb.AddRoleForUser(role, policy)
}
