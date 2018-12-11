package blacklist

import "github.com/globalsign/mgo/bson"

//BlacklistedIP holds information on a blacklisted IP address and
//the summary statistics on the host
type (
	BlacklistedIP struct {
		IP                string   `bson:"ip"`
		Connections       int      `bson:"conn"`
		UniqueConnections int      `bson:"uconn"`
		TotalBytes        int      `bson:"total_bytes"`
		Lists             []string `bson:"lists"`
		ConnectedHosts    []string `bson:",omitempty"`
	}

	//BlacklistedHostname holds information on a blacklisted hostname and
	//the summary statistics associated with the hosts behind the hostname
	BlacklistedHostname struct {
		Hostname          string   `bson:"hostname"`
		Connections       int      `bson:"conn"`
		UniqueConnections int      `bson:"uconn"`
		TotalBytes        int      `bson:"total_bytes"`
		Lists             []string `bson:"lists"`
		ConnectedHosts    []string `bson:",omitempty"`
	}

	//AnalysisInput contains the summary statistics of a unique connection
	AnalysisInput struct {
		IP                string   `bson:"ip"` // IP
		Connections       int      `bson:"conn_count"`
		UniqueConnections int      `bson:"uconn_count"`
		TotalBytes        int      `bson:"total_bytes"`
		AverageBytes      int      `bson:"avg_bytes"`
		Targets           []string `bson:"targets"`
	}

	//AnalysisOutput contains the summary statistics of a unique connection
	AnalysisOutput struct {
		IP                string   `bson:"ip"`
		Connections       int      `bson:"conn"`
		UniqueConnections int      `bson:"uconn"`
		TotalBytes        int      `bson:"total_bytes"`
		AverageBytes      int      `bson:"avg_bytes"`
		Lists             []string `bson:"lists"`
		Targets           []string `bson:"targets"`
	}

	//IPResult contains the summary of a result from the "ip" collection of rita-bl
	IPResult struct {
		ID        bson.ObjectId   `bson:"_id,omitempty"` // Unique Connection ID
		Index     string          `bson:"index"`         // Potentially malicious IP
		List      string          `bson:"list"`          // which blacklist ip was listed on
		ExtraData ExtraDataResult `bson:"extradata"`     // Associated data
	}

	//ExtraDataResult contains the structure of the extradata field in each document of the rita-bl ip collection
	ExtraDataResult struct {
		Date    string `bson:"date"`    // Date IP was added to blacklist
		Host    string `bson:"host"`    // IP in question
		Country string `bson:"country"` // Reported country of origin for IP
		ID      int32  `bson:"id"`      // not sure yet, but its there
	}
)
