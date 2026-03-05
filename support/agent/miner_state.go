package agent

import (
	"context"

	"github.com/post-quantumqoin/bitset"
	"github.com/post-quantumqoin/core-types/abi"
	"github.com/post-quantumqoin/core-types/big"
	"github.com/post-quantumqoin/core-types/dline"
	"github.com/post-quantumqoin/specs-contracts/contracts/builtin/miner"
	"github.com/post-quantumqoin/specs-contracts/contracts/util/adt"
	cid "github.com/ipfs/go-cid"
)

type MinerStateV4 struct {
	Root cid.Cid
	Ctx  context.Context
	st   *miner.State
}

func (m *MinerStateV4) state(store adt.Store) (*miner.State, error) {
	if m.st == nil {
		var st miner.State
		err := store.Get(m.Ctx, m.Root, &st)
		if err != nil {
			return nil, err
		}
		m.st = &st
	}
	return m.st, nil
}

func (m *MinerStateV4) HasSectorNo(store adt.Store, sectorNo abi.SectorNumber) (bool, error) {
	st, err := m.state(store)
	if err != nil {
		return false, err
	}
	return st.HasSectorNo(store, sectorNo)
}

func (m *MinerStateV4) FindSector(store adt.Store, sectorNo abi.SectorNumber) (uint64, uint64, error) {
	st, err := m.state(store)
	if err != nil {
		return 0, 0, err
	}
	return st.FindSector(store, sectorNo)
}

func (m *MinerStateV4) ProvingPeriodStart(store adt.Store) (abi.ChainEpoch, error) {
	st, err := m.state(store)
	if err != nil {
		return 0, err
	}
	return st.ProvingPeriodStart, nil
}

func (m *MinerStateV4) LoadSectorInfo(store adt.Store, sectorNo uint64) (SimSectorInfo, error) {
	st, err := m.state(store)
	if err != nil {
		return nil, err
	}
	sectors, err := st.LoadSectorInfos(store, bitfield.NewFromSet([]uint64{uint64(sectorNo)}))
	if err != nil {
		return nil, err
	}
	return &SectorInfoV4{info: sectors[0]}, nil
}

func (m *MinerStateV4) DeadlineInfo(store adt.Store, currEpoch abi.ChainEpoch) (*dline.Info, error) {
	st, err := m.state(store)
	if err != nil {
		return nil, err
	}
	return st.DeadlineInfo(currEpoch), nil
}

func (m *MinerStateV4) FeeDebt(store adt.Store) (abi.TokenAmount, error) {
	st, err := m.state(store)
	if err != nil {
		return big.Zero(), err
	}
	return st.FeeDebt, nil
}

func (m *MinerStateV4) LoadDeadlineState(store adt.Store, dlIdx uint64) (SimDeadlineState, error) {
	st, err := m.state(store)
	if err != nil {
		return nil, err
	}
	dls, err := st.LoadDeadlines(store)
	if err != nil {
		return nil, err
	}
	dline, err := dls.LoadDeadline(store, dlIdx)
	if err != nil {
		return nil, err
	}
	return &DeadlineStateV4{deadline: dline}, nil
}

type DeadlineStateV4 struct {
	deadline *miner.Deadline
}

func (d *DeadlineStateV4) LoadPartition(store adt.Store, partIdx uint64) (SimPartitionState, error) {
	part, err := d.deadline.LoadPartition(store, partIdx)
	if err != nil {
		return nil, err
	}
	return &PartitionStateV4{partition: part}, nil
}

type PartitionStateV4 struct {
	partition *miner.Partition
}

func (p *PartitionStateV4) Terminated() bitfield.BitField {
	return p.partition.Terminated
}

type SectorInfoV4 struct {
	info *miner.SectorOnChainInfo
}

func (s *SectorInfoV4) Expiration() abi.ChainEpoch {
	return s.info.Expiration
}

type MinerStateV5 struct {
	Root cid.Cid
	st   *miner.State
	Ctx  context.Context
}

func (m *MinerStateV5) state(store adt.Store) (*miner.State, error) {
	if m.st == nil {
		var st miner.State
		err := store.Get(m.Ctx, m.Root, &st)
		if err != nil {
			return nil, err
		}
		m.st = &st
	}
	return m.st, nil
}

func (m *MinerStateV5) HasSectorNo(store adt.Store, sectorNo abi.SectorNumber) (bool, error) {
	st, err := m.state(store)
	if err != nil {
		return false, err
	}
	return st.HasSectorNo(store, sectorNo)
}

func (m *MinerStateV5) FindSector(store adt.Store, sectorNo abi.SectorNumber) (uint64, uint64, error) {
	st, err := m.state(store)
	if err != nil {
		return 0, 0, err
	}
	return st.FindSector(store, sectorNo)
}

func (m *MinerStateV5) ProvingPeriodStart(store adt.Store) (abi.ChainEpoch, error) {
	st, err := m.state(store)
	if err != nil {
		return 0, err
	}
	return st.ProvingPeriodStart, nil
}

func (m *MinerStateV5) LoadSectorInfo(store adt.Store, sectorNo uint64) (SimSectorInfo, error) {
	st, err := m.state(store)
	if err != nil {
		return nil, err
	}
	sectors, err := st.LoadSectorInfos(store, bitfield.NewFromSet([]uint64{uint64(sectorNo)}))
	if err != nil {
		return nil, err
	}
	return &SectorInfoV5{info: sectors[0]}, nil
}

func (m *MinerStateV5) DeadlineInfo(store adt.Store, currEpoch abi.ChainEpoch) (*dline.Info, error) {
	st, err := m.state(store)
	if err != nil {
		return nil, err
	}
	return st.DeadlineInfo(currEpoch), nil
}

func (m *MinerStateV5) FeeDebt(store adt.Store) (abi.TokenAmount, error) {
	st, err := m.state(store)
	if err != nil {
		return big.Zero(), err
	}
	return st.FeeDebt, nil
}

func (m *MinerStateV5) LoadDeadlineState(store adt.Store, dlIdx uint64) (SimDeadlineState, error) {
	st, err := m.state(store)
	if err != nil {
		return nil, err
	}
	dls, err := st.LoadDeadlines(store)
	if err != nil {
		return nil, err
	}
	dline, err := dls.LoadDeadline(store, dlIdx)
	if err != nil {
		return nil, err
	}
	return &DeadlineStateV5{deadline: dline}, nil
}

type DeadlineStateV5 struct {
	deadline *miner.Deadline
}

func (d *DeadlineStateV5) LoadPartition(store adt.Store, partIdx uint64) (SimPartitionState, error) {
	part, err := d.deadline.LoadPartition(store, partIdx)
	if err != nil {
		return nil, err
	}
	return &PartitionStateV5{partition: part}, nil
}

type PartitionStateV5 struct {
	partition *miner.Partition
}

func (p *PartitionStateV5) Terminated() bitfield.BitField {
	return p.partition.Terminated
}

type SectorInfoV5 struct {
	info *miner.SectorOnChainInfo
}

func (s *SectorInfoV5) Expiration() abi.ChainEpoch {
	return s.info.Expiration
}
