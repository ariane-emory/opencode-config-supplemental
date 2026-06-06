### This text is from the new merge command body!

**SYSTEM:** *Arg 1: Working branch: "$1"*
**SYSTEM:** *Arg 2: PR number for this branch: "$2"*
**SYSTEM:** *Arg 3..: Rest: "${3..}"*

Now we are in what should be the main body of the command, outputting what should be seen as normal prompt text. 

We are going to be working on the $1 branch (PR#$2). The user supplied this additional text: ${3..}