CREATE MIGRATION m1wd47jitr24oh5yhrivss7l7qnld4r5cdwulpyezeanzv2akotzbq
    ONTO m1uiffr2mdojgldwu7lhlyom72x3hl2vud7ogcjrugvf57vxjytsmq
{
  ALTER TYPE default::Student {
      ALTER PROPERTY email {
          SET REQUIRED USING (<std::str>{'NA'});
      };
  };
};
