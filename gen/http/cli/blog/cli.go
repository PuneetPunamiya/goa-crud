// Code generated by goa v3.1.1, DO NOT EDIT.
//
// blog HTTP client CLI support package
//
// Command:
// $ goa gen crud/design

package cli

import (
	blogc "crud/gen/http/blog/client"
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `blog (create|list|remove|update)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` blog create --body '{
      "comments": [
         "Aperiam id corporis voluptatibus inventore.",
         "Sed et aut voluptatem et voluptas.",
         "Consequatur sunt asperiores natus iste eaque."
      ],
      "id": 3752768668,
      "name": "4dp"
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		blogFlags = flag.NewFlagSet("blog", flag.ContinueOnError)

		blogCreateFlags    = flag.NewFlagSet("create", flag.ExitOnError)
		blogCreateBodyFlag = blogCreateFlags.String("body", "REQUIRED", "")

		blogListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		blogRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		blogRemoveIDFlag = blogRemoveFlags.String("id", "REQUIRED", "ID of blog to remove")

		blogUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		blogUpdateBodyFlag = blogUpdateFlags.String("body", "REQUIRED", "")
		blogUpdateIDFlag   = blogUpdateFlags.String("id", "REQUIRED", "ID of blog to be updated")
	)
	blogFlags.Usage = blogUsage
	blogCreateFlags.Usage = blogCreateUsage
	blogListFlags.Usage = blogListUsage
	blogRemoveFlags.Usage = blogRemoveUsage
	blogUpdateFlags.Usage = blogUpdateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "blog":
			svcf = blogFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "blog":
			switch epn {
			case "create":
				epf = blogCreateFlags

			case "list":
				epf = blogListFlags

			case "remove":
				epf = blogRemoveFlags

			case "update":
				epf = blogUpdateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "blog":
			c := blogc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create":
				endpoint = c.Create()
				data, err = blogc.BuildCreatePayload(*blogCreateBodyFlag)
			case "list":
				endpoint = c.List()
				data = nil
			case "remove":
				endpoint = c.Remove()
				data, err = blogc.BuildRemovePayload(*blogRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = blogc.BuildUpdatePayload(*blogUpdateBodyFlag, *blogUpdateIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// blogUsage displays the usage of the blog command and its subcommands.
func blogUsage() {
	fmt.Fprintf(os.Stderr, `The blog service gives blog details.
Usage:
    %s [globalflags] blog COMMAND [flags]

COMMAND:
    create: Add new blog and return its ID.
    list: List all entries
    remove: Remove blog from storage
    update: Updating the existing blog

Additional help:
    %s blog COMMAND --help
`, os.Args[0], os.Args[0])
}
func blogCreateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog create -body JSON

Add new blog and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` blog create --body '{
      "comments": [
         "Aperiam id corporis voluptatibus inventore.",
         "Sed et aut voluptatem et voluptas.",
         "Consequatur sunt asperiores natus iste eaque."
      ],
      "id": 3752768668,
      "name": "4dp"
   }'
`, os.Args[0])
}

func blogListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog list

List all entries

Example:
    `+os.Args[0]+` blog list
`, os.Args[0])
}

func blogRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog remove -id UINT32

Remove blog from storage
    -id UINT32: ID of blog to remove

Example:
    `+os.Args[0]+` blog remove --id 3407509870
`, os.Args[0])
}

func blogUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] blog update -body JSON -id UINT32

Updating the existing blog
    -body JSON: 
    -id UINT32: ID of blog to be updated

Example:
    `+os.Args[0]+` blog update --body '{
      "comments": [
         "Necessitatibus id.",
         "Quos distinctio blanditiis cum totam molestiae dolorum.",
         "Impedit ipsam cupiditate soluta consequatur beatae soluta.",
         "Aut tempore dolore."
      ],
      "name": "Tenetur numquam iste et eos."
   }' --id 2497136214
`, os.Args[0])
}