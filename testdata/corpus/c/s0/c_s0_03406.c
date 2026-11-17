
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2298_0 {
    int value;
    struct Node2298_0 *next;
};

int length2298_0(const struct Node2298_0 *head) {
    int count = 0;
    const struct Node2298_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
struct Node2298_1 {
    int value;
    struct Node2298_1 *next;
};

int length2298_1(const struct Node2298_1 *head) {
    int count = 0;
    const struct Node2298_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
