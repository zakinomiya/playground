const std = @import("std");

fn solve(allocator: std.mem.Allocator, line: []u8, num: usize) !void {
    var m = std.ArrayList(u8).init(allocator);
    defer m.deinit();

    for (line) |c, i| {
        try m.insert(0, c);

        if (m.items.len == num) {
            var duplicate = false;
            for (m.items) |v, j| {
                var k: usize = j + 1;
                while (k < m.items.len) : (k += 1) {
                    duplicate = duplicate or m.items[k] == v;
                }
            }

            if (!duplicate) {
                std.debug.print("{d}\n", .{i + 1});
                return;
            }

            _ = m.pop();
        }
    }
}

pub fn main() !void {
    var arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    var allocator = arena.allocator();
    var buf = try allocator.alloc(u8, 5000);
    defer allocator.free(buf);

    var cwd = std.fs.cwd();
    const file = try cwd.readFile("in", buf);

    try solve(allocator, file, 4);
    try solve(allocator, file, 14);
}
