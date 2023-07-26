# Client-Server_System
a mimic to multiple clients, single server system using go Language

This is an example of Server Distribution, as the client talks "connect" to the master 
that doesn't have the data but returns to the client the place "on the internet" of these data
then the client goes and talks to every place that he got from the master and fetches the data from them "Slaves".

the next image illustrates the cycle of the system

![image](https://github.com/Marco-Emad/Client-Server_System/assets/56565607/77b6c622-6ac0-4b69-b60a-195c22a3ee5a)

Referenced from the Book: "Principles of Distributed Database Systems" By M. Tamer Özsu · Patrick Valduriez
