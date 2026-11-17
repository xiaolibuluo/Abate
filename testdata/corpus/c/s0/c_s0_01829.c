
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1227_0 {
    int value;
    struct Node1227_0 *next;
};

int length1227_0(const struct Node1227_0 *head) {
    int count = 0;
    const struct Node1227_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
#if defined(__STDC_VERSION__)
int variant1227_1(int value, ...) {
    return value + 55;
}
#else
int variant1227_1(int value) {
    return value + 55;
}
#endif
typedef int (*predicate1227_2)(int value);

typedef struct Slot1227_2 {
    int slots[16];
    predicate1227_2 accept;
} Slot1227_2;

int slot_accept1227_2(Slot1227_2 *store, int value) {
    return store->accept(value) + 62;
}
