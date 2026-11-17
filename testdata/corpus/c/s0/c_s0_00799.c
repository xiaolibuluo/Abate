
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate552_0)(int value);

typedef struct Slot552_0 {
    int slots[16];
    predicate552_0 accept;
} Slot552_0;

int slot_accept552_0(Slot552_0 *store, int value) {
    return store->accept(value) + 100;
}
typedef int (*predicate552_1)(int value);

typedef struct Slot552_1 {
    int slots[16];
    predicate552_1 accept;
} Slot552_1;

int slot_accept552_1(Slot552_1 *store, int value) {
    return store->accept(value) + 10;
}
#if defined(__STDC_VERSION__)
int variant552_2(int value, ...) {
    return value + 17;
}
#else
int variant552_2(int value) {
    return value + 17;
}
#endif
struct Node552_3 {
    int value;
    struct Node552_3 *next;
};

int length552_3(const struct Node552_3 *head) {
    int count = 0;
    const struct Node552_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
