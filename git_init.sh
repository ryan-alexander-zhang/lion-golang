
git init

git remote add origin https://github.com/ryan-alexander-zhang/lion-golang.git

git fetch origin

git pull origin main --allow-unrelated-histories

git add .

git commit -m "feat: initial commit"

git push origin main

# git rm --cached file
# git rm -r --cached dir
# git reset --soft HEAD~1
