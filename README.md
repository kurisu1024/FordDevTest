## Ask

Write a command line utility in Go which takes as arguments a list of directories. The program should output the sizes of each of the individual directories passed as well as a cumulative total. If a --recursive flag is provided, output the sizes of each of the individual directories passed and sub-directories recursively as well as a cumulative total. If a "--human" flag is passed, format the sizes to be human friendly by outputting the size in the most appropriate unit of bytes. For example, 304K for 304,000 bytes and 300M for 300000000 bytes.

### Things to iprove upon
- I didn't have time to create a proper test
suite for the main package, but did some manual testing
to see the cli working.  So first thing I would do is create a test
suite for the main package. 

- Next we could improve the performance of this cli
by adding some concurrency with goroutines and channel. 
- Expanding the test scenarios in the dir package 
would be nice to have as well.  


### Sample Output

```azure
C:\Users\OH Amity\GolandProjects\FordDevTest> go run main\main.go -human "\Users\OH Amity\GolandProjects"
path: \Users\OH Amity\GolandProjects size: 50.31 KiB
PS C:\Users\OH Amity\GolandProjects\FordDevTest> go run main\main.go -recursive "\Users\OH Amity\GolandProjects"
path: \Users\OH Amity\GolandProjects size: 51515
path: \Users\OH Amity\GolandProjects\FordDevTest size: 48366
path: \Users\OH Amity\GolandProjects\FordDevTest size: 48366
path: \Users\OH Amity\GolandProjects\FordDevTest\.git size: 35206
path: \Users\OH Amity\GolandProjects\FordDevTest\.git\COMMIT_EDITMSG size: 11
path: \Users\OH Amity\GolandProjects\FordDevTest\.git\COMMIT_EDITMSG size: 11
PS C:\Users\OH Amity\GolandProjects\FordDevTest> go run main\main.go -human -recursive "\Users\OH Amity\GolandProjects"
path: \Users\OH Amity\GolandProjects size: 50.31 KiB
path: \Users\OH Amity\GolandProjects\FordDevTest size: 47.23 KiB
path: \Users\OH Amity\GolandProjects\FordDevTest size: 47.23 KiB
path: \Users\OH Amity\GolandProjects\FordDevTest\.git size: 34.38 KiB
path: \Users\OH Amity\GolandProjects\FordDevTest\.git size: 34.38 KiB
path: \Users\OH Amity\GolandProjects\FordDevTest\.git\COMMIT_EDITMSG size: 11 bytes
path: \Users\OH Amity\GolandProjects\FordDevTest\.git\COMMIT_EDITMSG size: 11 bytes
PS C:\Users\OH Amity\GolandProjects\FordDevTest> go run main\main.go "\Users\OH Amity\GolandProjects"
path: \Users\OH Amity\GolandProjects size: 51515

```