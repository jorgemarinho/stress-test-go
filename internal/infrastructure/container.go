package infrastructure

import (
	"github.com/jorgemarinho/stress-test-go/internal/domain"
	"github.com/jorgemarinho/stress-test-go/internal/service"
)

type Container struct {
	loadTester      domain.LoadTester
	reportGenerator domain.ReportGenerator
	cli             CLI
}

func NewContainer() *Container {
	container := &Container{
		loadTester:      service.NewLoadTesterService(),
		reportGenerator: service.NewReportService(),
	}
	container.cli = NewCLI(container.loadTester, container.reportGenerator)
	return container
}

func (c *Container) GetLoadTester() domain.LoadTester {
	return c.loadTester
}

func (c *Container) GetReportGenerator() domain.ReportGenerator {
	return c.reportGenerator
}

func (c *Container) GetCLI() CLI {
	return c.cli
}
