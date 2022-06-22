package linkedList;

class Node {
    int data;
    Node next;

    Node(int data, Node next) {
        this.data = data;
        this.next = next;
    }
}

class MLinkedList {
    Node head;

    MLinkedList() {
        this.head = null;
    }

    public void printList() {
        if (this.head == null) {
            System.out.println("Linked list is empty");
            return;
        }

        Node itr = this.head;
        String llstring = "";
        while (itr != null) {
            llstring += Integer.toString(itr.data) + " ---> ";
            itr = itr.next;
        }

        System.out.println(llstring);
    }

    public int getLength() {
        int count = 0;
        Node itr = this.head;
        while (itr != null) {
            count += 1;
            itr = itr.next;
        }

        return count;
    }

    public void insertAtBeginning(int data) {
        Node temp = new Node(data, this.head);
        this.head = temp;
    }

    public void insertAtEnd(int data) {
        if (this.head == null) {
            this.head = new Node(data, null);
            return;
        }

        Node itr = this.head;
        while (itr.next != null) {
            itr = itr.next;
        }

        itr.next = new Node(data, null);
    }

    public void insertAt(int index, int data) throws IndexOutOfBoundsException {
        if (index < 0 || index > this.getLength()) {
            throw new IndexOutOfBoundsException("Invalid index");
        }

        if (index == 0) {
            Node temp = new Node(data, this.head);
            this.head = temp;
            return;
        }

        int count = 0;
        Node itr = this.head;
        while (itr != null) {
            if (count == index - 1) {
                Node temp = new Node(data, itr.next);
                itr.next = temp;
                break;
            }
            count += 1;
            itr = itr.next;
        }
    }

    public void removeAt(int index) throws IndexOutOfBoundsException {
        if (index < 0 || index >= this.getLength()) {
            throw new IndexOutOfBoundsException("Invalid index");
        }

        if (index == 0) {
            this.head = this.head.next;
            return;
        }

        int count = 0;
        Node itr = this.head;
        while (itr != null) {
            if (count == index - 1) {
                itr.next = itr.next.next;
                break;
            }
            count += 1;
            itr = itr.next;
        }
    }

    public void insertList(int[] dataList) {
        this.head = null;
        for (int data : dataList) {
            this.insertAtEnd(data);
        }
    }

}

class MyLinkedList {

    public static void main(String[] args) {
        MLinkedList ll = new MLinkedList();
        ll.insertAtBeginning(81);
        ll.insertAtBeginning(100);
        ll.insertAtBeginning(101);
        System.out.println(ll.getLength());
        ll.insertAtEnd(77);
        try {
            ll.insertAt(2, 99);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
        ll.printList();
        try {
            ll.removeAt(2);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
        ll.printList();
        int[] arr = { 1, 2, 3, 4, 5 };

        ll.insertList(arr);
        ll.printList();

    }
}