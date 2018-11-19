# Overview

This directory contains the get-set smart contract. You will deploy the smart contract via the Xooa dashboard. For a step-by-step getting started, refer to  <https://github.com/Xooa/samples/blob/master/README.md>

___


1. This app will be deployed on `Hyperledger Fabric`. It is written in `GoLang`.

2. While deploying the smart contract, click **Xooa Samples**. Click **Next**.

3. Select **Xooa-get-set** as the smart contract.

### Explore the get-set smart Contract end-points

1. Navigate to **Details** tab, Click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Navigate to **Smart Contract > Invoke Smart Contract Function** from the navigation pane.

  	**Invoke** will be used for making a change to the data inside the blockchain. **set** function will be used for invoking the smart contract.

  	**Query** can be used for reading the data from the blockchain only. **get** function will be used for querying the smart contract.

4. Type `set` in the fcn field.

5. In the **body** field, enter the data you want to store in blockchain in the following format:

  	`[ "<key>", "<value>" ]`

6. Click on **try**.

7. Congrats! You have saved the first data in the blockchain using **Xooa**.

8. To view your transaction as part of the blockchain, navigate to **Ledger** from your Xooa dashboard.

9. Navigate to **Transactions** tab.

10. You can expand the data field to see your transactions.