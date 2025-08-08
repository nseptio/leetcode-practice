import pandas as pd

def employee_bonus(employee: pd.DataFrame, bonus: pd.DataFrame) -> pd.DataFrame:
    merge_df = employee.merge(bonus, how='left', on='empId')
    return merge_df[(merge_df['bonus'] < 1000) | (merge_df['bonus'].isnull())][['name', 'bonus']]