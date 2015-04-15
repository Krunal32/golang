#Labs

Common repository for all lab exercises. This will eventually be populated with
the remaining lab exercises.

#IMPORTANT: Do not fork this repository. 

A writable repository has been created for you that you can push to. See the instructions below for details. Also, if you have trouble with github related issues, please go on working based on the pdf file until you can get help.

##Workflow and Handin Instructions:

For every lab, you are required to write a brief report in the form of answers
to questions in a text file (named ANSWERS.md) formatted as markdown. The labs
also requires that you submit your code files. Submission of these deliverables
shall be done through your personal git repository via a commit and push.

Below we describe the workflow that we expect you to follow:

0. A repository titled `username-labs` has been created for you, where `username` should be replaced with your own github username. You should not fork this repository!
1. Before continuing, download the pdf file and follow the instructions under Section 4. In particular you will need to do a few steps from here: http://golang.org/doc/code.html to set up your local go environment.
2. The recommended way is to set up github, is to use your ssh public-key. See the instructions in Section 3 of Lab 1 on how to set up your own ssh keys for your Unix account. We will use the same key for github in the next step.
4. When you have a ssh key pair set up on your Unix account, you should copy/paste the content of your `.ssh/id_rsa.pub` file (the public key part of the ssh key pair) into the ssh key entry box on github. You can use the command `cat $HOME/.ssh/id_rsa.pub` to display the content of the file. On github website, go to your user account page, click *Edit Profile*, in the left-hand menu click *SSH keys*, click *Add SSH Key*.
5. Next step is necessary to use the ssh key together with the `go get` command (it is really a work-around for a problem with the `go get` command.) (An alternative is to work directly with `git clone`, but we won't show that here.) You will need to add the following two lines to a file named `.gitconfig` in your `$HOME` directory; create the file if it does not already exist:

Add the following to `$HOME/.gitconfig` using an editor:

    [url "git@github.com:"]
        insteadOf = https://github.com/

This file may already contain other information, e.g. if you have set your name and email, in which case you should not overwrite these.

6. To get started with the go part of this lab, you can now use the `go get` command to clone the original `labs` repository. All students should now have *read access* to this repository. Here is how to do it:
  - On the command line enter:
  	`go get github.com/uis-dat320-fall2014/labs/lab2`.
    This will clone the original `labs` git repo (not your copy of it.) This is
    important because it means that you don't need to change the import path in
    the source files to use your own repository's path. That is, when
    you make a commit and push to submit your handin, you don't have to change
    this back to the original import path. For details, see the Import path
    issues of forked repository section below.
7. Next, run the following command: 
	`git remote add labs git@github.com:uis-dat320-fall2014/username-labs`
   where `username` should be replaced with your own github username.
8. This command adds your own `username-labs` repository as a remote repository. This means that once you've modified some files and committed the changes you can now run:
	`git push labs`
   to have them pushed up to your own `username-labs` repository on github.
9. If you make changes to your own `username-labs` repository using the github web interface, and want to pull those changes down to your own computer, you can run the command:
	`git pull labs master`
   In later labs, you will work in groups. This approach is also the way that you can download (pull) your group's code changes from github, assuming that another group member has previously pushed it out to github.
10. As time goes by we (the teaching staff) will be publishing updates to the original `labs` repo, e.g. new lab exercises. To see these updates, you will need to run the following command:
	`git pull origin master`

###Ready to submit?
1. When you are finished with all the exercises in the current lab, and wish to submit, then first make sure you commit your changes and write only the following: `username labX submission` in the first line of the commit message, where you replace `username` with your github username and `X` with the lab number. If there are any issue you want us to pay attention to, please add those comments after an empty line in the commit message.
2. Once you committed your changes locally, you must pushed your changes to github using: `git push labs`.
3. Now, we will be able to review your answers.

###Import path issues of forked/differently named repository

(you can ignore this section, and just follow the instructions above, but those interested in the technical understanding of what's going on here will want to read it.)

If your repository with Go code that has import paths that reference specific user accounts on github.com or other similar online services, you have two options to handle this depending on your intent. The approach that we take above is one approach, which allows you to continue to use the original source files' import paths, such as:

	import "github.com/uis-dat320-fall2014/labs/lab2/config"

This is useful if your intent is fix a bug or otherwise contribute code to the original project's code, e.g. through a pull request.

On the other hand, if your intent is to start your own forked version of a project, you will want to fix all the import paths in the project to something like:

	import "github.com/username/labs/lab2/config"

This can be fixed with easily accessible tools. However, since we expect you to submit your code back to us for review, you must follow the approach described above.

