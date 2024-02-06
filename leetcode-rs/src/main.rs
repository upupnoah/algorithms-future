use std::cell::RefCell;
use std::rc::Rc;

// 定义 User 结构体
#[derive(Debug)]
struct User {
    name: String,
    friends: RefCell<Vec<Rc<User>>>, // 使用 RefCell 来内部可变性
}

impl User {
    fn new(name: &str) -> Rc<Self> {
        Rc::new(User {
            name: name.to_string(),
            friends: RefCell::new(vec![]),
        })
    }

    // 添加朋友
    fn add_friend(&self, friend: Rc<User>) {
        self.friends.borrow_mut().push(friend);
    }

    // 打印用户及其朋友
    fn print_friends(&self) {
        let friend_names: Vec<_> = self
            .friends
            .borrow()
            .iter()
            .map(|f| f.name.clone())
            .collect();
        println!("{}'s friends: {:?}", self.name, friend_names);
    }
}

fn main() {
    let alice = User::new("Alice");
    let bob = User::new("Bob");
    let carol = User::new("Carol");

    // Alice 和 Bob 成为朋友
    alice.add_friend(bob.clone());
    // Alice 和 Carol 也成为朋友
    alice.add_friend(carol.clone());

    // 打印 Alice 的朋友
    alice.print_friends(); // 输出: Alice's friends: ["Bob", "Carol"]
}
