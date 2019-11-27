package dns

import "os"

//Modify with your Access Key Id and Access Key Secret
var (
	TestAccessKeyId     = os.Getenv("AccessKeyId")
	TestAccessKeySecret = os.Getenv("AccessKeySecret")
	TestDomainName      = os.Getenv("TopDomain")
	TestDomainGroupName = "testgroup"
	TestChanegGroupName = "testchangegroup"
)

var testClient *Client

func NewTestClient() *Client {
	if testClient == nil {
		testClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

// change DNSDefaultEndpoint to "http://alidns.aliyuncs.com"
func NewTestClientNew() *Client {
	if testClient == nil {
		testClient = NewClientNew(TestAccessKeyId, TestAccessKeySecret)
	}
	return testClient
}

var testDebugClient *Client

func NewTestClientForDebug() *Client {
	if testDebugClient == nil {
		testDebugClient = NewClient(TestAccessKeyId, TestAccessKeySecret)
		testDebugClient.SetDebug(true)
	}
	return testDebugClient
}
