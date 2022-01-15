# CaduceusTour

**"CaduseusTour"** is a chatbot for the Telegram messenger. First of all, the chatbot displays information about public procurement of the Russian Federation on behalf of public-law entities (federal and municipal institutions) in accordance with Russian Federal Law "44-ФЗ" . There are two options for requesting data:

1. state-owned enterprise details (choice of name and legal address);

The list of tenders is sorted by the date of order placement (the last element is the state order with the earliest placement date).

2. code of the federal subject of Russia;

The list of tenders is sorted in descending order of maximum price. Also available here is an option to get information on contracts placed on behalf of all public-law entities of the Russian Federation.

Once the request has been successfully processed, the following will be available: 
1. *Display a list of tenders*;

Technically, the result is presented as a queue whose size is set to 20 elements and where each element is a separate government contract and includes such information as the order value, service description, date of placement, selected site for auction/competition and so on. In the chat window it looks as follows: the user clicks the "next" button and receives a message from the bot with the next element of the queue, which contains the description of the government contract with the named data. When the user reaches the end of the queue or scrolls it (one of the bot's options when browsing the list), the queue is automatically updated (as long as there are contracts that have an earlier date for placing the order). In addition, all items in the current queue can be downloaded in a text document. 
When browsing an item, you can request a summary of the organisation - in the case of browsing by subject code, a summary will be made for the organisation whose information was sent by the bot in the last message.  

2. *Display a summary of the organisation*;
 
Another function of the chatbot is to produce a summary, where the average cost values and the sum of the maximum and minimum costs with a description of the services are calculated from the data on successfully completed public procurements of the public authority. The figures of the two organisations can then be compared - another chatbot function is responsible for that, which calculates the difference of all values for a specific period (determined automatically) and sends the result as two lists with a single message. Each list is a list with the name of the indicator and the corresponding value describing the positive difference - but the first list lists what the first organisation has an advantage in, and the second one lists what it is inferior in. A list of public procurements of the organisation against which the comparison was made can then be requested. 

____
### EXAMPLE

#### 1. Display information about state contracts from the State Hermitage Museum. 

+ 1.1. After launching the bot, a menu will open in chat with two buttons: "state-owned enterprise" and "federal subject" - here the first must be selected.
+ 1.2. Then a menu will open with the following buttons: "name" and "legal address" - these are search options by detailes, click on the first, in the window with the bot type in the message "Hermitage" and send it to the bot.
+ 1.3. When the request is successfully processed, the following menu will appear: "characteristic", "list of tenders", "back" - of all of these you need "list of purchases". Next the bot will send a message with information about the tender and the view menu will appear ("summary", "next", "document", "menu"). 
+ 1.4. To view other tenders, simply click "next" in the viewing menu. Or "document" - then the bot will offer you to download a text document that contains the information about 20 tenders. Then you can click "scroll through the list" and when the message with the state order numbered as "1 of 20" appears, you can download the next document for the next order.
+ 1.5. To obtain a summary in the view menu there is a button "summary" - simply click on it and the bot will send the ready summary, and the menu of viewing summary will open ("compare", "list of tenders", "menu")
+ 1.6. In order to compare the bulletin with another organization, there is a "compare" button - click on it and then the search options will be offered for the details: "name" and "legal address". This time we will choose "legal address" and enter "191011, Petersburg,  Ostrovski Square, 6" (address of Alexandrinsky Theatre). 
+ 1.7. Then two messages will come from the bot: one is a summary of the Alexandrinsky Theatre, and the second is the mapping itself.
+ 1.8. Next you can see a list of tenders of the Alexandrinsky Theatre by clicking on the "list of tenders" button in the menu.
+ 1.9. Or you can go to the main menu ("state-owned enterprise", "federal subject") by clicking on the "menu" button. 
#### 2. Display information on public contracts from all over the Russian Federation. 

+ 2.1. Press "federal subject" button in the main menu
+ 2.2. A new menu will open with buttons: "all", "subject code", "back". Press "all".
+ 2.3. After processing, the bot will send a message about the tenders with the highest price and the view menu will open ("summary", "next", "document", "menu"). The next tenders are arranged in descending order of price - there are "next" and "document" buttons to view them. If you click the second button, the bot offers you to download a text document with 20 tenders. The "feature" button shows a summary of the organisation whose information was sent by the bot in the last message.
 
 ____
### technical information

The connection to the Telegram messenger API is made using the Webhook method. 

The GosPlanAPI service is used to access and retrieve data. Including the mechanism of automatic renewal of the access key when its validity expires has been implemented. 

Repository: NoSQL database MongoDB.

Among the libraries involved there are: 

- goquery (site parsing (needed to identify a public entity by legal address);
- logrus (logging);
- goConvey (testing);

etc. 

#### TODO:
1. english localization
2. save summaries in the database 
3. redesign the search by legal address
 
