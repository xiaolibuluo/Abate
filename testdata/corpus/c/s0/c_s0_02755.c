
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1876_0)(int value);

typedef struct Slot1876_0 {
    int slots[16];
    predicate1876_0 accept;
} Slot1876_0;

int slot_accept1876_0(Slot1876_0 *store, int value) {
    return store->accept(value) + 46;
}
typedef int (*predicate1876_1)(int value);

typedef struct Slot1876_1 {
    int slots[16];
    predicate1876_1 accept;
} Slot1876_1;

int slot_accept1876_1(Slot1876_1 *store, int value) {
    return store->accept(value) + 53;
}
typedef int (*predicate1876_2)(int value);

typedef struct Slot1876_2 {
    int slots[16];
    predicate1876_2 accept;
} Slot1876_2;

int slot_accept1876_2(Slot1876_2 *store, int value) {
    return store->accept(value) + 60;
}
typedef int (*predicate1876_3)(int value);

typedef struct Slot1876_3 {
    int slots[16];
    predicate1876_3 accept;
} Slot1876_3;

int slot_accept1876_3(Slot1876_3 *store, int value) {
    return store->accept(value) + 67;
}
#if defined(__STDC_VERSION__)
int variant1876_4(int value, ...) {
    return value + 74;
}
#else
int variant1876_4(int value) {
    return value + 74;
}
#endif
struct Node1876_5 {
    int value;
    struct Node1876_5 *next;
};

int length1876_5(const struct Node1876_5 *head) {
    int count = 0;
    const struct Node1876_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
