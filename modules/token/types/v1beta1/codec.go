package v1beta1

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterInterface((*TokenI)(nil), nil)

	cdc.RegisterConcrete(&Token{}, "irismod/token/Token", nil)

	cdc.RegisterConcrete(&MsgIssueToken{}, "irismod/token/MsgIssueToken", nil)
	cdc.RegisterConcrete(&MsgEditToken{}, "irismod/token/MsgEditToken", nil)
	cdc.RegisterConcrete(&MsgMintToken{}, "irismod/token/MsgMintToken", nil)
	cdc.RegisterConcrete(&MsgBurnToken{}, "irismod/token/MsgBurnToken", nil)
	cdc.RegisterConcrete(&MsgTransferTokenOwner{}, "irismod/token/MsgTransferTokenOwner", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueToken{},
		&MsgEditToken{},
		&MsgMintToken{},
		&MsgBurnToken{},
		&MsgTransferTokenOwner{},
	)
	registry.RegisterInterface(
		"irismod.token.TokenI",
		(*TokenI)(nil),
		&Token{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}
