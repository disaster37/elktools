package elktools_kibana

import (
	"github.com/stretchr/testify/assert"
)

func (s *KBTestSuite) TestExportImportDashboards() {

	// Import all dashboards from fixtures
	err := importDashboards("../fixtures/saved_object/export.ndjson", "default", s.client)
	assert.NoError(s.T(), err)

	// Export all dashboard
	err = exportDashboards("/tmp/export.ndjson", "default", s.client)
	assert.NoError(s.T(), err)

}
