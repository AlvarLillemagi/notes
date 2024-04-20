Notes Tool

The Notes Tool is a command-line application designed to help users manage short single-line notes. Users can create collections of notes, open and view them, add new notes, or remove existing notes.


Usage:

To create or manage a collection, users run the tool with the desired collection name as an argument:

$ ./notestool [TAG]

For example, to manage notes in a collection named "coding_ideas", users would run:

$ ./notestool coding_ideas

Users can also use the help tag to display a brief help message that explains how to use the application:

$ ./notestool help


Once inside the tool, users can perform the following operations:

1.Display notes: View existing notes in the collection.

    *Example:
    
    $ ./notestool coding_ideas
    
    Welcome to the notes tool for coding_ideas

    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    1

    Notes:
    001 - note one
    002 - note two

2.Add a note: Append a new note to the collection.

    *Example:

    $ ./notestool coding_ideas
    
    Welcome to the notes tool for coding_ideas

    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    2

    Enter the note text:
    note three

    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    1

    Notes:
    001 - note one
    002 - note two
    003 - note three

3.Delete a note: Remove a specific note from the collection.

    *Example:

    $ ./notestool coding_ideas
    
    Welcome to the notes tool for coding_ideas

    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    3

    Enter the number of note to remove or 0 to cancel:
    3

    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    1

    Notes:
    001 - note one
    002 - note two

4.Exit the program: Terminate the tool.

    *Example:

    $ ./notestool coding_ideas
    
    Welcome to the notes tool for coding_ideas

    Select operation:
    1. Show notes.
    2. Add a note.
    3. Delete a note.
    4. Exit.

    4

Data Storage:

For each collection, a separate plain text file with the same name as the collection is created. Notes are stored in separate rows within this file. The data persists between runs of the tool.
