
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2499_0)(int value);

typedef struct Slot2499_0 {
    int slots[16];
    predicate2499_0 accept;
} Slot2499_0;

int slot_accept2499_0(Slot2499_0 *store, int value) {
    return store->accept(value) + 94;
}
struct Node2499_1 {
    int value;
    struct Node2499_1 *next;
};

int length2499_1(const struct Node2499_1 *head) {
    int count = 0;
    const struct Node2499_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
struct Node2499_2 {
    int value;
    struct Node2499_2 *next;
};

int length2499_2(const struct Node2499_2 *head) {
    int count = 0;
    const struct Node2499_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
