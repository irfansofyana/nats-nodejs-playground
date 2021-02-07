# Playground Area for learning `nats`

## How to Run

- Download & Install `nats` server
- Run `nats` server
- Run the application

If you are using docker:

- You can download the `nats` server by using:  `docker pull nats`
- Run the `nats` server by using command: `docker run --rm -it -p 4222:4222 nats`
- Check the `nats` server is running by `telnet localhost 4222`. You should see the output is something like this:

    ```
    Trying 127.0.0.1...
    Connected to localhost.
    Escape character is '^]'.
    INFO {"server_id":"NAE4EF5PDNNCEHW2P476MV6QK4TGJT2SXND5P5ZFCFE4DD3KNOQMPYCM","server_name":"NAE4EF5PDNNCEHW2P476MV6QK4TGJT2SXND5P5ZFCFE4DD3KNOQMPYCM","version":"2.1.9","proto":1,"git_commit":"7c76626","go":"go1.14.10","host":"0.0.0.0","port":4222,"max_payload":1048576,"client_id":1,"client_ip":"172.17.0.1"}
    ```

## Notes

- Official documentation of `nats`: https://docs.nats.io/
- npm `nats` page: https://www.npmjs.com/package/nats 
