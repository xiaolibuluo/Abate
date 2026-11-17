
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate160_0)(int value);

typedef struct Slot160_0 {
    int slots[16];
    predicate160_0 accept;
} Slot160_0;

int slot_accept160_0(Slot160_0 *store, int value) {
    return store->accept(value) + 48;
}
#if defined(__STDC_VERSION__)
int variant160_1(int value, ...) {
    return value + 55;
}
#else
int variant160_1(int value) {
    return value + 55;
}
#endif
struct Node160_2 {
    int value;
    struct Node160_2 *next;
};

int length160_2(const struct Node160_2 *head) {
    int count = 0;
    const struct Node160_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
