# MovieFetch

## Disclaimer
- I am NOT responisble for any legal issues or other you encounter with using this program... This program MAY list copyrighted material that could be illegal in your country...
- This program is in a pre-release stage so expect bugs and other issues...

## Intro
Movie fetch is a project of mine written in go to search for and then fetch movies from torrent sites.

## Stability / Testing
This program has tested with:
- VLC on WINDOWS
- QUICKTIME on MACOS

There is no garentee that other combinations will work but im hoping to test more soon.
For now I would recommend using these programs to play the movies.

## Prequisites
- Go (To build from source)
- One of the following media players '[vlc](https://www.videolan.org/vlc/)', '[mpv](https://mpv.io/)', or '[quicktime](https://support.apple.com/downloads/quicktime)' (MacOS version of Quicktime is the only one supported)

## Setup
1. Download a prebuilt binary or build from source. 
1. Create a config.toml with your settings (see config.toml.example included in download) you can get an api key from [tmdb](https://www.themoviedb.org/) or just use mine as its not rate limited.
1. Run the binary!

### To do
- [  ] Tv Show support
- [  ] Edit Searches
- [  ] Add more torrent sites
- [  ] Add more media players
- [  ] Add a gui (Long term)

My [Trello Board](https://trello.com/b/LUevlQih/moviefetch) (Updated more frequently + more stuff)

### External Packages:
Big thanks to:

[clearscreen | GeistInDerSH](https://www.github.com/GeistInDerSH/clearscreen)

[govalidator | asaskevich](https://www.github.com/asaskevich/govalidator)

[rain        | cenkalti](https://www.github.com/cenkalti/rain)

[color       | fatih](https://www.github.com/fatih/color)

[go-toml     | pelletier](https://www.github.com/pelletier/go-toml)

### License
[MIT](https://github.com/MD5-Hashm/moviefetch/blob/main/LICENCE)