module github.com/owasp-amass/amass/v4

go 1.23.1

require (
	github.com/99designs/gqlgen v0.17.55
	github.com/PuerkitoBio/goquery v1.10.0
	github.com/PuerkitoBio/purell v1.2.1
	github.com/adrg/strutil v0.3.1
	github.com/caffix/fullname_parser v0.0.0-20240817200809-1b9b04da88d0
	github.com/caffix/jarm-go v0.0.0-20240920030848-1c7ab2423494
	github.com/caffix/pipeline v0.2.4
	github.com/caffix/queue v0.2.0
	github.com/caffix/stringset v0.2.0
	github.com/cheggaaa/pb/v3 v3.1.5
	github.com/fatih/color v1.17.0
	github.com/geziyor/geziyor v0.0.0-20240812061556-229b8ca83ac1
	github.com/glebarez/sqlite v1.11.0
	github.com/go-ini/ini v1.67.0
	github.com/google/uuid v1.6.0
	github.com/gorilla/websocket v1.5.3
	github.com/hashicorp/go-multierror v1.1.1
	github.com/likexian/whois v1.15.5
	github.com/likexian/whois-parser v1.24.20
	github.com/miekg/dns v1.1.62
	github.com/nyaruka/phonenumbers v1.4.1
	github.com/openrdap/rdap v0.9.1
	github.com/owasp-amass/asset-db v0.9.0
	github.com/owasp-amass/open-asset-model v0.9.1
	github.com/owasp-amass/resolve v0.8.1
	github.com/rubenv/sql-migrate v1.7.0
	github.com/samber/slog-common v0.17.1
	github.com/samber/slog-syslog/v2 v2.5.0
	github.com/stretchr/testify v1.9.0
	github.com/tylertreat/BoomFilters v0.0.0-20210315201527-1a82519a3e43
	github.com/vektah/gqlparser/v2 v2.5.18
	github.com/yl2chen/cidranger v1.0.2
	go.uber.org/ratelimit v0.3.1
	golang.org/x/net v0.30.0
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/postgres v1.5.9
	gorm.io/gorm v1.25.12
	mvdan.cc/xurls/v2 v2.5.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/VividCortex/ewma v1.2.0 // indirect
	github.com/VividCortex/gohistogram v1.0.0 // indirect
	github.com/agnivade/levenshtein v1.2.0 // indirect
	github.com/alecthomas/kingpin/v2 v2.4.0 // indirect
	github.com/alecthomas/units v0.0.0-20240927000941-0f3dac36c52b // indirect
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/benbjohnson/clock v1.3.5 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chromedp/cdproto v0.0.0-20241014181340-cb3a7a1d51d7 // indirect
	github.com/chromedp/chromedp v0.11.0 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/d4l3k/messagediff v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/glebarez/go-sqlite v1.22.0 // indirect
	github.com/go-gorp/gorp/v3 v3.1.0 // indirect
	github.com/go-kit/kit v0.13.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.4.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.1 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/likexian/gokit v0.25.15 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.60.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/samber/lo v1.47.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/temoto/robotstxt v1.1.2 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/exp v0.0.0-20241009180824-f66d83c29e7c // indirect
	golang.org/x/mod v0.21.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	golang.org/x/tools v0.26.0 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
	gorm.io/datatypes v1.2.3 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	modernc.org/libc v1.61.0 // indirect
	modernc.org/mathutil v1.6.0 // indirect
	modernc.org/memory v1.8.0 // indirect
	modernc.org/sqlite v1.33.1 // indirect
)
