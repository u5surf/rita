package blacklist

import (
	"sync"

	"github.com/activecm/rita/config"
	"github.com/activecm/rita/database"
	"github.com/activecm/rita/datatypes/blacklist"
)

type (
	//writer simply writes AnalysisOutput objects to the beacons collection
	writer struct {
		source       bool
		db           *database.DB                   // provides access to MongoDB
		conf         *config.Config                 // contains details needed to access MongoDB
		writeChannel chan *blacklist.AnalysisOutput // holds analyzed data
		writeWg      sync.WaitGroup                 // wait for writing to finish
	}
)

//newWriter creates a writer object to write AnalysisOutput data to
//the beacons collection
func newWriter(source bool, db *database.DB, conf *config.Config) *writer {
	return &writer{
		source:       source,
		db:           db,
		conf:         conf,
		writeChannel: make(chan *blacklist.AnalysisOutput),
	}
}

//write queues up a AnalysisOutput to be written to the beacons collection
//Note: this function may block
func (w *writer) write(data *blacklist.AnalysisOutput) {
	w.writeChannel <- data
}

// close waits for the write threads to finish
func (w *writer) close() {
	close(w.writeChannel)
	w.writeWg.Wait()
}

// start kicks off a new write thread
func (w *writer) start() {
	w.writeWg.Add(1)
	go func() {
		ssn := w.db.Session.Copy()
		defer ssn.Close()

		for data := range w.writeChannel {
			if w.source {
				ssn.DB(w.db.GetSelectedDB()).C(w.conf.T.Blacklisted.SourceIPsTable).Insert(data)
			} else {
				ssn.DB(w.db.GetSelectedDB()).C(w.conf.T.Blacklisted.DestIPsTable).Insert(data)
			}
		}
		w.writeWg.Done()
	}()
}