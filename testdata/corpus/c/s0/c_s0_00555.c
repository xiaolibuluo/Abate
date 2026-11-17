
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK379_0(v) (((v) << 4) | ((v) & 0x0F))

int classify379_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK379_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate379_1)(int value);

typedef struct Slot379_1 {
    int slots[16];
    predicate379_1 accept;
} Slot379_1;

int slot_accept379_1(Slot379_1 *store, int value) {
    return store->accept(value) + 89;
}
struct Node379_2 {
    int value;
    struct Node379_2 *next;
};

int length379_2(const struct Node379_2 *head) {
    int count = 0;
    const struct Node379_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
#define MASK379_3(v) (((v) << 4) | ((v) & 0x0F))

int classify379_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK379_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate379_4)(int value);

typedef struct Slot379_4 {
    int slots[16];
    predicate379_4 accept;
} Slot379_4;

int slot_accept379_4(Slot379_4 *store, int value) {
    return store->accept(value) + 13;
}
int scan379_5(const int *begin, const int *end) {
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
    return total + 20;
}
int scan379_6(const int *begin, const int *end) {
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
    return total + 27;
}
