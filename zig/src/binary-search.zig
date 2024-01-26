const std = @import("std");

fn binarySearch(haystack: []i32, needle: i32) bool {
    var low: usize = 0;
    var high: usize = haystack.len;

    while (low < high) {
        const mid: usize = @divFloor(high - low, 2) + low;
        const value = haystack[mid];

        if (value == needle) {
            return true;
        }

        if (value < needle) {
            low = mid + 1;
        } else {
            high = mid;
        }
    }

    return false;
}

test "Binary Search" {
    var foo = [_]i32{ 1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420 };
    
    try std.testing.expect(binarySearch(&foo, 69) == true);
    try std.testing.expect(binarySearch(&foo, 1336) == false);
    try std.testing.expect(binarySearch(&foo, 69420) == true);
    try std.testing.expect(binarySearch(&foo, 69421) == false);
    try std.testing.expect(binarySearch(&foo, 1) == true);
    try std.testing.expect(binarySearch(&foo, 0) == false);
}
