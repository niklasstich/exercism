use std::fmt;
use std::fmt::{Display, Formatter};

const MINUTES_PER_HOUR: i32 = 60;
//this is probably redundant
const HOURS_PER_DAY: i32 = 24;
const MINUTES_PER_DAY: i32 = MINUTES_PER_HOUR * HOURS_PER_DAY;

//required traits for the test
#[derive(Debug, Eq, PartialEq)]
pub struct Clock {
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Clock {
            //use rem_euclid to wrap around to 0 at 24:00
            minutes: ((hours * MINUTES_PER_HOUR) + minutes).rem_euclid(MINUTES_PER_DAY),
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock {
            //use rem_euclid to wrap around to 0 at 24:00
            minutes: (self.minutes + minutes).rem_euclid(MINUTES_PER_DAY),
        }
    }

    //returns hour, minute in range 0-23 and 0-59 respectively
    pub fn time_in_range(&self) -> (i32, i32) {
        //use div_euclid instead of regular division to cover negative numbers
        let hour = self.minutes.div_euclid(MINUTES_PER_HOUR);
        let minute = self.minutes.rem_euclid(MINUTES_PER_HOUR);
        return (hour, minute);
    }
}

impl Display for Clock {
    fn fmt(&self, f: &mut Formatter<'_>) -> fmt::Result {
        //{:0>2} means fill in 0 to the left until length of formatted substring is at least 2 characters
        //{:a<4} would fill in a to the right until at least 4 characters
        let time = self.time_in_range();
        write!(f, "{:0>2}:{:0>2}", time.0, time.1)
    }
}
