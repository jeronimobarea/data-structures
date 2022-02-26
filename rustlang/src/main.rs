type Stack<T> = Vec<T>;

trait StackMethods<T> {
    fn pop(&mut self) -> Result<&T, String>;
    fn push(&mut self, item: T);
    fn empty(&self) -> bool;
    fn size(&self) -> i8;
    fn peek(&self) -> Result<&T, String>;
}

impl<T> StackMethods<T> for Stack<T> {
    fn pop(&mut self) -> Result<&T, String> {
        if self.empty() {
            Err(String::from("cannot remove an item stack is empty"))
        } else {
            let element: &T = match self.peek() {
                Ok(peek) => peek,
                Err(e) => return Err(e)
            };

            //let index: usize = (self.size() - 1) as usize;
            //self.remove(index);
            Ok(element)
        }
    }

    fn push(&mut self, _item: T) {
        let mut tmp: Vec<T> = vec![_item];
        self.append(&mut tmp);
    }

    fn empty(&self) -> bool {
        return self.size() == 0;
    }

    fn size(&self) -> i8 {
        return self.len() as i8;
    }

    fn peek(&self) -> Result<&T, String> {
        if self.empty() {
            Err(String::from("cannot retrieve any value stack is empty"))
        } else {
            let index: usize = (self.size() - 1) as usize;
            Ok(self.get(index).unwrap())
        }
    }
}

fn main() {
    let mut stack: Stack<i8> = vec![1, 2, 3];
    println!("{:?}", stack);

    stack.push(4);
    println!("{:?}", stack);

    stack.pop();
    println!("{:?}", stack);

    stack.pop();
    println!("{:?}", stack);
}
