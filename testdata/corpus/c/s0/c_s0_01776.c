
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1190_0)(int value);

typedef struct Slot1190_0 {
    int slots[16];
    predicate1190_0 accept;
} Slot1190_0;

int slot_accept1190_0(Slot1190_0 *store, int value) {
    return store->accept(value) + 52;
}
struct Node1190_1 {
    int value;
    struct Node1190_1 *next;
};

int length1190_1(const struct Node1190_1 *head) {
    int count = 0;
    const struct Node1190_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 59;
}
#if defined(__STDC_VERSION__)
int variant1190_2(int value, ...) {
    return value + 66;
}
#else
int variant1190_2(int value) {
    return value + 66;
}
#endif
typedef int (*predicate1190_3)(int value);

typedef struct Slot1190_3 {
    int slots[16];
    predicate1190_3 accept;
} Slot1190_3;

int slot_accept1190_3(Slot1190_3 *store, int value) {
    return store->accept(value) + 73;
}
struct Node1190_4 {
    int value;
    struct Node1190_4 *next;
};

int length1190_4(const struct Node1190_4 *head) {
    int count = 0;
    const struct Node1190_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
