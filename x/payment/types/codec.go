package types

import (
	"github.com/celestiaorg/celestia-app/app/encoding"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

var (
	ModuleCdc = codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgWirePayForData{}, "payment/WirePayForData", nil)
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgWirePayForData{},
	)

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgPayForData{},
	)

	registry.RegisterInterface(
		"cosmos.auth.v1beta1.BaseAccount",
		(*authtypes.AccountI)(nil),
	)

	registry.RegisterImplementations(
		(*authtypes.AccountI)(nil),
		&authtypes.BaseAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// makePaymentEncodingConfig uses the payment modules RegisterInterfaces
// function to create an encoding config for the payment module. This is useful
// so that we don't have to import the app package.
func makePaymentEncodingConfig() encoding.EncodingConfig {
	return encoding.MakeEncodingConfig(RegisterInterfaces)
}
