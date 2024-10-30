# Wasted-On-League-Terminal

Terminal based League of Legends playtime fetcher written in Go.

## Prerequisites

Go Version 1.22.7 or above.

## Building

1. Clone the Directory.
2. Navigate to the directory in the terminal. Then execute `go get owos02/Wasted-On-League-Terminal`.
3. Call `go build -o bin/`.

You should now have a Binary in the newly generated `/bin` folder that you can use with `./Wasted-On-League-Terminal`.

> [!NOTE]
> Depending on the OS you will need to make the Binary runable with: `chmod +x Wasted-On-League-Terminal`.

## Usage

```
---Wasted-On-League-Terminal---

Usage: WoL-Reborn [Username#GameTag|WoL-LINK] [Server]
Fetches time played in hours from WoL.gg.

Examples:
./Wasted-On-League-Terminal "the inescapable#EUW" "EUW"
./Wasted-On-League-Terminal "https://wol.gg/stats/euw/theinescapable-euw/"
```
