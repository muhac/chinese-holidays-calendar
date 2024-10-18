# To make it simple, copy the parser script
# to the current directory, and I only need
# to modify the titles of the output files.
mkdir docs
mkdir -p ../../docs/global/ca_on

cp -r ../../parser.cabal .
cp -r ../../parser .

sed -i \
    -e 's/Chinese Holidays/Ontario Statutory Holidays/' \
    -e 's/titleStatus status/"加拿大安大略省公共假日"/' \
    -e 's/name ++ show status/unwords name_en/' \
    -e 's/show status/name_cn ++ " "/' \
    -e 's/0xa95511fe/0xca4ada04;name_cn:name_en=splitOn "_" name/' \
    -e 's/import/import Data.List.Split (splitOn);import/' \
    ./parser/Main/Output.hs

cabal build
cabal run

mv docs/rest.ics ../../docs/global/ca_on/main.ics
mv docs/work.ics ../../docs/global/ca_on/rest.ics
