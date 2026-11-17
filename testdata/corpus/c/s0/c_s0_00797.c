
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node550_0 {
    int value;
    struct Node550_0 *next;
};

int length550_0(const struct Node550_0 *head) {
    int count = 0;
    const struct Node550_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
typedef int (*predicate550_1)(int value);

typedef struct Slot550_1 {
    int slots[16];
    predicate550_1 accept;
} Slot550_1;

int slot_accept550_1(Slot550_1 *store, int value) {
    return store->accept(value) + 81;
}
typedef int (*predicate550_2)(int value);

typedef struct Slot550_2 {
    int slots[16];
    predicate550_2 accept;
} Slot550_2;

int slot_accept550_2(Slot550_2 *store, int value) {
    return store->accept(value) + 88;
}
