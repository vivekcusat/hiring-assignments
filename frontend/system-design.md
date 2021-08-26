# Frontend system design interview

We would like to have a discussion around a hypothetical frontend application and the system design that supports it. We're interested in hearing how you think about composition - a fundamental in any react frontend - and what your process is when designing solutions to both technical challenges (such as state and data) and product challenges (such as the user experience).

We will not be writing any code, but we encourage the use of notes or high level psuedo code to help demonstrate your thought process and any significant or complex parts of the system.  It would be beneficial to share your screen during the meeting, so that we can see those notes.

We will ask questions throughout the session, but here are some core things to consider as a starting point:

- What key technologies, libraries or boilerplates are used to build the core of the application?
- How do those technology choices encourage a maintainable and scalable application?
- How is the application composed? What key pages, routes, hooks and components exist to address the requirements?
- How is state managed within the application? Consider both data fetched from the server and data held within the application. For example, your data fetching strategy might affect your choices around state.
- How can the system design and code support a good user experience?
- What performance considerations might need to be addressed in such a system?

## Application

The hypothetical application is:

An inbox for documents (invoices and receipts), much like an email client. We would like to:

- Show a list of all documents, highlighting new documents as they arrive
- Allow categorisation of documents
- View each document individually
    - Show the document itself (an image)
    - Allow associating the document to an existing accounting entry, removing the document from the inbox when it is attached. Entries are identified by an accounting year and an entry number

The user should be able to search and sort the documents in their inbox. They should also be able to delete and categorise multiple documents at the same time.
