<!--
order: 3
-->

# EndBlock

`Endblock()` is called at the end of every block. A check is performed to determine if it is the end of a epoch. This is determined by the function `IsLastBlockEpoch()`, which is called during each `EndBlock`. If it is the last block of the epoch, the following is performed during `EndBlock`:

1. Calculate the amount of coins to be sent to burner account and burned based on the spendable balance of the charity tax collector account

2. Disburse donations from the total spendable balance of the charity tax collector to the set `Charity` accounts.

3. Set the `TaxProceeds` to the store under the current epoch, and reset the current `TaxProceeds` for the next epoch.

4. Synchronize the params store `TaxCaps` with the KVStore `TaxCaps`. The params store `TaxCap` can be changed via governance proposal, but is not used for transactions. When `TaxCap` is updated via proposal, it will be put into effect only after it is synchronized at the end of a epoch. 

5. Emit events `EventPayout` and `EventFailedPayouts`(if applicable) for tracking purposes. 
