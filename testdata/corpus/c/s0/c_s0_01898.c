
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1276_0 {
    int value;
    struct Node1276_0 *next;
};

int length1276_0(const struct Node1276_0 *head) {
    int count = 0;
    const struct Node1276_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
typedef int (*predicate1276_1)(int value);

typedef struct Slot1276_1 {
    int slots[16];
    predicate1276_1 accept;
} Slot1276_1;

int slot_accept1276_1(Slot1276_1 *store, int value) {
    return store->accept(value) + 13;
}
#if defined(__STDC_VERSION__)
int variant1276_2(int value, ...) {
    return value + 20;
}
#else
int variant1276_2(int value) {
    return value + 20;
}
#endif
