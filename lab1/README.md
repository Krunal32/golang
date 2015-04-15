##Lab 1 instructions

Please consult the general workflow and handin instructions below. For lab 1, you should edit the file, ANSWERS.md as described within the file. Other source files may also be submitted as part of the pull request.

##Workflow and Handin Instructions:

For every lab, you are required to write a brief report, including answers to questions. The labs also requires that you submit your code. Submission of these deliverables shall be done through the git repository for each lab exercise. Below we describe the workflow that we expect you to follow. We provide a fairly detailed description below, so it may seem a bit complicated at first, but you’ll learn it quickly. And we will be around to help out, so if there are any questions about the workflow and submission process, please don’t hesitate to ask.

1. Click on the *Fork* button (upper right-hand corner of the page).
2. Follow the on-screen instructions.
3. Once you have a fork of the repository connected to your own user, you need to make a clone of the repository on your local machine so that you can edit the specific files for the lab. There are several ways to clone the repo. Here we only explain one way:
  - Open a terminal window on your local machine.
  - Navigate to a folder where you want to keep your opsys lab exercises, e.g. `cd opsys/labs/`.
  - Execute the git clone command as follows: `git clone https://github.com/username/labX.git` where you replace `username` with your github username and `X` with the lab number.
  - (*Instead of typing in the above URL in the terminal window, you can copy it to your clipboard from the github web page by clicking the copy-to-clipboard button shown in the right-hand margin menu. When you have the URL in the clipboard, simply type `git clone` followed by a space and paste the URL from the clipboard using CTRL-v and press enter.*)
4. You should now have a copy of the lab files that we have prepared. You can edit these however you want unless instructed otherwise, and you can create new files that you can add to your local fork of the git repo. Typically, you will want to add your own source code files that we are expecting you to deliver.
5. While you work on your submission, it is highly recommended that you add and commit your changes frequently to your local fork of the repo. These are the commands:
  - `git add ANSWERS.md mysrc.c`
  - `git commit`
6. You may also wish to push your changes to your github account as well, but this is not strictly necessary until you wish to submit. This is the command: `git push`.

###Ready to submit?
1. When you are finished with all the exercises in the lab, and wish to submit, then first make sure you have pushed your changes to github using: `git push`.
2. Next, go to original lab project (here) and click on the *Pull Requests* link in the right-hand margin menu, and on the page that comes up next, click on the green button titled *New Pull Request*.
3. The next page shows (at the top): *We’re showing branches in this repository, but you can also compare across forks*. Click the *compare across forks* link. In the dropdown menu that appears, select your own fork.
4. You will now see your own changes compared to any original files. On this page, select the green *Create pull request* button.
5. In the title field, write only the following: `username labX submission`, where you replace `username` with your github username and `X` with the lab number. If there are any issue you want us to pay attention to, please use the comment field. Otherwise, leave it blank.
6. Finally, click the green *Create pull request* button.
7. Now, we will be able to review your answers.

