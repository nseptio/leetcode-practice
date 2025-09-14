import pandas as pd

def game_analysis(activity: pd.DataFrame) -> pd.DataFrame:
    activity['row_num'] = activity.groupby('player_id')['event_date'].rank(method='first')
    result = activity[activity['row_num'] == 1][['player_id', 'event_date']]
    return result.rename(columns={"event_date": "first_login"})