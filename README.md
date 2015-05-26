#Go File System

> Operating System Class
> > Group 4 (Ju An, Zhang Yangkun, Zhang Tianyi, Xu Fangzhou)

## Test Cases
### Basic tests:
1. Insert 10 pair, read it back (see `insert_test()`)2. Restart backup, on successful restart  3. Delete 2 pair – without error returne (see `delete_test()`)4. Update 2 pair, read back the results (see `update_test()`)5. Restart primary, on successful restart (using `dump_test()`)
### Stress tests:
We use 10 processes to insert at the same time, and dump data to check if everything's right. Further more, we shut down and restart primary or backup several times during the process, and check the same thing. 
### Latency tests:
We contiunuously inserts 2000 valuses and record their latency, and do the same test on get operation. During the process, there should not be incorrect result, and output the latency.
## Notes:
1. The latency test actually does not restrict on latency, just to see how long it will delay.
2. We added 3 seconds delay every time we restart a server.
3. Our test code used a package in python called `requests`. Pleased use `sudo pip install requests` if not installed.
4. Our test program is a Python code located in `\bin\test`.