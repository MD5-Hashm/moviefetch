package torrent

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cenkalti/rain/torrent"
	"github.com/fatih/color"
)

type Stat struct {
	Status     string
	Downloaded int64
	Peers      int
	Total      int64
}

var Version string = "420"
var conf = torrent.Config{

	Database:                               "downloads/session.db",
	DataDir:                                "downloads",
	DataDirIncludesTorrentID:               true,
	PortBegin:                              50000,
	PortEnd:                                60000,
	MaxOpenFiles:                           10240,
	PEXEnabled:                             true,
	ResumeWriteInterval:                    30 * time.Second,
	PrivatePeerIDPrefix:                    "-MF" + Version + "-",
	PrivateExtensionHandshakeClientVersion: "MF " + Version,
	BlocklistUpdateInterval:                24 * time.Hour,
	BlocklistUpdateTimeout:                 10 * time.Minute,
	BlocklistEnabledForTrackers:            true,
	BlocklistEnabledForOutgoingConnections: true,
	BlocklistEnabledForIncomingConnections: true,
	BlocklistMaxResponseSize:               100 << 20,
	TorrentAddHTTPTimeout:                  30 * time.Second,
	MaxMetadataSize:                        30 << 20,
	MaxTorrentSize:                         10 << 20,
	MaxPieces:                              64 << 10,
	DNSResolveTimeout:                      5 * time.Second,
	SpeedLimitUpload:                       1,
	ResumeOnStartup:                        true,

	RPCEnabled:         true,
	RPCHost:            "127.0.0.1",
	RPCPort:            7246,
	RPCShutdownTimeout: 5 * time.Second,

	TrackerNumWant:              200,
	TrackerStopTimeout:          5 * time.Second,
	TrackerMinAnnounceInterval:  time.Minute,
	TrackerHTTPTimeout:          10 * time.Second,
	TrackerHTTPPrivateUserAgent: "Rain/" + Version,
	TrackerHTTPMaxResponseSize:  2 << 20,
	TrackerHTTPVerifyTLS:        true,

	DHTEnabled:             true,
	DHTHost:                "0.0.0.0",
	DHTPort:                7246,
	DHTAnnounceInterval:    30 * time.Minute,
	DHTMinAnnounceInterval: time.Minute,
	DHTBootstrapNodes: []string{
		"router.bittorrent.com:6881",
		"dht.transmissionbt.com:6881",
		"router.utorrent.com:6881",
		"dht.libtorrent.org:25401",
		"dht.aelitis.com:6881",
	},

	UnchokedPeers:                3,
	OptimisticUnchokedPeers:      1,
	MaxRequestsIn:                250,
	MaxRequestsOut:               250,
	DefaultRequestsOut:           50,
	RequestTimeout:               20 * time.Second,
	EndgameMaxDuplicateDownloads: 20,
	MaxPeerDial:                  80,
	MaxPeerAccept:                20,
	ParallelMetadataDownloads:    1,
	PeerConnectTimeout:           5 * time.Second,
	PeerHandshakeTimeout:         10 * time.Second,
	PieceReadTimeout:             30 * time.Second,
	MaxPeerAddresses:             2000,
	AllowedFastSet:               10,

	ReadCacheBlockSize: 128 << 10,
	ReadCacheSize:      256 << 20,
	ReadCacheTTL:       1 * time.Minute,
	ParallelReads:      1,
	ParallelWrites:     1,
	WriteCacheSize:     1 << 30,

	WebseedDialTimeout:             10 * time.Second,
	WebseedTLSHandshakeTimeout:     10 * time.Second,
	WebseedResponseHeaderTimeout:   10 * time.Second,
	WebseedResponseBodyReadTimeout: 10 * time.Second,
	WebseedRetryInterval:           time.Hour * 100,
	WebseedVerifyTLS:               true,
	WebseedMaxSources:              10,
	WebseedMaxDownloads:            3,
}

func getname(link string) []string {
	var titles []string
	titles = append(titles, strings.Split(link, `[Bitsearch.to]`)[1])
	return titles
}

func DownloadTorrentFile(link string) {
	response, err := http.Get(link)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	output, err := os.Create(getname(link)[0] + ".torrent")
	if err != nil {
		panic(err)
	}
	defer output.Close()
	_, err = io.Copy(output, response.Body)
	if err != nil {
		panic(err)
	}
}

func AddTorrent(torrentf string, out chan Stat) {
	DownloadTorrentFile(torrentf)
	var opt = torrent.AddTorrentOptions{
		ID:                getname(torrentf)[0],
		Stopped:           false,
		StopAfterDownload: true,
		StopAfterMetadata: false}
	if _, err := os.Stat("downloads/" + getname(torrentf)[0]); !os.IsNotExist(err) {
		err = os.RemoveAll("downloads/" + getname(torrentf)[0])
		if err != nil {
			color.Red("Failed to remove previous files from " + getname(torrentf)[0])
		}
	}
	ses, _ := torrent.NewSession(torrent.Config(conf))
	cont, _ := os.Open(getname(torrentf)[0] + ".torrent")
	defer cont.Close()
	tor, _ := ses.AddTorrent(cont, &opt)
	for range time.Tick(time.Second) {
		s := tor.Stats()
		if strings.Contains(s.Status.String(), "Stopped") {
			_ = ses.Close()
			//_ = os.Chdir("downloads")
			//_ = os.Remove("session.db")
			//_ = os.Chdir("..")
			break
		}
		if out != nil {
			out <- Stat{Status: s.Status.String(),
				Downloaded: s.Bytes.Completed,
				Peers:      s.Peers.Total,
				Total:      s.Bytes.Total}
		}
	}
}
