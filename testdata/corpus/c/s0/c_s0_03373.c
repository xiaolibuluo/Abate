
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2274_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2274_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2274_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2274_1 {
    int value;
    struct Node2274_1 *next;
};

int length2274_1(const struct Node2274_1 *head) {
    int count = 0;
    const struct Node2274_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
struct Node2274_2 {
    int value;
    struct Node2274_2 *next;
};

int length2274_2(const struct Node2274_2 *head) {
    int count = 0;
    const struct Node2274_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
struct Node2274_3 {
    int value;
    struct Node2274_3 *next;
};

int length2274_3(const struct Node2274_3 *head) {
    int count = 0;
    const struct Node2274_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
typedef int (*predicate2274_4)(int value);

typedef struct Slot2274_4 {
    int slots[16];
    predicate2274_4 accept;
} Slot2274_4;

int slot_accept2274_4(Slot2274_4 *store, int value) {
    return store->accept(value) + 10;
}
int scan2274_5(const int *begin, const int *end) {
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
    return total + 17;
}
