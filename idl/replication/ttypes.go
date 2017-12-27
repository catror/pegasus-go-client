// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package replication

import (
	"bytes"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/pegasus-kv/pegasus-go-client/idl/base"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = base.GoUnusedProtection__
var GoUnusedProtection__ int

// Attributes:
//  - Pid
//  - Ballot
//  - MaxReplicaCount
//  - Primary
//  - Secondaries
//  - LastDrops
//  - LastCommittedDecree
type PartitionConfiguration struct {
	Pid                 *base.Gpid         `thrift:"pid,1" json:"pid"`
	Ballot              int64              `thrift:"ballot,2" json:"ballot"`
	MaxReplicaCount     int32              `thrift:"max_replica_count,3" json:"max_replica_count"`
	Primary             *base.RPCAddress   `thrift:"primary,4" json:"primary"`
	Secondaries         []*base.RPCAddress `thrift:"secondaries,5" json:"secondaries"`
	LastDrops           []*base.RPCAddress `thrift:"last_drops,6" json:"last_drops"`
	LastCommittedDecree int64              `thrift:"last_committed_decree,7" json:"last_committed_decree"`
}

func NewPartitionConfiguration() *PartitionConfiguration {
	return &PartitionConfiguration{}
}

var PartitionConfiguration_Pid_DEFAULT *base.Gpid

func (p *PartitionConfiguration) GetPid() *base.Gpid {
	if !p.IsSetPid() {
		return PartitionConfiguration_Pid_DEFAULT
	}
	return p.Pid
}

func (p *PartitionConfiguration) GetBallot() int64 {
	return p.Ballot
}

func (p *PartitionConfiguration) GetMaxReplicaCount() int32 {
	return p.MaxReplicaCount
}

var PartitionConfiguration_Primary_DEFAULT *base.RPCAddress

func (p *PartitionConfiguration) GetPrimary() *base.RPCAddress {
	if !p.IsSetPrimary() {
		return PartitionConfiguration_Primary_DEFAULT
	}
	return p.Primary
}

func (p *PartitionConfiguration) GetSecondaries() []*base.RPCAddress {
	return p.Secondaries
}

func (p *PartitionConfiguration) GetLastDrops() []*base.RPCAddress {
	return p.LastDrops
}

func (p *PartitionConfiguration) GetLastCommittedDecree() int64 {
	return p.LastCommittedDecree
}
func (p *PartitionConfiguration) IsSetPid() bool {
	return p.Pid != nil
}

func (p *PartitionConfiguration) IsSetPrimary() bool {
	return p.Primary != nil
}

func (p *PartitionConfiguration) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.readField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.readField7(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *PartitionConfiguration) readField1(iprot thrift.TProtocol) error {
	p.Pid = &base.Gpid{}
	if err := p.Pid.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Pid), err)
	}
	return nil
}

func (p *PartitionConfiguration) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.Ballot = v
	}
	return nil
}

func (p *PartitionConfiguration) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.MaxReplicaCount = v
	}
	return nil
}

func (p *PartitionConfiguration) readField4(iprot thrift.TProtocol) error {
	p.Primary = &base.RPCAddress{}
	if err := p.Primary.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Primary), err)
	}
	return nil
}

func (p *PartitionConfiguration) readField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*base.RPCAddress, 0, size)
	p.Secondaries = tSlice
	for i := 0; i < size; i++ {
		_elem0 := &base.RPCAddress{}
		if err := _elem0.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem0), err)
		}
		p.Secondaries = append(p.Secondaries, _elem0)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *PartitionConfiguration) readField6(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*base.RPCAddress, 0, size)
	p.LastDrops = tSlice
	for i := 0; i < size; i++ {
		_elem1 := &base.RPCAddress{}
		if err := _elem1.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem1), err)
		}
		p.LastDrops = append(p.LastDrops, _elem1)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *PartitionConfiguration) readField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 7: ", err)
	} else {
		p.LastCommittedDecree = v
	}
	return nil
}

func (p *PartitionConfiguration) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("partition_configuration"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *PartitionConfiguration) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("pid", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:pid: ", p), err)
	}
	if err := p.Pid.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Pid), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:pid: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ballot", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:ballot: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.Ballot)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.ballot (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:ballot: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("max_replica_count", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:max_replica_count: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.MaxReplicaCount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.max_replica_count (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:max_replica_count: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("primary", thrift.STRUCT, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:primary: ", p), err)
	}
	if err := p.Primary.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Primary), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:primary: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("secondaries", thrift.LIST, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:secondaries: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Secondaries)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Secondaries {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:secondaries: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("last_drops", thrift.LIST, 6); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 6:last_drops: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.LastDrops)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.LastDrops {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 6:last_drops: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("last_committed_decree", thrift.I64, 7); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 7:last_committed_decree: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.LastCommittedDecree)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.last_committed_decree (7) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 7:last_committed_decree: ", p), err)
	}
	return err
}

func (p *PartitionConfiguration) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PartitionConfiguration(%+v)", *p)
}

