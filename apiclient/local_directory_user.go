package apiclient

import (
	"fmt"
)

func (client *Client) CreateLocalDirectoryUser(parent string, commandPayload map[string]map[string]interface{}) (*TurbotResourceMetadata, error) {
	query := createResourceMutation()
	responseData := &CreateResourceResponse{}
	commandMeta := map[string]string{
		"typeAka":  "tmod:@turbot/turbot-iam#/resource/types/localDirectoryUser",
		"parentId": parent,
	}
	variables := map[string]interface{}{
		"command": map[string]interface{}{
			"payload": commandPayload,
			"meta":    commandMeta,
		},
	}
	// execute api call
	if err := client.doRequest(query, variables, responseData); err != nil {
		return nil, fmt.Errorf("error creating folder: %s", err.Error())
	}
	return &responseData.Resource.Turbot, nil
}

func (client *Client) ReadLocalDirectoryUser(id string) (*LocalDirectoryUser, error) {
	// create a map of the properties we want the graphql query to return
	properties := map[string]string{
		"title":       "title",
		"parent":      "turbot.parentId",
		"email":       "email",
		"status":      "status",
		"displayName": "displayName",
		"givenName":   "givenName",
		"middleName":  "middleName",
		"familyName":  "familyName",
		"picture":     "picture",
	}
	query := readResourceQuery(id, properties)
	responseData := &ReadLocalDirectoryUserResponse{}
	// execute api call
	if err := client.doRequest(query, nil, responseData); err != nil {
		return nil, fmt.Errorf("error reading folder: %s", err.Error())
	}
	return &responseData.Resource, nil
}

func (client *Client) UpdateLocalDirectoryUserResource(id, parent string, commandPayload map[string]map[string]interface{}) (*TurbotResourceMetadata, error) {
	query := updateResourceMutation()
	responseData := &UpdateResourceResponse{}
	// add akas to turbotData
	commandPayload["turbotData"]["akas"] = []string{id}
	commandMeta := map[string]interface{}{
		"typeAka":  "tmod:@turbot/turbot-iam#/resource/types/localDirectoryUser",
		"parentId": parent,
	}
	variables := map[string]interface{}{
		"command": map[string]interface{}{
			"payload": commandPayload,
			"meta":    commandMeta,
		},
	}
	// execute api call
	if err := client.doRequest(query, variables, responseData); err != nil {
		return nil, fmt.Errorf("error creating folder: %s", err.Error())
	}
	return &responseData.Resource.Turbot, nil
}
