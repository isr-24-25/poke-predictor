from dotenv import load_dotenv
import os
import psycopg2
import pandas as pd

load_dotenv()
url = os.getenv("DB_URL")

def connect():
    try:
        conn = psycopg2.connect(url)
        return conn
    except psycopg2.Error as e:
        print(f"error connecting to database: {e}")
        
def get_all_pokemon():
    conn = connect()
    if not conn:
        raise ConnectionError("could not connect to database")
    
    query = """
        SELECT 
            b.id,
            b.name,
            b.hp,
            b.attack,
            b.defense,
            b.sp_attack,
            b.sp_defense,
            b.speed,
            array_agg(t.type) AS types
        FROM base_stats b
        JOIN pokemon_types_link pt ON b.id = pt.pokemon_id
        JOIN types t ON t.id = pt.type_id
        GROUP BY b.id, b.name, b.hp, b.attack, b.defense, b.sp_attack, b.sp_defense, b.speed
        ORDER BY b.id;
    """
    
    df = pd.read_sql(query, conn)
    df[['primary_type', 'secondary_type']] = df['types'].apply(lambda x: pd.Series(x + [None] * (2 - len(x))))
    df.drop(columns='types', inplace=True)
    
    return df