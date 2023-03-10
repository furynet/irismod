<!--
order: 3
-->

# Events

## Handlers

### MsgSwapOrder

| Type    | Attribute Key | Attribute Value |
| :------ | :------------ | :-------------- |
| swap    | amount        | {amount}        |
| swap    | sender        | {senderAddress} |
| swap    | recipient     | {recipient}     |
| swap    | is_buy_order  | {isBuyOrder}    |
| swap    | token_pair    | {tokenPair}     |
| message | module        | coinswap        |
| message | sender        | {senderAddress} |

### MsgAddLiquidity

| Type          | Attribute Key | Attribute Value |
| :------------ | :------------ | :-------------- |
| add_liquidity | sender        | {senderAddress} |
| add_liquidity | token_pair    | {tokenPair}     |
| message       | module        | coinswap        |
| message       | sender        | {senderAddress} |

### MsgAddUnilateralLiquidity

| Type                     | Attribute Key    | Attribute Value   |
|:-------------------------|:-----------------|:------------------|
| add_unilateral_liquidity | sender           | {senderAddress}   |
| add_unilateral_liquidity | token_unilateral | {tokenUnilateral} |
| message                  | module           | coinswap          |
| message                  | sender           | {senderAddress}   |

### MsgRemoveLiquidity

| Type             | Attribute Key | Attribute Value |
| :--------------- | :------------ | :-------------- |
| remove_liquidity | sender        | {senderAddress} |
| remove_liquidity | token_pair    | {tokenPair}     |
| message          | module        | coinswap        |
| message          | sender        | {senderAddress} |

### MsgRemoveUnilateralLiquidity

| Type                        | Attribute Key   | Attribute Value  |
|:----------------------------|:----------------|:-----------------|
| remove_unilateral_liquidity | sender          | {senderAddress}  |
| remove_unilateral_liquidity | token_unilateral| {tokenUnilateral}|
| message                     | module          | coinswap         |
| message                     | sender          | {senderAddress}  |
