# timeTracker
This is a Presence App ...


## Configuration
use `CONFIGOR_ENV` environment variable to specify which config to use. available configs are: 
1. `config.development.yml` for testing and development environment 
2. `confif.production.yml` for the production environment
please refer to [template file](config/config.template.yml) to see the available field for configuration

## configure facebox
1. create an account at https://machinebox.io/account and get your MB_KEY
2. `export MB_KEY="YOUR_MB_KEY" `
3. `docker run -d -p 8080:8080 -e "MB_KEY=$MB_KEY" machinebox/facebox`
4. install OpenCV : 
`brew install opencv3 `