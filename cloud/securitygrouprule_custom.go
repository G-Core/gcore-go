package cloud

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
)

// NewAndPoll creates a security group rule and polls for completion of the task.
// After the task completes, it fetches the parent security group and returns the created rule.
func (r *SecurityGroupRuleService) NewAndPoll(ctx context.Context, groupID string, params SecurityGroupRuleNewParams, opts ...option.RequestOption) (res *SecurityGroupRule, err error) {
	resource, err := r.New(ctx, groupID, params, opts...)
	if err != nil {
		return
	}

	opts = slices.Concat(r.Options, opts)
	precfg, err := requestconfig.PreRequestOptions(opts...)
	if err != nil {
		return
	}
	var getParams SecurityGroupGetParams
	requestconfig.UseDefaultParam(&params.ProjectID, precfg.CloudProjectID)
	requestconfig.UseDefaultParam(&params.RegionID, precfg.CloudRegionID)
	getParams.ProjectID = params.ProjectID
	getParams.RegionID = params.RegionID

	if len(resource.Tasks) != 1 {
		return nil, errors.New("expected exactly one task to be created")
	}
	taskID := resource.Tasks[0]

	tasks := NewTaskService(r.Options...)
	task, err := tasks.Poll(ctx, taskID, opts...)
	if err != nil {
		return
	}

	if !task.JSON.CreatedResources.Valid() || len(task.CreatedResources.SecurityGroupRules) != 1 {
		return nil, errors.New("expected exactly one security group rule to be created in a task")
	}
	ruleID := task.CreatedResources.SecurityGroupRules[0]

	sgService := NewSecurityGroupService(opts...)
	sg, err := sgService.Get(ctx, groupID, getParams, opts...)
	if err != nil {
		return
	}

	for _, rule := range sg.SecurityGroupRules {
		if rule.ID == ruleID {
			return &rule, nil
		}
	}

	return nil, fmt.Errorf("rule %s not found in security group %s after creation", ruleID, groupID)
}

// DeleteAndPoll deletes a security group rule and polls for completion.
func (r *SecurityGroupRuleService) DeleteAndPoll(ctx context.Context, ruleID string, params SecurityGroupRuleDeleteParams, opts ...option.RequestOption) error {
	resource, err := r.Delete(ctx, ruleID, params, opts...)
	if err != nil {
		return err
	}

	opts = slices.Concat(r.Options, opts)
	if len(resource.Tasks) == 0 {
		return errors.New("expected at least one task to be created")
	}
	taskID := resource.Tasks[0]

	tasks := NewTaskService(r.Options...)
	_, err = tasks.Poll(ctx, taskID, opts...)
	return err
}
