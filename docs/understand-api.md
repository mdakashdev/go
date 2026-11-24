1. net/http
   এটি Go-এর standard library। HTTP Server এবং API তৈরি করতে ব্যবহার হয়।
2. http.HandleFunc()
   এখানে:
     - /health = API route
     - w = Response writer
     - r = Incoming HTTP request
3. http.ListenAndServe()
   এটি তোমার HTTP server চালু করে port 8080-এ।




# Api endpoint 

/health
/about 
/worker


Build Our First JSON API