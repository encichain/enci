<!--
order: 2
-->

# State

## Key-Value Mapping

The following notation is used to describe key to value mapping:

`key -> value`

with `|` used to describe concatenation. Most keys will resemble the following format:

`0x00 | exampleBytes -> encoding_type(value_object)`

## TaxRateLimits

The constraints for the TaxRate and BurnRate. The minimum and maximum rate currently cannot be changed by governance (subject to change).

`0x01 -> ProtocolBuffer(TaxRateLimits)`

## TaxCaps

A `TaxCap` is the maximum amount of charity tax that can be charged for a single transaction for a specific denomination. A slice of `TaxCap` is stored in the module `GenesisState`. In the `KVStore` of the Charity module, a `denom` in the form of a string is mapped to a Cosmos SDK `Int`. TaxCaps can be updated via a governance proposal. 

The charity module param store holds a separate set of TaxCap objects, which is synchronized to the `KVStore` at the end of each period during `EndBlock`. When calculating taxes on transactions, the `TaxCap` is fetched from the `KVstore`.

`0x02 | denomBytes -> ProtocolBuffer(sdk.Int)`

## TaxProceeds

The amount of taxes collected during the current period represented by `sdk.Coins`. This is stored in a `CollectionPeriod` object at the end of each period during `EndBlock`, and the TaxProceeds value is reset. This is for tracking purposes, and the actual balance is in the charity tax collector account. 

`0x03 -> ProtocolBuffer(TaxProceeds{TaxProceeds: sdk.Coins})`

Stored under each period are the `TaxProceeds` collected during said period, under the following mapping:

`0x04 | periodBytes -> ProtocolBuffer(sdk.Coins)`

## Payouts

Payouts are mapped based on period for tracking purposes:

`0x05 | periodBytes  -> ProtocolBuffer([]Payout)`





