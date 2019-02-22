# Xooa Salesforce Smart Contract

This page provides an overview to Xooa Salesforce Smart Contract functionalities.
This smart contract is used in Xooa-Salesforce app.

This smart contract runs on `Hyperledger Fabric` and is written in `GoLang`.


## Overview

To test the names of the methods in the chaincode any method can be invoked with a single argument as "Xooa Test" and it will return a 200 success response saying "Method test call. Nothing will be committed to ledger". If the method does not exist in the chaincode an error response will be received.

This smart contract provides 3 functions:
  
  * storeData
  * retrieveData
  * deleteData


#### StoreData
StoreData method is used to store the salesforce objects in the ledger.

The storeData method is required to be called with 3 arguments. These arguments need to be the object type, object id and the object data. The first two arguments are used to create a composite key and the third argument is stored as the value json.

A success response returns the composite key.


#### RetrieveData
RetrieveData method is used to fetch value associated with the composite key in the ledger.

The retrieveData method requires 2 arguments - the object type and the object id. The ledger is searched by creating a composite key with the two values.

A success response returns the value associated with the key.


#### deleteData
DeleteData method is used to delete the state of a key from the ledger.

The deleteData method requires 2 arguments - the object type and the object id. The value in the ledger is deleted by using the composite key created with the two input values.

A success response returns the composite key.


## Deploy the Salesforce smart contract 
 
1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Salesforce** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from Dropbox App.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab in the App. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.



## Explore the get-set smart Contract end-points

This smart contract is meant to be used with Salesforce data as in the object id but can be tested with the explorer to get an overview.

1. Go to the **Details** tab, and then click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Go to **Smart Contract > Invoke Smart Contract Function**.

  	**Invoke** will write to the blockchain. The functions **storeData** and **deleteData** are used to invoke the smart contract.

  	**Query** will read data from the blockchain. The function **retrieveData** is used to query the smart contract.

4. In the **fcn** field, enter `storeData`.

5. In the **body** field, enter the data you want to store in the blockchain in the format:

  	`[ "<key>", "<value>" ]`

  Key can be replaced with the object id in Salesforce and value with a JSON containing object details.

6. Click **try**. 

> **Congrats!** You have saved data in a blockchain using **Xooa**.

7. To view your transaction as part of the blockchain, go to [https://xooa.com/blockchain](https://xooa.com/blockchain/ledger?utm_source=samplesRepo) or navigate to **Ledger** from your Xooa console.

8. Navigate to **Transactions** tab.

9. You can expand the data field to see your transactions.