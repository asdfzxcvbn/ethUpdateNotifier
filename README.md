# ethUpdateNotifier
my new update notifier, written for SultanMods/ETHSign in go!

## setup instructions
limited setup instructions cause i dont really care about this being FOSS in the first place, since the code quality isnt that good.

1. edit apps.json, it's a list of appstore links
2. edit consts.go as necssary
3. `go get -u`
4. `go build`
5. enjoy your very nice go binary (it will make `versions.db` at runtime)