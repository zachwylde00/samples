# Xooa Ethereum-Get-Set Smart Contract

This page provides an overview to Xooa-Ethereum-Get-Set Smart Contract functionalities.

This smart contract runs on `Ethereum` and is written in `Solidity`.


## Overview

This smart contract provides 2 functions:
  
  * Get
  * Set


#### Get

Get method is used to fetch the value associated with a key passed in the arguments.
This method expects a single argument as the key whose world state is required.
If a value is found for the key then it returns the value or else it returns an error message.


#### Set

Set method is used to store the key value pair in the ledger.
This method requires two arguments and takes the first as the key and second as value.
This method creates a transaction in the blockchain ledger and stores the key value pair.
If it succeeds in creating the transaction it returns a response with key value pair as payload or else an error response.



## Deploy the Xooa-Ethereum-Get-Set smart contract 
 
1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Ethereum-Get-Set** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from Dropbox App.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab in the App. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.



## Explore the Xooa-Ethereum-Get-Set smart Contract end-points

1. Go to the **Details** tab, and then click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Go to **Smart Contract > Invoke Smart Contract Function**.

  	**Invoke** will write to the blockchain. The function **set** is used to invoke the smart contract.

  	**Query** will read data from the blockchain. The function **get** is used to query the smart contract.

4. In the **fcn** field, enter `set`.

5. In the **body** field, enter the data you want to store in the blockchain in the format:

  	`[ "<key>", "<value>" ]`

6. Click **try**. 

> **Congrats!** You have saved data in a blockchain using **Xooa**.

7. To view your transaction as part of the blockchain, navigate to **Ledger** from your App Details page.

8. Navigate to **Transactions** tab.

9. You can expand the data field to see your transactions.