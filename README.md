
This page provides a step-by-step tutorial to deploy a sample app with Xooa's blockchain-as-a-service (BaaS).

# Overview

This repository contains the blockchain smart contract (henceforth smart contract). You will deploy the smart contract via the Xooa dashboard.

Using Xooa, you can provide a permanent cloud end-point for a smart contract, enabling cloud-to-cloud integration while maintaining the peer-to-peer capabilities of blockchain.

## Deploy the smart contract

 
1. Log in to Xooa using an exisitng GitHub, Google or Facebook account.

2. Go to **Apps**>**Deploy New**, if you looged in using Google or Facebook you will be prompted to connected GitHub account. This is required.

3. Search for the Github repository you want to use for deployment. For example, type **Xooa/samples** to deploy one of the sample provided by Xooa, then click **Select**, and then **Next**.

<img src="https://github.com/Xooa/samples/blob/master/screenshots/deploy.png" alt="HLD" width="500px"/>

4. Select the Smart Contract you want to deploy, and then click **Deploy**.

5. Relax:  Xooa is doing the blockchain heavy lifting. You will be redirected to app dashboard when the deployment completes.

6.  Navigate to **Identities** tab, click or tap **Add New**, enter name for Identity  and set permissions to Read+Write. Copy the **API Token** and save it in secure place. You will need when calling the APIs.

7. Copy and store the **API Token** values. You need this to authorize API requests.

___


## The Xooa app dashboard 

The dashboard consists of the following tabs:
<dl>
  <dt>Details</dt>
  <dd>Information about your app.</dd>
<dt>Identities</dt>
  <dd>Create or delete an identity. You can also specify access rights to different identities, thus controlling the access of endpoints for your smart contract</dd>
   <dt>Activities</dt>
  <dd>History of your app's activities./dd>
      <dt>Logs</dt>
 <dd>View smart contract events from the last 10 minutes.</dd>
  <dt>Manage</dt>
 <dd>Delete your app or update from github.</dd>
</dl>

___

## Explore the end-points for the smart Contract

1. Go to the **Details** tab, click **Explore API's**.

2. Enter **API Token** in the field in navigation pane.

3. Go to the **Smart Contract > Invoke Smart Contract Function** from the navigation pane.

<img src="https://github.com/Xooa/samples/blob/master/screenshots/invoke.png" alt="HLD" width="300px"/>

4. In the **body** field, enter the data in the format expected by the smart contract. 
For example, if using this repo **get-set** smart contract to store data in blockchain ledger using the **set** function, the body should be entered as: 

    `{ "args": [ "<key>", "<value>" ] }`

5. Click  **try**.
 * A response code of **200** indicates successful function call of the smart contract.
 * A response code of **202** indicates that your request is queued for processing. Final processing outcome may be obtained through **Result** end point by using **resultId** obtained in **response body**.
 * A response code of **400** indicates that you have a malformed request. Check the **body** field again.
 * A response code of **401** indicates that either you forgot to enter API token you have entered invalid API token.
Congrats! You have saved the first data in the blockchain using **Xooa**.

7. To view your transaction as part of the blockchain, from your Xooa dashboard, go to **Ledger**.

8. Go to the **Transactions** tab.
You can expand the data field to see your transactions.
