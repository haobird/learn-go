package Sample;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import java.util.*;

import static com.google.common.base.MoreObjects.toStringHelper;

@ThriftStruct("UserList")
public final class UserList
{
    public UserList() {
    }

    private List<User> userList;

    @ThriftField(value=1, name="userList", requiredness=Requiredness.REQUIRED)
    public List<User> getUserList() { return userList; }

    @ThriftField
    public void setUserList(final List<User> userList) { this.userList = userList; }

    private int page;

    @ThriftField(value=2, name="page", requiredness=Requiredness.REQUIRED)
    public int getPage() { return page; }

    @ThriftField
    public void setPage(final int page) { this.page = page; }

    private int limit;

    @ThriftField(value=3, name="limit", requiredness=Requiredness.REQUIRED)
    public int getLimit() { return limit; }

    @ThriftField
    public void setLimit(final int limit) { this.limit = limit; }

    @Override
    public String toString()
    {
        return toStringHelper(this)
            .add("userList", userList)
            .add("page", page)
            .add("limit", limit)
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

        UserList other = (UserList)o;

        return
            Objects.equals(userList, other.userList) &&
            Objects.equals(page, other.page) &&
            Objects.equals(limit, other.limit);
    }

    @Override
    public int hashCode() {
        return Arrays.deepHashCode(new Object[] {
            userList,
            page,
            limit
        });
    }
}
