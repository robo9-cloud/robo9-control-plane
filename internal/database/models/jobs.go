package models

import (
	"errors"
)

var (
	ErrorInvalidTaskType = errors.New("invalid task type")
)

type taskTypeEnum int64

const (
	TaskPriorityHighest   = 10
	TaskPriorityHigh      = 8
	TasksPriorityStandard = 5
	TasksPriorityLow      = 3
	TasksPriorityLowest   = 1

	TASK_TYPE_TERRAFORM_PLAN taskTypeEnum = iota + 1
	TASK_TYPE_TERRAFORM_APPLY
	TASK_TYPE_ANSIBLE
	TASK_TYPE_PACKER
	TASK_TYPE_CHEF
)

var taskTypes = map[string]taskTypeEnum{
	"TERRAFORM_PLAN_TASK":  TASK_TYPE_TERRAFORM_PLAN,
	"TERRAFORM_APPLY_TASK": TASK_TYPE_TERRAFORM_APPLY,
	"ANSIBLE_TASK":         TASK_TYPE_ANSIBLE,
	"PACKER_TASK":          TASK_TYPE_PACKER,
	"CHEF_TASK":            TASK_TYPE_CHEF,
}

func (j taskTypeEnum) String() string {
	switch j {
	case TASK_TYPE_TERRAFORM_PLAN:
		return "TERRAFORM_PLAN_TASK"
	case TASK_TYPE_TERRAFORM_APPLY:
		return "TERRAFORM_APPLY_TASK"
	case TASK_TYPE_ANSIBLE:
		return "ANSIBLE_TASK"
	case TASK_TYPE_PACKER:
		return "PACKER_TASK"
	case TASK_TYPE_CHEF:
		return "CHEF_TASK"
	default:
		return ""

	}
}
