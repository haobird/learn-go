package Sample;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import java.util.*;

@ThriftStruct("BizException")
public final class BizException extends RuntimeException
{
    private static final long serialVersionUID = 1L;

    public BizException() {
    }

    private int code;

    @ThriftField(value=1, name="code", requiredness=Requiredness.REQUIRED)
    public int getCode() { return code; }

    @ThriftField
    public void setCode(final int code) { this.code = code; }

    private String msg;

    @ThriftField(value=2, name="msg", requiredness=Requiredness.REQUIRED)
    public String getMsg() { return msg; }

    @ThriftField
    public void setMsg(final String msg) { this.msg = msg; }
}
