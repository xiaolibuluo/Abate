
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1660_0)(int value);

typedef struct Slot1660_0 {
    int slots[16];
    predicate1660_0 accept;
} Slot1660_0;

int slot_accept1660_0(Slot1660_0 *store, int value) {
    return store->accept(value) + 51;
}
struct Node1660_1 {
    int value;
    struct Node1660_1 *next;
};

int length1660_1(const struct Node1660_1 *head) {
    int count = 0;
    const struct Node1660_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
