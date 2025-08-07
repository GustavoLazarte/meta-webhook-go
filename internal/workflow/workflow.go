package workflow

import (
	"context"
	"errors"
	"sync"
	"webhook/internal/models"
)

// Step representa un paso de un workflow, extendido con lógica de dominio.
type Step struct {
	models.Step
	Action func(ctx context.Context, data map[string]interface{}) error
}

// Workflow representa un flujo de trabajo robusto y extensible.
type Workflow struct {
	models.Workflow
	Steps   []Step
	State   map[string]interface{}
	mu      sync.Mutex
	current int
}

// New crea un nuevo workflow de dominio a partir de datos y pasos.
func New(id, name string, steps []Step) *Workflow {
	return &Workflow{
		models.Workflow{
			ID:   id,
			Name: name,
		},
		steps,
		make(map[string]interface{}),
		sync.Mutex{},
		-1,
	}
}

// NewFromModel crea un workflow de dominio a partir de un modelo puro y acciones.
func NewFromModel(m models.Workflow, actions []func(ctx context.Context, data map[string]interface{}) error) *Workflow {
	var steps []Step
	for i, s := range m.Steps {
		var action func(ctx context.Context, data map[string]interface{}) error
		if i < len(actions) {
			action = actions[i]
		}
		steps = append(steps, Step{Step: s, Action: action})
	}
	return &Workflow{
		models.Workflow{
			ID:    m.ID,
			Name:  m.Name,
			Steps: m.Steps,
		},
		steps,
		make(map[string]interface{}),
		sync.Mutex{},
		-1,
	}
}

// NewStep crea un Step de dominio a partir de un modelo y una acción.
func NewStep(model models.Step, action func(ctx context.Context, data map[string]interface{}) error) Step {
	return Step{
		Step:   model,
		Action: action,
	}
}

// AddStep agrega un paso al workflow.
func (w *Workflow) AddStep(step Step) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Steps = append(w.Steps, step)
}

// NextStep avanza al siguiente paso y lo ejecuta.
func (w *Workflow) NextStep(ctx context.Context) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.current+1 >= len(w.Steps) {
		return errors.New("no hay más pasos en el workflow")
	}
	w.current++
	step := w.Steps[w.current]
	if step.Action != nil {
		return step.Action(ctx, w.State)
	}
	return nil
}

// CurrentStep devuelve el paso actual.
func (w *Workflow) CurrentStep() (Step, bool) {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.current >= 0 && w.current < len(w.Steps) {
		return w.Steps[w.current], true
	}
	return Step{}, false
}

// Reset reinicia el workflow al estado inicial.
func (w *Workflow) Reset() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.current = -1
	w.State = make(map[string]interface{})
}
