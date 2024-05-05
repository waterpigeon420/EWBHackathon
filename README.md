## Inspiration
EatYourFill

The inspiration behind our project was to create a user-friendly and intuitive platform that allows individuals to be aware of their food purchases, portions and sustainable choices to reduce household food waste. We identified that targetting individual household food waste would make a large difference in mitigating overall waste.

## What it does
Our app, EatYourFill, enables users to generate portioned recipes given their personal preference and bio-data. The portions and recipes are designed utilizing the Google Gemini AI model, given the users data especially physical exertion data, caloric intake data, BMI etc. The app provides exact ingredients and amounts for each recipe allowing the user to budget and plan for a low waste food purchase. A nutritional breakdown of all the ingredients is also provided. The app also provides a sustainability score for each recipe allowing the users to see the environmental impact of their dietary choices.

## How we built it
We built EatYourFill using the Google Gemini AI model as the core engine for prompting and generating recipe data. Additionally, we used golang's sdk for gemini to set up a simple local server to process the responses from the model. For the frontend, we utilized Flutter, an easy to deploy lightweight frontend engine. We then finally used the Edamam API to get a nutritional breakdown for our suggested ingredients.

## Challenges we ran into
One of the main challenges we encountered was configuring the prompts for the Google Gemini AI model to ensure accurate and relevant responses. Finding the right prompts that effectively capture the user's intent and generate meaningful outputs required experimentation and fine-tuning. Overcoming this challenge was crucial to delivering a seamless user experience.

## Accomplishments that we're proud of
We are proud to have developed a functional prototype of EatYourFill within the timeframe of the EWBHackathon. Despite facing challenges in prompt configuration, we were able to implement a robust solution that showcases the capabilities of natural language processing in recipe generation and sustainability analysis. Additionally, We were able to build and integrate a comprehensive Flutter frontend with our simple golang server to provide the user with all the relevant food, sustainability and nutritional information.

## What we learned
Through the process of building DataBuddy, we gained valuable insights into the intricacies of natural language processing and its applications in data analysis. We learned how to effectively utilize AI models like Google Gemini to interpret user queries and perform data-related tasks.

## What's next for EatYourFill
Moving forward, we plan to further refine and enhance EatYourFill's capabilities. This includes expanding the range of user data used to suggest recipes and potion and implementing advanced features for prediction of user eating and purchasing behaviours.