package Sample;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import java.util.*;

import static com.google.common.base.MoreObjects.toStringHelper;

@ThriftStruct("Response")
public final class Response
{
    public Response() {
    }

    private int errCode;

    @ThriftField(value=1, name="errCode", requiredness=Requiredness.REQUIRED)
    public int getErrCode() { return errCode; }

    @ThriftField
    public void setErrCode(final int errCode) { this.errCode = errCode; }

    private String errMsg;

    @ThriftField(value=2, name="errMsg", requiredness=Requiredness.REQUIRED)
    public String getErrMsg() { return errMsg; }

    @ThriftField
    public void setErrMsg(final String errMsg) { this.errMsg = errMsg; }

    private Map<String, String> data;

    @ThriftField(value=3, name="data", requiredness=Requiredness.REQUIRED)
    public Map<String, String> getData() { return data; }

    @ThriftField
    public void setData(final Map<String, String> data) { this.data = data; }

    @Override
    public String toString()
    {
        return toStringHelper(this)
            .add("errCode", errCode)
            .add("errMsg", errMsg)
            .add("data", data)
            .toString();
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (o == null || getClass() != o.getClass()) {
            return false;
        }

        Response other = (Response)o;

        return
            Objects.equals(errCode, other.errCode) &&
            Objects.equals(errMsg, other.errMsg) &&
            Objects.equals(data, other.data);
    }

    @Override
    public int hashCode() {
        return Arrays.deepHashCode(new Object[] {
            errCode,
            errMsg,
            data
        });
    }
}
