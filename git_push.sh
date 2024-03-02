#!/bin/sh
# ref: https://help.github.com/articles/adding-an-existing-project-to-github-using-the-command-line/
#
# Usage example: /bin/sh ./git_push.sh wing328 openapi-petstore-perl "minor update" "gitlab.com"

jjkirkpatrick=$1
spacetraders-sdk=$2
release_note=$3
git_host=$4

if [ "$git_host" = "" ]; then
    git_host="github.com"
    echo "[INFO] No command line input provided. Set \$git_host to $git_host"
fi

if [ "$jjkirkpatrick" = "" ]; then
    jjkirkpatrick="jjkirkpatrick"
    echo "[INFO] No command line input provided. Set \$jjkirkpatrick to $jjkirkpatrick"
fi

if [ "$spacetraders-sdk" = "" ]; then
    spacetraders-sdk="spacetraders-sdk"
    echo "[INFO] No command line input provided. Set \$spacetraders-sdk to $spacetraders-sdk"
fi

if [ "$release_note" = "" ]; then
    release_note="Minor update"
    echo "[INFO] No command line input provided. Set \$release_note to $release_note"
fi

# Initialize the local directory as a Git repository
git init

# Adds the files in the local repository and stages them for commit.
git add .

# Commits the tracked changes and prepares them to be pushed to a remote repository.
git commit -m "$release_note"

# Sets the new remote
git_remote=$(git remote)
if [ "$git_remote" = "" ]; then # git remote not defined
    
    if [ "$GIT_TOKEN" = "" ]; then
        echo "[INFO] \$GIT_TOKEN (environment variable) is not set. Using the git credential in your environment."
        git remote add origin https://${git_host}/${jjkirkpatrick}/${spacetraders-sdk}.git
    else
        git remote add origin https://${jjkirkpatrick}:"${GIT_TOKEN}"@${git_host}/${jjkirkpatrick}/${spacetraders-sdk}.git
    fi
    
fi

git pull origin master

# Pushes (Forces) the changes in the local repository up to the remote repository
echo "Git pushing to https://${git_host}/${jjkirkpatrick}/${spacetraders-sdk}.git"
git push origin master 2>&1 | grep -v 'To https'
