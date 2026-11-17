
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1911_0)(int value);

typedef struct Slot1911_0 {
    int slots[16];
    predicate1911_0 accept;
} Slot1911_0;

int slot_accept1911_0(Slot1911_0 *store, int value) {
    return store->accept(value) + 16;
}
struct Node1911_1 {
    int value;
    struct Node1911_1 *next;
};

int length1911_1(const struct Node1911_1 *head) {
    int count = 0;
    const struct Node1911_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
struct Node1911_2 {
    int value;
    struct Node1911_2 *next;
};

int length1911_2(const struct Node1911_2 *head) {
    int count = 0;
    const struct Node1911_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
