package fleet

import "context"

// LiveQueryStore defines an interface for storing and retrieving the status of
// live queries in the Fleet system.
type LiveQueryStore interface {
	// RunQuery starts a query with the given name and SQL, targeting the
	// provided host IDs.
	RunQuery(name, sql string, hostIDs []uint) error
	// StopQuery stops a running query with the given name. Hosts will no longer
	// receive the query after StopQuery has been called.
	StopQuery(name string) error
	// QueriesForHost returns the active queries for the given host ID. The
	// return value maps from query name to SQL.
	QueriesForHost(hostID uint) (map[string]string, error)
	// QueryCompletedByHost marks the query with the given name as completed by the
	// given host. After calling QueryCompleted, that query will no longer be
	// sent to the host.
	QueryCompletedByHost(name string, hostID uint) error
	// CleanupInactiveQueries removes any queries that are not in the provided
	// list of active campaign IDs. This is used via a cron job to regularly
	// cleanup any queries that may have failed to be stopped properly.
	CleanupInactiveQueries(ctx context.Context, activeCampaignIDs []uint) error
}
