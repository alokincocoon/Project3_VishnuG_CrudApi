CREATE MIGRATION m1uiffr2mdojgldwu7lhlyom72x3hl2vud7ogcjrugvf57vxjytsmq
    ONTO initial
{
  CREATE SCALAR TYPE default::studentNumber EXTENDING std::sequence;
  CREATE TYPE default::Student {
      CREATE REQUIRED PROPERTY sid: default::studentNumber;
      CREATE PROPERTY studentId := (('AL' ++ <std::str><std::int64>__source__.sid));
      CREATE INDEX ON (.studentId);
      CREATE PROPERTY email: std::str {
          CREATE CONSTRAINT std::exclusive;
      };
      CREATE REQUIRED PROPERTY name: std::str;
      CREATE PROPERTY phones: array<std::json>;
  };
};
