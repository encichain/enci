<!--
order: 0
title: Charity Overview
-->

# `x/charity`

## Abstract

This document specifies the Charity module of EnciChain. 

The Charity module forms the core of EnciChain, and fulfills a major goal of the blockchain ecosystem to contribute positively to the world. Stored by the module are beneficiary(charity) accounts approved by governance. The Charity module is responsible for holding taxes collected from transactions, and initiating payment(donations) to the stored beneficiary accounts at the end of each collection period. A percentage of the taxes burned at the end of each collection period before the payments are disbursed. 

In addition, the module also stores the transaction tax rate, the limits to the tax rates , and a hard cap on the amount of taxes that can be charged. With the exception of the tax rate constraints, the above can be changed via governance proposal. Tracking of collection period data is also performed by the Charity module.