
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1476_0 {
    int value;
    struct Node1476_0 *next;
};

int length1476_0(const struct Node1476_0 *head) {
    int count = 0;
    const struct Node1476_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
struct Node1476_1 {
    int value;
    struct Node1476_1 *next;
};

int length1476_1(const struct Node1476_1 *head) {
    int count = 0;
    const struct Node1476_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
typedef int (*predicate1476_2)(int value);

typedef struct Slot1476_2 {
    int slots[16];
    predicate1476_2 accept;
} Slot1476_2;

int slot_accept1476_2(Slot1476_2 *store, int value) {
    return store->accept(value) + 98;
}
