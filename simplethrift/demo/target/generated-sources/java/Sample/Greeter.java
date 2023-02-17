package Sample;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.service.*;
import com.google.common.util.concurrent.ListenableFuture;
import java.io.*;
import java.util.*;

@ThriftService("Greeter")
public interface Greeter
{
    @ThriftService("Greeter")
    public interface Async
    {
        @ThriftMethod(value = "SayHello",
                      exception = {
                          @ThriftException(type=BizException.class, id=1)
                      })
        ListenableFuture<Response> sayHello(
            @ThriftField(value=1, name="user", requiredness=Requiredness.REQUIRED) final User user
        );

        @ThriftMethod(value = "GetUser",
                      exception = {
                          @ThriftException(type=BizException.class, id=1)
                      })
        ListenableFuture<Response> getUser(
            @ThriftField(value=1, name="uid", requiredness=Requiredness.REQUIRED) final int uid
        );
    }
    @ThriftMethod(value = "SayHello",
                  exception = {
                      @ThriftException(type=BizException.class, id=1)
                  })
    Response sayHello(
        @ThriftField(value=1, name="user", requiredness=Requiredness.REQUIRED) final User user
    ) throws BizException, org.apache.thrift.TException;

    @ThriftMethod(value = "GetUser",
                  exception = {
                      @ThriftException(type=BizException.class, id=1)
                  })
    Response getUser(
        @ThriftField(value=1, name="uid", requiredness=Requiredness.REQUIRED) final int uid
    ) throws BizException, org.apache.thrift.TException;
}