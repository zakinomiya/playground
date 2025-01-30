const std = @import("std");

const DB = struct {
    inner: std.StringHashMap([]const u8),

    pub fn deinit(self: *DB) void {
        self.inner.deinit();
    }

    pub fn set(self: *DB, key: []const u8, value: []const u8) !void {
        try self.inner.put(key, value);
    }

    pub fn get(self: *DB, key: []const u8) ?[]const u8 {
        return self.inner.get(key);
    }

    pub fn del(self: *DB, key: []const u8) bool {
        return self.inner.remove(key);
    }
};

pub fn initDB(allocator: std.mem.Allocator) DB {
    return DB{ .inner = std.StringHashMap([]const u8).init(allocator) };
}

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();
    const stdin = std.io.getStdIn().reader();

    var allocator = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    std.debug.print("ZigDB\n", .{});
    var db = initDB(allocator.allocator());

    const MAX_LINE = 1024;
    while (true) {
        try stdout.print("> ", .{});

        var line_buffer: [MAX_LINE]u8 = undefined;
        const line = try stdin.readUntilDelimiterOrEof(&line_buffer, '\n') orelse {
            break;
        };
        var tokenizer = std.mem.tokenizeAny(u8, line, " ");
        const cmd = tokenizer.next() orelse {
            try stdout.print("invalid input", .{});
            continue;
        };

        switch (checkCommand(cmd)) {
            Action.Set => {
                const key = tokenizer.next() orelse {
                    try stdout.print("set <key> <value>", .{});
                    continue;
                };
                const value = tokenizer.next() orelse {
                    try stdout.print("set <key> <value>", .{});
                    continue;
                };
                db.set(key, value) catch |err| {
                    try stdout.print("error: {}\n", .{err});
                    continue;
                };
                std.debug.print("value: {s} is set for key: {s} \n", .{ value, key });
            },
            Action.Get => {
                const key = tokenizer.next() orelse {
                    try stdout.print("get <key>", .{});
                    continue;
                };
                const value = db.get(key) orelse {
                    try stdout.print("key not found\n", .{});
                    continue;
                };
                std.debug.print("value: {s} is found for key: {s} \n", .{ value, key });
            },
            Action.Del => {},
            Action.Exit => {},
            Action.Unknown => {
                try stdout.print("Unknown command\n", .{});
            },
        }
    }
}

const Action = enum {
    Unknown,
    Set,
    Get,
    Del,
    Exit,
};

pub fn checkCommand(command: []const u8) Action {
    if (std.mem.startsWith(u8, command, "set")) {
        return Action.Set;
    }
    if (std.mem.startsWith(u8, command, "get")) {
        return Action.Get;
    }
    if (std.mem.startsWith(u8, command, "del")) {
        return Action.Del;
    }
    if (std.mem.startsWith(u8, command, "exit")) {
        return Action.Exit;
    }
    return Action.Unknown;
}

test "test basic functions" {
    var db = initDB(std.testing.allocator);
    defer db.deinit();
    try db.set("hello", "world");

    const value = db.get("hello") orelse {
        return error.Error;
    };
    try std.testing.expectEqualStrings("world", value);

    if (!db.del("hello")) {
        return error.Error;
    }
    try std.testing.expect(db.get("hello"));
}
