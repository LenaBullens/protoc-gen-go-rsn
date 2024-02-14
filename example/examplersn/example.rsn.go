// Code generated by protoc-gen-go-rsn. DO NOT EDIT.
// versions:
// - protoc-gen-go-rsn (devel)
// - protoc            v4.25.1

package examplersn

import (
	errors "errors"
	fmt "fmt"
	regexp "regexp"
)

// LogEntryParentType indicates the possible parent types for resource 'LogEntry'.
type LogEntryParentType = string

const (
	// ProjectLogEntryParentType is a possible parent type for resource 'LogEntry', used for 'project'.
	ProjectLogEntryParentType LogEntryParentType = "project"
	// OrganizationLogEntryParentType is a possible parent type for resource 'LogEntry', used for 'organization'.
	OrganizationLogEntryParentType LogEntryParentType = "organization"
	// FolderLogEntryParentType is a possible parent type for resource 'LogEntry', used for 'folder'.
	FolderLogEntryParentType LogEntryParentType = "folder"
	// BillingAccountLogEntryParentType is a possible parent type for resource 'LogEntry', used for 'billingAccount'.
	BillingAccountLogEntryParentType LogEntryParentType = "billingAccount"
)

// LogEntryParentRsn is the representation of a 'LogEntry' parent resource name.
// This contains all possible identifiers for all parent types. The Type field implicitly declares the
// identifiers used for that context.
type LogEntryParentRsn struct {
	BillingAccountId string
	FolderId         string
	OrganizationId   string
	ProjectId        string
	Type             LogEntryParentType
}

// ResourceName formats r to a resource name.
func (r LogEntryParentRsn) ResourceName() string {
	switch r.Type {
	case ProjectLogEntryParentType:
		return fmt.Sprintf("projects/%s", r.ProjectId)
	case OrganizationLogEntryParentType:
		return fmt.Sprintf("organizations/%s", r.OrganizationId)
	case FolderLogEntryParentType:
		return fmt.Sprintf("folders/%s", r.FolderId)
	case BillingAccountLogEntryParentType:
		return fmt.Sprintf("billingAccounts/%s", r.BillingAccountId)
	}
	return ""
}

// IsZero reports whether r represents the zero resource name.
func (r LogEntryParentRsn) IsZero() bool {
	return len(r.BillingAccountId) == 0 && len(r.FolderId) == 0 && len(r.OrganizationId) == 0 && len(r.ProjectId) == 0 && len(r.Type) == 0
}

var (
	projectLogEntryParentPattern        = regexp.MustCompile(`^projects/([^/]+)$`)
	organizationLogEntryParentPattern   = regexp.MustCompile(`^organizations/([^/]+)$`)
	folderLogEntryParentPattern         = regexp.MustCompile(`^folders/([^/]+)$`)
	billingAccountLogEntryParentPattern = regexp.MustCompile(`^billingAccounts/([^/]+)$`)
)

// ParseLogEntryParentResourceName parses the given parent resource name to a parent resource name.
// Errors for unknown patterns.
func ParseLogEntryParentResourceName(resourceName string) (LogEntryParentRsn, error) {
	if match := projectLogEntryParentPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryParentRsn{
			ProjectId: match[1],
			Type:      ProjectLogEntryParentType,
		}, nil
	}
	if match := organizationLogEntryParentPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryParentRsn{
			OrganizationId: match[1],
			Type:           OrganizationLogEntryParentType,
		}, nil
	}
	if match := folderLogEntryParentPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryParentRsn{
			FolderId: match[1],
			Type:     FolderLogEntryParentType,
		}, nil
	}
	if match := billingAccountLogEntryParentPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryParentRsn{
			BillingAccountId: match[1],
			Type:             BillingAccountLogEntryParentType,
		}, nil
	}
	return LogEntryParentRsn{}, errors.New("invalid parent resource name for resource type 'LogEntry' in service 'logging.googleapis.com'")
}

// LogEntryRsn is the representation of a 'LogEntry' resource name.
type LogEntryRsn struct {
	Parent     LogEntryParentRsn
	LogEntryId string
}

// ResourceName formats r to a resource name.
func (r LogEntryRsn) ResourceName() string {
	switch r.Parent.Type {
	case ProjectLogEntryParentType:
		return fmt.Sprintf("projects/%s/logEntries/%s", r.Parent.ProjectId, r.LogEntryId)
	case OrganizationLogEntryParentType:
		return fmt.Sprintf("organizations/%s/logEntries/%s", r.Parent.OrganizationId, r.LogEntryId)
	case FolderLogEntryParentType:
		return fmt.Sprintf("folders/%s/logEntries/%s", r.Parent.FolderId, r.LogEntryId)
	case BillingAccountLogEntryParentType:
		return fmt.Sprintf("billingAccounts/%s/logEntries/%s", r.Parent.BillingAccountId, r.LogEntryId)
	}
	return ""
}

