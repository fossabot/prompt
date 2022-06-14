// Code generated by 'option-gen'. DO NOT EDIT.

package kube

import (
	prompt "github.com/paralus/prompt/pkg/prompt"
)

var labelOptions = []prompt.Suggest{
	{Text: "--all", Description: "Select all resources, including uninitialized ones, in the namespace of the specified resource types"},
	{Text: "--allow-missing-template-keys", Description: "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats."},
	{Text: "--dry-run", Description: "If true, only print the object that would be sent, without sending it."},
	{Text: "--field-selector", Description: "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type."},
	{Text: "-f", Description: "Filename, directory, or URL to files identifying the resource to update the labels"},
	{Text: "--filename", Description: "Filename, directory, or URL to files identifying the resource to update the labels"},
	{Text: "-k", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	{Text: "--kustomize", Description: "Process the kustomization directory. This flag can't be used together with -f or -R."},
	{Text: "--list", Description: "If true, display the labels for a given resource."},
	{Text: "--local", Description: "If true, label will NOT contact api-server but run locally."},
	{Text: "-o", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--output", Description: "Output format. One of: json|yaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-file."},
	{Text: "--overwrite", Description: "If true, allow labels to be overwritten, otherwise reject label updates that overwrite existing labels."},
	{Text: "--record", Description: "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists."},
	{Text: "-R", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--recursive", Description: "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory."},
	{Text: "--resource-version", Description: "If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource."},
	{Text: "-l", Description: "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)."},
	{Text: "--selector", Description: "Selector (label query) to filter on, not including uninitialized ones, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)."},
	{Text: "--template", Description: "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview]."},
}
