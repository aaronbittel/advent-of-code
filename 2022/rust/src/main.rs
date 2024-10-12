use chrono::{self, naive, Datelike, Local};

fn main() {
    let today = Local::now().date_naive();
    let christmas =
        naive::NaiveDate::from_ymd_opt(today.year(), 12, 24).expect("expect a valid year");
    let days_till_christmas = (christmas - today).num_days();
    println!("{} till Christmas!", days_till_christmas);
}
