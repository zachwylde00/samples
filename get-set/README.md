
This page provides a step-by-step tutorial to deploy a sample app with Xooa's blockchain-as-a-service (BaaS).

The repository used in this example is <https://github.com/Xooa/samples>

# Overview

This repository contains the blockchain smart contract (henceforth smart contract). You will deploy the chaincode via the Xooa dashboard.

Using Xooa, you can provide a permanent cloud end-point for a smart contract, enabling cloud-to-cloud integration while maintaining the peer-to-peer capabilities of blockchain.

## Deploy the get-set chaincode 

 
1. Log in to Xooa.

2. Navigate to **Apps**, then **Deploy New**.

3. Search for **Xooa/samples**, and click **Select**, and then **Next**.

4. Select **Xooa-get-set** as the Smart Contract, and then click **Deploy**.

5. Relax:  Xooa is doing the blockchain heavy lifting. You will be redirected to app dashboard when the deployment completes.

6.  Navigate to **Identities** tab, click **Show API Token**.

7. Copy and store the **API Token** values. You need this to authorize API requests.

___

## Explore the end-point for the smart Contract

1. Navigate to **Details** tab, click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Navigate to **Smart Contract > Invoke Smart Contract Function** from the navigation pane.

4. We will be calling **set** function from the smart contract. Type `set` in the fcn field.

5. In the **body** field, enter the data you want to store in blockchain in the following format:

`{
  "args": [
    "<key>",
    "<value>"
  ]
}`

6. Click on **try**.

7. Congrats! You have saved the first data in the blockchain using **XOOA**.

8. To view your transaction as part of the blockchain, navigate to **Ledger** from your Xooa dashboard.

9. Navigate to **Transactions** tab.

10. You can expand the data field to see your transactions.