// Attributes:
//  - AppName
//  - PartitionIndices
type QueryCfgRequest struct {
	AppName          string  `thrift:"app_name,1" json:"app_name"`
	PartitionIndices []int32 `thrift:"partition_indices,2" json:"partition_indices"`
}

func NewQueryCfgRequest() *QueryCfgRequest {
	return &QueryCfgRequest{}
}

func (p *QueryCfgRequest) GetAppName() string {
	return p.AppName
}

func (p *QueryCfgRequest) GetPartitionIndices() []int32 {
	return p.PartitionIndices
}
func (p *QueryCfgRequest) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *QueryCfgRequest) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.AppName = v
	}
	return nil
}

func (p *QueryCfgRequest) readField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]int32, 0, size)
	p.PartitionIndices = tSlice
	for i := 0; i < size; i++ {
		var _elem2 int32
		if v, err := iprot.ReadI32(); err != nil {
			return thrift.PrependError("error reading field 0: ", err)
		} else {
			_elem2 = v
		}
		p.PartitionIndices = append(p.PartitionIndices, _elem2)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *QueryCfgRequest) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("query_cfg_request"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *QueryCfgRequest) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("app_name", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:app_name: ", p), err)
	}
	if err := oprot.WriteString(string(p.AppName)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.app_name (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:app_name: ", p), err)
	}
	return err
}

func (p *QueryCfgRequest) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("partition_indices", thrift.LIST, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:partition_indices: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.I32, len(p.PartitionIndices)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.PartitionIndices {
		if err := oprot.WriteI32(int32(v)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T. (0) field write error: ", p), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:partition_indices: ", p), err)
	}
	return err
}

func (p *QueryCfgRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QueryCfgRequest(%+v)", *p)
}

// Attributes:
//  - Err
//  - AppID
//  - PartitionCount
//  - IsStateful
//  - Partitions
type QueryCfgResponse struct {
	Err            *base.ErrorCode           `thrift:"err,1" json:"err"`
	AppID          int32                     `thrift:"app_id,2" json:"app_id"`
	PartitionCount int32                     `thrift:"partition_count,3" json:"partition_count"`
	IsStateful     bool                      `thrift:"is_stateful,4" json:"is_stateful"`
	Partitions     []*PartitionConfiguration `thrift:"partitions,5" json:"partitions"`
}

func NewQueryCfgResponse() *QueryCfgResponse {
	return &QueryCfgResponse{}
}

var QueryCfgResponse_Err_DEFAULT *base.ErrorCode

func (p *QueryCfgResponse) GetErr() *base.ErrorCode {
	if !p.IsSetErr() {
		return QueryCfgResponse_Err_DEFAULT
	}
	return p.Err
}

func (p *QueryCfgResponse) GetAppID() int32 {
	return p.AppID
}

func (p *QueryCfgResponse) GetPartitionCount() int32 {
	return p.PartitionCount
}

func (p *QueryCfgResponse) GetIsStateful() bool {
	return p.IsStateful
}

func (p *QueryCfgResponse) GetPartitions() []*PartitionConfiguration {
	return p.Partitions
}
func (p *QueryCfgResponse) IsSetErr() bool {
	return p.Err != nil
}

func (p *QueryCfgResponse) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.readField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.readField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.readField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.readField5(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *QueryCfgResponse) readField1(iprot thrift.TProtocol) error {
	p.Err = &base.ErrorCode{}
	if err := p.Err.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Err), err)
	}
	return nil
}

func (p *QueryCfgResponse) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.AppID = v
	}
	return nil
}

func (p *QueryCfgResponse) readField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.PartitionCount = v
	}
	return nil
}

func (p *QueryCfgResponse) readField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 4: ", err)
	} else {
		p.IsStateful = v
	}
	return nil
}

func (p *QueryCfgResponse) readField5(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*PartitionConfiguration, 0, size)
	p.Partitions = tSlice
	for i := 0; i < size; i++ {
		_elem3 := &PartitionConfiguration{}
		if err := _elem3.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem3), err)
		}
		p.Partitions = append(p.Partitions, _elem3)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *QueryCfgResponse) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("query_cfg_response"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *QueryCfgResponse) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("err", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:err: ", p), err)
	}
	if err := p.Err.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Err), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:err: ", p), err)
	}
	return err
}

func (p *QueryCfgResponse) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("app_id", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:app_id: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.AppID)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.app_id (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:app_id: ", p), err)
	}
	return err
}

func (p *QueryCfgResponse) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("partition_count", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:partition_count: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.PartitionCount)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.partition_count (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:partition_count: ", p), err)
	}
	return err
}

func (p *QueryCfgResponse) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("is_stateful", thrift.BOOL, 4); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 4:is_stateful: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.IsStateful)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.is_stateful (4) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 4:is_stateful: ", p), err)
	}
	return err
}

func (p *QueryCfgResponse) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("partitions", thrift.LIST, 5); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 5:partitions: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Partitions)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.Partitions {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 5:partitions: ", p), err)
	}
	return err
}

func (p *QueryCfgResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("QueryCfgResponse(%+v)", *p)
}
