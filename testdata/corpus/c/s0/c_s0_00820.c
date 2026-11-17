
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node567_0 {
    int value;
    struct Node567_0 *next;
};

int length567_0(const struct Node567_0 *head) {
    int count = 0;
    const struct Node567_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
struct Node567_1 {
    int value;
    struct Node567_1 *next;
};

int length567_1(const struct Node567_1 *head) {
    int count = 0;
    const struct Node567_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
typedef int (*predicate567_2)(int value);

typedef struct Slot567_2 {
    int slots[16];
    predicate567_2 accept;
} Slot567_2;

int slot_accept567_2(Slot567_2 *store, int value) {
    return store->accept(value) + 18;
}
int scan567_3(const int *begin, const int *end) {
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
    return total + 25;
}
#if defined(__STDC_VERSION__)
int variant567_4(int value, ...) {
    return value + 32;
}
#else
int variant567_4(int value) {
    return value + 32;
}
#endif
struct Node567_5 {
    int value;
    struct Node567_5 *next;
};

int length567_5(const struct Node567_5 *head) {
    int count = 0;
    const struct Node567_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
