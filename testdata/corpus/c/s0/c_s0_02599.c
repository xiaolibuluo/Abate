
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1762_0)(int value);

typedef struct Slot1762_0 {
    int slots[16];
    predicate1762_0 accept;
} Slot1762_0;

int slot_accept1762_0(Slot1762_0 *store, int value) {
    return store->accept(value) + 19;
}
struct Node1762_1 {
    int value;
    struct Node1762_1 *next;
};

int length1762_1(const struct Node1762_1 *head) {
    int count = 0;
    const struct Node1762_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
