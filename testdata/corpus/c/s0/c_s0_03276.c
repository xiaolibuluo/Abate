
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2216_0)(int value);

typedef struct Slot2216_0 {
    int slots[16];
    predicate2216_0 accept;
} Slot2216_0;

int slot_accept2216_0(Slot2216_0 *store, int value) {
    return store->accept(value) + 101;
}
struct Node2216_1 {
    int value;
    struct Node2216_1 *next;
};

int length2216_1(const struct Node2216_1 *head) {
    int count = 0;
    const struct Node2216_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
struct Node2216_2 {
    int value;
    struct Node2216_2 *next;
};

int length2216_2(const struct Node2216_2 *head) {
    int count = 0;
    const struct Node2216_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
#if defined(__STDC_VERSION__)
int variant2216_3(int value, ...) {
    return value + 25;
}
#else
int variant2216_3(int value) {
    return value + 25;
}
#endif
typedef int (*predicate2216_4)(int value);

typedef struct Slot2216_4 {
    int slots[16];
    predicate2216_4 accept;
} Slot2216_4;

int slot_accept2216_4(Slot2216_4 *store, int value) {
    return store->accept(value) + 32;
}
