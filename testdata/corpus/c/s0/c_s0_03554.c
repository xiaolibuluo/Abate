
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2398_0 {
    int value;
    struct Node2398_0 *next;
};

int length2398_0(const struct Node2398_0 *head) {
    int count = 0;
    const struct Node2398_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
typedef int (*predicate2398_1)(int value);

typedef struct Slot2398_1 {
    int slots[16];
    predicate2398_1 accept;
} Slot2398_1;

int slot_accept2398_1(Slot2398_1 *store, int value) {
    return store->accept(value) + 49;
}
struct Node2398_2 {
    int value;
    struct Node2398_2 *next;
};

int length2398_2(const struct Node2398_2 *head) {
    int count = 0;
    const struct Node2398_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
