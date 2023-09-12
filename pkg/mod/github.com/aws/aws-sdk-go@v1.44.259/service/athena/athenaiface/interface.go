// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package athenaiface provides an interface to enable mocking the Amazon Athena service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package athenaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/athena"
)

// AthenaAPI provides an interface to enable mocking the
// athena.Athena service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Athena.
//	func myFunc(svc athenaiface.AthenaAPI) bool {
//	    // Make svc.BatchGetNamedQuery request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := athena.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockAthenaClient struct {
//	    athenaiface.AthenaAPI
//	}
//	func (m *mockAthenaClient) BatchGetNamedQuery(input *athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockAthenaClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type AthenaAPI interface {
	BatchGetNamedQuery(*athena.BatchGetNamedQueryInput) (*athena.BatchGetNamedQueryOutput, error)
	BatchGetNamedQueryWithContext(aws.Context, *athena.BatchGetNamedQueryInput, ...request.Option) (*athena.BatchGetNamedQueryOutput, error)
	BatchGetNamedQueryRequest(*athena.BatchGetNamedQueryInput) (*request.Request, *athena.BatchGetNamedQueryOutput)

	BatchGetPreparedStatement(*athena.BatchGetPreparedStatementInput) (*athena.BatchGetPreparedStatementOutput, error)
	BatchGetPreparedStatementWithContext(aws.Context, *athena.BatchGetPreparedStatementInput, ...request.Option) (*athena.BatchGetPreparedStatementOutput, error)
	BatchGetPreparedStatementRequest(*athena.BatchGetPreparedStatementInput) (*request.Request, *athena.BatchGetPreparedStatementOutput)

	BatchGetQueryExecution(*athena.BatchGetQueryExecutionInput) (*athena.BatchGetQueryExecutionOutput, error)
	BatchGetQueryExecutionWithContext(aws.Context, *athena.BatchGetQueryExecutionInput, ...request.Option) (*athena.BatchGetQueryExecutionOutput, error)
	BatchGetQueryExecutionRequest(*athena.BatchGetQueryExecutionInput) (*request.Request, *athena.BatchGetQueryExecutionOutput)

	CancelCapacityReservation(*athena.CancelCapacityReservationInput) (*athena.CancelCapacityReservationOutput, error)
	CancelCapacityReservationWithContext(aws.Context, *athena.CancelCapacityReservationInput, ...request.Option) (*athena.CancelCapacityReservationOutput, error)
	CancelCapacityReservationRequest(*athena.CancelCapacityReservationInput) (*request.Request, *athena.CancelCapacityReservationOutput)

	CreateCapacityReservation(*athena.CreateCapacityReservationInput) (*athena.CreateCapacityReservationOutput, error)
	CreateCapacityReservationWithContext(aws.Context, *athena.CreateCapacityReservationInput, ...request.Option) (*athena.CreateCapacityReservationOutput, error)
	CreateCapacityReservationRequest(*athena.CreateCapacityReservationInput) (*request.Request, *athena.CreateCapacityReservationOutput)

	CreateDataCatalog(*athena.CreateDataCatalogInput) (*athena.CreateDataCatalogOutput, error)
	CreateDataCatalogWithContext(aws.Context, *athena.CreateDataCatalogInput, ...request.Option) (*athena.CreateDataCatalogOutput, error)
	CreateDataCatalogRequest(*athena.CreateDataCatalogInput) (*request.Request, *athena.CreateDataCatalogOutput)

	CreateNamedQuery(*athena.CreateNamedQueryInput) (*athena.CreateNamedQueryOutput, error)
	CreateNamedQueryWithContext(aws.Context, *athena.CreateNamedQueryInput, ...request.Option) (*athena.CreateNamedQueryOutput, error)
	CreateNamedQueryRequest(*athena.CreateNamedQueryInput) (*request.Request, *athena.CreateNamedQueryOutput)

	CreateNotebook(*athena.CreateNotebookInput) (*athena.CreateNotebookOutput, error)
	CreateNotebookWithContext(aws.Context, *athena.CreateNotebookInput, ...request.Option) (*athena.CreateNotebookOutput, error)
	CreateNotebookRequest(*athena.CreateNotebookInput) (*request.Request, *athena.CreateNotebookOutput)

	CreatePreparedStatement(*athena.CreatePreparedStatementInput) (*athena.CreatePreparedStatementOutput, error)
	CreatePreparedStatementWithContext(aws.Context, *athena.CreatePreparedStatementInput, ...request.Option) (*athena.CreatePreparedStatementOutput, error)
	CreatePreparedStatementRequest(*athena.CreatePreparedStatementInput) (*request.Request, *athena.CreatePreparedStatementOutput)

	CreatePresignedNotebookUrl(*athena.CreatePresignedNotebookUrlInput) (*athena.CreatePresignedNotebookUrlOutput, error)
	CreatePresignedNotebookUrlWithContext(aws.Context, *athena.CreatePresignedNotebookUrlInput, ...request.Option) (*athena.CreatePresignedNotebookUrlOutput, error)
	CreatePresignedNotebookUrlRequest(*athena.CreatePresignedNotebookUrlInput) (*request.Request, *athena.CreatePresignedNotebookUrlOutput)

	CreateWorkGroup(*athena.CreateWorkGroupInput) (*athena.CreateWorkGroupOutput, error)
	CreateWorkGroupWithContext(aws.Context, *athena.CreateWorkGroupInput, ...request.Option) (*athena.CreateWorkGroupOutput, error)
	CreateWorkGroupRequest(*athena.CreateWorkGroupInput) (*request.Request, *athena.CreateWorkGroupOutput)

	DeleteDataCatalog(*athena.DeleteDataCatalogInput) (*athena.DeleteDataCatalogOutput, error)
	DeleteDataCatalogWithContext(aws.Context, *athena.DeleteDataCatalogInput, ...request.Option) (*athena.DeleteDataCatalogOutput, error)
	DeleteDataCatalogRequest(*athena.DeleteDataCatalogInput) (*request.Request, *athena.DeleteDataCatalogOutput)

	DeleteNamedQuery(*athena.DeleteNamedQueryInput) (*athena.DeleteNamedQueryOutput, error)
	DeleteNamedQueryWithContext(aws.Context, *athena.DeleteNamedQueryInput, ...request.Option) (*athena.DeleteNamedQueryOutput, error)
	DeleteNamedQueryRequest(*athena.DeleteNamedQueryInput) (*request.Request, *athena.DeleteNamedQueryOutput)

	DeleteNotebook(*athena.DeleteNotebookInput) (*athena.DeleteNotebookOutput, error)
	DeleteNotebookWithContext(aws.Context, *athena.DeleteNotebookInput, ...request.Option) (*athena.DeleteNotebookOutput, error)
	DeleteNotebookRequest(*athena.DeleteNotebookInput) (*request.Request, *athena.DeleteNotebookOutput)

	DeletePreparedStatement(*athena.DeletePreparedStatementInput) (*athena.DeletePreparedStatementOutput, error)
	DeletePreparedStatementWithContext(aws.Context, *athena.DeletePreparedStatementInput, ...request.Option) (*athena.DeletePreparedStatementOutput, error)
	DeletePreparedStatementRequest(*athena.DeletePreparedStatementInput) (*request.Request, *athena.DeletePreparedStatementOutput)

	DeleteWorkGroup(*athena.DeleteWorkGroupInput) (*athena.DeleteWorkGroupOutput, error)
	DeleteWorkGroupWithContext(aws.Context, *athena.DeleteWorkGroupInput, ...request.Option) (*athena.DeleteWorkGroupOutput, error)
	DeleteWorkGroupRequest(*athena.DeleteWorkGroupInput) (*request.Request, *athena.DeleteWorkGroupOutput)

	ExportNotebook(*athena.ExportNotebookInput) (*athena.ExportNotebookOutput, error)
	ExportNotebookWithContext(aws.Context, *athena.ExportNotebookInput, ...request.Option) (*athena.ExportNotebookOutput, error)
	ExportNotebookRequest(*athena.ExportNotebookInput) (*request.Request, *athena.ExportNotebookOutput)

	GetCalculationExecution(*athena.GetCalculationExecutionInput) (*athena.GetCalculationExecutionOutput, error)
	GetCalculationExecutionWithContext(aws.Context, *athena.GetCalculationExecutionInput, ...request.Option) (*athena.GetCalculationExecutionOutput, error)
	GetCalculationExecutionRequest(*athena.GetCalculationExecutionInput) (*request.Request, *athena.GetCalculationExecutionOutput)

	GetCalculationExecutionCode(*athena.GetCalculationExecutionCodeInput) (*athena.GetCalculationExecutionCodeOutput, error)
	GetCalculationExecutionCodeWithContext(aws.Context, *athena.GetCalculationExecutionCodeInput, ...request.Option) (*athena.GetCalculationExecutionCodeOutput, error)
	GetCalculationExecutionCodeRequest(*athena.GetCalculationExecutionCodeInput) (*request.Request, *athena.GetCalculationExecutionCodeOutput)

	GetCalculationExecutionStatus(*athena.GetCalculationExecutionStatusInput) (*athena.GetCalculationExecutionStatusOutput, error)
	GetCalculationExecutionStatusWithContext(aws.Context, *athena.GetCalculationExecutionStatusInput, ...request.Option) (*athena.GetCalculationExecutionStatusOutput, error)
	GetCalculationExecutionStatusRequest(*athena.GetCalculationExecutionStatusInput) (*request.Request, *athena.GetCalculationExecutionStatusOutput)

	GetCapacityAssignmentConfiguration(*athena.GetCapacityAssignmentConfigurationInput) (*athena.GetCapacityAssignmentConfigurationOutput, error)
	GetCapacityAssignmentConfigurationWithContext(aws.Context, *athena.GetCapacityAssignmentConfigurationInput, ...request.Option) (*athena.GetCapacityAssignmentConfigurationOutput, error)
	GetCapacityAssignmentConfigurationRequest(*athena.GetCapacityAssignmentConfigurationInput) (*request.Request, *athena.GetCapacityAssignmentConfigurationOutput)

	GetCapacityReservation(*athena.GetCapacityReservationInput) (*athena.GetCapacityReservationOutput, error)
	GetCapacityReservationWithContext(aws.Context, *athena.GetCapacityReservationInput, ...request.Option) (*athena.GetCapacityReservationOutput, error)
	GetCapacityReservationRequest(*athena.GetCapacityReservationInput) (*request.Request, *athena.GetCapacityReservationOutput)

	GetDataCatalog(*athena.GetDataCatalogInput) (*athena.GetDataCatalogOutput, error)
	GetDataCatalogWithContext(aws.Context, *athena.GetDataCatalogInput, ...request.Option) (*athena.GetDataCatalogOutput, error)
	GetDataCatalogRequest(*athena.GetDataCatalogInput) (*request.Request, *athena.GetDataCatalogOutput)

	GetDatabase(*athena.GetDatabaseInput) (*athena.GetDatabaseOutput, error)
	GetDatabaseWithContext(aws.Context, *athena.GetDatabaseInput, ...request.Option) (*athena.GetDatabaseOutput, error)
	GetDatabaseRequest(*athena.GetDatabaseInput) (*request.Request, *athena.GetDatabaseOutput)

	GetNamedQuery(*athena.GetNamedQueryInput) (*athena.GetNamedQueryOutput, error)
	GetNamedQueryWithContext(aws.Context, *athena.GetNamedQueryInput, ...request.Option) (*athena.GetNamedQueryOutput, error)
	GetNamedQueryRequest(*athena.GetNamedQueryInput) (*request.Request, *athena.GetNamedQueryOutput)

	GetNotebookMetadata(*athena.GetNotebookMetadataInput) (*athena.GetNotebookMetadataOutput, error)
	GetNotebookMetadataWithContext(aws.Context, *athena.GetNotebookMetadataInput, ...request.Option) (*athena.GetNotebookMetadataOutput, error)
	GetNotebookMetadataRequest(*athena.GetNotebookMetadataInput) (*request.Request, *athena.GetNotebookMetadataOutput)

	GetPreparedStatement(*athena.GetPreparedStatementInput) (*athena.GetPreparedStatementOutput, error)
	GetPreparedStatementWithContext(aws.Context, *athena.GetPreparedStatementInput, ...request.Option) (*athena.GetPreparedStatementOutput, error)
	GetPreparedStatementRequest(*athena.GetPreparedStatementInput) (*request.Request, *athena.GetPreparedStatementOutput)

	GetQueryExecution(*athena.GetQueryExecutionInput) (*athena.GetQueryExecutionOutput, error)
	GetQueryExecutionWithContext(aws.Context, *athena.GetQueryExecutionInput, ...request.Option) (*athena.GetQueryExecutionOutput, error)
	GetQueryExecutionRequest(*athena.GetQueryExecutionInput) (*request.Request, *athena.GetQueryExecutionOutput)

	GetQueryResults(*athena.GetQueryResultsInput) (*athena.GetQueryResultsOutput, error)
	GetQueryResultsWithContext(aws.Context, *athena.GetQueryResultsInput, ...request.Option) (*athena.GetQueryResultsOutput, error)
	GetQueryResultsRequest(*athena.GetQueryResultsInput) (*request.Request, *athena.GetQueryResultsOutput)

	GetQueryResultsPages(*athena.GetQueryResultsInput, func(*athena.GetQueryResultsOutput, bool) bool) error
	GetQueryResultsPagesWithContext(aws.Context, *athena.GetQueryResultsInput, func(*athena.GetQueryResultsOutput, bool) bool, ...request.Option) error

	GetQueryRuntimeStatistics(*athena.GetQueryRuntimeStatisticsInput) (*athena.GetQueryRuntimeStatisticsOutput, error)
	GetQueryRuntimeStatisticsWithContext(aws.Context, *athena.GetQueryRuntimeStatisticsInput, ...request.Option) (*athena.GetQueryRuntimeStatisticsOutput, error)
	GetQueryRuntimeStatisticsRequest(*athena.GetQueryRuntimeStatisticsInput) (*request.Request, *athena.GetQueryRuntimeStatisticsOutput)

	GetSession(*athena.GetSessionInput) (*athena.GetSessionOutput, error)
	GetSessionWithContext(aws.Context, *athena.GetSessionInput, ...request.Option) (*athena.GetSessionOutput, error)
	GetSessionRequest(*athena.GetSessionInput) (*request.Request, *athena.GetSessionOutput)

	GetSessionStatus(*athena.GetSessionStatusInput) (*athena.GetSessionStatusOutput, error)
	GetSessionStatusWithContext(aws.Context, *athena.GetSessionStatusInput, ...request.Option) (*athena.GetSessionStatusOutput, error)
	GetSessionStatusRequest(*athena.GetSessionStatusInput) (*request.Request, *athena.GetSessionStatusOutput)

	GetTableMetadata(*athena.GetTableMetadataInput) (*athena.GetTableMetadataOutput, error)
	GetTableMetadataWithContext(aws.Context, *athena.GetTableMetadataInput, ...request.Option) (*athena.GetTableMetadataOutput, error)
	GetTableMetadataRequest(*athena.GetTableMetadataInput) (*request.Request, *athena.GetTableMetadataOutput)

	GetWorkGroup(*athena.GetWorkGroupInput) (*athena.GetWorkGroupOutput, error)
	GetWorkGroupWithContext(aws.Context, *athena.GetWorkGroupInput, ...request.Option) (*athena.GetWorkGroupOutput, error)
	GetWorkGroupRequest(*athena.GetWorkGroupInput) (*request.Request, *athena.GetWorkGroupOutput)

	ImportNotebook(*athena.ImportNotebookInput) (*athena.ImportNotebookOutput, error)
	ImportNotebookWithContext(aws.Context, *athena.ImportNotebookInput, ...request.Option) (*athena.ImportNotebookOutput, error)
	ImportNotebookRequest(*athena.ImportNotebookInput) (*request.Request, *athena.ImportNotebookOutput)

	ListApplicationDPUSizes(*athena.ListApplicationDPUSizesInput) (*athena.ListApplicationDPUSizesOutput, error)
	ListApplicationDPUSizesWithContext(aws.Context, *athena.ListApplicationDPUSizesInput, ...request.Option) (*athena.ListApplicationDPUSizesOutput, error)
	ListApplicationDPUSizesRequest(*athena.ListApplicationDPUSizesInput) (*request.Request, *athena.ListApplicationDPUSizesOutput)

	ListApplicationDPUSizesPages(*athena.ListApplicationDPUSizesInput, func(*athena.ListApplicationDPUSizesOutput, bool) bool) error
	ListApplicationDPUSizesPagesWithContext(aws.Context, *athena.ListApplicationDPUSizesInput, func(*athena.ListApplicationDPUSizesOutput, bool) bool, ...request.Option) error

	ListCalculationExecutions(*athena.ListCalculationExecutionsInput) (*athena.ListCalculationExecutionsOutput, error)
	ListCalculationExecutionsWithContext(aws.Context, *athena.ListCalculationExecutionsInput, ...request.Option) (*athena.ListCalculationExecutionsOutput, error)
	ListCalculationExecutionsRequest(*athena.ListCalculationExecutionsInput) (*request.Request, *athena.ListCalculationExecutionsOutput)

	ListCalculationExecutionsPages(*athena.ListCalculationExecutionsInput, func(*athena.ListCalculationExecutionsOutput, bool) bool) error
	ListCalculationExecutionsPagesWithContext(aws.Context, *athena.ListCalculationExecutionsInput, func(*athena.ListCalculationExecutionsOutput, bool) bool, ...request.Option) error

	ListCapacityReservations(*athena.ListCapacityReservationsInput) (*athena.ListCapacityReservationsOutput, error)
	ListCapacityReservationsWithContext(aws.Context, *athena.ListCapacityReservationsInput, ...request.Option) (*athena.ListCapacityReservationsOutput, error)
	ListCapacityReservationsRequest(*athena.ListCapacityReservationsInput) (*request.Request, *athena.ListCapacityReservationsOutput)

	ListCapacityReservationsPages(*athena.ListCapacityReservationsInput, func(*athena.ListCapacityReservationsOutput, bool) bool) error
	ListCapacityReservationsPagesWithContext(aws.Context, *athena.ListCapacityReservationsInput, func(*athena.ListCapacityReservationsOutput, bool) bool, ...request.Option) error

	ListDataCatalogs(*athena.ListDataCatalogsInput) (*athena.ListDataCatalogsOutput, error)
	ListDataCatalogsWithContext(aws.Context, *athena.ListDataCatalogsInput, ...request.Option) (*athena.ListDataCatalogsOutput, error)
	ListDataCatalogsRequest(*athena.ListDataCatalogsInput) (*request.Request, *athena.ListDataCatalogsOutput)

	ListDataCatalogsPages(*athena.ListDataCatalogsInput, func(*athena.ListDataCatalogsOutput, bool) bool) error
	ListDataCatalogsPagesWithContext(aws.Context, *athena.ListDataCatalogsInput, func(*athena.ListDataCatalogsOutput, bool) bool, ...request.Option) error

	ListDatabases(*athena.ListDatabasesInput) (*athena.ListDatabasesOutput, error)
	ListDatabasesWithContext(aws.Context, *athena.ListDatabasesInput, ...request.Option) (*athena.ListDatabasesOutput, error)
	ListDatabasesRequest(*athena.ListDatabasesInput) (*request.Request, *athena.ListDatabasesOutput)

	ListDatabasesPages(*athena.ListDatabasesInput, func(*athena.ListDatabasesOutput, bool) bool) error
	ListDatabasesPagesWithContext(aws.Context, *athena.ListDatabasesInput, func(*athena.ListDatabasesOutput, bool) bool, ...request.Option) error

	ListEngineVersions(*athena.ListEngineVersionsInput) (*athena.ListEngineVersionsOutput, error)
	ListEngineVersionsWithContext(aws.Context, *athena.ListEngineVersionsInput, ...request.Option) (*athena.ListEngineVersionsOutput, error)
	ListEngineVersionsRequest(*athena.ListEngineVersionsInput) (*request.Request, *athena.ListEngineVersionsOutput)

	ListEngineVersionsPages(*athena.ListEngineVersionsInput, func(*athena.ListEngineVersionsOutput, bool) bool) error
	ListEngineVersionsPagesWithContext(aws.Context, *athena.ListEngineVersionsInput, func(*athena.ListEngineVersionsOutput, bool) bool, ...request.Option) error

	ListExecutors(*athena.ListExecutorsInput) (*athena.ListExecutorsOutput, error)
	ListExecutorsWithContext(aws.Context, *athena.ListExecutorsInput, ...request.Option) (*athena.ListExecutorsOutput, error)
	ListExecutorsRequest(*athena.ListExecutorsInput) (*request.Request, *athena.ListExecutorsOutput)

	ListExecutorsPages(*athena.ListExecutorsInput, func(*athena.ListExecutorsOutput, bool) bool) error
	ListExecutorsPagesWithContext(aws.Context, *athena.ListExecutorsInput, func(*athena.ListExecutorsOutput, bool) bool, ...request.Option) error

	ListNamedQueries(*athena.ListNamedQueriesInput) (*athena.ListNamedQueriesOutput, error)
	ListNamedQueriesWithContext(aws.Context, *athena.ListNamedQueriesInput, ...request.Option) (*athena.ListNamedQueriesOutput, error)
	ListNamedQueriesRequest(*athena.ListNamedQueriesInput) (*request.Request, *athena.ListNamedQueriesOutput)

	ListNamedQueriesPages(*athena.ListNamedQueriesInput, func(*athena.ListNamedQueriesOutput, bool) bool) error
	ListNamedQueriesPagesWithContext(aws.Context, *athena.ListNamedQueriesInput, func(*athena.ListNamedQueriesOutput, bool) bool, ...request.Option) error

	ListNotebookMetadata(*athena.ListNotebookMetadataInput) (*athena.ListNotebookMetadataOutput, error)
	ListNotebookMetadataWithContext(aws.Context, *athena.ListNotebookMetadataInput, ...request.Option) (*athena.ListNotebookMetadataOutput, error)
	ListNotebookMetadataRequest(*athena.ListNotebookMetadataInput) (*request.Request, *athena.ListNotebookMetadataOutput)

	ListNotebookSessions(*athena.ListNotebookSessionsInput) (*athena.ListNotebookSessionsOutput, error)
	ListNotebookSessionsWithContext(aws.Context, *athena.ListNotebookSessionsInput, ...request.Option) (*athena.ListNotebookSessionsOutput, error)
	ListNotebookSessionsRequest(*athena.ListNotebookSessionsInput) (*request.Request, *athena.ListNotebookSessionsOutput)

	ListPreparedStatements(*athena.ListPreparedStatementsInput) (*athena.ListPreparedStatementsOutput, error)
	ListPreparedStatementsWithContext(aws.Context, *athena.ListPreparedStatementsInput, ...request.Option) (*athena.ListPreparedStatementsOutput, error)
	ListPreparedStatementsRequest(*athena.ListPreparedStatementsInput) (*request.Request, *athena.ListPreparedStatementsOutput)

	ListPreparedStatementsPages(*athena.ListPreparedStatementsInput, func(*athena.ListPreparedStatementsOutput, bool) bool) error
	ListPreparedStatementsPagesWithContext(aws.Context, *athena.ListPreparedStatementsInput, func(*athena.ListPreparedStatementsOutput, bool) bool, ...request.Option) error

	ListQueryExecutions(*athena.ListQueryExecutionsInput) (*athena.ListQueryExecutionsOutput, error)
	ListQueryExecutionsWithContext(aws.Context, *athena.ListQueryExecutionsInput, ...request.Option) (*athena.ListQueryExecutionsOutput, error)
	ListQueryExecutionsRequest(*athena.ListQueryExecutionsInput) (*request.Request, *athena.ListQueryExecutionsOutput)

	ListQueryExecutionsPages(*athena.ListQueryExecutionsInput, func(*athena.ListQueryExecutionsOutput, bool) bool) error
	ListQueryExecutionsPagesWithContext(aws.Context, *athena.ListQueryExecutionsInput, func(*athena.ListQueryExecutionsOutput, bool) bool, ...request.Option) error

	ListSessions(*athena.ListSessionsInput) (*athena.ListSessionsOutput, error)
	ListSessionsWithContext(aws.Context, *athena.ListSessionsInput, ...request.Option) (*athena.ListSessionsOutput, error)
	ListSessionsRequest(*athena.ListSessionsInput) (*request.Request, *athena.ListSessionsOutput)

	ListSessionsPages(*athena.ListSessionsInput, func(*athena.ListSessionsOutput, bool) bool) error
	ListSessionsPagesWithContext(aws.Context, *athena.ListSessionsInput, func(*athena.ListSessionsOutput, bool) bool, ...request.Option) error

	ListTableMetadata(*athena.ListTableMetadataInput) (*athena.ListTableMetadataOutput, error)
	ListTableMetadataWithContext(aws.Context, *athena.ListTableMetadataInput, ...request.Option) (*athena.ListTableMetadataOutput, error)
	ListTableMetadataRequest(*athena.ListTableMetadataInput) (*request.Request, *athena.ListTableMetadataOutput)

	ListTableMetadataPages(*athena.ListTableMetadataInput, func(*athena.ListTableMetadataOutput, bool) bool) error
	ListTableMetadataPagesWithContext(aws.Context, *athena.ListTableMetadataInput, func(*athena.ListTableMetadataOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*athena.ListTagsForResourceInput) (*athena.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *athena.ListTagsForResourceInput, ...request.Option) (*athena.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*athena.ListTagsForResourceInput) (*request.Request, *athena.ListTagsForResourceOutput)

	ListTagsForResourcePages(*athena.ListTagsForResourceInput, func(*athena.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *athena.ListTagsForResourceInput, func(*athena.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListWorkGroups(*athena.ListWorkGroupsInput) (*athena.ListWorkGroupsOutput, error)
	ListWorkGroupsWithContext(aws.Context, *athena.ListWorkGroupsInput, ...request.Option) (*athena.ListWorkGroupsOutput, error)
	ListWorkGroupsRequest(*athena.ListWorkGroupsInput) (*request.Request, *athena.ListWorkGroupsOutput)

	ListWorkGroupsPages(*athena.ListWorkGroupsInput, func(*athena.ListWorkGroupsOutput, bool) bool) error
	ListWorkGroupsPagesWithContext(aws.Context, *athena.ListWorkGroupsInput, func(*athena.ListWorkGroupsOutput, bool) bool, ...request.Option) error

	PutCapacityAssignmentConfiguration(*athena.PutCapacityAssignmentConfigurationInput) (*athena.PutCapacityAssignmentConfigurationOutput, error)
	PutCapacityAssignmentConfigurationWithContext(aws.Context, *athena.PutCapacityAssignmentConfigurationInput, ...request.Option) (*athena.PutCapacityAssignmentConfigurationOutput, error)
	PutCapacityAssignmentConfigurationRequest(*athena.PutCapacityAssignmentConfigurationInput) (*request.Request, *athena.PutCapacityAssignmentConfigurationOutput)

	StartCalculationExecution(*athena.StartCalculationExecutionInput) (*athena.StartCalculationExecutionOutput, error)
	StartCalculationExecutionWithContext(aws.Context, *athena.StartCalculationExecutionInput, ...request.Option) (*athena.StartCalculationExecutionOutput, error)
	StartCalculationExecutionRequest(*athena.StartCalculationExecutionInput) (*request.Request, *athena.StartCalculationExecutionOutput)

	StartQueryExecution(*athena.StartQueryExecutionInput) (*athena.StartQueryExecutionOutput, error)
	StartQueryExecutionWithContext(aws.Context, *athena.StartQueryExecutionInput, ...request.Option) (*athena.StartQueryExecutionOutput, error)
	StartQueryExecutionRequest(*athena.StartQueryExecutionInput) (*request.Request, *athena.StartQueryExecutionOutput)

	StartSession(*athena.StartSessionInput) (*athena.StartSessionOutput, error)
	StartSessionWithContext(aws.Context, *athena.StartSessionInput, ...request.Option) (*athena.StartSessionOutput, error)
	StartSessionRequest(*athena.StartSessionInput) (*request.Request, *athena.StartSessionOutput)

	StopCalculationExecution(*athena.StopCalculationExecutionInput) (*athena.StopCalculationExecutionOutput, error)
	StopCalculationExecutionWithContext(aws.Context, *athena.StopCalculationExecutionInput, ...request.Option) (*athena.StopCalculationExecutionOutput, error)
	StopCalculationExecutionRequest(*athena.StopCalculationExecutionInput) (*request.Request, *athena.StopCalculationExecutionOutput)

	StopQueryExecution(*athena.StopQueryExecutionInput) (*athena.StopQueryExecutionOutput, error)
	StopQueryExecutionWithContext(aws.Context, *athena.StopQueryExecutionInput, ...request.Option) (*athena.StopQueryExecutionOutput, error)
	StopQueryExecutionRequest(*athena.StopQueryExecutionInput) (*request.Request, *athena.StopQueryExecutionOutput)

	TagResource(*athena.TagResourceInput) (*athena.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *athena.TagResourceInput, ...request.Option) (*athena.TagResourceOutput, error)
	TagResourceRequest(*athena.TagResourceInput) (*request.Request, *athena.TagResourceOutput)

	TerminateSession(*athena.TerminateSessionInput) (*athena.TerminateSessionOutput, error)
	TerminateSessionWithContext(aws.Context, *athena.TerminateSessionInput, ...request.Option) (*athena.TerminateSessionOutput, error)
	TerminateSessionRequest(*athena.TerminateSessionInput) (*request.Request, *athena.TerminateSessionOutput)

	UntagResource(*athena.UntagResourceInput) (*athena.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *athena.UntagResourceInput, ...request.Option) (*athena.UntagResourceOutput, error)
	UntagResourceRequest(*athena.UntagResourceInput) (*request.Request, *athena.UntagResourceOutput)

	UpdateCapacityReservation(*athena.UpdateCapacityReservationInput) (*athena.UpdateCapacityReservationOutput, error)
	UpdateCapacityReservationWithContext(aws.Context, *athena.UpdateCapacityReservationInput, ...request.Option) (*athena.UpdateCapacityReservationOutput, error)
	UpdateCapacityReservationRequest(*athena.UpdateCapacityReservationInput) (*request.Request, *athena.UpdateCapacityReservationOutput)

	UpdateDataCatalog(*athena.UpdateDataCatalogInput) (*athena.UpdateDataCatalogOutput, error)
	UpdateDataCatalogWithContext(aws.Context, *athena.UpdateDataCatalogInput, ...request.Option) (*athena.UpdateDataCatalogOutput, error)
	UpdateDataCatalogRequest(*athena.UpdateDataCatalogInput) (*request.Request, *athena.UpdateDataCatalogOutput)

	UpdateNamedQuery(*athena.UpdateNamedQueryInput) (*athena.UpdateNamedQueryOutput, error)
	UpdateNamedQueryWithContext(aws.Context, *athena.UpdateNamedQueryInput, ...request.Option) (*athena.UpdateNamedQueryOutput, error)
	UpdateNamedQueryRequest(*athena.UpdateNamedQueryInput) (*request.Request, *athena.UpdateNamedQueryOutput)

	UpdateNotebook(*athena.UpdateNotebookInput) (*athena.UpdateNotebookOutput, error)
	UpdateNotebookWithContext(aws.Context, *athena.UpdateNotebookInput, ...request.Option) (*athena.UpdateNotebookOutput, error)
	UpdateNotebookRequest(*athena.UpdateNotebookInput) (*request.Request, *athena.UpdateNotebookOutput)

	UpdateNotebookMetadata(*athena.UpdateNotebookMetadataInput) (*athena.UpdateNotebookMetadataOutput, error)
	UpdateNotebookMetadataWithContext(aws.Context, *athena.UpdateNotebookMetadataInput, ...request.Option) (*athena.UpdateNotebookMetadataOutput, error)
	UpdateNotebookMetadataRequest(*athena.UpdateNotebookMetadataInput) (*request.Request, *athena.UpdateNotebookMetadataOutput)

	UpdatePreparedStatement(*athena.UpdatePreparedStatementInput) (*athena.UpdatePreparedStatementOutput, error)
	UpdatePreparedStatementWithContext(aws.Context, *athena.UpdatePreparedStatementInput, ...request.Option) (*athena.UpdatePreparedStatementOutput, error)
	UpdatePreparedStatementRequest(*athena.UpdatePreparedStatementInput) (*request.Request, *athena.UpdatePreparedStatementOutput)

	UpdateWorkGroup(*athena.UpdateWorkGroupInput) (*athena.UpdateWorkGroupOutput, error)
	UpdateWorkGroupWithContext(aws.Context, *athena.UpdateWorkGroupInput, ...request.Option) (*athena.UpdateWorkGroupOutput, error)
	UpdateWorkGroupRequest(*athena.UpdateWorkGroupInput) (*request.Request, *athena.UpdateWorkGroupOutput)
}

var _ AthenaAPI = (*athena.Athena)(nil)
