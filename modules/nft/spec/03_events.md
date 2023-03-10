# Events

The nft module emits the following events:

## Handlers

### MsgIssueDenom

| Type        | Attribute Key | Attribute Value  |
|:------------|:--------------|:-----------------|
| issue_denom | denom_id      | {nftDenomID}     |
| issue_denom | denom_name    | {nftDenomName}   |
| issue_denom | creator       | {creatorAddress} |
| message     | module        | nft              |
| message     | sender        | {senderAddress}  |

### MsgTransferNFT

| Type         | Attribute Key | Attribute Value    |
|:-------------|:--------------|:-------------------|
| transfer_nft | token_id      | {tokenID}          |
| transfer_nft | denom_id      | {nftDenomID}       |
| transfer_nft | sender        | {senderAddress}    |
| transfer_nft | recipient     | {recipientAddress} |
| message      | module        | nft                |
| message      | sender        | {senderAddress}    |

### MsgEditNFT

| Type     | Attribute Key | Attribute Value |
|:---------|:--------------|:----------------|
| edit_nft | token_id      | {tokenID}       |
| edit_nft | denom_id      | {nftDenomID}    |
| edit_nft | token_uri     | {tokenURI}      |
| edit_nft | owner         | {ownerAddress}  |
| message  | module        | nft             |
| message  | sender        | {senderAddress} |

### MsgMintNFT

| Type     | Attribute Key | Attribute Value    |
|:---------|:--------------|:-------------------|
| mint_nft | token_id      | {tokenID}          |
| mint_nft | denom_id      | {nftDenomID}       |
| mint_nft | token_uri     | {tokenURI}         |
| mint_nft | recipient     | {recipientAddress} |
| message  | module        | nft                |
| message  | sender        | {senderAddress}    |

### MsgBurnNFTs

| Type     | Attribute Key | Attribute Value |
|:---------|:--------------|:----------------|
| burn_nft | denom_id      | {nftDenomID}    |
| burn_nft | token_id      | {tokenID}       |
| burn_nft | owner         | {ownerAddress}  |
| message  | module        | nft             |
| message  | sender        | {senderAddress} |

### MsgTransferDenom

| Type           | Attribute Key | Attribute Value    |
|:---------------|:--------------|:-------------------|
| transfer_denom | denom_id      | {nftDenomID}       |
| transfer_denom | sender        | {senderAddress}    |
| transfer_denom | recipient     | {recipientAddress} |
| message        | module        | nft                |
| message        | sender        | {senderAddress}    |