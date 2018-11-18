#! /usr/bin/env bash

stonetypes="darkstone stone lightstone"
earthtypes="darkearth earth lightearth clay mud"
dusttypes="darksand redsand lightsand dirt dust brick"

for t in $stonetypes; do
    for u in $stonetypes; do
        for v in $stonetypes; do
            echo "rect" $(( ( RANDOM % 3 )  + 1 )) $(( ( RANDOM % 3 )  + 1 )) $(( ( RANDOM % 5 )  + 3 )) $(( ( RANDOM % 5 )  + 3 )) "; $t $u $v;" | sed 's| ; |; |g' > assets/skel/rubble_misc_stone_"$t"_"$u"_"$v".txt
        done
    done
done

for t in $earthtypes; do
    for u in $earthtypes; do
        for v in $earthtypes; do
            echo "rect" $(( ( RANDOM % 3 )  + 1 )) $(( ( RANDOM % 3 )  + 1 )) $(( ( RANDOM % 5 )  + 3 )) $(( ( RANDOM % 5 )  + 3 )) "; $t $u $v;" | sed 's| ; |; |g' > assets/skel/rubble_misc_earth_"$t"_"$u"_"$v".txt
        done
    done
done

for t in $dusttypes; do
    for u in $dusttypes; do
        for v in $dusttypes; do
            echo "rect" $(( ( RANDOM % 3 )  + 1 )) $(( ( RANDOM % 3 )  + 1 )) $(( ( RANDOM % 5 )  + 3 )) $(( ( RANDOM % 5 )  + 3 )) "; $t $u $v;" | sed 's| ; |; |g' > assets/skel/rubble_misc_dust_"$t"_"$u"_"$v".txt
        done
    done
done
