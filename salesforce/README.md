# Xooa Salesforce Smart Contract

This page provides a step-by-step tutorial to integrate a Salesforce app with Xooa's blockchain-as-a-service (BaaS).

The repository used in this example is <https://github.com/Xooa/integrations>

## Overview

This repository contains the blockchain chaincode (henceforth chaincode) and the Salesforce app. You will deploy the chaincode via the Xooa console and the Salesforce app via the Salesforce setup page.

Using Xooa, you can provide a permanent cloud end-point for Salesforce, enabling cloud-to-cloud integration while maintaining the peer-to-peer capabilities of blockchain.

## Deploy the Salesforce chaincode 

1. Log in or create a Xooa account at [https://xooa.com/blockchain](https://xooa.com/blockchain?utm_source=samplesRepo)

2. Click **apps**, then **Deploy New**. 
If this is your first time deploying a chaincode app with Xooa, you will need to authorize Xooa with your GitHub account.

    a. Click **Connect to GitHub**.

    b. Follow the onscreen instructions to complete the authorization process.

3. Search for the **integrations** repo (or your fork).

4. Select the repo, and then click **Next**.

5. Enter a name and description for your app.

6. Select the branch (usually **master**) and **salesforce** as the chaincode, and then click **Deploy**.

7. Relax.  Xooa is doing the blockchain heavy lifting. You will be redirected to app dashboard when the deployment completes.

8.  On the **Identities** tab, click **Show API Token**.

9. Copy and store the **App ID** and **API Token** values. You need these to authorize API requests in your **Salesforce** app.

___

## Developer account and installation
A developer account is required for running the salesforce app from the IDE. 

1. Visit developer.salesforce.com and complete the process to get a developer account.

2. https://login.salesforce.com/packaging/installPackage.apexp?p0=04t6F0000039CP0&isdtp=p1 Open this link and login to begin the instllation process. 

3. After finishing the installation, open App Launcher and click on Xooa Blockchain to open the app.

4. In the first page named 'Settings', under Xooa Settings section, enter the **App ID** and **API Token**. For function name, enter **saveObjectData** as the value and click `Save`. The fuction name parameter corresponds to the function described in the chaincode we deployed earlier and can be changed as per your (i.e., chaincode developer's) need.

5. On the right section of page under Salesforce Settings section, select all the triggers you like and hit `Save`.

6. This completes the basic setup for using Salesforce with Xooa. The triggers and their operations can be changed dynamically on the Settings page to enable/disable saving data to Xooa.


## Troubleshooting
You can open `Developer Console` while being logged in to Salesforce and check the **Logs** section entries to troubleshoot the errors.
