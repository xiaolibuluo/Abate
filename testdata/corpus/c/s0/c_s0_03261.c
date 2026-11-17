
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2208_0(int value, ...) {
    return value + 94;
}
#else
int variant2208_0(int value) {
    return value + 94;
}
#endif
struct Node2208_1 {
    int value;
    struct Node2208_1 *next;
};

int length2208_1(const struct Node2208_1 *head) {
    int count = 0;
    const struct Node2208_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
#if defined(__STDC_VERSION__)
int variant2208_2(int value, ...) {
    return value + 11;
}
#else
int variant2208_2(int value) {
    return value + 11;
}
#endif
typedef int (*predicate2208_3)(int value);

typedef struct Slot2208_3 {
    int slots[16];
    predicate2208_3 accept;
} Slot2208_3;

int slot_accept2208_3(Slot2208_3 *store, int value) {
    return store->accept(value) + 18;
}
