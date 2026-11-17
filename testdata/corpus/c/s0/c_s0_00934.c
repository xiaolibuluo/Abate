
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate641_0)(int value);

typedef struct Slot641_0 {
    int slots[16];
    predicate641_0 accept;
} Slot641_0;

int slot_accept641_0(Slot641_0 *store, int value) {
    return store->accept(value) + 93;
}
struct Node641_1 {
    int value;
    struct Node641_1 *next;
};

int length641_1(const struct Node641_1 *head) {
    int count = 0;
    const struct Node641_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
