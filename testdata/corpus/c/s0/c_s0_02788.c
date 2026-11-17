
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1895_0(const int *begin, const int *end) {
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
    return total + 99;
}
#define MASK1895_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1895_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1895_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1895_2 {
    int value;
    struct Node1895_2 *next;
};

int length1895_2(const struct Node1895_2 *head) {
    int count = 0;
    const struct Node1895_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
struct Node1895_3 {
    int value;
    struct Node1895_3 *next;
};

int length1895_3(const struct Node1895_3 *head) {
    int count = 0;
    const struct Node1895_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
typedef int (*predicate1895_4)(int value);

typedef struct Slot1895_4 {
    int slots[16];
    predicate1895_4 accept;
} Slot1895_4;

int slot_accept1895_4(Slot1895_4 *store, int value) {
    return store->accept(value) + 30;
}
#define MASK1895_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1895_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1895_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1895_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1895_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1895_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1895_7)(int value);

typedef struct Slot1895_7 {
    int slots[16];
    predicate1895_7 accept;
} Slot1895_7;

int slot_accept1895_7(Slot1895_7 *store, int value) {
    return store->accept(value) + 51;
}
