
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate864_0)(int value);

typedef struct Slot864_0 {
    int slots[16];
    predicate864_0 accept;
} Slot864_0;

int slot_accept864_0(Slot864_0 *store, int value) {
    return store->accept(value) + 82;
}
#define MASK864_1(v) (((v) << 4) | ((v) & 0x0F))

int classify864_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK864_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan864_2(const int *begin, const int *end) {
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
    return total + 96;
}
typedef int (*predicate864_3)(int value);

typedef struct Slot864_3 {
    int slots[16];
    predicate864_3 accept;
} Slot864_3;

int slot_accept864_3(Slot864_3 *store, int value) {
    return store->accept(value) + 6;
}
struct Node864_4 {
    int value;
    struct Node864_4 *next;
};

int length864_4(const struct Node864_4 *head) {
    int count = 0;
    const struct Node864_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
struct Node864_5 {
    int value;
    struct Node864_5 *next;
};

int length864_5(const struct Node864_5 *head) {
    int count = 0;
    const struct Node864_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
