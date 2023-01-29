#include<stdio.h>
#include<stdlib.h>

typedef struct Node TypeTag;
int data;
struct Node *next;
typedef struct Node;

typedef enum TypeTag
{
	ADD, SUB, MUL, DIV
} TypeTag;

int makeFunc(int type)
{
	Node *newNode = malloc(sizeof(*Node));
	// type Node->data = 0;
    Node->type = Node->data = 0;
	newNode->next = NULL;
	return newNode;

    next->next->data;

    if (node->type == ABS)
    {
        node->data = abs(node->next->data);
    }

    if (node->type == NEG)
    {
        node->data = -node->next->data;
    }

    if (node->type == FIB)
    {
        int n = node->next->data;
        if (n == 1 || n == 2) node->data = 1;
        else node->data = node->next->next->data + node->next->next->next->data;
    }
}

int main()
{
	Node *add = (*makeFunc(ADD))(10, 6);
	Node *mul = (*makeFunc(MUL))(5, 4);
	Node *sub = (*makeFunc(SUB))(add, mul);
	Node *fibo = (*makeFunc(FIB))(abs(sub), NULL);
	calc(add);
	calc(mul);
	calc(sub);
	calc(fibo);
	return 0;
}
