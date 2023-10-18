FROM node:18.16.1

WORKDIR /app
COPY package*.json ./

# Install production dependencies.
RUN npm ci

# Copy local code to the container image.
COPY . ./

RUN npm run build

CMD [ "npm", "run", "production" ]