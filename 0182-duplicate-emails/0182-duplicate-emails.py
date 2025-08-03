import pandas as pd

def duplicate_emails(person: pd.DataFrame) -> pd.DataFrame:

    df_duplicated_email = person[person.duplicated(subset='email', keep='first')]
    df_unique_email = df_duplicated_email.drop_duplicates(subset='email')
    return df_unique_email[['email']]