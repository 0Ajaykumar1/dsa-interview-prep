class Node:
    def __init__(self, data=None, next=None, prev=None):
        self.data = data
        self.next = next
        self.prev = prev


class DoubleLinkedList:
    def __init__(self):
        self.head = None

    def print_forward(self):
        if self.head is None:
            print("Linked list is empty")
            return

        itr = self.head
        llstring = ""
        while itr:
            llstring += str(itr.data) + " ---> "
            itr = itr.next

        print(llstring)

    def print_backward(self):
        if self.head is None:
            print("Linked list is empty")
            return

        lastNode = self.head
        while lastNode.next:
            lastNode = lastNode.next

        itr = lastNode
        llstring = ""
        while itr:
            llstring += str(itr.data) + " ---> "
            itr = itr.prev

        print(llstring)

    def get_length(self):
        count = 0
        itr = self.head
        while itr:
            count += 1
            itr = itr.next
        return count

    def insert_at_beginning(self, data):
        if self.head is None:
            self.head = Node(data, None, None)
            return
        temp = Node(data, self.head, self.head.prev)
        self.head.prev = temp
        self.head = temp

    def insert_at_end(self, data):
        if self.head is None:
            self.head = Node(data, None, None)
            return

        itr = self.head
        while itr.next:
            itr = itr.next

        temp = Node(data, None, itr)
        itr.next = temp

    def insert_at(self, index, data):
        if index < 0 or index > self.get_length():
            raise Exception("Invalid index")

        if index == 0:
            temp = Node(data, self.head, self.head.prev)
            self.head.prev = temp
            self.head = temp
            return

        itr = self.head
        count = 0
        while itr:
            if count == index - 1:
                temp = Node(data, itr.next, itr)
                if temp.next:
                    itr.next.prev = temp
                itr.next = temp
                break

            itr = itr.next
            count += 1

    def remove_at(self, index):
        if index < 0 or index >= self.get_length():
            raise Exception("Invalid index")

        if index == 0:
            self.head.next.prev = self.head.prev
            self.head = self.head.next

        itr = self.head
        count = 0
        while itr:
            if count == index - 1:
                if itr.next.next:
                    itr.next.next.prev = itr
                itr.next = itr.next.next
                break
            itr = itr.next
            count += 1

    def insert_values(self, data_list):
        self.head = None

        for data in data_list:
            self.insert_at_end(data)

    def insert_after_value(self, data_after, data):
        if self.head is None:
            return

        if self.head.data == data_after:
            temp = Node(data, self.head.next, self.head)
            self.head.next.prev = temp
            self.head.next = temp
            return

        itr = self.head

        while itr:
            if(itr.data == data_after):
                temp = Node(data, itr.next, itr)
                itr.next.prev = temp
                itr.next = temp
                break
            itr = itr.next

    def remove_by_value(self, data):
        if self.head is None:
            return
        if self.head.data == data:
            self.head.next.prev = self.head.prev
            self.head = self.head.next
            return

        itr = self.head

        while itr.next:
            if itr.next.data == data:
                if itr.next.next:
                    itr.next.next.prev = itr
                itr.next = itr.next.next
                break
            itr = itr.next

        itr = itr.next


if __name__ == "__main__":
    ll = DoubleLinkedList()
    ll.insert_at_beginning(82)
    ll.insert_at_beginning(81)
    ll.insert_at_beginning(80)
    ll.insert_at_beginning(79)
    ll.insert_at_end(83)
    ll.remove_at(3)
    ll.print_forward()
    ll.print_backward()
    ll.insert_after_value(81, 82)
    ll.print_forward()
    ll.print_backward()
    ll.insert_after_value(81, 100)
    ll.print_forward()
    ll.print_backward()
    ll.remove_by_value(100)
    ll.print_forward()
    ll.print_backward()
