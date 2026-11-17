
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate737_0)(int value);

typedef struct Slot737_0 {
    int slots[16];
    predicate737_0 accept;
} Slot737_0;

int slot_accept737_0(Slot737_0 *store, int value) {
    return store->accept(value) + 80;
}
struct Node737_1 {
    int value;
    struct Node737_1 *next;
};

int length737_1(const struct Node737_1 *head) {
    int count = 0;
    const struct Node737_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
