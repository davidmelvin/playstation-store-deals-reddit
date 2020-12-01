# Playstation Store Deals Reddit Bot

This reddit bot is designed to see when there are new deals pages from [store.playstation.com](store.playstation.com) posted in the [r/ps4deals](http://reddit.com/r/ps4deals) subreddit and automatically reply with a comment generating an easy to read to list of what the games are and what the sale prices are. This is to eliminate the manual work many kind people already do in the subreddit.

## Usage
1. Copy [`.env.sample`](.env.sample) and rename it to `.env` and fill it with the correct values
1. Run `source .env` in the root directory to set the env vars for the bot to authenticate
1. Run `go run .` or `go build`
## To-do
1. Use a reddit api wrapper to listen to get the latest posts on r/ps4deals
1. Consider saving the last timestamp we searched for posts in DynamoDB or something similar to avoid extra work between invocations. Think about DB structure. Consider storing: timestamp, urls used, time comment was added. Consider having to re-run on the same post and marking a comment as sucessful or not
1. Check if there is a r/ps5deals and how that would differ
1. use a reddit api wrapper to post comments
1. Upload to AWS Lambda and test invocations
1. Consider different regions
1. Improve stability
1. Automatically revoke the oauth token when we don't need it anymore
1. Add/improve error handling and logging
1. make CSV and upload to Google Sheets and post that link?
1. Figure out how to support links outside https://store.playstation.com/en-us/category/3fc38af7-0e2c-4de6-a585-3e562e54b81e/1 format. Maybe just accept mentions with those links as comments, and I or someone else can trigger it manually. If it's a link like https://store.playstation.com/en-us/view/25d9b52a-7dcf-11ea-acb6-06293b18fe04/e62084eb-203f-11eb-aadc-062143ad1e8d, can crawl that page for the "See all" link to the correct page
1. Are PSN Plus vs normal discounts different? Where are the differences stored in next.js data? Example of a deals page with a separate PS plus price: https://store.playstation.com/en-us/category/30826e82-088f-4cc2-aaa4-81507aa31353/1. The PS Plus discounts are marked as upsells:
```
"$Product:UP4133-CUSA09936_00-MUDRUNNERUS00000:en-us.price": {
                "basePrice": "$34.99",
                "discountedPrice": "$8.74",
                "discountText": "-75%",
                "isFree": false,
                "serviceBranding": {
                    "type": "json",
                    "json": []
                },
                "upsellServiceBranding": {
                    "type": "json",
                    "json": [
                        "PS_PLUS"
                    ]
                },
                "upsellText": "Save 5% more",
                "__typename": "SkuPrice"
            },
            ```
