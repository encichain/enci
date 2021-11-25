<!--
order: 5
-->

# Parameters

Charity module parameters are handled by the Cosmos SDK x/params module, and parameter key pairs are stored in a module specific `subspace` of the parameter store.

The Charity module contains the following parameters:

| Key                | Type          | Example (JSON format)                                |
| ------------------ | ------------- | -----------------------------------------------------|
| Charities          | []Charity     | [{charity_name: "", accAddress: "", checksum: ""}]   |
| TaxRate            | string(Dec)   | "0.005000000000000000"                               |
| TaxCaps            | []TaxCap      | [{denom: "uenci", Cap: "1000000"}]                   |
| BurnRate           | string(Dec)   | "0.035000000000000000"                               |


## Param Changes
All parameters can be changed via governance proposal. Param changes to the Charity module are not validated by the app before they are set. It is up to the community to vet change proposals to ensure no state breaking changes occur.

## Charities

The `Charities` param is stored in the param store to enable change via governance `param-change` proposal. It is a slice of `Charity` objects. This structure allows more than one charity to be set. In a param change proposal, the proposer must specify the entire slice of `Charity`. Only basic validation checks are done during module genesis and when the param pair is set. It is during disbursement of tax payments that each `Charity` in the list is checked for validity. 

## TaxRate

`TaxRate` determines the tax that is be collected from each applicable transaction message. This tax is not determined by the local validator, and is instead a universal tax. Tax fees are added onto Gas fees for a transaction, and summed as `Fee`. The handling of tax calculation and deduction is done via the `AnteHandler`. `TaxRate` is constrained by `TaxRateLimits`, which determine the minimum rate and the maximum rate that can be set. 

## TaxCaps
The `TaxCaps` in the parameter store are not used directly by the app. Instead, a separate copy in the KVStore is fetched instead to be used for tax calculation. However, the KVStore tax caps are synchronized with the parameter store `TaxCaps`. As such, `TaxCaps` updated via param-change proposals do not take effect until the KVStore is updated. 

## BurnRate

`BurnRate` determines the percentage of spendable balance of the charity tax collector account that is to be sent to the module `Burn` account at the end of each period. `BurnRate` is also constrained. 