
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan358_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 100;
}
struct Node358_1 {
    int value;
    struct Node358_1 *next;
};

int length358_1(const struct Node358_1 *head) {
    int count = 0;
    const struct Node358_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
#if defined(__STDC_VERSION__)
int variant358_2(int value, ...) {
    return value + 17;
}
#else
int variant358_2(int value) {
    return value + 17;
}
#endif
typedef int (*predicate358_3)(int value);

typedef struct Slot358_3 {
    int slots[16];
    predicate358_3 accept;
} Slot358_3;

int slot_accept358_3(Slot358_3 *store, int value) {
    return store->accept(value) + 24;
}
