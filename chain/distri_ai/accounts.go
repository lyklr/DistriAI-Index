// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package distri_ai

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type AiModel struct {
	Owner      ag_solanago.PublicKey
	Name       string
	Framework  uint8
	License    uint8
	Type1      uint8
	Type2      uint8
	Tags       string
	CreateTime int64
	UpdateTime int64
}

var AiModelDiscriminator = [8]byte{42, 206, 111, 34, 42, 121, 50, 138}

func (obj AiModel) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(AiModelDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Framework` param:
	err = encoder.Encode(obj.Framework)
	if err != nil {
		return err
	}
	// Serialize `License` param:
	err = encoder.Encode(obj.License)
	if err != nil {
		return err
	}
	// Serialize `Type1` param:
	err = encoder.Encode(obj.Type1)
	if err != nil {
		return err
	}
	// Serialize `Type2` param:
	err = encoder.Encode(obj.Type2)
	if err != nil {
		return err
	}
	// Serialize `Tags` param:
	err = encoder.Encode(obj.Tags)
	if err != nil {
		return err
	}
	// Serialize `CreateTime` param:
	err = encoder.Encode(obj.CreateTime)
	if err != nil {
		return err
	}
	// Serialize `UpdateTime` param:
	err = encoder.Encode(obj.UpdateTime)
	if err != nil {
		return err
	}
	return nil
}

func (obj *AiModel) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(AiModelDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[42 206 111 34 42 121 50 138]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Framework`:
	err = decoder.Decode(&obj.Framework)
	if err != nil {
		return err
	}
	// Deserialize `License`:
	err = decoder.Decode(&obj.License)
	if err != nil {
		return err
	}
	// Deserialize `Type1`:
	err = decoder.Decode(&obj.Type1)
	if err != nil {
		return err
	}
	// Deserialize `Type2`:
	err = decoder.Decode(&obj.Type2)
	if err != nil {
		return err
	}
	// Deserialize `Tags`:
	err = decoder.Decode(&obj.Tags)
	if err != nil {
		return err
	}
	// Deserialize `CreateTime`:
	err = decoder.Decode(&obj.CreateTime)
	if err != nil {
		return err
	}
	// Deserialize `UpdateTime`:
	err = decoder.Decode(&obj.UpdateTime)
	if err != nil {
		return err
	}
	return nil
}

type Dataset struct {
	Owner      ag_solanago.PublicKey
	Name       string
	Scale      uint8
	License    uint8
	Type1      uint8
	Type2      uint8
	Tags       string
	CreateTime int64
	UpdateTime int64
}

var DatasetDiscriminator = [8]byte{242, 85, 87, 90, 234, 188, 241, 17}

func (obj Dataset) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(DatasetDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Name` param:
	err = encoder.Encode(obj.Name)
	if err != nil {
		return err
	}
	// Serialize `Scale` param:
	err = encoder.Encode(obj.Scale)
	if err != nil {
		return err
	}
	// Serialize `License` param:
	err = encoder.Encode(obj.License)
	if err != nil {
		return err
	}
	// Serialize `Type1` param:
	err = encoder.Encode(obj.Type1)
	if err != nil {
		return err
	}
	// Serialize `Type2` param:
	err = encoder.Encode(obj.Type2)
	if err != nil {
		return err
	}
	// Serialize `Tags` param:
	err = encoder.Encode(obj.Tags)
	if err != nil {
		return err
	}
	// Serialize `CreateTime` param:
	err = encoder.Encode(obj.CreateTime)
	if err != nil {
		return err
	}
	// Serialize `UpdateTime` param:
	err = encoder.Encode(obj.UpdateTime)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Dataset) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(DatasetDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[242 85 87 90 234 188 241 17]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Name`:
	err = decoder.Decode(&obj.Name)
	if err != nil {
		return err
	}
	// Deserialize `Scale`:
	err = decoder.Decode(&obj.Scale)
	if err != nil {
		return err
	}
	// Deserialize `License`:
	err = decoder.Decode(&obj.License)
	if err != nil {
		return err
	}
	// Deserialize `Type1`:
	err = decoder.Decode(&obj.Type1)
	if err != nil {
		return err
	}
	// Deserialize `Type2`:
	err = decoder.Decode(&obj.Type2)
	if err != nil {
		return err
	}
	// Deserialize `Tags`:
	err = decoder.Decode(&obj.Tags)
	if err != nil {
		return err
	}
	// Deserialize `CreateTime`:
	err = decoder.Decode(&obj.CreateTime)
	if err != nil {
		return err
	}
	// Deserialize `UpdateTime`:
	err = decoder.Decode(&obj.UpdateTime)
	if err != nil {
		return err
	}
	return nil
}

