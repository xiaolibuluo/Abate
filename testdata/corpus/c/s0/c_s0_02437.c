
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1652_0(int value, ...) {
    return value + 44;
}
#else
int variant1652_0(int value) {
    return value + 44;
}
#endif
#define MASK1652_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1652_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1652_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1652_2)(int value);

typedef struct Slot1652_2 {
    int slots[16];
    predicate1652_2 accept;
} Slot1652_2;

int slot_accept1652_2(Slot1652_2 *store, int value) {
    return store->accept(value) + 58;
}
struct Node1652_3 {
    int value;
    struct Node1652_3 *next;
};

int length1652_3(const struct Node1652_3 *head) {
    int count = 0;
    const struct Node1652_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
int scan1652_4(const int *begin, const int *end) {
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
    return total + 72;
}
struct Node1652_5 {
    int value;
    struct Node1652_5 *next;
};

int length1652_5(const struct Node1652_5 *head) {
    int count = 0;
    const struct Node1652_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
int scan1652_6(const int *begin, const int *end) {
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
    return total + 86;
}
#define MASK1652_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1652_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1652_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
