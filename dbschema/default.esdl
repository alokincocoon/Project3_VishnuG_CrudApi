module default {
    scalar type studentNumber extending sequence;
    type Student {
        required sid: studentNumber;
        studentId := 'AL' ++ <str><int64>__source__.sid;
        required name: str;
        required email: str {
            constraint exclusive;
        };
        phones: array<json>;
        index on (.studentId);
    }
}