type Machine struct {
	// Owner of this machine.
	Owner ag_solanago.PublicKey

	// UUID of this machine.
	Uuid [16]uint8

	// The metadata by json format of this machine.
	Metadata string

	// The status of this machine.
	Status MachineStatus

	// The price of this machine.
	Price uint64

	// The maximum number of hours the machine can be rent.
	MaxDuration uint32

	// The GB amount of this machine's avaliable disk.
	Disk uint32

	// The total number of completed orders for this machine.
	CompletedCount uint32

	// The total number of failed orders for this machine.
	FailedCount uint32

	// Computing power score of this machine.
	Score uint8

	// Total claimed periodic rewards.
	ClaimedPeriodicRewards uint64

	// Total claimed task rewards.
	ClaimedTaskRewards uint64
	OrderPda           ag_solanago.PublicKey
}

var MachineDiscriminator = [8]byte{25, 102, 22, 13, 58, 243, 138, 79}

func (obj Machine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MachineDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `MaxDuration` param:
	err = encoder.Encode(obj.MaxDuration)
	if err != nil {
		return err
	}
	// Serialize `Disk` param:
	err = encoder.Encode(obj.Disk)
	if err != nil {
		return err
	}
	// Serialize `CompletedCount` param:
	err = encoder.Encode(obj.CompletedCount)
	if err != nil {
		return err
	}
	// Serialize `FailedCount` param:
	err = encoder.Encode(obj.FailedCount)
	if err != nil {
		return err
	}
	// Serialize `Score` param:
	err = encoder.Encode(obj.Score)
	if err != nil {
		return err
	}
	// Serialize `ClaimedPeriodicRewards` param:
	err = encoder.Encode(obj.ClaimedPeriodicRewards)
	if err != nil {
		return err
	}
	// Serialize `ClaimedTaskRewards` param:
	err = encoder.Encode(obj.ClaimedTaskRewards)
	if err != nil {
		return err
	}
	// Serialize `OrderPda` param:
	err = encoder.Encode(obj.OrderPda)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Machine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MachineDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[25 102 22 13 58 243 138 79]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `MaxDuration`:
	err = decoder.Decode(&obj.MaxDuration)
	if err != nil {
		return err
	}
	// Deserialize `Disk`:
	err = decoder.Decode(&obj.Disk)
	if err != nil {
		return err
	}
	// Deserialize `CompletedCount`:
	err = decoder.Decode(&obj.CompletedCount)
	if err != nil {
		return err
	}
	// Deserialize `FailedCount`:
	err = decoder.Decode(&obj.FailedCount)
	if err != nil {
		return err
	}
	// Deserialize `Score`:
	err = decoder.Decode(&obj.Score)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedPeriodicRewards`:
	err = decoder.Decode(&obj.ClaimedPeriodicRewards)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedTaskRewards`:
	err = decoder.Decode(&obj.ClaimedTaskRewards)
	if err != nil {
		return err
	}
	// Deserialize `OrderPda`:
	err = decoder.Decode(&obj.OrderPda)
	if err != nil {
		return err
	}
	return nil
}

type MachineNew struct {
	Owner ag_solanago.PublicKey
	Uuid  [16]uint8

	// The metadata by json format of this machine.
	Metadata string

	// The status of this machine.
	Status MachineStatus

	// The price of this machine.
	Price uint64

	// The maximum number of hours the machine can be rent.
	MaxDuration uint32

	// The GB amount of this machine's avaliable disk.
	Disk           uint32
	CompletedCount uint32
	FailedCount    uint32
	Score          uint8

	// Total claimed periodic rewards.
	ClaimedPeriodicRewards uint64

	// Total claimed task rewards.
	ClaimedTaskRewards uint64
	OrderPda           ag_solanago.PublicKey
}

var MachineNewDiscriminator = [8]byte{27, 50, 112, 47, 115, 74, 107, 33}

