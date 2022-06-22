function Node(data = null, next = null) {
    this.data = data;
    this.next = next;
}

function LinkedList() {
    this.head = null;

    this.print = function () {
        if (this.head == null) {
            console.log("Linked list is empty");
            return;
        }

        let itr = this.head;
        let llstring = "";
        while (itr) {
            llstring += itr.data + " ---> "
            itr = itr.next;
        }

        console.log(llstring);
    }
    this.getLength = function () {
        let count = 0;
        let itr = this.head;
        while (itr) {
            count += 1;
            itr = itr.next;
        }
        return count;
    }
    this.insertAtBeginning = function (data) {
        let temp = new Node(data, this.head);
        this.head = temp;
    }
    this.insertAtEnd = function (data) {
        if (this.head == null) {
            this.head = new Node(data, null);
            return;
        }
        let itr = this.head;
        while (itr.next) {
            itr = itr.next;
        }
        itr.next = new Node(data, null);
    }
    this.insertAt = function (index, data) {
        if (index < 0 || index > this.getLength()) {
            console.error("Invalid index");
            return;
        }

        if (index == 0) {
            this.insertAtBeginning(data);
            return;
        }

        let count = 0;
        let itr = this.head;


        while (itr) {
            if (count == index - 1) {
                let temp = new Node(data, itr.next);
                itr.next = temp;
                break;
            }
            itr = itr.next;
            count += 1;
        }
    }

    this.removeAt = function (index) {
        if (index < 0 || index >= this.getLength()) {
            console.error("Invalid index");
            return;
        }

        if (index == 0) {
            this.head = this.head.next;
            return;
        }

        let itr = this.head;
        let count = 0;
        while (itr) {
            if (count == index - 1) {
                itr.next = itr.next.next;
                break;
            }

            itr = itr.next;
            count += 1;
        }
    }

    this.insertValues = function (dataList) {
        this.head = null;
        dataList.forEach(element => {
            this.insertAtEnd(element);
        });
    }
}



let ll = new LinkedList();
ll.insertAtBeginning(85);
ll.insertAtBeginning(32);
console.log(ll.getLength());
ll.print();
ll.insertAtEnd(100);
console.log(ll.getLength());
ll.print();

ll.insertAt(1, 21);
ll.print();
ll.removeAt(1);
ll.print();

ll.insertValuesAtEnd(["banana", "mango", "grapes", "orange"])
ll.print();