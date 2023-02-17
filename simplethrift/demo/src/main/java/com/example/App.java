package com.example;

import org.apache.thrift.TException;
import org.apache.thrift.protocol.TBinaryProtocol;
import org.apache.thrift.protocol.TProtocol;
import org.apache.thrift.transport.TSocket;
import org.apache.thrift.transport.TTransport;
import Sample.*;

/**
 * Hello world!
 */
public final class App {
    private App() {
    }

    /**
     * Says hello to the world.
     * @param args The arguments of the program.
     */
    public static void main(String[] args) {
        System.out.println("Hello World!");

        try {
            TTransport transport;
            transport = new TSocket("localhost", 9092);
            transport.open();
            TProtocol protocol = new TBinaryProtocol(transport);

            User u = new User();
            u.setName("ddd");
            System.out.println(u.getName());


            // u.name = "";
            // u.avatar = "xxx";
            // u.address = "f";
            // u.mobile = "f";
            // Response resp = client.SayHello(u);
            // System.out.println(resp);

            transport.close();
        
        } catch (TException x) {
            // x.printStackTrace();
            System.out.println(x);
        }
    }
}
