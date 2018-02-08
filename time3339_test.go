package src

import (
	"testing"
)

var rfc3339_tests = []struct {
	in      string
	unixOut int64
}{
	{"2015-10-05T17:17:02+00:00", 1444065422},
	{"2015-10-05T17:18:08+00:00", 1444065488},
	{"2015-10-05T17:15:08+01:00", 1444061708},
	{"2075-10-05T17:15:08+01:10", 3337517108},
}

var yyyymmddTests = []struct {
	in      string
	unixOut int64
}{
	{"2011-08-11", 1313020800},
	{"2013-02-12", 1360627200},
	{"2222-01-03", 7952515200},
	{"1999-10-10", 939513600},
}

var salesForceTests = []struct {
	in      string
	unixOut int64
}{
	{"2015-10-05T18:17:04.000Z", 1444069024}, //1
	{"2015-10-05T16:15:52.000Z", 1444061752}, //2
	{"2015-10-05T17:17:02.000Z", 1444065422}, //100
	{"2015-10-05T17:17:02.120Z", 1444065422},
	{"2015-10-05T17:17:02.001Z", 1444065422},
	{"2003-10-05T17:17:02.001Z", 1065374222},
	{"2003-10-05T17:17:02.021Z", 1065374222},
	//ISO 8601 time (with colon)
	{"2018-02-08T19:17:32+01:00", 1518113852},
	{"2012-02-08T19:17:32+01:20", 1328723852},
	{"2001-02-08T09:17:32+00:00", 981623852},
	{"2031-02-08T09:17:32+02:00", 1928301452},
	{"1001-02-08T09:17:32+01:10", -30575375548},
	{"2321-02-02T09:14:32+00:01", 11079278012},
	{"2391-02-08T09:17:32+09:00", 13288753052},
}

func TestSalesforceToRFC3339(t *testing.T) {
	for _, tt := range salesForceTests {
		mytime, err := NewFromSalesforceFormat(tt.in)

		if err != nil {
			t.Fatal(err)
		}

		rfc3339 := mytime.ToRFC3339()

		newTime, err := NewFromRFC3339(rfc3339)

		if err != nil {
			t.Fatal(err)
		}

		if newTime.ToUnix() != mytime.ToUnix() {
			t.Errorf("SalesforceFormat converted to rfc3339 does not have the same value => has %d, want %d", newTime.ToUnix(), mytime.ToUnix())
		}
	}
}
func TestRFC3339ToSalesforce(t *testing.T) {
	for _, tt := range rfc3339_tests {
		mytime, err := NewFromRFC3339(tt.in)

		if err != nil {
			t.Fatal(err)
		}

		salesforce := mytime.ToSalesforceFormat()

		newTime, err := NewFromSalesforceFormat(salesforce)

		if err != nil {
			t.Fatal(err)
		}

		if newTime.ToUnix() != mytime.ToUnix() {
			t.Errorf("rfc3339 converted to salesforceformat does not have the same value => has %d, want %d", newTime.ToUnix(), mytime.ToUnix())
		}
	}
}

func TestSalesforceSameInputAsOutput(t *testing.T) {
	for _, tt := range salesForceTests {
		mytime, err := NewFromSalesforceFormat(tt.in)

		if err != nil {
			t.Fatal(err)
		}

		output := mytime.ToSalesforceFormat()
		if output != tt.in {
			//the output might have additional zeros stripped away. Check if the input and output values are the same.
			n, _ := NewFromSalesforceFormat(output)

			if n.ToUnix() != mytime.ToUnix() {
				t.Errorf("NewFromSalesforceFormat(%s).ToSalesforceFormat() => has %s, want %s", tt.in, output, tt.in)
			}
		}
	}
}

func TestRFC3999(t *testing.T) {
	for _, tt := range rfc3339_tests {
		mytime, _ := NewFromRFC3339(tt.in)
		if mytime.ToUnix() != tt.unixOut {
			t.Errorf("NewFromRFC3339(%s) => has %d, want %d", tt.in, mytime.ToUnix(), tt.unixOut)
		}
	}
}

func TestYYYYMMDD(t *testing.T) {
	for _, tt := range yyyymmddTests {
		mytime, err := NewFromYYYMMDD(tt.in)

		if err != nil {
			t.Fatal(err)
		}

		if mytime.ToUnix() != tt.unixOut {
			t.Errorf("NewFromYYYMMDD(%s) => has %d, want %d", tt.in, mytime.ToUnix(), tt.unixOut)
		}
	}
}

func TestSalesForceInput(t *testing.T) {
	for _, tt := range salesForceTests {
		mytime, err := NewFromSalesforceFormat(tt.in)

		if err != nil {
			t.Fatal(err)
		}

		if mytime.ToUnix() != tt.unixOut {
			t.Errorf("NewFromSalesforceFormat(%s) => has %d, want %d", tt.in, mytime.ToUnix(), tt.unixOut)
		}
	}
}
