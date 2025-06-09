from db import get_all_pokemon
import random
import pandas as pd

def flatten(team_df, prefix):
    flattened = {}
    for i, row in enumerate(team_df.itertuples(), start=1):
        flattened.update({
            f"{prefix}{i}_id": row.id,
            f"{prefix}{i}_name": row.name,
            f"{prefix}{i}_hp": row.hp,
            f"{prefix}{i}_attack": row.attack,
            f"{prefix}{i}_defense": row.defense,
            f"{prefix}{i}_sp_attack": row.sp_attack,
            f"{prefix}{i}_sp_defense": row.sp_defense,
            f"{prefix}{i}_speed": row.speed,
            f"{prefix}{i}_primary_type": row.primary_type,
            f"{prefix}{i}_secondary_type": row.secondary_type
        })
    return pd.DataFrame([flattened])


def get_damage(attacker_stats, defender_stats):
    attack = attacker_stats[0]
    special_attack = attacker_stats[1]
    defense = defender_stats[0]
    special_defense = defender_stats[1]
    
    if attack >= special_attack:
        attack_stat = attack
        defense_stat = defense
    else:
        attack_stat = special_attack
        defense_stat = special_defense
        
    critical_hit = 2 if random.randint(0, 1) >= 0.0625 else 1
    random = random.randint(217, 255) // 255
    damage = (((2 * critical_hit / 5 + 2) * attack_stat / defense_stat) / 50) + 2 * random
    
    return damage

def extract_team(flat_df, prefix):
    flat_row = flat_df.iloc[0]
    team = []
    for i in range(1, 7):
        id_col = f"{prefix}{i}_id"
        if id_col in flat_row and pd.notna(flat_row[id_col]):
            stats = {
                "attack": flat_row[f"{prefix}{i}_attack"],
                "sp_attack": flat_row[f"{prefix}{i}_sp_attack"],
                "defense": flat_row[f"{prefix}{i}_defense"],
                "sp_defense": flat_row[f"{prefix}{i}_sp_defense"],
                "hp": flat_row[f"{prefix}{i}_hp"],
                "name": flat_row[f"{prefix}{i}_name"]
            }
            team.append(stats)
    return team

def get_damage(attacker, defender):
    attack_stat = max(attacker["attack"], attacker["sp_attack"])
    defense_stat = max(defender["defense"], defender["sp_defense"])

    critical = 2 if random.random() < 0.1 else 1
    base_damage = (((2 * critical / 5 + 2) * attack_stat / defense_stat) / 50) + 2
    damage = base_damage * random.randint(217, 255) / 255
    return damage

def simulate_battle(team1, team2):
    p1 = random.choice(team1).copy()
    p2 = random.choice(team2).copy()

    hp1 = p1["hp"]
    hp2 = p2["hp"]

    while hp1 > 0 and hp2 > 0:
        hp2 -= get_damage(p1, p2)
        if hp2 <= 0:
            return 1
        hp1 -= get_damage(p2, p1)
        if hp1 <= 0:
            return 2

    return 1 if hp1 > 0 else 2

def generate_matchups(num_samples):
    df = get_all_pokemon()
    training_rows = []
    
    for _ in range(num_samples):
        p1_team = [df.sample(1).iloc[0] for _ in range(6)]
        p2_team = [df.sample(1).iloc[0] for _ in range(6)]

        p1_flat = flatten(pd.DataFrame(p1_team), prefix="p")
        p2_flat = flatten(pd.DataFrame(p2_team), prefix="opp")
        
        full_row = pd.concat([p1_flat, p2_flat], axis=1)
        winner = simulate_battle(extract_team(p1_flat, "p"), extract_team(p2_flat, "opp"))
        full_row = pd.concat([p1_flat, p2_flat], axis=1)
        full_row["winner"] = pd.Series([winner])
        training_rows.append(full_row)
        
    return pd.concat(training_rows, ignore_index=True)