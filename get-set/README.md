# Overview

This repository contains the get-set smart contract.  Deploy the smart contract via the Xooa console, refer to  <https://github.com/Xooa/samples/blob/master/README.md> to get started, deploying **Xooa-get-set** from this repository.

Here, you'll get to know: how to invoke a smart contract, write data to the blockchain and how to retrieve that data.

**Project documentation:** <https://docs.xooa.com>
___

This app runs on `Hyperledger Fabric`. It is written in `GoLang`.

## Deploy the SmartThings smart contract 


1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Get-Set** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from SmartApp.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.

## Explore the get-set smart Contract end-points

1. Go to the **Details** tab, and then click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Go to **Smart Contract > Invoke Smart Contract Function**.

  	**Invoke** will write to the blockchain. The function **set** will invoke the smart contract.

  	**Query** will read data from the blockchain. The function **get** will query the smart contract.

4. In the **fcn** field, enter `set`.

5. In the **body** field, enter the data you want to store in the blockchain in the format:

  	`[ "<key>", "<value>" ]`

6. Click **try**. 

> **Congrats!** You have saved data in a blockchain using **Xooa**.

7. To view your transaction as part of the blockchain, go to [https://xooa.com/blockchain](https://xooa.com/blockchain/ledger?source=samplesRepo) or navigate to **Ledger** from your Xooa console.

8. Navigate to **Transactions** tab.

9. You can expand the data field to see your transactions.