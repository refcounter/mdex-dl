# Mangadex Downloader 
![default workflow](https://github.com/refcounter/mdex-dl/actions/workflows/go.yml/badge.svg) ![build workflow](https://github.com/refcounter/mdex-dl/actions/workflows/release.yml/badge.svg)

As the name implies, it is a downloader for [mangadex](https://mangadex.org). I know, there are a bunch of such programs already, but this one emerged from my particular needs. I was using [mangadex-downloader](https://mangadex-downloader.rtfd.io/) --which worked fine; until my network conditions got worse and i couldn't get past the errors (idk why python doesn't work under scarce circunstances (ie, bandwidth)), plus my pc broke and i was forced to my old x86 potato...

## Instalation
### From Source
```bash
git clone https://github.com/refcounter/mdex-dl.git
cd mdex-dl
go build 
mv mdex-dl ~/.local/bin/
```

## Usage
Should be as strainghtforward as running the included binary
```bash
mdex-dl --help
mdex-dl v0.0.1 - A simple mangadex downloader

Flags:

  -ds
    	Data Saver mode (default true)
  -ec int
    	End Chapter
  -help
    	Get help on the 'mdex-dl' command.
  -lang string
    	Translator language (default "en")
  -sc int
    	Start Chapter
  -url string
    	Manga's url from https://mangadex.org

```

## License
mdex-dl is under MIT License. See [LICENSE](./LICENSE)
