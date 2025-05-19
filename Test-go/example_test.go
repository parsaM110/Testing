package test

import (
	"testing"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/ozontech/allure-go/pkg/framework/runner"
)

func TestRunner(t *testing.T) {
	runner.Run(t, "My first test", func(t provider.T) {
		t.NewStep("My First Step!")
	})
	runner.Run(t, "My second test", func(t provider.T) {
		t.WithNewStep("My Second Step!", func(sCtx provider.StepCtx) {
			sCtx.NewStep("My First SubStep!")
		})
	})
}
