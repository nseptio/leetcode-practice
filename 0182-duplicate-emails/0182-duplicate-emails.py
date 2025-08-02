import pandas as pd

def duplicate_emails(person: pd.DataFrame) -> pd.DataFrame:

    person = person[person.duplicated(subset='email', keep='first')]
    df_unique_email = person.drop_duplicates(subset='email')
    return df_unique_email[['email']]