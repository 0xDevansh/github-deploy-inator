
<h1 align="center">
    <img alt="Project logo" width="100" src="docs/logo.png" align="center"> GitHub Deploy-inator</h1>

<div align="center">
    <a href="https://discord.gg/qJnrRvt7wW">
        <img alt="Discord" src="https://img.shields.io/discord/873232757508157470?color=%235865F2&label=support&style=for-the-badge">
    </a>
</div>

> Automatic deployment app based on GitHub webhooks 

## 💡 Motivation

I code and maintain a lot of Discord bots and other projects, all
running on my web server. Every time I pushed an update to a bot, I'd
have to SSH in, pull the code from Github, build it and restart the
process. Although I managed to boil this down into a sweet `yarn deploy`,
that still needed me to SSH into the server. I tried implementing git hooks
but to no avail.

And then, this project was born.

## 🛠️ Installation and Setup

Note to future me: Write out this part

## 📝 config.json

All the required data must be provided in a config.json file, placed in the current
working directory.

### Format

- `port`: The port on which the application will listen for webhooks
    - type: `string`
    - format: `":DDDD"`, where D is a digit
    - example: `":8000"`, `":440"`
      <br><br>
  
- `endpoint`: The endpoint where the webhooks will be sent
    - type: `string`
    - format: `"/*"`
    - example: `"/webhooks/github"`, `"/github/listener"`
      <br><br>
  
- `listeners`: Settings for individual listeners
  - type: `Listener[]`

#### Listener

- `name`: [required] A unique name for the listener. This is mentioned when a webhook is received, executed or failed.
    - type: `string`
    - example: `my-chat-app` (try not to include spaces)
      <br><br>
    
- `repository`: [required] The full name of the repository for which this webhook will be executed.
  - type: `string`
  - format: `"author-name/repository-name"`
  - example: `"DeathVenom54/github-deploy-inator"`
    <br><br>

- `directory`: [required] The absolute path to the directory (folder) where the command will be executed.
  - type: `string`
  - example: `"E:/projects/github-deploy-inator"`, `"/home/dv/projects/github-deploy-inator"`
    <br><br>
  
- `command`: [required] The command to run when the webhook is received.
  - type: `string`
  - example: `"yarn deploy"`, `"git pull origin main"`
  <br><br>
- `secret`: The secret token set for your webhook. This makes sure that the webhook is from GitHub and is highly recommended to set.
  - type: `string`
  - example: `j4g34O3TK2JF4jrnjrkj34nt3i4`
    <br><br>
  
- `branch`: Execute the command only if the push was to this branch.
  - type: `string`
  - example: `"main"`, `"dev"`
    <br><br>

- `allowedPushers`: Execute the command only if this array contains the pusher's GitHub username
  - type: `string[]`
  - example: `["DeathVenom54", "webnoob"]`
    <br><br>

- `notifyDiscord`: If you want to receive a notification on Discord (via webhook)
  - type: `boolean`
    <br><br>

- `discord`: This contains information needed for sending Discord notifications
  - type: `Discord`

#### Discord

- `webhook`: [required] The url of the webhook where notifications should be sent
  - type: `string`
  - example: `"https://discord.com/api/webhooks/938275411766720533/s4nhfM-8XH1hMu9WYqSBUFaSD_erXSn6qqfdazzieCwtlINZho4teSvdlnEYgBM1E1IO"`
    <br><br>

- `notifyBeforeRun`: Whether a notification should be sent before running the command
  - type: `boolean`
    <br><br>

- `sendOutput`: Whether the notification should contain the output sent by the command
  - type: `boolean`

## 💻 Contributing to this project

If you find any bug in this project, have a suggestion or wish to contribute to it, feel free to [open an issue](https://github.com/DeathVenom54/github-deploy-inator/issues/new).

