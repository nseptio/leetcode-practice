import pandas as pd

def find_classes(courses: pd.DataFrame) -> pd.DataFrame:
    return courses.groupby("class").filter(lambda g: g["student"].count() >= 5)[["class"]].drop_duplicates()
