// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package securitylake provides the client and types for making API
// requests to Amazon Security Lake.
//
// Amazon Security Lake is in preview release. Your use of the Security Lake
// preview is subject to Section 2 of the Amazon Web Services Service Terms
// (http://aws.amazon.com/service-terms/)("Betas and Previews").
//
// Amazon Security Lake is a fully managed security data lake service. You can
// use Security Lake to automatically centralize security data from cloud, on-premises,
// and custom sources into a data lake that's stored in your Amazon Web Servicesaccount.
// Amazon Web Services Organizations is an account management service that lets
// you consolidate multiple Amazon Web Services accounts into an organization
// that you create and centrally manage. With Organizations, you can create
// member accounts and invite existing accounts to join your organization. Security
// Lake helps you analyze security data for a more complete understanding of
// your security posture across the entire organization. It can also help you
// improve the protection of your workloads, applications, and data.
//
// The data lake is backed by Amazon Simple Storage Service (Amazon S3) buckets,
// and you retain ownership over your data.
//
// Amazon Security Lake integrates with CloudTrail, a service that provides
// a record of actions taken by a user, role, or an Amazon Web Services service
// in Security Lake CloudTrail captures API calls for Security Lake as events.
// The calls captured include calls from the Security Lake console and code
// calls to the Security Lake API operations. If you create a trail, you can
// enable continuous delivery of CloudTrail events to an Amazon S3 bucket, including
// events for Security Lake. If you don't configure a trail, you can still view
// the most recent events in the CloudTrail console in Event history. Using
// the information collected by CloudTrail you can determine the request that
// was made to Security Lake, the IP address from which the request was made,
// who made the request, when it was made, and additional details. To learn
// more about Security Lake information in CloudTrail, see the Amazon Security
// Lake User Guide (https://docs.aws.amazon.com/security-lake/latest/userguide/securitylake-cloudtrail.html).
//
// Security Lake automates the collection of security-related log and event
// data from integrated Amazon Web Services and third-party services. It also
// helps you manage the lifecycle of data with customizable retention and replication
// settings. Security Lake converts ingested data into Apache Parquet format
// and a standard open-source schema called the Open Cybersecurity Schema Framework
// (OCSF).
//
// Other Amazon Web Services and third-party services can subscribe to the data
// that's stored in Security Lake for incident response and security data analytics.
//
// See https://docs.aws.amazon.com/goto/WebAPI/securitylake-2018-05-10 for more information on this service.
//
// See securitylake package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/securitylake/
//
// # Using the Client
//
// To contact Amazon Security Lake with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Security Lake client SecurityLake for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/securitylake/#New
package securitylake
