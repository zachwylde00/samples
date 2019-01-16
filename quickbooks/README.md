
This page provides a step-by-step tutorial to integrate the Quickbooks app with Xooa's blockchain platform-as-a-service (PaaS).

# Overview

This repository contains the blockchain smart contract (chaincode). You'll deploy the smart contract via the Xooa console.

Using Xooa, provide a permanent cloud end-point for the QuickBooks app, enable cloud-to-cloud integration while maintaining the peer-to-peer capabilities of blockchain.

## Deploy the QuickBooks smart contract 

 
1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-QuickBooks** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from SmartApp.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.

___

## Connect your company to Xooa Blockchain

1. Go to <https://xooa.com/integrations/quickbooks>. Click **Sign in with Intuit**.

2. Click **Connect to QuickBooks**.

3. Select the company you want to configure and invoke updates to *Xooa smart contract*.

4. Click **Connect**.

5. Enter **API token** you have noted earlier. Enter **function name** to call when any change happens.

	> The app is now configured and ready to receive any event from *QuickBooks* and logs to *Xooa blockchain*.

## Log and view data in Xooa Blockchain

1. Log in to the company you have just connected to the Xooa app.

2. Create or update an entity. Wait for few minutes as it can take upto 10 minutes to log data in the blolckchain.

3. Go to [https://xooa.com/blockchain](https://xooa.com/blockchain/ledger?utm_source=samplesRepo) > Transactions.

4. Expand the **Data** field of latest transaction from the ledger.

	> Congrats! You have logged your QuickBooks entry in the ledger successfully.

<img src="https://github.com/Xooa/samples/blob/master/images/qb_ledger.png" alt="ledger" width="900px"/>
