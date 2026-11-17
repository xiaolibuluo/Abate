
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2101_0)(int value);

typedef struct Slot2101_0 {
    int slots[16];
    predicate2101_0 accept;
} Slot2101_0;

int slot_accept2101_0(Slot2101_0 *store, int value) {
    return store->accept(value) + 61;
}
struct Node2101_1 {
    int value;
    struct Node2101_1 *next;
};

int length2101_1(const struct Node2101_1 *head) {
    int count = 0;
    const struct Node2101_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
