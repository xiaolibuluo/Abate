
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1714_0 {
    int value;
    struct Node1714_0 *next;
};

int length1714_0(const struct Node1714_0 *head) {
    int count = 0;
    const struct Node1714_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
typedef int (*predicate1714_1)(int value);

typedef struct Slot1714_1 {
    int slots[16];
    predicate1714_1 accept;
} Slot1714_1;

int slot_accept1714_1(Slot1714_1 *store, int value) {
    return store->accept(value) + 81;
}
#if defined(__STDC_VERSION__)
int variant1714_2(int value, ...) {
    return value + 88;
}
#else
int variant1714_2(int value) {
    return value + 88;
}
#endif
