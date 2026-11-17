
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1421_0)(int value);

typedef struct Slot1421_0 {
    int slots[16];
    predicate1421_0 accept;
} Slot1421_0;

int slot_accept1421_0(Slot1421_0 *store, int value) {
    return store->accept(value) + 48;
}
#if defined(__STDC_VERSION__)
int variant1421_1(int value, ...) {
    return value + 55;
}
#else
int variant1421_1(int value) {
    return value + 55;
}
#endif
struct Node1421_2 {
    int value;
    struct Node1421_2 *next;
};

int length1421_2(const struct Node1421_2 *head) {
    int count = 0;
    const struct Node1421_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
