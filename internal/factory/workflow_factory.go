package factory

import (
	"context"
	"webhook/internal/models"
	"webhook/internal/workflow"
)

// WorkflowFactory permite crear workflows de dominio a partir de modelos y acciones.
type WorkflowFactory struct{}

// NewWorkflow crea un workflow de dominio a partir de un modelo y acciones.
func (f WorkflowFactory) NewWorkflow(m models.Workflow, actions []func(ctx context.Context, data map[string]interface{}) error) *workflow.Workflow {
	return workflow.NewFromModel(m, actions)
}

// NewStep crea un Step de dominio a partir de un modelo y una acción.
func (f WorkflowFactory) NewStep(model models.Step, action func(ctx context.Context, data map[string]interface{}) error) workflow.Step {
	return workflow.NewStep(model, action)
}
