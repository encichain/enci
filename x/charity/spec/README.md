<!--
order: 0
title: Charity Overview
-->

# `x/charity`

## Abstract

This document specifies the Charity module of EnciChain. 

The Charity module forms the core of EnciChain, and fulfills a major goal of the blockchain ecosystem to contribute positively to the world. Stored by the module are beneficiary(charity) accounts approved by governance. The Charity module is responsible for holding taxes collected from transactions, and initiating payment(donations) with the collected taxes to the stored beneficiary accounts at the end of each collection epoch. A percentage of the taxes burned at the end of each collection epoch before the payments are disbursed. 

In addition, the module also stores the transaction tax rate, the limits to the tax rates , and a hard cap on the amount of taxes that can be charged. With the exception of the tax rate constraints, the above can be changed via governance proposal. Tracking of collection epoch data is also performed by the Charity module.

## Contents
 
1. **[Concepts](01_concepts.md)**
    - [Charities](01_concepts.md#Charities)
    - [TaxRateLimits](01_concepts.md#TaxRateLimits)
    - [Payouts](01_concepts.md#Payouts)
    - [Collection Epochs](01_concepts.md#Collection_Epochs)
2. **[State](02_state.md)**
    - [Key-Value Mapping](02_state.md#Key-Value-Mapping)
    - [TaxRateLimits](02_state.md#TaxRateLimits)
    - [TaxCaps](02_state.md#TaxCaps)
    - [TaxProceeds](02_state.md#TaxProceeds)
    - [Payouts](02_state.md#Payouts)
3. **[EndBlock](03_endblock.md)**
4. **[Events](04_events.md)**
5. **[Params](05_params.md)**
    - [Param Changes](05_params.md#Param-Changes)
    - [Charities](05_params.md#Charities)
    - [TaxRate](05_params.md#TaxRate)
    - [TaxCaps](05_params.md#TaxCaps)
    - [BurnRate](05_params.md#BurnRate)
