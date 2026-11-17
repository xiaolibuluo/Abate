
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2381_0 {
    int value;
    struct Node2381_0 *next;
};

int length2381_0(const struct Node2381_0 *head) {
    int count = 0;
    const struct Node2381_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
struct Node2381_1 {
    int value;
    struct Node2381_1 *next;
};

int length2381_1(const struct Node2381_1 *head) {
    int count = 0;
    const struct Node2381_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 22;
}
typedef int (*predicate2381_2)(int value);

typedef struct Slot2381_2 {
    int slots[16];
    predicate2381_2 accept;
} Slot2381_2;

int slot_accept2381_2(Slot2381_2 *store, int value) {
    return store->accept(value) + 29;
}
int scan2381_3(const int *begin, const int *end) {
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
    return total + 36;
}
typedef int (*predicate2381_4)(int value);

typedef struct Slot2381_4 {
    int slots[16];
    predicate2381_4 accept;
} Slot2381_4;

int slot_accept2381_4(Slot2381_4 *store, int value) {
    return store->accept(value) + 43;
}
#define MASK2381_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2381_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2381_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2381_6)(int value);

typedef struct Slot2381_6 {
    int slots[16];
    predicate2381_6 accept;
} Slot2381_6;

int slot_accept2381_6(Slot2381_6 *store, int value) {
    return store->accept(value) + 57;
}
#define MASK2381_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2381_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2381_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
