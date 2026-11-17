
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate741_0)(int value);

typedef struct Slot741_0 {
    int slots[16];
    predicate741_0 accept;
} Slot741_0;

int slot_accept741_0(Slot741_0 *store, int value) {
    return store->accept(value) + 35;
}
#if defined(__STDC_VERSION__)
int variant741_1(int value, ...) {
    return value + 42;
}
#else
int variant741_1(int value) {
    return value + 42;
}
#endif
struct Node741_2 {
    int value;
    struct Node741_2 *next;
};

int length741_2(const struct Node741_2 *head) {
    int count = 0;
    const struct Node741_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