func (obj MachineNew) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(MachineNewDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `MaxDuration` param:
	err = encoder.Encode(obj.MaxDuration)
	if err != nil {
		return err
	}
	// Serialize `Disk` param:
	err = encoder.Encode(obj.Disk)
	if err != nil {
		return err
	}
	// Serialize `CompletedCount` param:
	err = encoder.Encode(obj.CompletedCount)
	if err != nil {
		return err
	}
	// Serialize `FailedCount` param:
	err = encoder.Encode(obj.FailedCount)
	if err != nil {
		return err
	}
	// Serialize `Score` param:
	err = encoder.Encode(obj.Score)
	if err != nil {
		return err
	}
	// Serialize `ClaimedPeriodicRewards` param:
	err = encoder.Encode(obj.ClaimedPeriodicRewards)
	if err != nil {
		return err
	}
	// Serialize `ClaimedTaskRewards` param:
	err = encoder.Encode(obj.ClaimedTaskRewards)
	if err != nil {
		return err
	}
	// Serialize `OrderPda` param:
	err = encoder.Encode(obj.OrderPda)
	if err != nil {
		return err
	}
	return nil
}

func (obj *MachineNew) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(MachineNewDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[27 50 112 47 115 74 107 33]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `MaxDuration`:
	err = decoder.Decode(&obj.MaxDuration)
	if err != nil {
		return err
	}
	// Deserialize `Disk`:
	err = decoder.Decode(&obj.Disk)
	if err != nil {
		return err
	}
	// Deserialize `CompletedCount`:
	err = decoder.Decode(&obj.CompletedCount)
	if err != nil {
		return err
	}
	// Deserialize `FailedCount`:
	err = decoder.Decode(&obj.FailedCount)
	if err != nil {
		return err
	}
	// Deserialize `Score`:
	err = decoder.Decode(&obj.Score)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedPeriodicRewards`:
	err = decoder.Decode(&obj.ClaimedPeriodicRewards)
	if err != nil {
		return err
	}
	// Deserialize `ClaimedTaskRewards`:
	err = decoder.Decode(&obj.ClaimedTaskRewards)
	if err != nil {
		return err
	}
	// Deserialize `OrderPda`:
	err = decoder.Decode(&obj.OrderPda)
	if err != nil {
		return err
	}
	return nil
}

type Order struct {
	// UUID of this order.
	OrderId [16]uint8

	// The buyer of this order.
	Buyer ag_solanago.PublicKey

	// The seller of this order.
	Seller ag_solanago.PublicKey

	// The machine id of this order.
	MachineId [16]uint8

	// The price of this order.
	Price uint64

	// The duration hours of this order.
	Duration uint32

	// The total amount of this order.
	Total uint64

	// The metadata by json format of this order.
	Metadata string

	// The status of this order.
	Status OrderStatus

	// The order time of this order.
	OrderTime int64

	// The start time of this order.
	StartTime int64

	// The refund time of this order.
	RefundTime int64
}

var OrderDiscriminator = [8]byte{134, 173, 223, 185, 77, 86, 28, 51}

func (obj Order) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(OrderDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `OrderId` param:
	err = encoder.Encode(obj.OrderId)
	if err != nil {
		return err
	}
	// Serialize `Buyer` param:
	err = encoder.Encode(obj.Buyer)
	if err != nil {
		return err
	}
	// Serialize `Seller` param:
	err = encoder.Encode(obj.Seller)
	if err != nil {
		return err
	}
	// Serialize `MachineId` param:
	err = encoder.Encode(obj.MachineId)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `Duration` param:
	err = encoder.Encode(obj.Duration)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	// Serialize `OrderTime` param:
	err = encoder.Encode(obj.OrderTime)
	if err != nil {
		return err
	}
	// Serialize `StartTime` param:
	err = encoder.Encode(obj.StartTime)
	if err != nil {
		return err
	}
	// Serialize `RefundTime` param:
	err = encoder.Encode(obj.RefundTime)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Order) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(OrderDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[134 173 223 185 77 86 28 51]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `OrderId`:
	err = decoder.Decode(&obj.OrderId)
	if err != nil {
		return err
	}
	// Deserialize `Buyer`:
	err = decoder.Decode(&obj.Buyer)
	if err != nil {
		return err
	}
	// Deserialize `Seller`:
	err = decoder.Decode(&obj.Seller)
	if err != nil {
		return err
	}
	// Deserialize `MachineId`:
	err = decoder.Decode(&obj.MachineId)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `Duration`:
	err = decoder.Decode(&obj.Duration)
	if err != nil {
		return err
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	// Deserialize `OrderTime`:
	err = decoder.Decode(&obj.OrderTime)
	if err != nil {
		return err
	}
	// Deserialize `StartTime`:
	err = decoder.Decode(&obj.StartTime)
	if err != nil {
		return err
	}
	// Deserialize `RefundTime`:
	err = decoder.Decode(&obj.RefundTime)
	if err != nil {
		return err
	}
	return nil
}

type OrderNew struct {
	OrderId [16]uint8

	// The buyer of this order.
	Buyer ag_solanago.PublicKey

	// The seller of this order.
	Seller ag_solanago.PublicKey

	// The machine id of this order.
	MachineId [16]uint8

	// The price of this order.
	Price uint64

	// The duration hours of this order.
	Duration uint32

	// The total amount of this order.
	Total uint64

	// The metadata by json format of this order.
	Metadata string

	// The status of this order.
	Status OrderStatus

	// The order time of this order.
	OrderTime int64

	// The start time of this order.
	StartTime int64

	// The refund time of this order.
	RefundTime int64
}

var OrderNewDiscriminator = [8]byte{6, 203, 52, 103, 118, 10, 252, 19}

func (obj OrderNew) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(OrderNewDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `OrderId` param:
	err = encoder.Encode(obj.OrderId)
	if err != nil {
		return err
	}
	// Serialize `Buyer` param:
	err = encoder.Encode(obj.Buyer)
	if err != nil {
		return err
	}
	// Serialize `Seller` param:
	err = encoder.Encode(obj.Seller)
	if err != nil {
		return err
	}
	// Serialize `MachineId` param:
	err = encoder.Encode(obj.MachineId)
	if err != nil {
		return err
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `Duration` param:
	err = encoder.Encode(obj.Duration)
	if err != nil {
		return err
	}
	// Serialize `Total` param:
	err = encoder.Encode(obj.Total)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	// Serialize `Status` param:
	err = encoder.Encode(obj.Status)
	if err != nil {
		return err
	}
	// Serialize `OrderTime` param:
	err = encoder.Encode(obj.OrderTime)
	if err != nil {
		return err
	}
	// Serialize `StartTime` param:
	err = encoder.Encode(obj.StartTime)
	if err != nil {
		return err
	}
	// Serialize `RefundTime` param:
	err = encoder.Encode(obj.RefundTime)
	if err != nil {
		return err
	}
	return nil
}

func (obj *OrderNew) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(OrderNewDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[6 203 52 103 118 10 252 19]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `OrderId`:
	err = decoder.Decode(&obj.OrderId)
	if err != nil {
		return err
	}
	// Deserialize `Buyer`:
	err = decoder.Decode(&obj.Buyer)
	if err != nil {
		return err
	}
	// Deserialize `Seller`:
	err = decoder.Decode(&obj.Seller)
	if err != nil {
		return err
	}
	// Deserialize `MachineId`:
	err = decoder.Decode(&obj.MachineId)
	if err != nil {
		return err
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `Duration`:
	err = decoder.Decode(&obj.Duration)
	if err != nil {
		return err
	}
	// Deserialize `Total`:
	err = decoder.Decode(&obj.Total)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	// Deserialize `Status`:
	err = decoder.Decode(&obj.Status)
	if err != nil {
		return err
	}
	// Deserialize `OrderTime`:
	err = decoder.Decode(&obj.OrderTime)
	if err != nil {
		return err
	}
	// Deserialize `StartTime`:
	err = decoder.Decode(&obj.StartTime)
	if err != nil {
		return err
	}
	// Deserialize `RefundTime`:
	err = decoder.Decode(&obj.RefundTime)
	if err != nil {
		return err
	}
	return nil
}

type Reward struct {
	// Reward period.
	Period uint32

	// Start time of this reward period.
	StartTime int64

	// Reward pool in this period.
	Pool uint64

	// Participating machine number in this period.
	MachineNum uint32

	// Periodic reward per machine in this period.
	UnitPeriodicReward uint64

	// Task number in this period.
	TaskNum uint32

	// Task reward per task in this period.
	UnitTaskReward uint64
}

var RewardDiscriminator = [8]byte{174, 129, 42, 212, 190, 18, 45, 34}

func (obj Reward) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RewardDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Period` param:
	err = encoder.Encode(obj.Period)
	if err != nil {
		return err
	}
	// Serialize `StartTime` param:
	err = encoder.Encode(obj.StartTime)
	if err != nil {
		return err
	}
	// Serialize `Pool` param:
	err = encoder.Encode(obj.Pool)
	if err != nil {
		return err
	}
	// Serialize `MachineNum` param:
	err = encoder.Encode(obj.MachineNum)
	if err != nil {
		return err
	}
	// Serialize `UnitPeriodicReward` param:
	err = encoder.Encode(obj.UnitPeriodicReward)
	if err != nil {
		return err
	}
	// Serialize `TaskNum` param:
	err = encoder.Encode(obj.TaskNum)
	if err != nil {
		return err
	}
	// Serialize `UnitTaskReward` param:
	err = encoder.Encode(obj.UnitTaskReward)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Reward) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RewardDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[174 129 42 212 190 18 45 34]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Period`:
	err = decoder.Decode(&obj.Period)
	if err != nil {
		return err
	}
	// Deserialize `StartTime`:
	err = decoder.Decode(&obj.StartTime)
	if err != nil {
		return err
	}
	// Deserialize `Pool`:
	err = decoder.Decode(&obj.Pool)
	if err != nil {
		return err
	}
	// Deserialize `MachineNum`:
	err = decoder.Decode(&obj.MachineNum)
	if err != nil {
		return err
	}
	// Deserialize `UnitPeriodicReward`:
	err = decoder.Decode(&obj.UnitPeriodicReward)
	if err != nil {
		return err
	}
	// Deserialize `TaskNum`:
	err = decoder.Decode(&obj.TaskNum)
	if err != nil {
		return err
	}
	// Deserialize `UnitTaskReward`:
	err = decoder.Decode(&obj.UnitTaskReward)
	if err != nil {
		return err
	}
	return nil
}

type RewardMachine struct {
	Period uint32

	// Machine owner.
	Owner ag_solanago.PublicKey

	// Machine id.
	MachineId [16]uint8

	// Task number submited in this period.
	TaskNum uint32

	// Reward has been claimed.
	Claimed bool
}

var RewardMachineDiscriminator = [8]byte{106, 87, 186, 254, 4, 139, 144, 74}

func (obj RewardMachine) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(RewardMachineDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Period` param:
	err = encoder.Encode(obj.Period)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `MachineId` param:
	err = encoder.Encode(obj.MachineId)
	if err != nil {
		return err
	}
	// Serialize `TaskNum` param:
	err = encoder.Encode(obj.TaskNum)
	if err != nil {
		return err
	}
	// Serialize `Claimed` param:
	err = encoder.Encode(obj.Claimed)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RewardMachine) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(RewardMachineDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[106 87 186 254 4 139 144 74]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Period`:
	err = decoder.Decode(&obj.Period)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `MachineId`:
	err = decoder.Decode(&obj.MachineId)
	if err != nil {
		return err
	}
	// Deserialize `TaskNum`:
	err = decoder.Decode(&obj.TaskNum)
	if err != nil {
		return err
	}
	// Deserialize `Claimed`:
	err = decoder.Decode(&obj.Claimed)
	if err != nil {
		return err
	}
	return nil
}

type Task struct {
	// Task id
	Uuid [16]uint8

	// Period
	Period uint32

	// Machine owner.
	Owner ag_solanago.PublicKey

	// Machine id of this task.
	MachineId [16]uint8

	// The metadata by json format of this task.
	Metadata string
}

var TaskDiscriminator = [8]byte{79, 34, 229, 55, 88, 90, 55, 84}

func (obj Task) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TaskDiscriminator[:], false)
	if err != nil {
		return err
	}
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `Period` param:
	err = encoder.Encode(obj.Period)
	if err != nil {
		return err
	}
	// Serialize `Owner` param:
	err = encoder.Encode(obj.Owner)
	if err != nil {
		return err
	}
	// Serialize `MachineId` param:
	err = encoder.Encode(obj.MachineId)
	if err != nil {
		return err
	}
	// Serialize `Metadata` param:
	err = encoder.Encode(obj.Metadata)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Task) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TaskDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[79 34 229 55 88 90 55 84]",
				fmt.Sprint(discriminator[:]))
		}
	}
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `Period`:
	err = decoder.Decode(&obj.Period)
	if err != nil {
		return err
	}
	// Deserialize `Owner`:
	err = decoder.Decode(&obj.Owner)
	if err != nil {
		return err
	}
	// Deserialize `MachineId`:
	err = decoder.Decode(&obj.MachineId)
	if err != nil {
		return err
	}
	// Deserialize `Metadata`:
	err = decoder.Decode(&obj.Metadata)
	if err != nil {
		return err
	}
	return nil
}
