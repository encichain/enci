<!--
order: 1
-->

# Concepts

## Charities

A charity is a beneficiary account stored by the module as a param. These accounts are determined by governance proposals. Once a valid charity is set, the charity will receive payments in the form of donations from transaction taxes collected during each collection period.

The structure of a `Charity` object is as follows:

```go
type Charity struct {
	CharityName string // Name of the charity
	AccAddress  string // Account address in form of Bech32 string
	Checksum    string // SHA-256 checksum of CharityName + AccAddress
}
```
## Payouts

A `Payout` is an object representing a successful donation to a set charity account. At the end of each collection period during `EndBlock`, payments are disbursed from the charity tax collector account, with the payment quantity determined by dividing the total spendable balance of the collector account by the number of charities stored. 

## Collection Periods

A collection period is a period of time in which taxes were collected from transactions. These periods are defined for the purpose of determining the frequency of payments to charities. 
In addition, a percentage of the taxes collected are burned at the end of each period, which is determined by the BurnRate. 


## Updating Params

As a decentralized blockchain, certain params of the blockchain can be updated and changed via community governance proposals. This includes the charities, the `TaxRate`, `BurnRate`, and `TaxCaps`. 