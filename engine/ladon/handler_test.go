package ladon

import (
	"fmt"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ory/keto/sdk/go/keto/client"
	"github.com/ory/keto/sdk/go/keto/client/engines"
	"github.com/ory/keto/sdk/go/keto/models"

	"github.com/gobuffalo/packr"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/urfave/negroni"

	"github.com/ory/herodot"
	"github.com/ory/keto/engine"
	"github.com/ory/keto/storage"
)

func nc(t *testing.T, u string) *client.OryKeto {
	uu, err := url.ParseRequestURI(u)
	require.NoError(t, err)

	return client.NewHTTPClientWithConfig(nil, &client.TransportConfig{
		Host:     uu.Host,
		BasePath: uu.Path,
		Schemes:  []string{uu.Scheme},
	})
}

func TestAllowed(t *testing.T) {
	box := packr.NewBox("./rego")
	compiler, err := engine.NewCompiler(box, logrus.New())
	require.NoError(t, err)

	s := storage.NewMemoryManager()
	sh := storage.NewHandler(s, herodot.NewJSONWriter(nil))
	e := engine.NewEngine(compiler, herodot.NewJSONWriter(nil))
	le := NewEngine(s, sh, e, herodot.NewJSONWriter(nil))

	n := negroni.Classic()
	r := httprouter.New()
	le.Register(r)
	n.UseHandler(r)

	ts := httptest.NewServer(n)
	defer ts.Close()

	cl := nc(t, ts.URL)
	for _, f := range EnabledFlavors {
		t.Run(fmt.Sprintf("flavor=%s", f), func(t *testing.T) {
			t.Run(fmt.Sprint("action=create"), func(t *testing.T) {
				for _, p := range policies[f] {
					t.Run(fmt.Sprintf("policy=%s", p.ID), func(t *testing.T) {
						_, err := cl.Engines.UpsertOryAccessControlPolicy(engines.NewUpsertOryAccessControlPolicyParams().WithFlavor(f).WithBody(toSwaggerPolicy(p)))
						require.NoError(t, err)
					})
				}
				for _, r := range roles[f] {
					t.Run(fmt.Sprintf("role=%s", r.ID), func(t *testing.T) {
						_, err := cl.Engines.UpsertOryAccessControlPolicyRole(engines.NewUpsertOryAccessControlPolicyRoleParams().WithFlavor(f).WithBody(toSwaggerRole(r)))
						require.NoError(t, err)
					})
				}
			})

			t.Run("action=authorize", func(t *testing.T) {
				for k, c := range requests[f] {
					t.Run(fmt.Sprintf("case=%d", k), func(t *testing.T) {
						d, err := cl.Engines.DoOryAccessControlPoliciesAllow(engines.NewDoOryAccessControlPoliciesAllowParams().WithFlavor(f).WithBody(&c.req))
						if c.allowed {
							require.NoError(t, err)
							assert.Equal(t, c.allowed, *d.Payload.Allowed)
						} else {
							require.IsType(t, engines.NewDoOryAccessControlPoliciesAllowForbidden(), err)
							assert.Equal(t, c.allowed, *(err.(*engines.DoOryAccessControlPoliciesAllowForbidden).Payload.Allowed))
						}
					})
				}
			})
		})
	}
}

func TestValidatePolicy(t *testing.T) {
	_, err := validatePolicy(Policy{})
	require.Error(t, err)

	_, err = validatePolicy(Policy{Effect: "bar"})
	require.Error(t, err)

	p, err := validatePolicy(Policy{Effect: "allow"})
	require.NoError(t, err)
	assert.NotEmpty(t, p.ID)

	p, err = validatePolicy(Policy{Effect: "deny", ID: "foo"})
	require.NoError(t, err)
	assert.Equal(t, "foo", p.ID)
}

func crudts() *httptest.Server {
	s := storage.NewMemoryManager()
	sh := storage.NewHandler(s, herodot.NewJSONWriter(nil))
	e := NewEngine(s, sh, nil, herodot.NewJSONWriter(nil))
	r := httprouter.New()
	e.Register(r)
	return httptest.NewServer(r)
}

func toSwaggerPolicy(p Policy) *models.OryAccessControlPolicy {
	return &models.OryAccessControlPolicy{
		Actions:     p.Actions,
		ID:          p.ID,
		Resources:   p.Resources,
		Subjects:    p.Subjects,
		Effect:      p.Effect,
		Conditions:  p.Conditions,
		Description: p.Description,
	}
}

