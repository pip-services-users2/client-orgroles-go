package version1

import (
	"encoding/json"

	"github.com/pip-services-users2/client-orgroles-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := make(map[string]string, 0)

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	r := make(map[string]any, 0)

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func toUserRoles(obj []*protos.UserRoles) []*UserRolesV1 {
	if obj == nil {
		return nil
	}

	result := make([]*UserRolesV1, len(obj))

	for _, item := range obj {
		role := NewUserRolesV1(item.Id, item.Roles, convert.DateTimeConverter.ToDateTime(item.UpdateTime))
		result = append(result, role)
	}

	return result
}

func fromUserRoles(roles []*UserRolesV1) []*protos.UserRoles {
	if roles == nil {
		return nil
	}

	result := make([]*protos.UserRoles, len(roles))

	for _, item := range roles {
		role := &protos.UserRoles{
			Id:         item.Id,
			Roles:      item.Roles,
			UpdateTime: convert.StringConverter.ToString(item.UpdateTime),
		}
		result = append(result, role)
	}

	return result
}
