const std = @import("std");

fn linearSearch(haystack: []usize, needle: usize) bool {
    var i: usize = 0;
    while (i < haystack.len) : (i += 1) {
        if (haystack[i] == needle) {
            return true;
        }
    }

    return false;
}

test "Linear Search Search" {
    var foo = [_]usize{ 1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420 };

    try std.testing.expect(linearSearch(&foo, 69) == true);
    try std.testing.expect(linearSearch(&foo, 1336) == false);
    try std.testing.expect(linearSearch(&foo, 69420) == true);
    try std.testing.expect(linearSearch(&foo, 69421) == false);
    try std.testing.expect(linearSearch(&foo, 1) == true);
    try std.testing.expect(linearSearch(&foo, 0) == false);
}