// IsZero reports whether r represents the zero resource name.
func (r LogEntryRsn) IsZero() bool {
	return r.Parent.IsZero() && len(r.LogEntryId) == 0
}

var (
	projectLogEntryRsnPattern        = regexp.MustCompile(`^projects/([^/]+)/logEntries/([^/]+)$`)
	organizationLogEntryRsnPattern   = regexp.MustCompile(`^organizations/([^/]+)/logEntries/([^/]+)$`)
	folderLogEntryRsnPattern         = regexp.MustCompile(`^folders/([^/]+)/logEntries/([^/]+)$`)
	billingAccountLogEntryRsnPattern = regexp.MustCompile(`^billingAccounts/([^/]+)/logEntries/([^/]+)$`)
)

// ParseLogEntryResourceName parses the given resource name to a resource name.
// Errors for unknown patterns.
func ParseLogEntryResourceName(resourceName string) (LogEntryRsn, error) {
	if match := projectLogEntryRsnPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryRsn{
			Parent: LogEntryParentRsn{
				ProjectId: match[1],
				Type:      ProjectLogEntryParentType,
			},
			LogEntryId: match[2],
		}, nil
	}
	if match := organizationLogEntryRsnPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryRsn{
			Parent: LogEntryParentRsn{
				OrganizationId: match[1],
				Type:           OrganizationLogEntryParentType,
			},
			LogEntryId: match[2],
		}, nil
	}
	if match := folderLogEntryRsnPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryRsn{
			Parent: LogEntryParentRsn{
				FolderId: match[1],
				Type:     FolderLogEntryParentType,
			},
			LogEntryId: match[2],
		}, nil
	}
	if match := billingAccountLogEntryRsnPattern.FindStringSubmatch(resourceName); match != nil {
		return LogEntryRsn{
			Parent: LogEntryParentRsn{
				BillingAccountId: match[1],
				Type:             BillingAccountLogEntryParentType,
			},
			LogEntryId: match[2],
		}, nil
	}
	return LogEntryRsn{}, errors.New("invalid resource name for resource type 'LogEntry' in service 'logging.googleapis.com'")
}

// BookParentType indicates the possible parent types for resource 'Book'.
type BookParentType = string

const (
	// PublisherBookParentType is a possible parent type for resource 'Book', used for 'publisher'.
	PublisherBookParentType BookParentType = "publisher"
)

// BookParentRsn is the representation of a 'Book' parent resource name.
// This contains all possible identifiers for all parent types. The Type field implicitly declares the
// identifiers used for that context.
type BookParentRsn struct {
	PublisherId string
	Type        BookParentType
}

// ResourceName formats r to a resource name.
func (r BookParentRsn) ResourceName() string {
	switch r.Type {
	case PublisherBookParentType:
		return fmt.Sprintf("publishers/%s", r.PublisherId)
	}
	return ""
}

// IsZero reports whether r represents the zero resource name.
func (r BookParentRsn) IsZero() bool {
	return len(r.PublisherId) == 0 && len(r.Type) == 0
}

var (
	publisherBookParentPattern = regexp.MustCompile(`^publishers/([^/]+)$`)
)

// ParseBookParentResourceName parses the given parent resource name to a parent resource name.
// Errors for unknown patterns.
func ParseBookParentResourceName(resourceName string) (BookParentRsn, error) {
	if match := publisherBookParentPattern.FindStringSubmatch(resourceName); match != nil {
		return BookParentRsn{
			PublisherId: match[1],
			Type:        PublisherBookParentType,
		}, nil
	}
	return BookParentRsn{}, errors.New("invalid parent resource name for resource type 'Book' in service 'library.googleapis.com'")
}

// BookRsn is the representation of a 'Book' resource name.
type BookRsn struct {
	Parent BookParentRsn
	BookId string
}

// ResourceName formats r to a resource name.
func (r BookRsn) ResourceName() string {
	switch r.Parent.Type {
	case PublisherBookParentType:
		return fmt.Sprintf("publishers/%s/books/%s", r.Parent.PublisherId, r.BookId)
	}
	return ""
}

// IsZero reports whether r represents the zero resource name.
func (r BookRsn) IsZero() bool {
	return r.Parent.IsZero() && len(r.BookId) == 0
}

var (
	publisherBookRsnPattern = regexp.MustCompile(`^publishers/([^/]+)/books/([^/]+)$`)
)

// ParseBookResourceName parses the given resource name to a resource name.
// Errors for unknown patterns.
func ParseBookResourceName(resourceName string) (BookRsn, error) {
	if match := publisherBookRsnPattern.FindStringSubmatch(resourceName); match != nil {
		return BookRsn{
			Parent: BookParentRsn{
				PublisherId: match[1],
				Type:        PublisherBookParentType,
			},
			BookId: match[2],
		}, nil
	}
	return BookRsn{}, errors.New("invalid resource name for resource type 'Book' in service 'library.googleapis.com'")
}
