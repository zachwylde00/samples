
This page provides a step-by-step tutorial to integrate a Google forms add-on with Xooa's blockchain platform-as-a-service (PaaS).

# Overview

This repository contains a Google forms add-on and the blockchain smart contract (chaincode) for that add-on. With these instructions, you can deploy the smart contract via the Xooa console and the Google forms add-on via the Google script editor.

## Deploy the Google forms smart contract 

1. Follow the instructions here: https://docs.xooa.com/start.html#deploy-the-smart-contract-app, selecting the **Xooa-Google-Forms** as the smart contract.

2. Record the **API Token** when it is shown: you will need it to authorize API requests from SmartApp.

   > **Tip:**  to regenerate the API token: 
   >
   > 1. Go to the **Identities** tab. 
   > 2. Next to the ID, select **Actions**.
   > 3. Select **Regenerate API Token**, and then select **Generate**.

___

## Set up the Xooa Google forms add-on

1. Go to <https://docs.google.com/forms> and start creating a new form.

2. From within your new form, click the **More** menu and select **Script editor**. If you're presented with a welcome screen, click **Blank Project**.

3. Create a new HTML file by selecting the menu item **File > New > HTML file**. Name this file **Sidebar** (Apps Script adds the .html extension automatically).

4. Repeat the previous step to create one more HTML file named **About.html**.

5. Replace the content of those files with the code from the repo respectively

6. Select the menu item **File > Save all**. Name your new script "Xooa blockchain logger" and click **OK**.

When you've completed this process, you'll have a project with 1 script file and 2 HTML files.

## Trying out the Xooa Google forms add-on

1. Switch back to your form. Add a text question to your form. Under Question Title, enter 'Email Address'. You can create other form items if you like.

2. After a few seconds, a **Xooa blockchain logger** sub-menu will appear under the **Add-ons extension** menu. (If you chose a different name for your script, that name appears instead.) Click **Add-ons > Xooa blockchain logger**, and in the resulting dialog click Configure logger.

3. **Xooa API Token** in the sidebar that has popped up on the screen.

4. Change the **Function Name** if you've changed it in the smart contract. Otherwise, keep the default name.

5. Click **Save**.

6. Your settings are now saved. The form is now configured to log data in **Xooa blockchain**.

