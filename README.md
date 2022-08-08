# ruby-rce

This dangerous project is intended to run on a private network behind a firewall!
Only authorized users should be able to cause interactions with this program!

This program provides a simple API for validation and execution of ruby scripts.
It is designed to facilitate script execution for Ruby on Rails apps where employees
create ruby scripts for customer support purposes.

The main web application should use it to do the following (via protected routes):

1. analyze and validate scripts (internally, Ripper will check the script)
2. execute a script (now known as a "job") with options such as dry_run
3. list running and completed jobs
4. obtain exit statuses and logs for jobs

It remains the main application's responsibility to:

1. Organize scripts
2. Utilize validation API to guard script creation to ensure validity
3. Utilize validation API parameter analysis response during script creation
4. Organize targets (whichever machine is running this program is considered a target)
5. Utilize the targets' Jobs API to provide show status and results

All data is stored on the filesystem in the configured locations.

Please configure:

1. Path to ruby executable (for using Ripper)
2. Path to for file storage 

---

Future:

- https://github.com/distribworks/dkron Distributed, fault tolerant job scheduling system