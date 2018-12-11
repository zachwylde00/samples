/**
 * @OnlyCurrentDoc
 *
 * The above comment directs Apps Script to limit the scope of file
 * access for this add-on. It specifies that this add-on will only
 * attempt to read or modify the files in which the add-on is used,
 * and not all of the user's files. The authorization request message
 * presented to users will reflect this limited scope.
 */

/**
 *  Xooa Google-forms add-on
 *
 *  Copyright 2018 Xooa
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at:
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 *  for the specific language governing permissions and limitations under the License.
 *  
 *  Author: Arisht Jain
 *  Last Modified: 08/10/2018
 */

/**
 * A global constant String holding the title of the add-on. This is
 * used to identify the add-on in the notification emails.
 */
var ADDON_TITLE = 'Xooa blockchain';

/**
 * Adds a custom menu to the active form to show the add-on sidebar.
 *
 * @param {object} e The event parameter for a simple onOpen trigger. To
 *     determine which authorization mode (ScriptApp.AuthMode) the trigger is
 *     running in, inspect e.authMode.
 */
function onOpen(e) {
  FormApp.getUi()
      .createAddonMenu()
      .addItem('Configure logger', 'showSidebar')
      .addItem('About', 'showAbout')
      .addToUi();
}

/**
 * Runs when the add-on is installed.
 *
 * @param {object} e The event parameter for a simple onInstall trigger. To
 *     determine which authorization mode (ScriptApp.AuthMode) the trigger is
 *     running in, inspect e.authMode. (In practice, onInstall triggers always
 *     run in AuthMode.FULL, but onOpen triggers may be AuthMode.LIMITED or
 *     AuthMode.NONE).
 */
function onInstall(e) {
  onOpen(e);
}

/**
 * Opens a sidebar in the form containing the add-on's user interface for
 * configuring the notifications this add-on will produce.
 */
function showSidebar() {
  var ui = HtmlService.createHtmlOutputFromFile('Sidebar')
      .setSandboxMode(HtmlService.SandboxMode.IFRAME)
      .setTitle(ADDON_TITLE);
  FormApp.getUi().showSidebar(ui);
}

/**
 * Opens a purely-informational dialog in the form explaining details about
 * this add-on.
 */
function showAbout() {
  var ui = HtmlService.createHtmlOutputFromFile('About')
      .setSandboxMode(HtmlService.SandboxMode.IFRAME)
      .setWidth(420)
      .setHeight(270);
  FormApp.getUi().showModalDialog(ui, 'About ' + ADDON_TITLE);
}

/**
 * Save sidebar settings to this form's Properties, and update the onFormSubmit
 * trigger as needed.
 *
 * @param {Object} settings An Object containing key-value
 *      pairs to store.
 */
function saveSettings(settings) {
   PropertiesService.getDocumentProperties().setProperties(settings);
   adjustFormSubmitTrigger();
}

/**
 * Queries the User Properties and adds additional data required to populate
 * the sidebar UI elements.
 *
 * @return {Object} A collection of Property values and
 *     related data used to fill the configuration sidebar.
 */
function getSettings() {
  var settings = PropertiesService.getDocumentProperties().getProperties();

  // Get text field items in the form and compile a list
  //   of their titles and IDs.
  var form = FormApp.getActiveForm();
  var textItems = form.getItems(FormApp.ItemType.TEXT);
  settings.textItems = [];
  for (var i = 0; i < textItems.length; i++) {
    settings.textItems.push({
      title: textItems[i].getTitle(),
      id: textItems[i].getId()
    });
  }
  return settings;
}

/**
 * Adjust the onFormSubmit trigger based on user's requests.
 */
function adjustFormSubmitTrigger() {
  var form = FormApp.getActiveForm();
  var triggers = ScriptApp.getUserTriggers(form);
  var settings = PropertiesService.getDocumentProperties();

  // Create a new trigger if required; delete existing trigger
  //   if it is not needed.
  var existingTrigger = null;
  for (var i = 0; i < triggers.length; i++) {
    if (triggers[i].getEventType() == ScriptApp.EventType.ON_FORM_SUBMIT) {
      existingTrigger = triggers[i];
      break;
    }
  }
  if (existingTrigger) {
    ScriptApp.deleteTrigger(existingTrigger);
  }
  var trigger = ScriptApp.newTrigger('sendResponse')
  .forForm(form)
  .onFormSubmit()
  .create();
}


// Invokes the Smart Contract to log response in Xooa blockchain
function sendResponse() {
  var form = FormApp.getActiveForm();
  var settings = PropertiesService.getDocumentProperties();
  var formResponses = form.getResponses();
  if(formResponses.length > 0){
    var formResponse = formResponses[formResponses.length-1];
    var itemResponses = formResponse.getItemResponses();
    var responseKey = formResponse.getId();
    if(formResponse.getRespondentEmail()) {
      responseKey = formResponse.getRespondentEmail();
    } else {
      responseKey = formResponse.getId();
    }
    
    // Create a JSON value to log in Xooa blockchain
    var value = '{';
    for (var j = 0; j < itemResponses.length; j++) {
      var itemResponse = itemResponses[j];
      value += '\"' + itemResponse.getItem().getTitle() + '\":\"' + itemResponse.getResponse() + '\"';
      if(j<itemResponses.length-1) {
        value += ',';
      }
    }
    value += '}';
    
    var json = [
        responseKey,
        value
      ]

    Logger.log(json);
    console.log(json);
      
    // Call invoke API to log response in Xooa blockchain
    var options = {
      'method': 'post',
      'contentType': 'application/json',
      'headers': { 
        "Authorization": "Bearer " + settings.getProperty('apiToken')
      },
      'payload': JSON.stringify(json),
      'muteHttpExceptions': true
    }
    
    var url = "https://api.xooa.com/api/v1/invoke/" + settings.getProperty('functionName')
    var response = UrlFetchApp.fetch(url, options);
    Logger.log(response.getContentText())
    console.log(response.getContentText())
    
    if(response.getResponseCode() == 202) {
      // Invoke request was not completed immediately. Going to call results API.
      var sleepTime = 3000
      var requestCount = 5
      var i = 0
      var responseStatus = 202
      var url1 = "https://api.xooa.com/api/v1/results/"+JSON.parse(response.getContentText()).resultId
      var options1 = {
      'method': 'get',
      'contentType': 'application/json',
      'headers': { 
        "Authorization": "Bearer " + settings.getProperty('apiToken')
      },
      'muteHttpExceptions': true
      }
      
      while (i < requestCount && responseStatus == 202) {
        // Invoke request has not been processed yet.
        Utilities.sleep(sleepTime);
        i++;
        var response1 = UrlFetchApp.fetch(url1, options1);
        if(response1.getResponseCode() == 202) {
          continue;
        } else {
          // Invoke request is completed.
          responseStatus = response1.getResponseCode();
          Logger.log(response1.getContentText())
          console.log(response1.getContentText())
        }
      }
      Logger.log("Request to results API completed. Api ran for "+ i + " times.")
      console.log("Request to results API completed. Api ran for "+ i + " times.")
    }
  } else {
    Logger.log("No responses received.")
    console.log("No responses received.")
  }
}