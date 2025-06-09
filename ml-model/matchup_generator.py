from db import get_all_pokemon
import random
import pandas as pd

def flatten(df):
    size = 6
    flattened_data = {}
    
    for i in range(size):
        if i < len(df):
            p = df.iloc[i]
            for col in p.index:
                flattened_data[f"p{i+1}_{col}"] = p[col]
        else:
            for col in df.columns:
                flattened_data[f"p{i+1}_{col}"] = None
    
    return pd.DataFrame([flattened_data])

def generate_matchups(num_samples):
    df = get_all_pokemon()
    train_df = []
    
    for i in range(num_samples):
        p1_pokemon = []
        p2_pokemon = []
        pokemon_no = random.randint(1, 6)
        for _ in range(pokemon_no):  
            a = random.randint(0, len(df) - 1)
            p1 = df.iloc[a]
            p1_pokemon.append(p1)
            
        pokemon_no = random.randint(1, 6)
        for _ in range(pokemon_no):
            a = random.randint(0, len(df) - 1)
            p2 = df.iloc[a]
            p2_pokemon.append(p2)
            
        p1_df = pd.DataFrame(p1_pokemon)
        p2_df = pd.DataFrame(p2_pokemon)
        
        p1_flattened = flatten(p1_df)
        p2_flattened = flatten(p2_df)

        merged_df = pd.concat([p1_flattened, p2_flattened], axis=1)
        train_df.append(merged_df)
        
    return pd.concat(train_df, ignore_index=True)