<div align="center">
<img src="LOGO.png" alt="hack-browser-data logo" />
</div> 


# HackBrowserData

[![Lint](https://github.com/moonD4rk/HackBrowserData/actions/workflows/lint.yml/badge.svg)](https://github.com/moonD4rk/HackBrowserData/actions/workflows/lint.yml) [![build](https://github.com/moonD4rk/HackBrowserData/actions/workflows/build.yml/badge.svg)](https://github.com/moonD4rk/HackBrowserData/actions/workflows/build.yml) [![Release](https://github.com/moonD4rk/HackBrowserData/actions/workflows/release.yml/badge.svg)](https://github.com/moonD4rk/HackBrowserData/actions/workflows/release.yml)

[中文说明](https://github.com/moonD4rk/HackBrowserData/blob/master/README_ZH.md)

`HackBrowserData` is a command-line tool for decrypting and exporting browser data  ( passwords, history, cookies, bookmarks, credit cards, download records, localStorage and extension ) from the browser. It supports the most popular browsers on the market and runs on Windows, macOS and Linux.

> Disclaimer: This tool is only intended for security research. Users are responsible for all legal and related liabilities resulting from the use of this tool. The original author does not assume any legal responsibility.

## Supported Browser

### Windows
| Browser            | Password | Cookie | Bookmark | History |
|:-------------------|:--------:|:------:|:--------:|:-------:|
| Google Chrome      |    ✅     |   ✅    |    ✅     |    ✅    |
| Google Chrome Beta |    ✅     |   ✅    |    ✅     |    ✅    |
| Chromium           |    ✅     |   ✅    |    ✅     |    ✅    |
| Microsoft Edge     |    ✅     |   ✅    |    ✅     |    ✅    |
| 360 Speed          |    ✅     |   ✅    |    ✅     |    ✅    |
| QQ                 |    ✅     |   ✅    |    ✅     |    ✅    |
| Brave              |    ✅     |   ✅    |    ✅     |    ✅    |
| Opera              |    ✅     |   ✅    |    ✅     |    ✅    |
| OperaGX            |    ✅     |   ✅    |    ✅     |    ✅    |
| Vivaldi            |    ✅     |   ✅    |    ✅     |    ✅    |
| Yandex             |    ✅     |   ✅    |    ✅     |    ✅    |
| CocCoc             |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Beta       |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Dev        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox ESR        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Nightly    |    ✅     |   ✅    |    ✅     |    ✅    |
| Internet Explorer  |    ❌     |   ❌    |    ❌     |    ❌    |


### MacOS

Based on Apple's security policy, some browsers **require a current user password** to decrypt.

| Browser            | Password | Cookie | Bookmark | History |
|:-------------------|:--------:|:------:|:--------:|:-------:|
| Google Chrome      |    ✅     |   ✅    |    ✅     |    ✅    |
| Google Chrome Beta |    ✅     |   ✅    |    ✅     |    ✅    |
| Chromium           |    ✅     |   ✅    |    ✅     |    ✅    |
| Microsoft Edge     |    ✅     |   ✅    |    ✅     |    ✅    |
| Brave              |    ✅     |   ✅    |    ✅     |    ✅    |
| Opera              |    ✅     |   ✅    |    ✅     |    ✅    |
| OperaGX            |    ✅     |   ✅    |    ✅     |    ✅    |
| Vivaldi            |    ✅     |   ✅    |    ✅     |    ✅    |
| CocCoc             |    ✅     |   ✅    |    ✅     |    ✅    |
| Yandex             |    ✅     |   ✅    |    ✅     |    ✅    |
| Arc                |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Beta       |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Dev        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox ESR        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Nightly    |    ✅     |   ✅    |    ✅     |    ✅    |
| Safari             |    ❌     |   ❌    |    ❌     |    ❌    |

### Linux

| Browser            | Password | Cookie | Bookmark | History |
|:-------------------|:--------:|:------:|:--------:|:-------:|
| Google Chrome      |    ✅     |   ✅    |    ✅     |    ✅    |
| Google Chrome Beta |    ✅     |   ✅    |    ✅     |    ✅    |
| Chromium           |    ✅     |   ✅    |    ✅     |    ✅    |
| Microsoft Edge Dev |    ✅     |   ✅    |    ✅     |    ✅    |
| Brave              |    ✅     |   ✅    |    ✅     |    ✅    |
| Opera              |    ✅     |   ✅    |    ✅     |    ✅    |
| Vivaldi            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox            |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Beta       |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Dev        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox ESR        |    ✅     |   ✅    |    ✅     |    ✅    |
| Firefox Nightly    |    ✅     |   ✅    |    ✅     |    ✅    |


## Getting started

### Install

Installation of `HackBrowserData` is dead-simple, just download [the release for your system](https://github.com/moonD4rk/HackBrowserData/releases) and run the binary.

> In some situations, this security tool will be treated as a virus by Windows Defender or other antivirus software and can not be executed. The code is all open source, you can modify and compile by yourself.

### Building from source

only support `go 1.18+` with go generics

```bash
$ git clone https://github.com/moonD4rk/HackBrowserData

$ cd HackBrowserData/cmd/hack-browser-data

$ CGO_ENABLED=1 go build
```

### Cross compile

Need install target OS's `gcc` library, here's an example of use `Mac` building for `Windows` and `Linux`

#### For Windows

```shell
brew install mingw-w64

CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build
```

#### For Linux

````shell
brew install FiloSottile/musl-cross/musl-cross

CC=x86_64-linux-musl-gcc CXX=x86_64-linux-musl-g++ GOARCH=amd64 GOOS=linux CGO_ENABLED=1 go build -ldflags "-linkmode external -extldflags -static"
````

### Run

You can double-click to run, or use command line.

```
PS C:\test> .\hack-browser-data.exe -h
NAME:
   hack-browser-data - Export password|bookmark|cookie|history|credit card|download|localStorage|extension from browser

USAGE:
   [hack-browser-data -b chrome -f json -dir results -cc]
   Export all browingdata(password/cookie/history/bookmark) from browser
   Github Link: https://github.com/moonD4rk/HackBrowserData

VERSION:
   0.5.0

GLOBAL OPTIONS:
   --verbose, --vv                   verbose (default: false)
   --compress, --zip                 compress result to zip (default: false)
   --browser value, -b value         available browsers: all|brave|chrome|chrome-beta|chromium|coccoc|edge|firefox|opera|opera-gx|vivaldi|yandex (default: "all")
   --results-dir value, --dir value  export dir (default: "results")
   --format value, -f value          file name csv|json (default: "csv")
   --profile-path value, -p value    custom profile dir path, get with chrome://version
   --full-export, --full             is export full browsing data (default: true)
   --help, -h                        show help
   --version, -v                     print the version


PS C:\test> .\hack-browser-data.exe -b all -f json --dir results -zip
[NOTICE] [browser.go:46,pickChromium] find browser Chrome success  
[NOTICE] [browser.go:46,pickChromium] find browser Microsoft Edge success  
[NOTICE] [browsingdata.go:59,Output] output to file results/microsoft_edge_download.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/microsoft_edge_password.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/microsoft_edge_creditcard.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/microsoft_edge_bookmark.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/microsoft_edge_cookie.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/microsoft_edge_history.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_history.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_download.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_password.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_creditcard.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_bookmark.json success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_cookie.json success  

```
### Run with custom browser profile folder

If you want to export data from a custom browser profile folder, you can use the `-p` parameter to specify the path of the browser profile folder. PS: use double quotes to wrap the path.
```
PS C:\Users\User\Desktop> .\hack-browser-data.exe -b chrome -p "C:\Users\User\AppData\Local\Microsoft\Edge\User Data\Default"

[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_creditcard.csv success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_bookmark.csv success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_cookie.csv success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_history.csv success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_download.csv success  
[NOTICE] [browsingdata.go:59,Output] output to file results/chrome_password.csv success  
```

### Some other projects based on HackBrowserData
[Sharp-HackBrowserData](https://github.com/S3cur3Th1sSh1t/Sharp-HackBrowserData)

[Reflective-HackBrowserData](https://github.com/idiotc4t/Reflective-HackBrowserData)


## Contributors

![](/CONTRIBUTORS.svg)

## Stargazers over time
[![Star History Chart](https://api.star-history.com/svg?repos=moond4rk/hackbrowserdata&type=Date)](https://github.com/moond4rk/HackBrowserData)


## 404StarLink 2.0 - Galaxy
`HackBrowserData` is a part of 404Team [StarLink-Galaxy](https://github.com/knownsec/404StarLink2.0-Galaxy), if you have any questions about `HackBrowserData` or want to find a partner to communicate with，please refer to the [Starlink group](https://github.com/knownsec/404StarLink2.0-Galaxy#community).
<a href="https://github.com/knownsec/404StarLink2.0-Galaxy" target="_blank"><img src="https://raw.githubusercontent.com/knownsec/404StarLink-Project/master/logo.png" align="middle"/></a>

##  JetBrains OS licenses
``HackBrowserData`` had been being developed with `GoLand` IDE under the **free JetBrains Open Source license(s)** granted by JetBrains s.r.o., hence I would like to express my thanks here.

<a href="https://www.jetbrains.com/?from=HackBrowserData" target="_blank"><img src="https://raw.githubusercontent.com/moonD4rk/staticfiles/master/picture/jetbrains-variant-4.png" width="256" align="middle"/></a>

