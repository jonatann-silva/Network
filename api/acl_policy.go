package api

import (
	"path"

	"jonatann-silva/network/dsn/structs"
)

const (
	aclPoliciesPath = "/api/acl/policies"
)

// ACLPolicies is a handle to the ACL policies API
type ACLPolicies struct {
	client *Client
}

// ACLPolicies returns a handle on the ACL policies endpoints.
func (c *Client) ACLPolicies() *ACLPolicies {
	return &ACLPolicies{client: c}
}

// Create :
func (p *ACLPolicies) Upsert(policy *structs.ACLPolicy) (*structs.ACLPolicy, error) {

	var rcvPolicy *structs.ACLPolicy
	err := p.client.createResource(aclPoliciesPath, policy, &rcvPolicy)
	if err != nil {
		return nil, err
	}

	return policy, nil
}

// Delete :
func (p *ACLPolicies) Delete(name string) error {

	err := p.client.deleteResource(name, aclPoliciesPath, nil)
	if err != nil {
		return err
	}

	return nil
}

// Get :
func (p *ACLPolicies) Get(name string) (*structs.ACLPolicy, error) {

	out := &structs.ACLPolicy{}
	err := p.client.getResource(aclPoliciesPath, name, out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// List :
func (p *ACLPolicies) List() ([]*structs.ACLPolicyListStub, error) {

	var items []*structs.ACLPolicyListStub
	err := p.client.listResources(path.Join(aclPoliciesPath, "/"), nil, &items)
	if err != nil {
		return nil, err
	}

	return items, nil
}
