# Assignment used for Data Scientist Position

This task is intended for candidates applying for a Data Science position at e-conomic. The assignment contains real data and directly reflects the actual challenges we face at e-conomic when we are trying to apply machine learning to the field of automating accounting processes.

To limit the workload the task has been split into two parts. In the first part you are expected to write code, build predictive models, measure the quality of these models and document your results. In the second part you are not expected to implement a predictive model, but instead you are expected to devise and propose a theoretical solution to the problem.

## The problem
Companies often have a lot of expenses. Each payment of these expenses is in accounting terms called a financial transaction which should be entered into the financial books. But where? This is the job of the accountant to figure out. An expense for paint might go to the account for renovation, and an expense for a taxi ride might go to the account for travel expenses.

In e-conomic it is possible to import your bank statements and use these to create the financial transactions, so all you have to do is to import the bank statements, and decide on which account you want to book each expense ..and you’re done.

But what if we could make this process even smoother by learning how the expenses should be booked? Wouldn’t that be great? If we know that a company's books looks like this:

    “Cleaning service & Co”	        --> Maintenance account
    “Andersen’s cleaning service”   --> Maintenance account
    “McDonald's Nørrebro”           --> Employee catering account

We might just go ahead and help the company do future bookings of _“McDonald's Nørrebro”_, _“Copenhagen Cleaning Company”_ or maybe we might even be able to help with _“Burger King Kastrup Airport”_, if we look at how other companies usually do their bookkeeping.

__A few things to notice:__

* The same text from the bank statements might lead to two different accounts. An example of this is that you might buy both food for your employees (which you might want to put on the employee catering account) and chocolate (which you might want to put on the gifts and flowers account) in the same shop.
* Each company have their own chart of accounts. However many companies have a chart of account that varies only a little from the e-conomic default chart of accounts.
* Two companies that have the same chart of account might not agree how certain expenses should be booked. An example of this might be sandwiches bought for a meeting with business partners. One company might always put this on the employee catering account, another one might always put it on the meetings account, and a third might even have a sandwiches for meetings with business partners account.
* Each account has a number. There are no restrictions on these account numbers besides that an account number is unique within a company. So the travel expense account in one company might have the same number as the employee catering account in another company.

## Part 1 - Per company prediction

How can we help the accountant do the bookkeeping of the company’s expenses? To start out simple and to limit the amount of workload, we avoid the challenge of companies having different chart of accounts by making per company predictions, ie. train a model for each company and make each prediction based on only the company's own model.

In this part you are expected to write code, build predictive models, measure the quality of these models and document your results.

Your solution to this task you will be judged on criteria related to the product that you will be delivering when working at e-conomic, eg. among other things: coding style, approach to problem solving, guiding the research by data driven decisions, documentation of your results, choice of quality measure, etc.

## Part 2 - Can we do better?

Alright, in Part 1 above you developed a predictive model to predict accounts based on each company’s own data. What if we wanted to utilize all the data we have across all 100.000 companies? Then we would also be able to help in situations where we get new kind of expenses that we have never seen before within a company. But how should we even tackle the problem of constructing such a model?

In this part you are not expected to implement a predictive model, but instead you are expected to devise and propose a theoretical solution to the problem.

Your solution to this task you will be judged on criteria related to the value that you are bringing to the team, eg. creativity, knowledge within mathematical modeling, communicating your ideas, ability to discuss pros and cons of different solutions, etc.

## Description of data
The dataset consist of expenses from 100 random companies. For each company we provide all expenses that was booked in e-conomic. The rows are sorted first by date, then by company id.

Due to privacy reasons the amounts has been bucketed and the texts has been obfuscated using the following function:

// TODO CODE HERE

Here is a small sample of what the data looks like:

// TODO DATA SAMPLE HERE

// TODO MENTION WHERE THE DATA IS LOCATED

## Got stuck?
You can always email Helge (helge.munk.jacobsen@visma.com) and ask for advice or just ask question to ensure you correctly understood the task. This will not be seen as a sign of weakness, to the contrary it shows that fully understanding the problem is important to you.

## Suggestions for improvements?

Please help us improve this assignment. Feel free to make a pull request!
