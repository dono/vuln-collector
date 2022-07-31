# vuln-collector
Dump vulnerability data output by [vulsio/gost](https://github.com/vulsio/gost) and [vulsio/go-cve-dictionary](https://github.com/vulsio/go-cve-dictionary) into a single SQLite file

## Requires
- SQLite3
- git
- gcc
- go v1.7.1 or later

## Supports
- [Red Hat](https://access.redhat.com/security/security-updates/)
- [Debian](https://security-tracker.debian.org/tracker/)
- [Ubuntu](https://ubuntu.com/security/cves)
- [NVD](https://nvd.nist.gov/vuln/search)


## Usage
- `git clone https://github.com/dono/vuln-collector.git`
- `go build cmd/vuln-collector`
- `./vuln-collector -out ./vulns.sqlite3`

## Notes
- Total DB data is about 1.4Total DB data is about 1.2 GB as of July 2022. GB as of July 2022.
