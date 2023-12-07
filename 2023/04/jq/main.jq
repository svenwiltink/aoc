split("\n") | 
    map(
        select(.!="") | 
        split("|") as $raw | 
        {
            count: 1,
            winning: [$raw[0] | split(":")[1] | split(" ")[] | select(. != "") | tonumber], 
            ours: [$raw[1] | split(" ")[] | select(. != "") | tonumber]    
        } | 
        .ours - (.ours - .winning) | 
        length -1 | 
        select(.>=0) | 
        exp2
    ) | add
