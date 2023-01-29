# task-1-Backend

## Pseudocode for task

``` Include the stdio.h and stdlib.h header files.

typedef struct Node TypeTag type; int data; struct Node *next; typedef struct Node TypeTag

typedef enum TypeTag { ADD, SUB, MUL, DIV } TypeTag;

Node* makeFunc (int type) { \s Node newNode = (Node)malloc(sizeof(Node)); \s new Node->type equals type Node->data = 0; newNode->next = NULL; return newNode; }

next->next->data; } else if (node->type == ABS) { node->data = abs(node->next->data); } else if (node->type == NEG) { node->data = -node->next->data; } else if (node->type == FIB) { int n = node->next->data; if (n == 1 || n == 2) node->data = 1; else node->data = node->next->next->data + node->next->next->next->data; } }

int main () { Node *add = (*makeFunc(ADD))(10, 6); Node *mul = (*makeFunc(MUL))(5, 4); Node *sub = (*makeFunc(SUB))(add, mul); Node *fibo = (*makeFunc(FIB))(abs(sub), NULL); calc(add); calc(mul); calc(sub); calc(fibo); return 0; }
```

## Step-by-step explanation

- The above code creates a structure for a node in a linked list along with an enumerated type for the kind of operation to be done on the node. The makeFunc method builds a new node with the type and data supplied, while the calc function calculates the outcome of the operation indicated by the type tag. The main function then constructs a linked list of nodes representing the actions to be executed and runs the calc function on each node to compute the final result.

- The primary function begins by creating a new node for the addition operation with the data fields set to 10 and 6. The identical process is then repeated for the multiplication operation, with the data fields set to 5 and 4. Next, it generates a new node for the subtraction operation, populating its data fields with the results of the addition and multiplication operations. It then generates a new node for the fibonacci operation with the result of the subtraction operation as the data field.

- The calc function calculates the outcomes of operations defined by type tags. For addition and multiplication, it just calculates and saves the result in the data field of the node. It calculates the result of the subtraction operation and saves it in the node's data field. Using the results of the addition and multiplication operations, the nth fibonacci number is calculated for the fibonacci operation and stored in the data field of the node.

- The main function displays the results of the actions after all nodes have been constructed and the calc function has been run on each of them. For addition, multiplication, and subtraction, the result stored in the node's data field is printed. The fibonacci operation outputs the nth fibonacci number contained in the data field of the node.
This program use a linked list to depict an execution sequence. The makeFunc method produces a new node for each operation, while the calc function calculates the outcomes of the operations and puts them in the data field of the newly created node. The outcomes of operations are then printed by the main function.

- This application calculates the Fibonacci numbers using dynamic programming. The makeFunc method produces a new node for each operation, while the calc function calculates the outcomes of the operations and puts them in the data field of the newly created node. The outcomes of operations are then printed by the main function.