package index

var _ Backend = &esIndex{}

func testEsClient() (*esIndex, error) {
	return createESIndex(esHost, index)
}
