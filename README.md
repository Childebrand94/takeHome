# NumInfo Analyser


## Introduction

NumInfo Analyser is a web application designed to provide detailed insights into text messages by analyzing 
phone numbers and message content. The core functionality revolves around extracting and displaying information 
such as the country, region, operator, and the most specific prefix associated with a given phone number. 
Additionally, the application highlights and makes URLs within the message body clickable, enhancing user 
interaction with the message content.


## Libraries Used

- xcurl :
    I opted for xurls due to its robust regex capabilities, 
    making it highly effective at identifying URLs within text. This 
    precision ensures that all types of URLs, regardless of their format 
    or complexity, are accurately detected and processed.

- DomPurify :
    For sanitizing the HTML, DOMPurify was the clear choice. 
    It's a trusted library known for its thorough cleansing of HTML content, 
    effectively stripping out any potentially malicious code. This safeguard 
    is crucial for maintaining the security of the application, especially 
    when rendering user-generated links, protecting against XSS (Cross-Site Scripting) 
    attacks without compromising the integrity of the content.

Requirements Checklist

 - [x] API Endpoint for Phone Number Analysis
 - [x] Efficient Prefix Matching
 - [x] URL Detection and Conversion
 - [x] Data Sanitization
 - [x] Error Handling and Validation
 - [x] Unit and Integration Tests
 - [x] Documentation
    

## Prerequisites

- Go
- TypeScript

## Installation

### First, clone the repository to your local machine:

```bash
git clone https://github.com/Childebrand94/takeHome.git
cd takeHome
```
### Setting Up the Backend (Go)

Navigate to the backend directory and build the Go application:
```
cd backend
go cmd/build
```
To start the backend server, run:
```
./main
```
## Setting Up the Frontend (React TypeScript)

Navigate to the frontend directory from the root of the project:
```
npm install
```
Start the frontend application:
```
npm start
```

     
