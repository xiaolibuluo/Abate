
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node794_0 {
    int value;
    struct Node794_0 *next;
};

int length794_0(const struct Node794_0 *head) {
    int count = 0;
    const struct Node794_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
#if defined(__STDC_VERSION__)
int variant794_1(int value, ...) {
    return value + 52;
}
#else
int variant794_1(int value) {
    return value + 52;
}
#endif
typedef int (*predicate794_2)(int value);

typedef struct Slot794_2 {
    int slots[16];
    predicate794_2 accept;
} Slot794_2;

int slot_accept794_2(Slot794_2 *store, int value) {
    return store->accept(value) + 59;
}
#if defined(__STDC_VERSION__)
int variant794_3(int value, ...) {
    return value + 66;
}
#else
int variant794_3(int value) {
    return value + 66;
}
#endif
