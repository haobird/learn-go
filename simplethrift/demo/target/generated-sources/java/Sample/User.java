package Sample;

import com.facebook.swift.codec.*;
import com.facebook.swift.codec.ThriftField.Requiredness;
import com.facebook.swift.codec.ThriftField.Recursiveness;
import java.util.*;

import static com.google.common.base.MoreObjects.toStringHelper;

@ThriftStruct("User")
public final class User
{
    public User() {
    }

    private int id;

    @ThriftField(value=1, name="id", requiredness=Requiredness.REQUIRED)
    public int getId() { return id; }

    @ThriftField
    public void setId(final int id) { this.id = id; }

    private String name;

    @ThriftField(value=2, name="name", requiredness=Requiredness.REQUIRED)
    public String getName() { return name; }

    @ThriftField
    public void setName(final String name) { this.name = name; }

    private String avatar;

    @ThriftField(value=3, name="avatar", requiredness=Requiredness.REQUIRED)
    public String getAvatar() { return avatar; }

    @ThriftField
    public void setAvatar(final String avatar) { this.avatar = avatar; }

    private String address;

    @ThriftField(value=4, name="address", requiredness=Requiredness.REQUIRED)
    public String getAddress() { return address; }

    @ThriftField
    public void setAddress(final String address) { this.address = address; }

    private String mobile;

    @ThriftField(value=5, name="mobile", requiredness=Requiredness.REQUIRED)
    public String getMobile() { return mobile; }

    @ThriftField
    public void setMobile(final String mobile) { this.mobile = mobile; }

    @Override
    public String toString()
    {
        return toStringHelper(this)
            .add("id", id)
            .add("name", name)
            .add("avatar", avatar)
            .add("address", address)
            .add("mobile", mobile)
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

        User other = (User)o;

        return
            Objects.equals(id, other.id) &&
            Objects.equals(name, other.name) &&
            Objects.equals(avatar, other.avatar) &&
            Objects.equals(address, other.address) &&
            Objects.equals(mobile, other.mobile);
    }

    @Override
    public int hashCode() {
        return Arrays.deepHashCode(new Object[] {
            id,
            name,
            avatar,
            address,
            mobile
        });
    }
}
