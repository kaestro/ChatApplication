# Start from the official Node.js LTS image
FROM node:lts

# Set the working directory
WORKDIR /chatApplication

# Copy the current directory contents into the container
COPY . .

# Install any needed packages specified in package.json
# RUN npm install

# Make port 80 available to the world outside this container
EXPOSE 80

# Run the app when the container launches
CMD ["node", "app.js"]