func fromSwaggerPolicy(p models.OryAccessControlPolicy) Policy {
	return Policy{
		Actions:     p.Actions,
		ID:          p.ID,
		Resources:   p.Resources,
		Subjects:    p.Subjects,
		Effect:      p.Effect,
		Conditions:  p.Conditions,
		Description: p.Description,
	}
}

func toSwaggerRole(r Role) *models.OryAccessControlPolicyRole {
	return &models.OryAccessControlPolicyRole{
		Members: r.Members,
		ID:      r.ID,
	}
}

func fromSwaggerRole(r models.OryAccessControlPolicyRole) Role {
	return Role{
		Members: r.Members,
		ID:      r.ID,
	}
}

func TestPolicyCRUD(t *testing.T) {
	ts := crudts()
	defer ts.Close()

	c := nc(t, ts.URL)
	for _, f := range EnabledFlavors {
		for l, p := range policies[f] {
			_, err := c.Engines.GetOryAccessControlPolicy(engines.NewGetOryAccessControlPolicyParams().WithFlavor(f).WithID(p.ID))
			require.Error(t, err)

			_, err = c.Engines.UpsertOryAccessControlPolicy(engines.NewUpsertOryAccessControlPolicyParams().WithFlavor(f).WithBody(toSwaggerPolicy(p)))
			require.NoError(t, err)

			o, err := c.Engines.GetOryAccessControlPolicy(engines.NewGetOryAccessControlPolicyParams().WithFlavor(f).WithID(p.ID))
			require.NoError(t, err)
			assert.Equal(t, p, fromSwaggerPolicy(*o.Payload))

			limit, offset := int64(100), int64(0)
			os, err := c.Engines.ListOryAccessControlPolicies(engines.NewListOryAccessControlPoliciesParams().WithFlavor(f).WithLimit(&limit).WithOffset(&offset))
			require.NoError(t, err)

			var ps Policies
			for _, v := range os.Payload {
				ps = append(ps, fromSwaggerPolicy(*v))
			}

			assert.Equal(t, ps, policies[f][:l+1])
		}

		for _, p := range policies[f] {
			_, err := c.Engines.DeleteOryAccessControlPolicy(engines.NewDeleteOryAccessControlPolicyParams().WithFlavor(f).WithID(p.ID))
			require.NoError(t, err)

			_, err = c.Engines.GetOryAccessControlPolicy(engines.NewGetOryAccessControlPolicyParams().WithFlavor(f).WithID(p.ID))
			require.Error(t, err)
		}
	}
}

func TestRoleCRUD(t *testing.T) {
	ts := crudts()
	defer ts.Close()

	c := nc(t, ts.URL)
	for _, f := range EnabledFlavors {
		for l, r := range roles[f] {
			_, err := c.Engines.GetOryAccessControlPolicyRole(engines.NewGetOryAccessControlPolicyRoleParams().WithFlavor(f).WithID(r.ID))
			require.Error(t, err)

			ou, err := c.Engines.UpsertOryAccessControlPolicyRole(engines.NewUpsertOryAccessControlPolicyRoleParams().WithFlavor(f).WithBody(toSwaggerRole(r)))
			require.NoError(t, err)
			require.EqualValues(t, r, fromSwaggerRole(*ou.Payload))

			o, err := c.Engines.GetOryAccessControlPolicyRole(engines.NewGetOryAccessControlPolicyRoleParams().WithFlavor(f).WithID(r.ID))
			require.NoError(t, err)
			require.EqualValues(t, r, fromSwaggerRole(*o.Payload))

			limit, offset := int64(100), int64(0)
			os, err := c.Engines.ListOryAccessControlPolicyRoles(engines.NewListOryAccessControlPolicyRolesParams().WithFlavor(f).WithLimit(&limit).WithOffset(&offset))
			require.NoError(t, err)

			var ps Roles
			for _, v := range os.Payload {
				ps = append(ps, fromSwaggerRole(*v))
			}

			assert.Equal(t, ps, roles[f][:l+1])
		}

		for _, r := range roles[f] {
			_, err := c.Engines.DeleteOryAccessControlPolicyRole(engines.NewDeleteOryAccessControlPolicyRoleParams().WithFlavor(f).WithID(r.ID))
			require.NoError(t, err)

			_, err = c.Engines.GetOryAccessControlPolicyRole(engines.NewGetOryAccessControlPolicyRoleParams().WithFlavor(f).WithID(r.ID))
			require.Error(t, err)
		}
	}
}
