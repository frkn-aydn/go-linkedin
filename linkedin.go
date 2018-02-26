package GoLinkedin

import "strings"

type LinkedIn struct {
	ClientID     string
	ClientSecret string
	CallbackURL  string
	Scope        string
	APIVersion   string
	APICalls     map[string]string
	BaseURL      string
}

func New(ClientID, ClientSecret, CallbackURL, APIVersion string, Scope []string) (*LinkedIn, string) {
	if ClientID == "" || ClientSecret == "" || CallbackURL == "" || APIVersion == "" || len(Scope) == 0 {
		return nil, "Missing fields."
	}
	o := &LinkedIn{}
	o.ClientID = ClientID
	o.ClientSecret = ClientSecret
	o.CallbackURL = CallbackURL
	o.Scope = strings.Join(Scope, " ")
	o.APIVersion = APIVersion
	o.BaseURL = "https://api.linkedin.com"
	o.APICalls = make(map[string]string)

	// BASIC APIS
	o.APICalls["myProfile"] = "people/~:(first-name,last-name,headline,picture-urls::(original),id,location,industry,current-share,num-connections,num-connections-capped,summary,specialties,positions,email-address)"
	o.APICalls["myConnections"] = "people/~/connections"
	o.APICalls["myNetworkShares"] = "people/~/shares"
	o.APICalls["myNetworksUpdates"] = "people/~/network/updates"
	o.APICalls["myNetworkUpdates"] = "people/~/network/updates?scope=self"

	// PEOPLE SEARCH APIS
	o.APICalls["peopleSearchWithKeywords"] = "people-search:(people:(id,first-name,last-name,picture-url,headline),num-results,facets)?keywords=Hacker+in+Residence"
	o.APICalls["peopleSearchWithFacets"] = "people-search:(people,facets)?facet=location,us:84"

	// GROUPS APIS
	o.APICalls["myGroups"] = "people/~/group-memberships?membership-state=member"
	o.APICalls["groupSuggestions"] = "people/~/suggestions/groups"
	o.APICalls["groupPosts"] = "groups/12345/posts:(title,summary,creator)?order=recency"
	o.APICalls["groupDetails"] = "groups/12345:(id,name,short-description,description,posts)"

	// COMPANY APIS
	o.APICalls["myFollowingCompanies"] = "people/~/following/companies"
	o.APICalls["myFollowCompanySuggestions"] = "people/~/suggestions/to-follow/companies"
	o.APICalls["companyDetails"] = "companies/1337:(id,name,description,industry,logo-url)"
	o.APICalls["companySearch"] = "company-search:(companies,facets)?facet=location,us:84"

	// JOBS APIS
	o.APICalls["myJobSuggestions"] = "people/~/suggestions/job-suggestions"
	o.APICalls["myJobBookmarks"] = "people/~/job-bookmarks"
	o.APICalls["jobDetails"] = "jobs/1452577:(id,company:(name),position:(title))"
	o.APICalls["jobSearch"] = "job-search:(jobs,facets)?facet=location,us:84"
	return o, ""
}
