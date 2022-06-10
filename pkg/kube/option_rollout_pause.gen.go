// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/paralus/prompt/pkg/prompt"
)

var rolloutPauseOptions = []prompt.Suggest{
	{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to get from a server."},
	{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	{Text: "-o", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}
