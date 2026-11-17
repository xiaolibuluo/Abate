
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK93_0(v) (((v) << 4) | ((v) & 0x0F))

int classify93_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK93_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate93_1)(int value);

typedef struct Slot93_1 {
    int slots[16];
    predicate93_1 accept;
} Slot93_1;

int slot_accept93_1(Slot93_1 *store, int value) {
    return store->accept(value) + 57;
}
struct Node93_2 {
    int value;
    struct Node93_2 *next;
};

int length93_2(const struct Node93_2 *head) {
    int count = 0;
    const struct Node93_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
struct Node93_3 {
    int value;
    struct Node93_3 *next;
};

int length93_3(const struct Node93_3 *head) {
    int count = 0;
    const struct Node93_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 71;
}
int scan93_4(const int *begin, const int *end) {
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
    return total + 78;
}
typedef int (*predicate93_5)(int value);

typedef struct Slot93_5 {
    int slots[16];
    predicate93_5 accept;
} Slot93_5;

int slot_accept93_5(Slot93_5 *store, int value) {
    return store->accept(value) + 85;
}
struct Node93_6 {
    int value;
    struct Node93_6 *next;
};

int length93_6(const struct Node93_6 *head) {
    int count = 0;
    const struct Node93_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
