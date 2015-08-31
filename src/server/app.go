package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bmizerany/lpx"
	"github.com/docopt/docopt-go"
	"log"
	"net/http"
	"os"
	"strconv"
)

type LogLine struct {
	PrivalVersion string `json:"priv"`
	Time          string `json:"time"`
	HostName      string `json:"hostname"`
	Name          string `json:"name"`
	ProcID        string `json:"procid"`
	MsgID         string `json:"msgid"`
	Data          string `json:"data"`
}

func NewLogLineFromLpx(lp *lpx.Reader) *LogLine {
	hdr := lp.Header()
	data := lp.Bytes()
	return &LogLine{
		string(hdr.PrivalVersion),
		string(hdr.Time),
		string(hdr.Hostname),
		string(hdr.Name),
		string(hdr.Procid),
		string(hdr.Msgid),
		string(data),
	}
}

var logsCh chan *LogLine

const LOGSCH_BUFFER = 100

func receiveLogs() {
	for line := range logsCh {
		err := handleLog(line)
		if err != nil {
			log.Fatalf("Error handling log: %v", err)
		}
	}
}

func handleLog(line *LogLine) error {
	var err error
	data, err := json.Marshal(&line)
	if err != nil {
		log.Fatalf("JSON error: %v", err)
	}
	_, err = fmt.Println(string(data))
	return err
}

func handlerLogs(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "POST":
		// Handle logplex frames
		lp := lpx.NewReader(bufio.NewReader(r.Body))
		for lp.Next() {
			logsCh <- NewLogLineFromLpx(lp)
		}
	case r.Method == "GET":
		// Stream received logs back to the client
	}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	usage := `vessel
Usage:
  vessel -p PORT
Options:
  -p PORT --port=PORT    HTTP port to listen
`

	arguments, err := docopt.Parse(usage, nil, true, "vessel", false)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	portString := arguments["--port"].(string)
	port, err := strconv.Atoi(portString)
	if err != nil || port == 0 {
		fmt.Fprintf(os.Stderr, "err: invalid port %s\n", portString)
		os.Exit(2)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)

	logsCh = make(chan *LogLine, LOGSCH_BUFFER)
	go receiveLogs()

	http.HandleFunc("/logs", handlerLogs)
	staticFs := http.FileServer(http.Dir("static"))
	http.Handle("/", staticFs)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "http server failure: %v\n", err)
		os.Exit(2)
	}
}
