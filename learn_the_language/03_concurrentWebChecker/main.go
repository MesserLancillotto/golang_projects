package main

import (
	"fmt"
	"net/http"
	"time"
	"sync"
)

var urlList []string = []string{
	// Fondazioni e Organizzazioni FOSS / Digital Rights
	"fsf.org",
	"eff.org",
	"mozilla.org",
	"apache.org",
	"linuxfoundation.org",
	"kde.org",
	"gnome.org",
	"documentfoundation.org",
	"eclipse.org",
	"python.org",
	"rust-lang.org",
	"golang.org",
	"gnu.org",
	"openbsd.org",
	"freebsd.org",
	"netbsd.org",
	"softwarefreedom.org",
	"fdroid.org",
	"videolan.org",
	"matrix.org",

	// Distribuzioni Linux e Sistemi Operativi
	"archlinux.org",
	"debian.org",
	"ubuntu.com",
	"fedoraproject.org",
	"alpinelinux.org",
	"gentoo.org",
	"nixos.org",
	"manjaro.org",
	"linuxmint.com",
	"opensuse.org",
	"postmarketos.org",
	"tails.boum.org",
	"qubes-os.org",
	"rockylinux.org",
	"almalinux.org",

	// Enciclopedie, Archivi e Informazione Libera
	"wikipedia.org",
	"wikimedia.org",
	"wikibooks.org",
	"wiktionary.org",
	"archive.org",
	"gutenberg.org",
	"openstreetmap.org",
	"wikileaks.org",
	"curlie.org",
	"openlibrary.org",
	"wikiquote.org",
	"wikisource.org",
	"wikinews.org",
	"commons.wikimedia.org",
	"wikidata.org",

	// Software di Grafica, Multimedia e Produttività FOSS
	"krita.org",
	"blender.org",
	"gimp.org",
	"inkscape.org",
	"libreoffice.org",
	"audacityteam.org",
	"obs-project.com",
	"handbrake.fr",
	"scribus.net",
	"freecad.org",
	"godotengine.org",
	"darktable.org",
	"rawtherapee.com",
	"shotcut.org",
	"kdenlive.org",

	// Strumenti di Sviluppo, Browser e Sicurezza FOSS
	"curl.se",
	"git-scm.com",
	"vim.org",
	"neovim.io",
	"torproject.org",
	"signal.org",
	"keepassxc.org",
	"veracrypt.fr",
	"syncthing.net",
	"wireguard.com",
	"wireshark.org",
	"7-zip.org",
	"ffmpeg.org",
	"sqlite.org",
	"postgresql.org",
	"mariadb.org",
	"redis.io",
	"nginx.org",
	"haproxy.org",
	"caddyserver.com",

	// Piattaforme, News e Community Open Source
	"mastodon.social",
	"pixelfed.org",
	"peertube.tv",
	"discourse.org",
	"osnews.com",
	"phoronix.com",
	"lwn.net",
	"slashdot.org",
	"linux.org",
	"linuxquestions.org",
	"itsfoss.com",
	"opensource.com",
	"frenzy.dev",
	"hacker-news.firebaseio.com",
	"fedoramagazine.org",
}

type UrlAndState struct {
	Url string
	State int
	Error error
}

func checkStatus(url string, wg *sync.WaitGroup, channel chan<-UrlAndState) {
	if wg != nil {
		defer wg.Done()
	}
	var client http.Client = http.Client{Timeout : 10 * time.Second}
	response, err := client.Head("https://" + url)
	if err != nil {
		channel <- UrlAndState{Url : url, State : 0, Error : err}
		return
	}
	defer response.Body.Close()
	channel <- UrlAndState{Url : url, State : response.StatusCode, Error : nil}
}

func main() {
	var wg sync.WaitGroup
	var channel chan UrlAndState = make(chan UrlAndState, len(urlList))

	wg.Add(len(urlList))

	for _, url := range urlList {
		go checkStatus(url, &wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for r := range channel {
		fmt.Printf("%s\t%d\n", r.Url, r.State)
	}
}