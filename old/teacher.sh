#!/usr/bin/env bash
INTERVIEW_ID=$(grep "SEE INTERVIEW #" streets/Buckingham_Place | cut -d '#' -f 2)
echo $INTERVIEW_ID
cat interviews/interview-$INTERVIEW_ID
echo $MAIN_SUSPECT
