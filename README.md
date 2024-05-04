## Inspiration
EWB project.

The inspiration behind our project was to create a user-friendly and intuitive platform that allows individuals to interact with their data using natural language. We aimed to harness the power of AI to simplify data analysis and empower users to derive insights from their data effortlessly.

## What it does
Our app, DataBuddy, enables users to chat with their data using natural language queries. Utilizing the Google Gemini AI model, DataBuddy interprets user inputs and provides responses based on the actions and analyses the data can perform. Users can ask questions, issue commands, or request specific analyses, receiving instant feedback in a conversational manner.

## How we built it
We built DataBuddy using the Google Gemini AI model as the core engine for natural language processing. Additionally, we leveraged Langchain's create_pandas_dataframe_agent to handle data manipulation and analysis tasks efficiently. For the frontend, we utilized Streamlit, a powerful Python library for building interactive web applications with minimal code.

Here is a flow diagram on how it works
![Here](https://cdn.discordapp.com/attachments/1223088660572803147/1235828913805066300/Screenshot_from_2024-05-03_06-31-13.png?ex=6635cadd&is=6634795d&hm=8999933cac3bf1ffaa5cb31ce7e330a4e3a4270602958af975104e3460bd43c0&)

## Challenges we ran into
One of the main challenges we encountered was configuring the prompts for the Google Gemini AI model to ensure accurate and relevant responses. Finding the right prompts that effectively capture the user's intent and generate meaningful outputs required experimentation and fine-tuning. Overcoming this challenge was crucial to delivering a seamless user experience.

## Accomplishments that we're proud of
We are proud to have developed a functional prototype of DataBuddy within the timeframe of the Google AI Hackathon. Despite facing challenges in prompt configuration, we were able to implement a robust solution that showcases the capabilities of natural language processing in data analysis. Additionally, integrating Langchain's create_pandas_dataframe_agent and Streamlit frontend allowed us to create a cohesive and user-friendly application.

## What we learned
Through the process of building DataBuddy, we gained valuable insights into the intricacies of natural language processing and its applications in data analysis. We learned how to effectively utilize AI models like Google Gemini to interpret user queries and perform data-related tasks. Additionally, we enhanced our skills in frontend development using Streamlit, enabling us to create interactive and visually appealing interfaces.

## What's next for DataBuddy
Moving forward, we plan to further refine and enhance DataBuddy's capabilities. This includes expanding the range of supported data analyses and implementing advanced features for data visualization and exploration. Additionally, we aim to optimize the prompt configuration process to improve the accuracy and relevance of responses, also, to make a more robust agent that can go on to automate the who business analysis tasks. Ultimately, our goal is to continue evolving DataBuddy into a powerful and indispensable tool for data-driven decision-making.