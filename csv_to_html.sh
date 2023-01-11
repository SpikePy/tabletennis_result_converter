#!/usr/bin/env sh

# 1. go to https://tttv.click-tt.de
# 2. login with mail address
# 3. switch to tab "Downloads"
# 4. download "Vereinsspielplan (csv)"
# 5. run script

###[ VARIABLE DEFINITION ]######################################################

csv_original=Vereinsspielplan_*.csv
csv_converted=/tmp/results_converted.csv


###[ CONVERT/FIX CHARACTER ENCODING ]###########################################

iconv -f ISO-8859-1 -t UTF-8 ${csv_original} --output="${csv_converted}"
dos2unix -q "${csv_converted}"


###[ convert csv to html ]######################################################

# UNPLAYED MATCHES
html_result=result_unplayed.html
grep --extended-regexp ';0;0$' "${csv_converted}" | awk --field-separator=';' '
    BEGIN {print "<table>\n  <tr>\n    <th>Termin</th>\n    <th>Staffel</th>\n    <th>Heim-Verein</th>\n    <th>Gast-Verein</th>\n  </tr>"}
    {
        if (NR>1) {
            print "  <tr>\n    <td>"$1"</td>\n    <td>"$8"</td>\n    <td>"$21"</td>\n    <td>"$27"</td>\n  </tr>"
        }
    }
    END {print "</table>"}' > "${html_result}"

# PLAYED MATCHES
html_result=result_played.html
grep --extended-regexp ';[1-9];[1-9]$' "${csv_converted}" | awk --field-separator=';' '
    BEGIN {print "<table>\n  <tr>\n    <th>Termin</th>\n    <th>Staffel</th>\n    <th>Heim-Verein</th>\n    <th>Gast-Verein</th>\n    <th>Ergebnis</th>\n  </tr>"}
    {
        if (NR>1) {
            print "  <tr>\n    <td>"$1"</td>\n    <td>"$8"</td>\n    <td>"$21"</td>\n    <td>"$27"</td>\n    <td>"$28":"$29"</td>\n  </tr>"
        }
    }
    END {print "</table>"}' > "${html_result}"

# cleanup
rm "${csv_converted}"
