
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node199_0 {
    int value;
    struct Node199_0 *next;
};

int length199_0(const struct Node199_0 *head) {
    int count = 0;
    const struct Node199_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 70;
}
