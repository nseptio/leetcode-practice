import pandas as pd
import numpy as np

def second_highest_salary(employee: pd.DataFrame) -> pd.DataFrame:
    employee["rank_salary"] = employee["salary"].rank(method="dense", ascending=False)
    second_highest = employee[employee["rank_salary"] == 2]
    if second_highest.empty:
        return pd.DataFrame({"SecondHighestSalary": [np.nan]})
    return second_highest.head(1)[["salary"]].rename(columns={"salary": "SecondHighestSalary"})