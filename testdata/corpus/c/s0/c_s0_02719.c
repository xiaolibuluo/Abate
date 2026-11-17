
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1848_0)(int value);

typedef struct Slot1848_0 {
    int slots[16];
    predicate1848_0 accept;
} Slot1848_0;

int slot_accept1848_0(Slot1848_0 *store, int value) {
    return store->accept(value) + 70;
}
#define MASK1848_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1848_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1848_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1848_2 {
    int value;
    struct Node1848_2 *next;
};

int length1848_2(const struct Node1848_2 *head) {
    int count = 0;
    const struct Node1848_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
struct Node1848_3 {
    int value;
    struct Node1848_3 *next;
};

int length1848_3(const struct Node1848_3 *head) {
    int count = 0;
    const struct Node1848_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
struct Node1848_4 {
    int value;
    struct Node1848_4 *next;
};

int length1848_4(const struct Node1848_4 *head) {
    int count = 0;
    const struct Node1848_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
int scan1848_5(const int *begin, const int *end) {
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
    return total + 8;
}
#if defined(__STDC_VERSION__)
int variant1848_6(int value, ...) {
    return value + 15;
}
#else
int variant1848_6(int value) {
    return value + 15;
}
#endif
struct Node1848_7 {
    int value;
    struct Node1848_7 *next;
};

int length1848_7(const struct Node1848_7 *head) {
    int count = 0;
    const struct Node1848_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
