# Playstation Store Deals Reddit Bot

This reddit bot is designed to see when there are new deals pages from [store.playstation.com](store.playstation.com) posted in the [r/ps4deals](reddit.com/r/ps4deals) subreddit and automatically reply with a comment generating an easy to read to list of what the games are and what the sale prices are. This is to eliminate the manual work many kind people already do in the subreddit.

## To-do
1. Use a reddit api wrapper to listen to get the latest posts on r/ps4deals
1. Consider saving the last timestamp we searched for posts in DynamoDB or something similar to avoid extra work between invocations. Think about DB structure. Consider storing: timestamp, urls used, time comment was added. Consider having to re-run on the same post and marking a comment as sucessful or not
1. Check if there is a r/ps5deals and how that would differ
1. use a reddit api wrapper to post comments
1. Upload to AWS Lambda and test invocations
1. Consider different regions
1. Improve stability