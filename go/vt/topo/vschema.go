/*
Copyright 2019 The Vitess Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package topo

import (
	"context"
	"path"

	"google.golang.org/protobuf/proto"

	"vitess.io/vitess/go/vt/log"
	"vitess.io/vitess/go/vt/vterrors"

	vschemapb "vitess.io/vitess/go/vt/proto/vschema"
)

// SaveVSchema saves a Vschema. A valid Vschema should be passed in. It does not verify its correctness.
// If the VSchema is empty, just remove it.
func (ts *Server) SaveVSchema(ctx context.Context, keyspace string, vschema *vschemapb.Keyspace) error {
	nodePath := path.Join(KeyspacesPath, keyspace, VSchemaFile)
	data, err := vschema.MarshalVT()
	if err != nil {
		return err
	}

	_, err = ts.globalCell.Update(ctx, nodePath, data, nil)
	if err != nil {
		log.Errorf("failed to update vschema for keyspace %s: %v", keyspace, err)
	} else {
		log.Infof("successfully updated vschema for keyspace %s: %+v", keyspace, vschema)
	}
	return err
}

// DeleteVSchema delete the keyspace if it exists
func (ts *Server) DeleteVSchema(ctx context.Context, keyspace string) error {
	log.Infof("deleting vschema for keyspace %s", keyspace)
	nodePath := path.Join(KeyspacesPath, keyspace, VSchemaFile)
	return ts.globalCell.Delete(ctx, nodePath, nil)
}

// GetVSchema fetches the vschema from the topo.
func (ts *Server) GetVSchema(ctx context.Context, keyspace string) (*vschemapb.Keyspace, error) {
	nodePath := path.Join(KeyspacesPath, keyspace, VSchemaFile)
	data, _, err := ts.globalCell.Get(ctx, nodePath)
	if err != nil {
		return nil, err
	}
	var vs vschemapb.Keyspace
	err = proto.Unmarshal(data, &vs)
	if err != nil {
		return nil, vterrors.Wrapf(err, "bad vschema data: %q", data)
	}
	return &vs, nil
}

// EnsureVSchema makes sure that a vschema is present for this keyspace or creates a blank one if it is missing
func (ts *Server) EnsureVSchema(ctx context.Context, keyspace string) error {
	vschema, err := ts.GetVSchema(ctx, keyspace)
	if err != nil && !IsErrType(err, NoNode) {
		log.Infof("error in getting vschema for keyspace %s: %v", keyspace, err)
	}
	if vschema == nil || IsErrType(err, NoNode) {
		err = ts.SaveVSchema(ctx, keyspace, &vschemapb.Keyspace{
			Sharded:  false,
			Vindexes: make(map[string]*vschemapb.Vindex),
			Tables:   make(map[string]*vschemapb.Table),
		})
		if err != nil {
			log.Errorf("could not create blank vschema: %v", err)
			return err
		}
	}
	return nil
}

// SaveRoutingRules saves the routing rules into the topo.
func (ts *Server) SaveRoutingRules(ctx context.Context, routingRules *vschemapb.RoutingRules) error {
	data, err := routingRules.MarshalVT()
	if err != nil {
		return err
	}

	if len(data) == 0 {
		// No vschema, remove it. So we can remove the keyspace.
		if err := ts.globalCell.Delete(ctx, RoutingRulesFile, nil); err != nil && !IsErrType(err, NoNode) {
			return err
		}
		return nil
	}

	_, err = ts.globalCell.Update(ctx, RoutingRulesFile, data, nil)
	return err
}

// GetRoutingRules fetches the routing rules from the topo.
func (ts *Server) GetRoutingRules(ctx context.Context) (*vschemapb.RoutingRules, error) {
	rr := &vschemapb.RoutingRules{}
	data, _, err := ts.globalCell.Get(ctx, RoutingRulesFile)
	if err != nil {
		if IsErrType(err, NoNode) {
			return rr, nil
		}
		return nil, err
	}
	err = rr.UnmarshalVT(data)
	if err != nil {
		return nil, vterrors.Wrapf(err, "bad routing rules data: %q", data)
	}
	return rr, nil
}

// SaveShardRoutingRules saves the shard routing rules into the topo.
func (ts *Server) SaveShardRoutingRules(ctx context.Context, shardRoutingRules *vschemapb.ShardRoutingRules) error {
	data, err := shardRoutingRules.MarshalVT()
	if err != nil {
		return err
	}

	if len(data) == 0 {
		if err := ts.globalCell.Delete(ctx, ShardRoutingRulesFile, nil); err != nil && !IsErrType(err, NoNode) {
			return err
		}
		return nil
	}

	_, err = ts.globalCell.Update(ctx, ShardRoutingRulesFile, data, nil)
	return err
}

// GetShardRoutingRules fetches the shard routing rules from the topo.
func (ts *Server) GetShardRoutingRules(ctx context.Context) (*vschemapb.ShardRoutingRules, error) {
	srr := &vschemapb.ShardRoutingRules{}
	data, _, err := ts.globalCell.Get(ctx, ShardRoutingRulesFile)
	if err != nil {
		if IsErrType(err, NoNode) {
			return srr, nil
		}
		return nil, err
	}
	err = srr.UnmarshalVT(data)
	if err != nil {
		return nil, vterrors.Wrapf(err, "invalid shard routing rules: %q", data)
	}
	return srr, nil
}

// CreateKeyspaceRoutingRules wraps the underlying Conn.Create.
func (ts *Server) CreateKeyspaceRoutingRules(ctx context.Context, value *vschemapb.KeyspaceRoutingRules) error {
	data, err := value.MarshalVT()
	if err != nil {
		return err
	}
	if _, err := ts.globalCell.Create(ctx, ts.GetKeyspaceRoutingRulesPath(), data); err != nil {
		return err
	}
	return nil
}

// SaveKeyspaceRoutingRules saves the given routing rules proto in the topo at
// the defined path.
// It does NOT delete the file if you have requested to save empty routing rules
// (effectively deleting all routing rules in the file). This makes it different
// from the other routing rules (table and shard) save functions today. This is
// done as it simplifies the interactions with this key/file so that the typical
// access pattern is:
//   - If the file exists, we can lock it, read it, modify it, and save it back.
//   - If the file does not exist, we can create it and save the new rules.
//   - If multiple callers are racing to create the file, only one will succeed
//     and all other callers can simply retry once as the file will now exist.
//
// We can revisit this in the future and align things as we add locking and other
// topo server features to the other types of routing rules. We may then apply
// this new model used for keyspace routing rules to the other routing rules, or
// we may come up with a better model and apply it to the keyspace routing rules
// as well.
func (ts *Server) SaveKeyspaceRoutingRules(ctx context.Context, rules *vschemapb.KeyspaceRoutingRules) error {
	data, err := rules.MarshalVT()
	if err != nil {
		return err
	}
	_, err = ts.globalCell.Update(ctx, ts.GetKeyspaceRoutingRulesPath(), data, nil)
	return err
}

func (ts *Server) GetKeyspaceRoutingRules(ctx context.Context) (*vschemapb.KeyspaceRoutingRules, error) {
	rules := &vschemapb.KeyspaceRoutingRules{}
	data, _, err := ts.globalCell.Get(ctx, ts.GetKeyspaceRoutingRulesPath())
	if err != nil {
		if IsErrType(err, NoNode) {
			return nil, nil
		}
		return nil, err
	}
	err = rules.UnmarshalVT(data)
	if err != nil {
		return nil, vterrors.Wrapf(err, "bad keyspace routing rules data: %q", data)
	}
	return rules, nil
}

// GetMirrorRules fetches the mirror rules from the topo.
func (ts *Server) GetMirrorRules(ctx context.Context) (*vschemapb.MirrorRules, error) {
	rr := &vschemapb.MirrorRules{}
	data, _, err := ts.globalCell.Get(ctx, MirrorRulesFile)
	if err != nil {
		if IsErrType(err, NoNode) {
			return rr, nil
		}
		return nil, err
	}
	err = rr.UnmarshalVT(data)
	if err != nil {
		return nil, vterrors.Wrapf(err, "bad mirror rules data: %q", data)
	}
	return rr, nil
}

// SaveMirrorRules saves the mirror rules into the topo.
func (ts *Server) SaveMirrorRules(ctx context.Context, mirrorRules *vschemapb.MirrorRules) error {
	data, err := mirrorRules.MarshalVT()
	if err != nil {
		return err
	}

	if len(data) == 0 {
		// No vschema, remove it. So we can remove the keyspace.
		if err := ts.globalCell.Delete(ctx, MirrorRulesFile, nil); err != nil && !IsErrType(err, NoNode) {
			return err
		}
		return nil
	}

	_, err = ts.globalCell.Update(ctx, MirrorRulesFile, data, nil)
	return err
}
