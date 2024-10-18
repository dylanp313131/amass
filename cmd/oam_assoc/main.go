// Copyright © by Jeff Foley 2017-2024. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.
// SPDX-License-Identifier: Apache-2.0

// oam_assoc: Analyze collected OAM data to identify assets associated with the seed data
//
//	+----------------------------------------------------------------------------+
//	| ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░  OWASP Amass  ░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░ |
//	+----------------------------------------------------------------------------+
//	|      .+++:.            :                             .+++.                 |
//	|    +W@@@@@@8        &+W@#               o8W8:      +W@@@@@@#.   oW@@@W#+   |
//	|   &@#+   .o@##.    .@@@o@W.o@@o       :@@#&W8o    .@#:  .:oW+  .@#+++&#&   |
//	|  +@&        &@&     #@8 +@W@&8@+     :@W.   +@8   +@:          .@8         |
//	|  8@          @@     8@o  8@8  WW    .@W      W@+  .@W.          o@#:       |
//	|  WW          &@o    &@:  o@+  o@+   #@.      8@o   +W@#+.        +W@8:     |
//	|  #@          :@W    &@+  &@+   @8  :@o       o@o     oW@@W+        oW@8    |
//	|  o@+          @@&   &@+  &@+   #@  &@.      .W@W       .+#@&         o@W.  |
//	|   WW         +@W@8. &@+  :&    o@+ #@      :@W&@&         &@:  ..     :@o  |
//	|   :@W:      o@# +Wo &@+        :W: +@W&o++o@W. &@&  8@#o+&@W.  #@:    o@+  |
//	|    :W@@WWWW@@8       +              :&W@@@@&    &W  .o#@@W&.   :W@WWW@@&   |
//	|      +o&&&&+.                                                    +oooo.    |
//	+----------------------------------------------------------------------------+
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"github.com/caffix/stringset"
	"github.com/fatih/color"
	"github.com/owasp-amass/amass/v4/config"
	"github.com/owasp-amass/amass/v4/utils"
	"github.com/owasp-amass/amass/v4/utils/afmt"
	assetdb "github.com/owasp-amass/asset-db"
	dbt "github.com/owasp-amass/asset-db/types"
	"github.com/owasp-amass/open-asset-model/domain"
	oamreg "github.com/owasp-amass/open-asset-model/registration"
)

const (
	timeFormat = "01/02 15:04:05 2006 MST"
	usageMsg   = "[options] [-since '" + timeFormat + "'] " + "-d domain"
)

type assocArgs struct {
	Domains *stringset.Set
	Since   string
	Options struct {
		NoColor bool
		Silent  bool
	}
	Filepaths struct {
		ConfigFile string
		Directory  string
		Domains    string
	}
}

