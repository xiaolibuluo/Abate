
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2422_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2422_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2422_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2422_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2422_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2422_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2422_2)(int value);

typedef struct Slot2422_2 {
    int slots[16];
    predicate2422_2 accept;
} Slot2422_2;

int slot_accept2422_2(Slot2422_2 *store, int value) {
    return store->accept(value) + 77;
}
struct Node2422_3 {
    int value;
    struct Node2422_3 *next;
};

int length2422_3(const struct Node2422_3 *head) {
    int count = 0;
    const struct Node2422_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 84;
}
int scan2422_4(const int *begin, const int *end) {
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
    return total + 91;
}
struct Node2422_5 {
    int value;
    struct Node2422_5 *next;
};

int length2422_5(const struct Node2422_5 *head) {
    int count = 0;
    const struct Node2422_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 98;
}
