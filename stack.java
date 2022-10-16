class Stk{
    int stk[];
    int capacity;
    int top;
    Stk(int size){
        stk= new int[size];
        capacity=size;
        top=-1;
    }

    void push(int elem){
        if(isFull())
        System.out.println("Stack is full");
        else
        stk[++top]=elem;
    }

    int pop(){
        if(isEmpty()){
        System.out.println("Stack is empty");
            System.exit(0);
        }
         
        return stk[top--];
        
    }

    boolean isFull(){
        return top==capacity-1;
    }

    boolean isEmpty(){
        return top==-1;
    }

    void display(){
        if(isEmpty()){
            System.out.println("Empty");
        }
        else{
        for(int i=top;i>=0;i--)
        System.out.println(stk[i]);
        }
    }

    int peek(){
        return stk[top];
    }
}

public class stack {
    public static void main(String[] args){
        Stk s=new Stk(4);
        s.push(1);
        s.push(2);
        s.push(3);
        s.push(4);
        s.push(5);
        s.display();
        System.out.println(s.pop());
        

    }
}
