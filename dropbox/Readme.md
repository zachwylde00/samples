# Xooa Dropbox Smart Contract

This page provides an overview to the Xooa Dropbox Smart Contract functionalities.
This smart contract is used in Xooa-Dropbox app.

This smart contract runs on `Hyperledger Fabric` and is written in `GoLang`.


## Overview

This smart contract provides 5 functions:

  * SaveProfile
  * GetProfile
  * SaveNewEvent
  * GetEntityDetails
  * GetHistoryForEntity
    

#### saveProfile

This function is used to save the user profile in the blockchain. 
The profile details contains the latest Dropbox OAuth Key and the latest Dropbox cursor for the account id. The profile can also contain the Xooa API Token for the Xooa App to whcich the user wants to connect to process the data.
This function takes in two arguments in an array of strings - the first argument is the account id for the user and the second is the profile json.
This function is called after every webhook event processing for each account to update the new cursor associated with the account id.


#### getProfile

This function is used to retreive the user's profile's latest state from the ledger.
This function takes in the account id as its only argument.
This function is called whenever a webhook request is received from dropbox with a list of account ids which have changes in their account. All these account ids are queried one by one to get the latest cursor from the profile in order to find out the changes from the previous state


#### saveNewEvent

This function is used to store the details for each of the changes detected in the webhook request. 
This function takes two arguments in an array of strings - the first argument is the unique document id for the dropbox document and the second argument is the data (mostly metadata) associated with the document.
Whenever a webhook request is received all the account ids are processed for changes using the cursor stored in the profile. For all the documents associated with the account that have changes saveNewEvent is called to store the new state.


#### getEntityDetails

This function is used to get the latest details associated with the document.
This function takes in the document id as its only argument.


#### getHistoryForEntity

This function is used to get the complete history for a id from the blockchain.
This function takes in only one argument - the key for which to retreive the history.


## Deploy the Dropbox smart contract 
 
1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Dropbox** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from Dropbox App.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab in the App. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.



## Explore the get-set smart Contract end-points

This smart contract is meant to be used with dropbox data as in the account id or the file id in the dropbox but can be tested with the explorer to get an overview.

1. Go to the **Details** tab, and then click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Go to **Smart Contract > Invoke Smart Contract Function**.

  	**Invoke** will write to the blockchain. The function **saveProfile** and **saveNewEvent** are used to invoke the smart contract.

  	**Query** will read data from the blockchain. The function **getProfile** and **getEntityDetails** is used to query the smart contract.

4. In the **fcn** field, enter `saveProfile`.

5. In the **body** field, enter the data you want to store in the blockchain in the format:

  	`[ "<key>", "<value>" ]`

  Key can be replaced with the account id in dropbox and value with a JSON containing account id and OAuth token associated with the account .

6. Click **try**. 

> **Congrats!** You have saved data in a blockchain using **Xooa**.

7. To view your transaction as part of the blockchain, go to [https://xooa.com/blockchain](https://xooa.com/blockchain/ledger?utm_source=samplesRepo) or navigate to **Ledger** from your Xooa console.

8. Navigate to **Transactions** tab.

9. You can expand the data field to see your transactions.