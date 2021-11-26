select ai.animal_id, ai.animal_type, ai.`name`

from animal_ins ai
join animal_outs ao on ai.animal_id = ao.animal_id

where   ai.sex_upon_intake like 'Intact%'
and     (ao.sex_upon_outcome = 'Spayed Female' or ao.sex_upon_outcome = 'Neutered Male')