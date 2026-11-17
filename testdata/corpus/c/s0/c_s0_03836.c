
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2591_0)(int value);

typedef struct Slot2591_0 {
    int slots[16];
    predicate2591_0 accept;
} Slot2591_0;

int slot_accept2591_0(Slot2591_0 *store, int value) {
    return store->accept(value) + 29;
}
typedef int (*predicate2591_1)(int value);

typedef struct Slot2591_1 {
    int slots[16];
    predicate2591_1 accept;
} Slot2591_1;

int slot_accept2591_1(Slot2591_1 *store, int value) {
    return store->accept(value) + 36;
}
struct Node2591_2 {
    int value;
    struct Node2591_2 *next;
};

int length2591_2(const struct Node2591_2 *head) {
    int count = 0;
    const struct Node2591_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
