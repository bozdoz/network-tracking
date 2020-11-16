# network-tracking

Runs a cron job to update network status to google spreadsheet; tracking (significant) downtime.

### Getting Started

1. Create `credentials.json` in root directory:
  - Credentials.json must be of type "service_account"
  - See: https://console.cloud.google.com/iam-admin/serviceaccounts/details/

2. Create `.env` in root with SPREADSHEET_ID=abcdef
  - get spreadsheet id from spreadsheet URL: 
    - https://docs.google.com/spreadsheets/d/{spreadsheetID}/edit
  
3. Develop in VSCode Remote Containers (ms-vscode-remote.remote-containers)


### Building & Running

1. SPREADSHEET_ID is defined (however you want)
2. `docker build -t network-cron --build-arg SPREADSHEET_ID .`
3. `docker run --rm -d --name network-cron network-cron`

or:

1. SPREADSHEET_ID is defined in .env
2. `docker-compose up -d`