func main() {
	var args assocArgs
	var help1, help2, verbose bool
	assocCommand := flag.NewFlagSet("assoc", flag.ContinueOnError)

	args.Domains = stringset.New()
	defer args.Domains.Close()

	assocBuf := new(bytes.Buffer)
	assocCommand.SetOutput(assocBuf)

	assocCommand.BoolVar(&help1, "h", false, "Show the program usage message")
	assocCommand.BoolVar(&help2, "help", false, "Show the program usage message")
	assocCommand.BoolVar(&verbose, "v", false, "Show additional information about the associated assets")
	assocCommand.Var(args.Domains, "d", "Domain names separated by commas (can be used multiple times)")
	assocCommand.StringVar(&args.Since, "since", "", "Exclude all assets discovered before (format: "+timeFormat+")")
	assocCommand.BoolVar(&args.Options.NoColor, "nocolor", false, "Disable colorized output")
	assocCommand.BoolVar(&args.Options.Silent, "silent", false, "Disable all output during execution")
	assocCommand.StringVar(&args.Filepaths.ConfigFile, "config", "", "Path to the YAML configuration file")
	assocCommand.StringVar(&args.Filepaths.Directory, "dir", "", "Path to the directory containing the graph database")
	assocCommand.StringVar(&args.Filepaths.Domains, "df", "", "Path to a file providing registered domain names")

	var usage = func() {
		afmt.G.Fprintf(color.Error, "Usage: %s %s\n\n", path.Base(os.Args[0]), usageMsg)
		assocCommand.PrintDefaults()
		afmt.G.Fprintln(color.Error, assocBuf.String())
	}

	if len(os.Args) < 2 {
		usage()
		return
	}
	if err := assocCommand.Parse(os.Args[1:]); err != nil {
		afmt.R.Fprintf(color.Error, "%v\n", err)
		os.Exit(1)
	}
	if help1 || help2 {
		usage()
		return
	}
	if args.Options.NoColor {
		color.NoColor = true
	}
	if args.Options.Silent {
		color.Output = io.Discard
		color.Error = io.Discard
	}
	if args.Filepaths.Domains != "" {
		list, err := config.GetListFromFile(args.Filepaths.Domains)
		if err != nil {
			afmt.R.Fprintf(color.Error, "Failed to parse the domain names file: %v\n", err)
			os.Exit(1)
		}
		args.Domains.InsertMany(list...)
	}
	if args.Domains.Len() == 0 {
		afmt.R.Fprintln(color.Error, "No root domain names were provided")
		os.Exit(1)
	}

	var err error
	var start time.Time
	if args.Since != "" {
		start, err = time.Parse(timeFormat, args.Since)
		if err != nil {
			afmt.R.Fprintf(color.Error, "%s is not in the correct format: %s\n", args.Since, timeFormat)
			os.Exit(1)
		}
	}

	cfg := config.NewConfig()
	// Check if a configuration file was provided, and if so, load the settings
	if err := config.AcquireConfig(args.Filepaths.Directory, args.Filepaths.ConfigFile, cfg); err == nil {
		if args.Filepaths.Directory == "" {
			args.Filepaths.Directory = cfg.Dir
		}
		if args.Domains.Len() == 0 {
			args.Domains.InsertMany(cfg.Domains()...)
		}
	} else if args.Filepaths.ConfigFile != "" {
		afmt.R.Fprintf(color.Error, "Failed to load the configuration file: %v\n", err)
		os.Exit(1)
	}
	// Connect with the graph database containing the enumeration data
	db := utils.OpenGraphDatabase(cfg)
	if db == nil {
		afmt.R.Fprintln(color.Error, "Failed to connect with the database")
		os.Exit(1)
	}

	for _, name := range args.Domains.Slice() {
		for i, assoc := range getAssociations(name, start, db) {
			if i != 0 {
				fmt.Println()
			}

			var rel string
			switch v := assoc.Asset.(type) {
			case *oamreg.DomainRecord:
				rel = "registrant_contact"
				afmt.G.Fprintln(color.Output, v.Domain)
				if verbose {
					fmt.Fprintf(color.Output, "%s%s\n%s%s\n", afmt.Blue("Name: "),
						afmt.Green(v.Name), afmt.Blue("Expiration: "), afmt.Green(v.ExpirationDate))
				}
			case *oamreg.AutnumRecord:
				rel = "registrant"
				afmt.G.Fprintln(color.Output, v.Handle)
				if verbose {
					fmt.Fprintf(color.Output, "%s%s\n%s%s\n%s%s\n", afmt.Blue("Name: "), afmt.Green(v.Name),
						afmt.Blue("Status: "), afmt.Green(v.Status[0]), afmt.Blue("Updated: "), afmt.Green(v.UpdatedDate))
				}
			case *oamreg.IPNetRecord:
				rel = "registrant"
				afmt.G.Fprintln(color.Output, v.CIDR.String())
				if verbose {
					fmt.Fprintf(color.Output, "%s%s\n%s%s\n%s%s\n", afmt.Blue("Name: "), afmt.Green(v.Name),
						afmt.Blue("Status: "), afmt.Green(v.Status[0]), afmt.Blue("Updated: "), afmt.Green(v.UpdatedDate))
				}
			}

			if verbose {
				afmt.B.Fprintln(color.Output, "Registrant: ")
				printContactInfo(assoc, rel, start, db)
				fmt.Println()
			}
		}
	}
}

func printContactInfo(assoc *dbt.Asset, regrel string, since time.Time, db *assetdb.AssetDB) {
	var contact *dbt.Asset
	if rels, err := db.OutgoingRelations(assoc, since, regrel); err == nil && len(rels) > 0 {
		if a, err := db.FindById(rels[0].ToAsset.ID, since); err == nil && a != nil {
			contact = a
		}
	}
	if contact == nil {
		return
	}

	for _, out := range []string{"person", "organization", "location", "phone", "email"} {
		if rels, err := db.OutgoingRelations(contact, since, out); err == nil && len(rels) > 0 {
			for _, rel := range rels {
				if a, err := db.FindById(rel.ToAsset.ID, since); err == nil && a != nil {
					fmt.Fprintf(color.Output, "%s%s%s\n",
						afmt.Blue(string(a.Asset.AssetType())), afmt.Blue(": "), afmt.Green(a.Asset.Key()))
				}
			}
		}
	}
}

func getAssociations(name string, since time.Time, db *assetdb.AssetDB) []*dbt.Asset {
	if !since.IsZero() {
		since = since.UTC()
	}

	var results []*dbt.Asset
	fqdns, err := db.FindByContent(&domain.FQDN{Name: name}, since)
	if err != nil || len(fqdns) == 0 {
		return results
	}

	var assets []*dbt.Asset
	for _, fqdn := range fqdns {
		if rels, err := db.OutgoingRelations(fqdn, since, "registration"); err == nil && len(rels) > 0 {
			for _, rel := range rels {
				if a, err := db.FindById(rel.ToAsset.ID, since); err == nil && a != nil {
					assets = append(assets, a)
				}
			}
		}
	}

	set := stringset.New()
	defer set.Close()

	for _, asset := range assets {
		set.Insert(asset.ID)
	}

	for findings := assets; len(findings) > 0; {
		assets = findings
		findings = []*dbt.Asset{}

		for _, a := range assets {
			if rels, err := db.OutgoingRelations(a, since, "associated_with"); err == nil && len(rels) > 0 {
				for _, rel := range rels {
					asset, err := db.FindById(rel.ToAsset.ID, since)
					if err != nil || asset == nil {
						continue
					}

					if !set.Has(asset.ID) {
						set.Insert(asset.ID)
						findings = append(findings, asset)
						results = append(results, asset)
					}
				}
			}
		}
	}
	return results
}
